package main

import (
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	router := gin.New()
	logHandler, err := NewLogHandler()
	if err != nil {
		log.Fatal(err)
	}

	router.POST("/log", logHandler.WriteLog)

	err = router.Run(":80")
	if err != nil {
		log.Fatal(err)
	}

}
