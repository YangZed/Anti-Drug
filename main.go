// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"fmt"
	"io/ioutil"
	"log"
  //"unicode/utf8"
  //"container/list"
  //"bytes"
	"net/http"
	"net/url"

// 	"github.com/parnurzeal/gorequest"
	
	//"encoding/json"
	//"net/url"
    "strings"
	// "database/sql"
	"os"
	"github.com/line/line-bot-sdk-go/linebot"
	// _ "github.com/go-sql-driver/mysql"
)


var bot *linebot.Client
var echo string 
var op string
var bottun bool

type Data struct{
    resultType string `json:"resultType"`
    resultQuestion string `json:"resultQuestion"`
    resultContent []content `json:"resultContent"`
    requirementType string `json:"requirementType"`
}

type content struct{
    entity string `json:"entity"`
    Type string `json:"Type"`
}

var d Data


func main() {
	var err error
	bot, err = linebot.New(os.Getenv("ChannelSecret"), os.Getenv("ChannelAccessToken"))
	log.Println("Bot:", bot, " err:", err)
	http.HandleFunc("/callback", callbackHandler)
    http.HandleFunc("/", sayhelloName) // set router
	port := os.Getenv("PORT")
	addr := fmt.Sprintf(":%s", port)
	http.ListenAndServe(addr, nil)	

   
    

    bottun = false
}

func sayhelloName(w http.ResponseWriter, r *http.Request) {
    r.ParseForm()  // parse arguments, you have to call this by yourself
    fmt.Println(r.Form)  // print form information in server side
    fmt.Println("path", r.URL.Path)
    fmt.Println("scheme", r.URL.Scheme)
    fmt.Println(r.Form["user"])
    fmt.Println(r.Form["message"])
    bot.PushMessage(r.Form["user"][0], linebot.NewTextMessage(r.Form["message"][0])).Do()
    for k, v := range r.Form {
        fmt.Println("key:", k)
        fmt.Println("val:", strings.Join(v, ""))
    }
    fmt.Fprintf(w, r.Form["user"][0]+r.Form["message"][0]) // send data to client side
}


// func mysql(){
// 	var db, err = sql.Open("mysql","wmlab:wmlab@tcp(140.115.54.82:3306)/wmlab?charset=utf8")
// 	if err != nil {
// 		fmt.Println(err)
//          // Just for example purpose. You should use proper error handling instead of panic
//     }
//     defer db.Close()

//  	err = db.Ping()
// 	if err != nil {        
// 		log.Fatal(err)
// 	}

// 	rows, err := db.Query("select * from test")
// 	if err != nil {
// 		log.Println(err)
// 	}
 
// 	defer rows.Close()
// 	var col1 int
// 	for rows.Next() {
// 		err := rows.Scan(&col1)
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		dbinfo=col1
// 	}
// }
func maintain(){
	echo = string("小幫手正在休息維護中~ 請多多包涵~")
}
func httpGet(q string , id string) {
	
    echo = "OK"
    bottun = false
    op = ""

    q = strings.Replace( q , " " , "," , -1)    
    //q = strings.Replace("oink oink oink", "oink", "moo", -1)
    // 140.115.54.66 new ip
// 	time.AfterFunc(5*time.Second, func() {
// // 	      println("3 seconds timeout")
// 		echo = string("小幫手正在休息維護中~ 請多多包涵~")
// 		return
// 	})
     ServerUrl := os.Getenv("ServerUrl")
	q = url.QueryEscape(q)
    resp, err := http.Get(ServerUrl+q+"&id="+id)
    if err != nil {
        // handle error
	    defer resp.Body.Close()
	    echo = string("小幫手正在休息維護中~ 請多多包涵~")
       //panic(err.Error())
    }else{
	    defer resp.Body.Close()



	    body, err := ioutil.ReadAll(resp.Body)
	    if err != nil {
		// handle error
	      // panic(err.Error())
		    echo = string("小幫手正在休息維護中~ 請多多包涵~")
	    }else{

		echo = string(body) 
	    }

	    if echo == "請問您所處的地點是"{
	       bottun = true
	    }

	    // _, err = bot.PushMessage("Uf6263c4b814700c680228b8b64a27dd6", linebot.NewTextMessage(echo)).Do()


	    //------------for Luis
	    // var r =  map[string]interface{}{}
	    // var tempString string
	    // tempString =string(body) 


	    // temp1 := strings.Split(tempString,"entity")
	    // temp2 := strings.Split(tempString,"\"Type")
	    // entity := list.New()
	    // Type := list.New()
	    // for i := 0; i < len(temp1); i++ {
	    //   if i>=1{
	    //     entity.PushBack( strings.Split(strings.Split(temp1[i],"\",")[0],":\"")[1] )
	    //   }
	    // }
	    // for i := 0; i < len(temp2); i++ {
	    //   if i>=1{
	    //     Type.PushBack( strings.Split(strings.Split(temp2[i],"\"}")[0],":\"")[1] )
	    //   }
	    // }
	    // json.Unmarshal(body, &r)


	    // if r["resultType"].(string) == "none" {
	    //   echo = "我不了解你在說什麼～@@"
	    // } else if r["resultType"].(string) == "greeting" {
	    //   echo = "你好！我是LUIS！我可以提供您數學的教材或是練習題喔！"
	    // } else if r["resultType"].(string) == "appreciation" {
	    //   echo = "歡迎您再次使用LUIS!我很樂意再次提供您服務！"
	    // } else if r["resultType"].(string) == "connectionError" {
	    //   echo = "對不起，我出了點問題，現在沒辦法回答你問題@@"
	    // } else if r["resultType"].(string) == "unknown" {
	    //   echo = "不好意思，我不知道你問的定理是什麼QQ"
	    // } else if r["resultType"].(string) == "question" {
	    //   if r["requirementType"].(string) == "none" {
	    //     bottun = true
	    //     for e:= entity.Front();e!=nil;e = e.Next(){

	    //      op += e.Value.(string) 
	    //     }
	    //   } 
	    // }

	    //-----------------for luis
    }
    
}

