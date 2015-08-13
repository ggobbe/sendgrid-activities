package main

import (
	"flag"
	"fmt"
	"stuff/sendgrid-activities/api"
)

func main() {
	const required = "REQUIRED"

	apiUser := flag.String("username", required, "Username to connect to the SendGrid API")
	apiKey := flag.String("password", required, "Password to connect to the SendGrid API")
	flag.Parse()

	if *apiUser == required || *apiKey == required {
		fmt.Println("Username and password are required.")
		return
	}

	sg := api.NewSendGridClient(*apiUser, *apiKey)

	bounces, err := sg.GetBounces()
	if err != nil {
		fmt.Println("The response of the Bounces API doesn't match the expected format.\nThis can happen if you have entered the wrong credentials.")
	} else {
		printActivities(bounces, "bounce")
	}

	blocks, err := sg.GetBlocks()
	if err != nil {
		fmt.Println("The response of the Blocks API doesn't match the expected format.\nThis can happen if you have entered the wrong credentials.")
	} else {
		printActivities(blocks, "blocks")
	}

	invalidEmails, err := sg.GetInvalidEmails()
	if err != nil {
		fmt.Println("The response of the InvalidEmails API doesn't match the expected format.\nThis can happen if you have entered the wrong credentials.")
	} else {
		printActivities(invalidEmails, "bounce")
	}
}

func printActivities(activities []*api.Activity, activityType string) {
	if len(activities) == 0 {
		return
	}

	for _, activity := range activities {
		fmt.Printf("\"%s\",\"%s\",\"%s\",\"%s\",\"%s\"\n", activityType, activity.Created, activity.Status, activity.Email, activity.Reason)
	}
}
