package provider

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
)

type remoteCaller struct {
	url string
}

type Response struct {
	Status  string
	Message string
}

func NewProvider(url string) ExternalCall {
	return &remoteCaller{url: url}
}

// Call makes the remote api call and returns the result
func (c *remoteCaller) Call() (interface{}, error) {
	resp, err := http.Get(c.url)
	if err != nil {
		log.Println("call to url:", c.url, " returned an error:", err)
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil,errors.New("error from gateway")
	}

	var providerResponse Response

	err = json.NewDecoder(resp.Body).Decode(&providerResponse)
	if err != nil {
		log.Println("Error reading providerResponse body ", err)
		return nil,err
	}

	return providerResponse, nil
}
