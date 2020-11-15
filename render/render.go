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

type calendarInfo struct {
	Time string
	Date string
}

type weatherInfo struct {
	Temperature string
	Conditions  string
}

type babyInfo struct {
	Days int
}

type renderInfo struct {
	Todo     []todoistTask
	Calendar calendarInfo
	Weather  weatherInfo
	Baby     babyInfo
}

type weatherApiFeatureTemp struct {
	Value float64 `json:"value"`
}

type weatherApiFeatureProps struct {
	TextDescription string                `json:"textDescription"`
	Temperature     weatherApiFeatureTemp `json:"temperature"`
}

type weatherApiFeature struct {
	Properties weatherApiFeatureProps `json:"properties"`
}

type weatherApiInfo struct {
	Features []weatherApiFeature `json:"features"`
}

func render(info renderInfo) error {
	dat, err := ioutil.ReadFile(os.Getenv("TEMPLATE"))
	if err != nil {
		return err
	}

	t, err := template.New("render").Parse(string(dat))
	if err != nil {
		return err
	}

	err = t.Execute(os.Stdout, info)
	if err != nil {
		return err
	}
	return nil
}

func getTodos() ([]todoistTask, error) {
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

func daysToBaby() babyInfo {
	newYork, _ := time.LoadLocation("America/New_York")
	dueDate := time.Date(2021, 06, 15, 0, 0, 0, 0, newYork)
	duration := time.Until(dueDate)
	days := int(duration.Hours() / 24.0)
	return babyInfo{days}
}

func getCalendarInfo() calendarInfo {
	_, month, day := time.Now().Date()
	hour, min, _ := time.Now().Clock()
	var hourStr string
	var minStr string
	var amPm string
	if hour > 12 {
		hour = hour - 12
		amPm = "pm"
	} else {
		amPm = "am"
	}
	hourStr = fmt.Sprintf("%d", hour)
	if min < 10 {
		minStr = fmt.Sprintf("0%d", min)
	} else {
		minStr = fmt.Sprintf("%d", min)
	}
	return calendarInfo{
		Time: fmt.Sprintf("%s:%s%s", hourStr, minStr, amPm),
		Date: fmt.Sprintf("%s %d", month.String(), day),
	}
}

func getWeatherInfo() (weatherInfo, error) {
	client := &http.Client{}
	url := fmt.Sprintf("https://api.weather.gov/zones/forecast/%s/observations", os.Getenv("WEATHER_ZONE"))
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return weatherInfo{}, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return weatherInfo{}, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return weatherInfo{}, err
	}
	var apiInfo weatherApiInfo
	err = json.Unmarshal(body, &apiInfo)
	if err != nil {
		return weatherInfo{}, err
	}
	if len(apiInfo.Features) == 0 {
		return weatherInfo{}, nil
	}
	tempF := (apiInfo.Features[0].Properties.Temperature.Value * (9.0 / 5.0)) + 32.0
	return weatherInfo{
		Temperature: fmt.Sprintf("%0.0f", tempF),
		Conditions:  apiInfo.Features[0].Properties.TextDescription,
	}, nil
}

func main() {
	items, err := getTodos()
	if err != nil {
		log.Fatal(err)
		return
	}
	weatherInfo, err := getWeatherInfo()
	if err != nil {
		log.Fatal(err)
		return
	}
	err = render(renderInfo{
		Todo:     items,
		Calendar: getCalendarInfo(),
		Weather:  weatherInfo,
		Baby:     daysToBaby(),
	})
	if err != nil {
		log.Fatal(err)
		return
	}
}
