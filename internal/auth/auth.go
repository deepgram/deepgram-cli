package auth

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"deepgram-cli/internal/common"
)

type DeviceCodeResponse struct {
	DeviceCode      string `json:"device_code"`
	VerificationURI string `json:"verification_uri"`
	ExpiresIn       int    `json:"expires_in"`
	Interval        int    `json:"interval"`
}

type DeviceCodeRequest struct {
	ClientId string   `json:"client_id"`
	Hostname string   `json:"hostname"`
	Scopes   []string `json:"scopes"`
}

type AccessTokenResponse struct {
	AccessToken string `json:"access_token"`
	ProjectId   string `json:"project_id"`
}

type AccessTokenRequest struct {
	ClientId   string `json:"client_id"`
	Hostname   string `json:"hostname"`
	DeviceCode string `json:"device_code"`
}

type ErrorResponse struct {
	Error            string `json:"error"`
	ErrorDescription string `json:"error_description"`
}

func RequestDeviceCode(client_id string, hostname string, scopes []string) (DeviceCodeResponse, error) {
	var (
		response    DeviceCodeResponse
		requestArgs DeviceCodeRequest
	)

	requestArgs = DeviceCodeRequest{
		ClientId: client_id,
		Hostname: hostname,
		Scopes:   scopes,
	}

	postData, _ := json.Marshal(requestArgs)
	url := fmt.Sprintf("%s/api/auth/device/code", common.BaseUrl)

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(postData))
	if err != nil {
		return response, err
	}

	if resp.StatusCode != http.StatusCreated {
		return response, ErrUnknown
	}

	defer resp.Body.Close() //skipcq: GO-S2307

	json.NewDecoder(resp.Body).Decode(&response)

	return response, nil
}

func PollForAccessToken(client_id string, hostname string, device_code string, interval int) (*AccessTokenResponse, error) {
	var (
		response      AccessTokenResponse
		errorResponse ErrorResponse
	)

	queryParams := url.Values{}
	queryParams.Set("device_code", device_code)
	queryParams.Set("client_id", client_id)
	queryParams.Set("hostname", hostname)

	url := fmt.Sprintf("%s/api/auth/device/token?%s", common.BaseUrl, queryParams.Encode())

	for {
		resp, err := http.Get(url)
		if err != nil {
			return nil, err
		}

		defer resp.Body.Close()

		if resp.StatusCode == http.StatusCreated {
			if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
				return nil, err
			}

			return &response, nil
		}

		if resp.StatusCode != http.StatusNotFound {
			return nil, fmt.Errorf("%s, description: %s", errorResponse.Error, errorResponse.ErrorDescription)
		}

		if err := json.NewDecoder(resp.Body).Decode(&errorResponse); err != nil {
			return nil, err
		}

		time.Sleep(time.Duration(interval) * time.Second)
	}
}
