package odata

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"
)

type Client struct {
	username string
	password string
	endpoint string
	client   *http.Client
}

// Return initialized odata client
func NewClient(username, password, endpoint string) *Client {
	return &Client{username, password, endpoint, &http.Client{}}
}

func (c *Client) get(url string) (string, error) {
	req, err := http.NewRequest("GET", c.endpoint+url, nil)
	if err != nil {
		return "", err
	}
	req.SetBasicAuth(c.username, c.password)
	resp, err := c.client.Do(req)
	if err != nil {
		return "", err
	}
	body, _ := ioutil.ReadAll(resp.Body)
	if resp.StatusCode/100 == 2 {
		return string(body), nil
	}
	return "", errors.New(string(body))
}

func (c *Client) post(url string, data []byte) (string, error) {
	req, err := http.NewRequest("POST", c.endpoint+url, bytes.NewReader(data))
	if err != nil {
		return "", err
	}
	req.SetBasicAuth(c.username, c.password)
	resp, err := c.client.Do(req)
	if err != nil {
		return "", err
	}
	body, _ := ioutil.ReadAll(resp.Body)
	if resp.StatusCode/100 == 2 {
		return string(body), nil
	}
	return "", errors.New(string(body))
}

func (c *Client) patch(url string, data []byte) (string, error) {
	req, err := http.NewRequest("PATCH", c.endpoint+url, bytes.NewReader(data))
	if err != nil {
		return "", err
	}
	req.SetBasicAuth(c.username, c.password)
	resp, err := c.client.Do(req)
	if err != nil {
		return "", err
	}
	body, _ := ioutil.ReadAll(resp.Body)
	if resp.StatusCode/100 == 2 {
		return string(body), nil
	}
	return "", errors.New(string(body))
}

func (c *Client) delete(url string) error {
	req, err := http.NewRequest("DELETE", c.endpoint+url, nil)
	if err != nil {
		return err
	}

	req.SetBasicAuth(c.username, c.password)
	resp, err := c.client.Do(req)
	if err != nil {
		return err
	}
	body, _ := ioutil.ReadAll(resp.Body)
	if len(body) != 0 {
		return errors.New(string(body))
	}
	return nil
}
