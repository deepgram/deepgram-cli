package auth

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"deepgram-cli/internal/common"
)

type Session struct {
	DeviceCode      string `json:"device_code"`
	UserCode        string `json:"user_code"`
	VerificationURI string `json:"verification_uri"`
	ExpiresIn       int    `json:"expires_in"`
	Interval        int    `json:"interval"`
}

func RequestDeviceCode(ppid int, hostname string, scope []string) (Session, error) {
	var (
		result Session
	)

	args := make(map[string]interface{})
	args["id"] = ppid
	args["hostname"] = hostname
	args["scopes"] = scope

	postData, _ := json.Marshal(args)

	fmt.Println(string(postData))

	url := fmt.Sprintf("%s/api/auth/device/code", common.BaseUrl)

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(postData))
	if err != nil {
		return result, err
	}

	if resp.StatusCode != 201 {
		return result, ErrUnknown
	}

	defer resp.Body.Close() //skipcq: GO-S2307

	json.NewDecoder(resp.Body).Decode(&result)

	return result, nil
}
