package main

import (
	"fmt"
	"net/http"
	"net/url"
)

func main() {
	base, err := url.Parse("https://www.phpsouthwales.uk/jsonapi/node/talk")
	if err != nil {
		return
	}

	params := url.Values{}
	params.Add("fields[node--person", "title")
	params.Add("fields[node--talk", "title,field_speakers")
	params.Add("include", "field_speakers")
	base.RawQuery = params.Encode()

	response, err := http.Get(base.String())
	if err != nil {
		fmt.Println("There was an error")
	}

	fmt.Println(response)
}
