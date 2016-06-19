package druid

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

type Client struct {
	URL string
}

type DruidError struct {
	Error string `json:"error"`
}

func New(url string) *Client {
	return &Client{url}
}

func (c *Client) RunQuery(query *AggregationQuery) ([]byte, error) {
	jsonBody, err := query.GetJSON()

	if err != nil {
		return nil, err
	}

	readerBody := bytes.NewReader(jsonBody)

	resp, err := http.Post(c.URL, "application/json", readerBody)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if resp.StatusCode != 200 {
		var queryError DruidError
		err = json.Unmarshal(body, &queryError)
		if err != nil {
			return nil, err
		}
		return nil, errors.New(queryError.Error)
	}

	return body, err
}
