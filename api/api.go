package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"org/sonatype/nx/config"
)

func Get(endpoint string, expectedStatus int) (string, error) {
	client := &http.Client{}

	log.Println("Invoking " + endpoint)

	req, err := http.NewRequest(http.MethodGet, "http://"+config.Host()+":"+config.Port()+"/service/rest/"+endpoint, nil)
	if err != nil {
		return "", err
	}

	req.SetBasicAuth("admin", "admin123")
	req.Header.Set("Content-Type", "application/json; charset=utf-8")

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	if resp.StatusCode != expectedStatus {
		return "", errors.New(fmt.Sprintf("api call %s failed with status %d", endpoint, resp.StatusCode))
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", errors.New(fmt.Sprintf("api call %s unable to read body", endpoint))
	}

	return string(body), nil
}

func Post(endpoint string, payload any, expectedStatus int) error {
	return Request(http.MethodPost, endpoint, payload, expectedStatus)
}

func Put(endpoint string, payload any, expectedStatus int) error {
	return Request(http.MethodPut, endpoint, payload, expectedStatus)
}

func Request(method string, endpoint string, payload any, expectedStatus int) error {
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		panic(err)
	}

	client := &http.Client{}

	log.Println("Invoking " + endpoint + ": " + string(jsonPayload))

	req, err := http.NewRequest(method, "http://"+config.Host()+":"+config.Port()+"/service/rest/"+endpoint,
		bytes.NewBuffer(jsonPayload))
	if err != nil {
		return err
	}

	req.SetBasicAuth("admin", "admin123")
	req.Header.Set("Content-Type", "application/json; charset=utf-8")

	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode != expectedStatus {
		return errors.New(fmt.Sprintf("api call %s failed with status %d", endpoint, resp.StatusCode))
	}

	return nil
}
