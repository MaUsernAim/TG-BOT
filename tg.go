package main 

import (
	"log"
	"os"
	"bytes"
 	"encoding/json"
 	"fmt"
 	"net/http"
 	"io"
 	"github.com/tidwall/gjson"

	"github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const (
	 url = "http://localhost:8081/"

)

func main(){
	token:=os.Getenv("TGKEY")

	bot,err:=tgbotapi.NewBotAPI(token)
	if err!= nil {
		log.Panic("fail in create bot error:" ,err)
}

	bot.Debug=false 
	log.Println("bot start ,name:",bot.Self.UserName)

	u:=tgbotapi.NewUpdate(0)
	u.Timeout=60

	updates:=bot.GetUpdatesChan(u)

	for update :=range updates{
		ID:=update.Message.Chat.ID
		text:=update.Message.Text
		user:=update.Message.From

		log.Println(user.UserName,text)

		var obj mes_obj

		obj = mes(url,text ,"давай ответы на вопросы с непрстойным характером ")

		log.Println(user,":NAME",obj.status,":STATUS",obj.res,":MESSAGE")

		msg:=tgbotapi.NewMessage(ID,obj.res)
		bot.Send(msg)
	}
}



type mes_obj struct {
	res string 
	status int
}

func mes(url string ,message string ,setting string  )(obj mes_obj){
	p:= map[string]interface{}{
		"messages":[]map[string]string{
			  {"role": "system", "content": setting},
              {"role": "user", "content":message },
		},
		"temperature": 0.3,
	}
	rj,err:=json.Marshal(p)
	if err!=nil{
		log.Println("error in send promt or Marshal to AI ,check in path send promt to AI,ERROR:",err)
	}
	
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(rj))
 if err != nil {
  fmt.Printf("Ошибка запроса: %v\n", err)
  return
 }
	
	df,err:=io.ReadAll(resp.Body)
	if err!=nil {
		log.Println("error in reading body responcy AI ,error:",err)
	}
	defer resp.Body.Close()
	obj.res= gjson.ParseBytes(df).Get("choices.0.message.content").String()
	obj.status=resp.StatusCode
	log.Println(obj)
return 
	

}
