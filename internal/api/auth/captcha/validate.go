package captcha

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"
)

var (
	SecretKey       = os.Getenv("hcaptchaSecret") // environmental variable containing secret.
	VerifyURL       = "https://hcaptcha.com/siteverify"
	request_counter int
)

type hCaptchaResponse struct {
	Success bool `json:"success"`
	// other fields like error codes can be added as needed
}

func counter() {
	request_counter++
	if request_counter%10 == 0 {
		currentTime := time.Now().Format("15:04-02/01/2006")
		fmt.Printf("%d Captcha requests %v\n", request_counter, currentTime)
	}
}

func VerifyHCaptcha(token string) (bool, error) {
	counter()
	data := url.Values{}
	data.Set("secret", SecretKey)
	data.Set("response", token)

	// Create a new POST request with the data
	req, err := http.NewRequest("POST", VerifyURL, bytes.NewBufferString(data.Encode()))
	if err != nil {
		return false, err
	}

	// Set the appropriate headers
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// Create a new HTTP client and send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return false, err
	}

	// Parse the JSON response
	fmt.Printf("captcha: %s\n", string(body))
	var hCaptchaResp hCaptchaResponse
	err = json.Unmarshal(body, &hCaptchaResp)
	if err != nil {
		return false, err
	}

	return hCaptchaResp.Success, nil
}

func test() {
	token := "your_h-captcha-response_token" // Replace with the actual token from the client

	success, err := VerifyHCaptcha(token)
	if err != nil {
		log.Fatalf("Error verifying hCaptcha: %v", err)
	}

	if success {
		fmt.Println("hCaptcha verification succeeded")
	} else {
		fmt.Println("hCaptcha verification failed")
	}
}
