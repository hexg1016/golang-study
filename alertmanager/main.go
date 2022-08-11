package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type AlertMessage struct {
	Labels      AlertLabels      ` json:"labels"`
	Annotations AlertAnnotations ` json:"annotations"`
}
type AlertLabels struct {
	AlertName string ` json:"alertname"`
	Service   string ` json:"service"`
	Severity  string ` json:"severity"`
	Instance  string ` json:"instance"`
}

type AlertAnnotations struct {
	Title       string ` json:"title"`
	Description string ` json:"description"`
	Summary     string ` json:"summary"`
}

func main() {
	var arr []AlertMessage
	message := AlertMessage{
		Labels: AlertLabels{
			AlertName: "programPush5",
			Severity:  "critical",
			Service:   "game",
			Instance:  "test",
		},
		Annotations: AlertAnnotations{
			Title:       "协程崩溃1",
			Description: "xiecheng bengkui",
			Summary:     "ssss",
		},
	}
	arr = append(arr, message)
	bytes, err := json.Marshal(&arr)
	if err != nil {
		return
	}
	fmt.Println(string(bytes))
	resp, err := http.Post("http://localhost:9093/api/v2/alerts", "application/json", strings.NewReader(string(bytes)))
	if err == nil {
		defer resp.Body.Close()
		body, _ := ioutil.ReadAll(resp.Body)
		fmt.Println(string(body))
	}
}
