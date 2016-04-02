package api

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

const bouncesAPI = "https://api.sendgrid.com/api/bounces.get.json?"
const blocksAPI = "https://api.sendgrid.com/api/blocks.get.json?"
const invalidEmailsAPI = "https://api.sendgrid.com/api/invalidemails.get.json?"

// Activity type will contain a bounced email details
type Activity struct {
	Status  string `json:"status"`
	Created string `json:"created"`
	Reason  string `json:"reason"`
	Email   string `json:"email"`
}

// Error type will contain the error details
type Error struct {
	Error string `json:"error"`
}

// GetBounces retrieves bounced emails as per https://sendgrid.com/docs/API_Reference/Web_API/bounces.html
func (sg *SGClient) GetBounces() ([]Activity, error) {
	return sg.getActivities(bouncesAPI)
}

// GetBlocks retrieves bounced emails as per https://sendgrid.com/docs/API_Reference/Web_API/blocks.html
func (sg *SGClient) GetBlocks() ([]Activity, error) {
	return sg.getActivities(blocksAPI)
}

// GetInvalidEmails retrieves bounced emails as per https://sendgrid.com/docs/API_Reference/Web_API/invalid_emails.html
func (sg *SGClient) GetInvalidEmails() ([]Activity, error) {
	return sg.getActivities(invalidEmailsAPI)
}

func (sg *SGClient) getActivities(apiURL string) ([]Activity, error) {
	resp, err := http.Get(apiURL + "api_user=" + sg.apiUser + "&api_key=" + sg.apiKey + "&date=1")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var activities []Activity
	err = json.Unmarshal(body, &activities)
	if err != nil {
		var error Error
		err = json.Unmarshal(body, &error)
		if err != nil {
			return nil, err
		}
		return nil, errors.New(error.Error)
	}

	return activities, nil
}
