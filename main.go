package main

import "github.com/gin-gonic/gin"

func main() {

	router := gin.Default()
	router.GET("/reply",Reply)
	router.POST("/newping",Newping)
	_ = router.Run(Port)

}



