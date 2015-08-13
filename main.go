package main

import (
	"flag"
	"fmt"
	"sendgrid-activities/api"
	"sort"
	"strings"
)

const (
	all           = "all"
	bounces       = "bounces"
	blocks        = "blocks"
	invalidEmails = "invalidEmails"
)

func main() {
	const required = "REQUIRED"

	apiUser := flag.String("username", required, "Username to connect to the SendGrid API")
	apiKey := flag.String("password", required, "Password to connect to the SendGrid API")
	activityType := flag.String("type", "all", "Types of activities to retrieve (all, bounces, blocks, invalidEmails)")
	flag.Parse()

	if *apiUser == required || *apiKey == required {
		fmt.Println("Username and password are required")
		return
	}

	sg := api.NewSendGridClient(*apiUser, *apiKey)

	switch strings.ToLower(*activityType) {
	case all:
		printAll(sg)
	case bounces:
		activities, err := sg.GetBounces()
		printActivities(activities, err)
	case blocks:
		activities, err := sg.GetBlocks()
		printActivities(activities, err)
	case strings.ToLower(invalidEmails):
		activities, err := sg.GetInvalidEmails()
		printActivities(activities, err)
	default:
		fmt.Printf("Invalid activity type: %s\n", *activityType)
	}

}

func printAll(sg *api.SGClient) {
	var activities []api.Activity
	bouncesActivities, err := sg.GetBounces()
	if err == nil {
		activities = append(activities, bouncesActivities...)
	} else {
		printError(err)
	}

	blocksActivities, err := sg.GetBlocks()
	if err == nil {
		activities = append(activities, blocksActivities...)
	} else {
		printError(err)
	}

	invalidEmailsActivities, err := sg.GetInvalidEmails()
	if err == nil {
		activities = append(activities, invalidEmailsActivities...)
	} else {
		printError(err)
	}

	printActivities(activities, nil)
}

func printActivities(activities []api.Activity, err error) {
	if err != nil {
		printError(err)
	}

	if len(activities) == 0 {
		fmt.Println("No activities retrieved.")
		return
	}

	sort.Sort(api.ByDate(activities))

	fmt.Printf("\"Created\",\"Status\",\"Email\",\"Reason\"\n")
	for _, activity := range activities {
		fmt.Printf("\"%s\",\"%s\",\"%s\",\"%s\"\n", activity.Created, activity.Status, activity.Email, activity.Reason)
	}
}

func printError(err error) {
	fmt.Printf("Error whilst retrieving activities: %s\n", err)
	fmt.Printf("This usually happens if you have entered the wrong credentials.\n")
}
