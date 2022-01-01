package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/pkg/errors"
)


func HappyNewYear() (string, error) {
	url := "https://coronavirusapifr.herokuapp.com/data/live/france"
	resp, err := http.Get(url)
	if err != nil || resp.StatusCode != 200 {
		return "", errors.Wrap(err, "cannot track Covid19")
	}
	defer resp.Body.Close()

	var m []map[string]interface{}
	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(bodyBytes, &m)
	if err != nil || len(m) == 0 {
		return "", errors.Wrap(err, "cannot get Covid19 data")
	}

	if v, ok := m[0]["hosp"].(float64); ok {
		if v < 10 {
			wish := fmt.Sprintf("Happy new year %d", time.Now().Year())
			return wish, nil
		}
	}

	return "Ah shit, here we go again.", nil
}

func main() {
	res, err := HappyNewYear()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println(res)
}