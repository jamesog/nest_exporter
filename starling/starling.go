package starling

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"runtime/debug"
	"strings"
	"time"
)

type Client struct {
	APIKey    string
	BaseURL   string
	userAgent string
}

func NewClient(baseURL, key string) Client {
	version := "devel"
	if bi, ok := debug.ReadBuildInfo(); ok {
		version = bi.Main.Version
	}

	return Client{
		APIKey: key,
		// Remove any trailing / from the URL; Starling's API doesn't handle extra slahes properly
		BaseURL:   strings.TrimSuffix(baseURL, "/"),
		userAgent: fmt.Sprintf("nest_exporter/%s", version),
	}
}

var (
	errBadRequest   = errors.New("bad request")
	errUnauthorized = errors.New("unauthorized")
)

func (c Client) makeRequest(method, endpoint string, payload []byte) ([]byte, error) {
	qs := url.Values{}
	qs.Set("key", c.APIKey)
	client := http.Client{
		Timeout: 5 * time.Second,
	}
	ep, err := url.Parse(c.BaseURL)
	if err != nil {
		return nil, fmt.Errorf("couldn't parse base URL: %w", err)
	}
	ep.Path += endpoint
	ep.RawQuery = qs.Encode()

	req, err := http.NewRequest(method, ep.String(), bytes.NewReader(payload))
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", c.userAgent)

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(resp.Body)
	defer func() { err = resp.Body.Close() }()

	switch resp.StatusCode {
	// We return the body for 400 and 401 to parse the JSON response from the Starling API.
	case http.StatusBadRequest:
		return body, errBadRequest
	case http.StatusUnauthorized:
		return body, errUnauthorized
	default:
		return body, err
	}
}

type Status struct {
	APIVersion      float64 `json:"apiVersion"`
	APIReady        bool    `json:"apiReady"`
	ConnectedToNest bool    `json:"connectedToNest"`
	AppName         string  `json:"appName"`
	Permissions     struct {
		Read   bool `json:"read"`
		Write  bool `json:"write"`
		Camera bool `json:"camera"`
	} `json:"permissions"`
}

func (c Client) Status() (*Status, error) {
	resp, err := c.makeRequest("GET", "/status", nil)
	if err != nil {
		if (errors.Is(err, errUnauthorized) || errors.Is(err, errBadRequest)) && resp != nil {
			status := Error{}
			err := json.Unmarshal(resp, &status)
			if err != nil {
				return nil, err
			}
			return nil, errors.New(status.Message)
		}
		return nil, err
	}
	// Device properties have a response object that has a status and
	status := Status{}
	err = json.Unmarshal(resp, &status)
	if err != nil {
		status := Error{}
		err = json.Unmarshal(resp, &status)
		if err != nil {
			return nil, err
		}
		return nil, errors.New(status.Message)
	}
	return &status, nil
}

type Devices struct {
	Status  string             `json:"status"`
	Devices []CommonProperties `json:"devices"`
}

func (c Client) Devices() (*Devices, error) {
	resp, err := c.makeRequest("GET", "/devices", nil)
	if err != nil {
		return nil, err
	}
	devices := Devices{}
	err = json.Unmarshal(resp, &devices)
	if err != nil {
		return nil, err
	}
	return &devices, nil
}

type Device struct {
	Status     string `json:"status"`
	Properties struct {
		CommonProperties
	} `json:"properties"`
}

func (c Client) deviceProperties(id string) ([]byte, error) {
	return c.makeRequest("GET", "/devices/"+id, nil)
}

type ThermostatResponse struct {
	Status     string               `json:"status"`
	Properties ThermostatProperties `json:"properties"`
}

func (c Client) ThermostatProperties(id string) (*ThermostatProperties, error) {
	resp, err := c.deviceProperties(id)
	if err != nil {
		return nil, err
	}

	device := ThermostatResponse{}
	err = json.Unmarshal(resp, &device)
	if err != nil {
		return nil, err
	}

	return &device.Properties, nil
}

type ProtectResponse struct {
	Status     string            `json:"status"`
	Properties ProtectProperties `json:"properties"`
}

func (c Client) ProtectProperties(id string) (*ProtectProperties, error) {
	resp, err := c.deviceProperties(id)
	if err != nil {
		return nil, err
	}

	device := ProtectResponse{}
	err = json.Unmarshal(resp, &device)
	if err != nil {
		return nil, err
	}

	return &device.Properties, nil
}

type Error struct {
	Status  string `json:"status,omitempty"`
	Code    string `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}
