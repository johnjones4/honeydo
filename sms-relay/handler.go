package main

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"net/url"
	"os"
	"sort"
	"strconv"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type todoistTask struct {
	Content   string `json:"content"`
	ProjectID int    `json:"project_id"`
}

func makeTodoistTask(content string) error {
	projectID, err := strconv.Atoi(os.Getenv("TODOIST_PROJECT_ID"))
	if err != nil {
		return err
	}
	task := todoistTask{Content: content, ProjectID: projectID}
	requestBody, err := json.Marshal(task)
	if err != nil {
		return err
	}
	client := &http.Client{}
	req, err := http.NewRequest("POST", "https://api.todoist.com/rest/v1/tasks", bytes.NewBuffer(requestBody))
	if err != nil {
		return err
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+os.Getenv("TODOIST_TOKEN"))
	_, err = client.Do(req)
	if err != nil {
		return err
	}
	return nil
}

func generateTwilioSignature(authToken string, URL string, postForm url.Values) string {
	keys := make([]string, 0, len(postForm))
	for key := range postForm {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	str := URL
	for _, key := range keys {
		str += key + postForm[key][0]
	}
	mac := hmac.New(sha1.New, []byte(authToken))
	mac.Write([]byte(str))
	expectedMac := mac.Sum(nil)
	return base64.StdEncoding.EncodeToString(expectedMac)
}

func validateTwilioRequest(authToken string, URL string, request events.APIGatewayProxyRequest, formValues url.Values) error {
	expectedTwilioSignature := generateTwilioSignature(authToken, URL, formValues)
	if request.Headers["X-Twilio-Signature"] != expectedTwilioSignature {
		return errors.New("Bad X-Twilio-Signature")
	}
	return nil
}

func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	formValues, err := url.ParseQuery(request.Body)
	if err != nil {
		log.Fatal(err)
		return events.APIGatewayProxyResponse{Body: "", StatusCode: 500}, err
	}
	err = validateTwilioRequest(os.Getenv("TWILIO_AUTH_TOKEN"), os.Getenv("ENDPOINT_URL"), request, formValues)
	if err != nil {
		log.Fatal(err)
		return events.APIGatewayProxyResponse{Body: "", StatusCode: 500}, err
	}
	err = makeTodoistTask(formValues["Body"][0])
	if err != nil {
		log.Fatal(err)
		return events.APIGatewayProxyResponse{Body: "", StatusCode: 500}, err
	}
	return events.APIGatewayProxyResponse{Body: "", StatusCode: 200}, nil
}

func main() {
	lambda.Start(Handler)
}