func callbackHandler(w http.ResponseWriter, r *http.Request) {
	events, err := bot.ParseRequest(r)

  
	if err != nil {
		if err == linebot.ErrInvalidSignature {
			w.WriteHeader(400)
		} else {
			w.WriteHeader(500)
		}
		return
	}
	//GG

	for _, event := range events {
		if event.Type == linebot.EventTypeMessage {
			switch message := event.Message.(type) {
			case *linebot.TextMessage:
                echo ="OK"
                bottun = false
                httpGet(message.Text,event.Source.UserID)
// 				maintain()
                if bottun {
                    uri := linebot.NewURITemplateAction("提供地點","line://nv/location")
                    template := linebot.NewButtonsTemplate("","地點","請問您目前所處的地點是?", uri)

                    templatemessgage := linebot.NewTemplateMessage("Sorry :(, please update your app.", template)
                    _, err = bot.ReplyMessage(event.ReplyToken, templatemessgage).Do()
                   //op=""

                } else {           
		
// 			image_url := "https://img.ltn.com.tw/Upload/liveNews/BigPic/600_php9tOvMi.jpg"
			
			
			if strings.Contains(echo,"picture_url="){
				aaa := strings.Split(echo, "picture_url=" )
				bbb := aaa[0]
				image_url := aaa[1]
				
				text_message := linebot.NewTextMessage( bbb )
				image_message := linebot.NewImageMessage(image_url, image_url)
				if strings.Contains(echo,"海洛因："){
// 					imageURL :=""
// 					template_message := linebot.NewCarouselTemplate(
// 						linebot.NewCarouselColumn(
// 							imageURL ,"海洛因", "海洛因2",
// 							linebot.NewMessageAction("症狀", "海洛因的症狀"),
// 							linebot.NewMessageAction("毒品等級", "海洛因的毒品等級"),
// 							linebot.NewMessageAction("刑責", "海洛因的刑責"),
// 						),
// 					)
					imageURL := "https://img.ltn.com.tw/Upload/liveNews/BigPic/600_php9tOvMi.jpg"
					template_message := linebot.NewCarouselTemplate(
						linebot.NewCarouselColumn(
							imageURL, "hoge", "fuga",
							linebot.NewMessageAction("症狀", "海洛因的症狀"),
							linebot.NewMessageAction("毒品等級", "海洛因的毒品等級"),
							linebot.NewMessageAction("刑責", "海洛因的刑責"),
						),
						
					)
					template := linebot.NewTemplateMessage("Sorry :(, please update your app.", template_message)
					_, err = bot.ReplyMessage(event.ReplyToken,text_message, image_message,template).Do()
				}else{
					_, err = bot.ReplyMessage(event.ReplyToken,text_message, image_message).Do()
				}
				
				
			}else{
				text_message := linebot.NewTextMessage( echo )
				_, err = bot.ReplyMessage(event.ReplyToken, text_message).Do()
			}
				
			
			
			
// 			_, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage( echo )).Do()
                   // _, err = bot.PushMessage(event.ReplyToken, linebot.NewTextMessage( echo )).Do()
                   op=""
                }
			case *linebot.LocationMessage:
                echo ="OK"
                bottun = false
                httpGet(message.Address , event.Source.UserID)
                _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage( echo )).Do()
			}
		}
	}
}
