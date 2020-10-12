package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

type todoistTask struct {
	Content string         `json:"content"`
	Due     todoistDuedate `json:"due"`
}

type todoistDuedate struct {
	Date   string `json:"date"`
	Amount string
}

func render(items []todoistTask) error {
	dat, err := ioutil.ReadFile(os.Getenv("TEMPLATE"))
	if err != nil {
		return err
	}

	t, err := template.New("render").Parse(string(dat))
	if err != nil {
		return err
	}

	data := struct {
		Items []todoistTask
	}{Items: items}
	err = t.Execute(os.Stdout, data)
	if err != nil {
		return err
	}
	return nil
}

func getItems() ([]todoistTask, error) {
	client := &http.Client{}
	url := fmt.Sprintf("https://api.todoist.com/rest/v1/tasks?project_id=%s", os.Getenv("TODOIST_PROJECT_ID"))
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return []todoistTask{}, err
	}
	req.Header.Add("Authorization", "Bearer "+os.Getenv("TODOIST_TOKEN"))
	resp, err := client.Do(req)
	if err != nil {
		return []todoistTask{}, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []todoistTask{}, err
	}
	items := make([]todoistTask, 0)
	err = json.Unmarshal(body, &items)
	if err != nil {
		return []todoistTask{}, err
	}
	for i := 0; i < len(items); i++ {
		if items[i].Due.Date != "" {
			dateStrct, err := time.Parse("2006-01-02", items[i].Due.Date)
			if err == nil {
				items[i].Due.Amount = fmt.Sprintf("%d Days", (dateStrct.Unix()-time.Now().Unix())/60/60/24)
			}
		}
	}
	return items, nil
}

func main() {
	items, err := getItems()
	if err != nil {
		log.Fatal(err)
		return
	}
	err = render(items)
	if err != nil {
		log.Fatal(err)
		return
	}
}
