package main

import (
	"./mongo"
	"./util"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
)


var items []mongo.Ping
var Txt = ""


func Reply(ctx *gin.Context) {
	ctx.JSON(http.StatusOK,gin.H{
		"code":200,
		"success":true,
		"txt":Txt,
		"relevantPing":items,
	})
}

func Newping(ctx *gin.Context) {
	body,_ := ioutil.ReadAll(ctx.Request.Body)
	Txt = string(body)
	pingMap := make(map[string]string)
	err := json.Unmarshal([]byte(Txt),&pingMap)
	if err!=nil {
		log.Fatal(err)
	}
	keyWords := util.ExtractKeyWords(Txt)
	items = mongo.FindRelevantPing(keyWords)
	//log.Println(keyWords)
}


