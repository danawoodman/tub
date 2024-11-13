package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type (
	Gizwits struct {
		appID string
	}
	Header struct {
		Key   string
		Value string
	}
)

func NewGizwits(appID string) *Gizwits {
	return &Gizwits{appID: appID}
}

func (self *Gizwits) AuthRequest(method, path string, payload any, token string) (*http.Response, error) {
	headers := []Header{{"X-Gizwits-User-Token", token}}
	return self.Request(method, path, payload, headers)
}

func (self *Gizwits) Request(method, path string, payload any, headers []Header) (*http.Response, error) {
	base := "https://usapi.gizwits.com/app"

	var body io.Reader
	if payload != nil {
		reqBytes, err := json.Marshal(payload)
		if err != nil {
			return nil, fmt.Errorf("error marshaling request body: %w", err)
		}
		body = bytes.NewBuffer(reqBytes)
	}

	req, err := http.NewRequest(method, base+path, body)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	req.Header.Set("X-Gizwits-Application-Id", self.appID)
	req.Header.Set("Content-Type", "application/json")

	for _, header := range headers {
		req.Header.Set(header.Key, header.Value)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("[%s] %s\n", method, path)
		fmt.Println("  Headers:")
		for k, v := range req.Header {
			fmt.Printf("    %s: %s\n", k, v)
		}
		if body != nil {
			fmt.Println("  Body:")
			fmt.Printf("    %+v\n", payload)
		}
		return nil, fmt.Errorf("ERROR: [%d] %s", resp.StatusCode, resp.Status)
	}

	return resp, nil
}
