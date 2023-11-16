package bluewonder

import (
	"fmt"
	"io"
	"net/http"
)

const DefaultRestUrl string = "https://mockapigwd.azurewebsites.net/"

type Client struct {
	HttpClient *http.Client
	ApiKey     string
	Host       string
	Base       string
}

func NewClient(apiKey string) *Client {
	return &Client{
		HttpClient: http.DefaultClient,
		ApiKey:     apiKey,
	}
}

func (c *Client) NewRequest(path string) (*http.Request, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/%s", DefaultRestUrl, path), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

func (c *Client) DoRequest(req *http.Request) ([]byte, error) {
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.ApiKey))
	res, err := c.HttpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	if res.StatusCode == http.StatusOK || res.StatusCode == http.StatusNoContent {
		return body, err
	} else {
		return nil, fmt.Errorf("status: %d, body: %s", res.StatusCode, body)
	}
}
