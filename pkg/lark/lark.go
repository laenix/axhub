package lark

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func SendMessage(apiUrl, axhubUrl, target, time string) {
	client := &http.Client{}
	var req *http.Request
	// reqdata := `{
	// 	"msg_type": "text",
	// 	"content": {"text": "` + "消息通知:" + msg + `"}
	// }`
	reqdata := `{
		"msg_type": "interactive",
		"card": {
			"config": {
				"wide_screen_mode": true
			},
			"header": {
				"title": {
					"tag": "plain_text",
					"content": "🟢推送了产品图"
				},
				"template": "blue"
			},
			"elements": [{
					"tag": "div",
					"text": {
						"content": "⚙ 目标仓库:  ` + target + `  \n📅 推送时间: ` + time + `",
						"tag": "lark_md"
					}
				},
				{
					"actions": [{
						"tag": "button",
						"text": {
							"content": "查看产品图",
							"tag": "plain_text"
						},
						"type": "primary",
						"url": "` + axhubUrl + `"
					}],
					"tag": "action"
				}
			]
		}
	}`
	req, err := http.NewRequest("POST", apiUrl, strings.NewReader(reqdata))
	if err != nil {
		log.Fatalln(err)
	}

	req.Header.Add("contentType", "application/json")
	req.Header.Add("Accept", "application/json, text/plain, */*")

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	html, _ := goquery.NewDocumentFromReader(resp.Body)
	fmt.Println(html.Text())
}
