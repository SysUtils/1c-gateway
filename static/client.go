package odata

import (
	"bytes"
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

// Return initialized odata client
func NewClient(username, password, endpoint string, poolSize, queueSize int, transport http.RoundTripper) *Client {
	metric := prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "one_c_http_parallel_requests",
		Help: "1c load from this service",
	})
	prometheus.MustRegister(metric)
	httpClient := &http.Client{}
	if transport != nil {
		httpClient.Transport = transport
	}
	client := &Client{username: username, password: password, endpoint: endpoint, client: httpClient, metrics: metric}
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
		c.metrics.Inc()
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
		c.metrics.Dec()
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
		return nil, errors.New(resp.Status + "\nBody:\n" + string(body))
	}
}
