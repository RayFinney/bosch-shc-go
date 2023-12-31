package bosch_shc_go

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

type Client struct {
	httpClient          *http.Client
	shcIp               string
	shcPort             int
	validateCertificate bool
	apiVersion          string
}

// NewClient creates a new Client, apiKey and httpClient are optional.
func NewClient(httpClient *http.Client, options Options) BoschShcGo {
	c := Client{
		httpClient:          httpClient,
		shcIp:               options.ShcIp,
		shcPort:             options.ShcPort,
		validateCertificate: options.ValidateCertificate,
		apiVersion:          options.ApiVersion,
	}
	if c.httpClient == nil {
		t := http.DefaultTransport.(*http.Transport).Clone()
		t.MaxIdleConns = 100
		t.MaxConnsPerHost = 100
		t.MaxIdleConnsPerHost = 100

		c.httpClient = &http.Client{
			Timeout:   60 * time.Second,
			Transport: t,
		}
	}
	if c.shcPort == 0 {
		c.shcPort = 8444
	}
	if c.apiVersion == "" {
		c.apiVersion = "1.0"
	}

	return &c
}

func (c Client) getUrl(suffix string) string {
	return fmt.Sprintf("https://%s:%d/smarthome%s", c.shcIp, c.shcPort, suffix)
}

func (c Client) setHeader(req *http.Request) {
	req.Header.Set("Accept", "application/json")
	req.Header.Set("api-version", c.apiVersion)
}

func checkForErrorsInResponse(body []byte) error {
	advErr := Error{}
	err := json.Unmarshal(body, &advErr)
	if err != nil {
		log.Println("Error unmarshalling error response:", err)
	}
	return errors.New(advErr.ErrorCode)
}

func (c Client) GetDevices() ([]Device, error) {
	req, err := http.NewRequest(http.MethodGet, c.getUrl("/devices"), nil)
	if err != nil {
		return nil, err
	}
	c.setHeader(req)
	response, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Println("Error closing response body:", err)
		}
	}(response.Body)
	if response.StatusCode != http.StatusOK {
		return nil, checkForErrorsInResponse(body)
	}
	devices := []Device{}
	err = json.Unmarshal(body, &devices)
	if err != nil {
		return nil, err
	}
	return devices, nil
}

func (c Client) GetDevice(id string) (Device, error) {
	req, err := http.NewRequest(http.MethodGet, c.getUrl(fmt.Sprintf("/devices/%s", id)), nil)
	if err != nil {
		return Device{}, err
	}
	c.setHeader(req)
	response, err := c.httpClient.Do(req)
	if err != nil {
		return Device{}, err
	}
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return Device{}, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Println("Error closing response body:", err)
		}
	}(response.Body)
	if response.StatusCode != http.StatusOK {
		return Device{}, checkForErrorsInResponse(body)
	}
	device := Device{}
	err = json.Unmarshal(body, &device)
	if err != nil {
		return Device{}, err
	}
	return device, nil
}

func (c Client) GetRooms() ([]Room, error) {
	req, err := http.NewRequest(http.MethodGet, c.getUrl("/rooms"), nil)
	if err != nil {
		return nil, err
	}
	c.setHeader(req)
	response, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Println("Error closing response body:", err)
		}
	}(response.Body)
	if response.StatusCode != http.StatusOK {
		return nil, checkForErrorsInResponse(body)
	}
	rooms := []Room{}
	err = json.Unmarshal(body, &rooms)
	if err != nil {
		return nil, err
	}
	return rooms, nil
}

func (c Client) GetRoom(id string) (Room, error) {
	req, err := http.NewRequest(http.MethodGet, c.getUrl(fmt.Sprintf("/rooms/%s", id)), nil)
	if err != nil {
		return Room{}, err
	}
	c.setHeader(req)
	response, err := c.httpClient.Do(req)
	if err != nil {
		return Room{}, err
	}
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return Room{}, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Println("Error closing response body:", err)
		}
	}(response.Body)
	if response.StatusCode != http.StatusOK {
		return Room{}, checkForErrorsInResponse(body)
	}
	room := Room{}
	err = json.Unmarshal(body, &room)
	if err != nil {
		return Room{}, err
	}
	return room, nil
}

func (c Client) GetScenarios() ([]Scenario, error) {
	req, err := http.NewRequest(http.MethodGet, c.getUrl("/scenarios"), nil)
	if err != nil {
		return nil, err
	}
	c.setHeader(req)
	response, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Println("Error closing response body:", err)
		}
	}(response.Body)
	if response.StatusCode != http.StatusOK {
		return nil, checkForErrorsInResponse(body)
	}
	scenarios := []Scenario{}
	err = json.Unmarshal(body, &scenarios)
	if err != nil {
		return nil, err
	}
	return scenarios, nil
}

func (c Client) GetScenario(id string) (Scenario, error) {
	req, err := http.NewRequest(http.MethodGet, c.getUrl(fmt.Sprintf("/scenarios/%s", id)), nil)
	if err != nil {
		return Scenario{}, err
	}
	c.setHeader(req)
	response, err := c.httpClient.Do(req)
	if err != nil {
		return Scenario{}, err
	}
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return Scenario{}, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Println("Error closing response body:", err)
		}
	}(response.Body)
	if response.StatusCode != http.StatusOK {
		return Scenario{}, checkForErrorsInResponse(body)
	}
	scenario := Scenario{}
	err = json.Unmarshal(body, &scenario)
	if err != nil {
		return Scenario{}, err
	}
	return scenario, nil
}

func (c Client) TriggerScenario(id string) error {
	req, err := http.NewRequest(http.MethodPost, c.getUrl(fmt.Sprintf("/scenarios/%s/triggers", id)), nil)
	if err != nil {
		return err
	}
	c.setHeader(req)
	response, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Println("Error closing response body:", err)
		}
	}(response.Body)
	if response.StatusCode != http.StatusAccepted {
		return checkForErrorsInResponse(body)
	}
	return nil
}

func (c Client) GetMessages() ([]Message, error) {
	req, err := http.NewRequest(http.MethodGet, c.getUrl("/messages"), nil)
	if err != nil {
		return nil, err
	}
	c.setHeader(req)
	response, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Println("Error closing response body:", err)
		}
	}(response.Body)
	if response.StatusCode != http.StatusOK {
		return nil, checkForErrorsInResponse(body)
	}
	messages := []Message{}
	err = json.Unmarshal(body, &messages)
	if err != nil {
		return nil, err
	}
	return messages, nil
}
