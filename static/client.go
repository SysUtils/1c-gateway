package odata

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/prometheus/client_golang/prometheus"
	"io/ioutil"
	"log"
	"net/http"
	"sync/atomic"
	"time"
)

type Client struct {
	queue    chan Request
	counter  int32
	username string
	password string
	endpoint string
	client   *http.Client
	PoolSize int
	metrics  prometheus.Gauge
}

type Request struct {
	httpRequest *http.Request
	error       chan error
	response    chan *http.Response
}

type ErrorResponse struct {
	Error `json:"odata.error"`
}

type Error struct {
	Code    string `json:"code"`
	Message struct {
		Lang  string `json:"lang"`
		Value string `json:"value"`
	} `json:"message"`
	HttpCode int `json:"http_code"`
}

func (err Error) Error() string {
	return err.Message.Value
}

func (err Error) Extensions() map[string]interface{} {
	return map[string]interface{}{
		"code":     err.Code,
		"message":  err.Message.Value,
		"httpCode": err.HttpCode,
	}
}

// Return initialized odata client
func NewClient(username, password, endpoint string, poolSize, queueSize int, transport http.RoundTripper, gauge prometheus.Gauge) *Client {
	httpClient := &http.Client{}
	if transport != nil {
		httpClient.Transport = transport
	}
	client := &Client{username: username, password: password, endpoint: endpoint, client: httpClient, metrics: gauge}
	client.queue = make(chan Request, queueSize)
	for i := 0; i < poolSize; i++ {
		go client.worker()
	}
	return client
}

func (c *Client) get(url string) ([]byte, error) {
	req, err := http.NewRequest("GET", c.endpoint+url, nil)
	if err != nil {
		return nil, err
	}
	req.SetBasicAuth(c.username, c.password)
	return c.do(req)
}

func (c *Client) post(url string, data []byte) ([]byte, error) {
	req, err := http.NewRequest("POST", c.endpoint+url, bytes.NewReader(data))
	if err != nil {
		return nil, err
	}
	req.SetBasicAuth(c.username, c.password)
	return c.do(req)
}

func (c *Client) patch(url string, data []byte) ([]byte, error) {
	req, err := http.NewRequest("PATCH", c.endpoint+url, bytes.NewReader(data))
	if err != nil {
		return nil, err
	}
	req.SetBasicAuth(c.username, c.password)
	return c.do(req)
}

func (c *Client) delete(url string) error {
	req, err := http.NewRequest("DELETE", c.endpoint+url, nil)
	if err != nil {
		return err
	}
	req.SetBasicAuth(c.username, c.password)
	body, err := c.do(req)
	if err != nil {
		return err
	}
	if len(body) != 0 {
		return errors.New(string(body))
	}
	return nil
}

func (c *Client) worker() {
	if c.queue == nil {
		log.Panicln("service not initialized")
	}
	for work := range c.queue {
		atomic.AddInt32(&c.counter, 1)
		if c.metrics != nil {
			c.metrics.Inc()
		}

		resp, err := c.client.Do(work.httpRequest)
		if err != nil {
			if work.error != nil {
				work.error <- err
			} else {
				log.Panicf("error: %v, request: %v", err, work.httpRequest)
			}
		}
		if work.response != nil {
			work.response <- resp
		}
		atomic.AddInt32(&c.counter, -1)
		if c.metrics != nil {
			c.metrics.Dec()
		}
	}
}

func (c *Client) Close() {
	for {
		time.Sleep(time.Microsecond)
		if len(c.queue) == 0 {
			time.Sleep(time.Nanosecond)
			if atomic.LoadInt32(&c.counter) == 0 {
				break
			}
		}
	}
}

func (c *Client) do(request *http.Request) ([]byte, error) {
	respChan := make(chan *http.Response)
	errChan := make(chan error)
	c.queue <- Request{
		response:    respChan,
		error:       errChan,
		httpRequest: request,
	}
	select {
	case err := <-errChan:
		return nil, err
	case resp := <-respChan:
		defer resp.Body.Close()
		body, _ := ioutil.ReadAll(resp.Body)
		if resp.StatusCode/100 == 2 {
			return body, nil
		}
		errResponse := ErrorResponse{}
		if err := json.Unmarshal(body, &errResponse); err != nil {
			return nil, errors.New(string(body))
		}

		errResponse.Error.HttpCode = resp.StatusCode
		return nil, errResponse.Error
	}
}
