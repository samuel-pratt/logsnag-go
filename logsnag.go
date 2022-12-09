package logsnag

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

type LogSnag struct {
	Token   string
	Project string
}

func (logsnag *LogSnag) GetProject() string {
	return logsnag.Project
}

func (logsnag *LogSnag) Publish(channel string, event string, icon string, tags map[string]any, notify bool) bool {
	url := "https://api.logsnag.com/v1/log"
	method := "POST"

	var pairs []string

	for key, value := range tags {
		pairs = append(pairs, fmt.Sprintf(`%s: %v`, key, value))
	}

	// Join the slice with commas and enclose the whole thing in curly braces to create the JSON string
	description := strings.Join(pairs, ", ")

	payload := strings.NewReader(`{
		"project": "` + logsnag.GetProject() + `",
		"channel": "` + channel + `",
		"event": "` + event + `",
		"description": "` + description + `",
		"icon": "` + icon + `",
		"notify": "` + strconv.FormatBool(notify) + `"
	}`)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return false
	}

	req.Header.Add("Authorization", "Bearer "+logsnag.Token)
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return false
	}
	defer res.Body.Close()

	_, err = ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return false
	}

	return true
}

func (logsnag *LogSnag) Insight(title string, value string, icon string) bool {
	url := "https://api.logsnag.com/v1/insight"
	method := "POST"

	payload := strings.NewReader(`{
		"title":` + title + `,
		"value":` + value + `,
		"icon":` + icon + `,
	}`)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		fmt.Println(err)
		return false
	}

	req.Header.Add("Authorization", "Bearer "+logsnag.Token)
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return false
	}
	defer res.Body.Close()

	_, err = ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return false
	}

	return true
}

func NewLogSnag(token string, project string) LogSnag {
	return LogSnag{
		Token:   token,
		Project: project,
	}
}
