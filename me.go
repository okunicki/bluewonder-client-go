package bluewonder

import (
	"bytes"
	"encoding/json"
	"net/http"
)

const MeUrl string = DefaultRestUrl + "/me"

type MeResponse struct {
	Name   string   `json:"name"`
	Type   string   `json:"type"`
	Groups []string `json:"groups"`
}

func (c *Client) GetMe() (*MeResponse, error) {
	req, err := http.NewRequest("GET", MeUrl, nil)
	if err != nil {
		return nil, err
	}

	res, err := c.DoRequest(req)
	if err != nil {
		return nil, err
	}

	var me MeResponse
	err = json.Unmarshal(res, &me)
	if err != nil {
		return nil, err
	}

	return &me, nil
}

func (c *Client) UpdateMe(updatedMe *MeResponse) error {
	j, err := json.Marshal(updatedMe)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", MeUrl, bytes.NewBuffer(j))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	_, err = c.DoRequest(req)
	if err != nil {
		return err
	}

	return nil
}
