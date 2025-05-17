package reymit

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type Response struct {
	Ok   bool `json:"ok"`
	Data Data `json:"data"`
}

type Data struct {
	Donates []Donate `json:"donates"`
}

type Donate struct {
	ID                  string  `json:"id"`
	UserID              string  `json:"user_id"`
	Authority           string  `json:"authority"`
	RefID               string  `json:"ref_id"`
	Time                float64 `json:"time"`
	Amount              float64 `json:"amount"`
	Currency            string  `json:"currency"`
	TomanAmount         int     `json:"toman_amount"`
	Name                string  `json:"name"`
	Description         string  `json:"description"`
	CensoredName        string  `json:"censored_name"`
	CensoredDescription string  `json:"censored_description"`
	Phone               string  `json:"phone"`
	Email               string  `json:"email"`
	ForGoal             string  `json:"for_goal"`
}

func GetLastDonations(token string) (*Response, error) {
	// Create HTTP client with timeout
	client := &http.Client{
		Timeout: time.Second * 10,
	}

	// Create the request
	url := fmt.Sprintf("https://api.reymit.ir/user/%s/donates/last-donates", token)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}

	// Make the request
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error making request: %v", err)
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %v", err)
	}

	// Check status code
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	// Parse JSON response
	var response Response
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, fmt.Errorf("error parsing JSON: %v", err)
	}

	return &response, nil
}
