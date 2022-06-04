package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/jedib0t/go-pretty/v6/table"
)

func DisplayLeaderboard() {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"Speaker name", "Number of talks"})
	t.AppendRow([]interface{}{"Oliver Davies", 5})
	t.Render()
}

func GetApiData() (string, error) {
	apiEndpointUrl, err := getApiEndpointUrl()

	if err != nil {
		return "", err
	}

	response, err := http.Get(apiEndpointUrl)

	if err != nil {
		fmt.Println("There was an error")
		return "", err
	}

	defer response.Body.Close()

	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
		return "", err
	}

	return string(responseBody), nil
}

func getApiEndpointUrl() (string, error) {
	base, err := url.Parse("https://www.phpsouthwales.uk/jsonapi/node/talk")

	if err != nil {
		return "", err
	}

	params := url.Values{}
	params.Add("fields[node--person", "title")
	params.Add("fields[node--talk", "title,field_speakers")
	params.Add("include", "field_speakers")
	base.RawQuery = params.Encode()

	return base.String(), nil
}

func main() {
	apiData, err := GetApiData()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(apiData)

	// TODO: populate with data from the API.
	DisplayLeaderboard()
}
