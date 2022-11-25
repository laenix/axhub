package axhub

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/PuerkitoBio/goquery"
)

type AxhubModel struct {
	Code int `json:"code"`
	Data struct {
		Pid            int       `json:"pid"`
		Name           string    `json:"name"`
		Path           string    `json:"path"`
		ShareMode      int       `json:"shareMode"`
		Type           int       `json:"type"`
		Software       int       `json:"software"`
		IsPublic       bool      `json:"isPublic"`
		IsExpress      bool      `json:"isExpress"`
		SidebarStatus  int       `json:"sidebarStatus"`
		Engine         int       `json:"engine"`
		IsAutogenerate bool      `json:"isAutogenerate"`
		GenerateStatus int       `json:"generateStatus"`
		GenerateTime   time.Time `json:"generateTime"`
		IsOwner        bool      `json:"isOwner"`
		Role           string    `json:"role"`
		Right          string    `json:"right"`
		IsLogined      bool      `json:"isLogined"`
	} `json:"data"`
}

func GetAxhubInfo(url string) (string, string) {
	client := &http.Client{}
	var req *http.Request
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("Accept", "application/json, text/plain, */*")
	if err != nil {
		log.Fatalln(err)
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()
	html, _ := goquery.NewDocumentFromReader(resp.Body)
	var axhubModel AxhubModel
	json.Unmarshal([]byte(html.Text()), &axhubModel)
	return axhubModel.Data.Name, axhubModel.Data.GenerateTime.String()
}
