package main

import (
	"net/url"
	"os"

	"github.com/jedib0t/go-pretty/v6/table"
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

	// TODO: populate with data from the API.
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"Speaker name", "Number of talks"})
	t.AppendRow([]interface{}{"Oliver Davies", 5})
	t.Render()
}
