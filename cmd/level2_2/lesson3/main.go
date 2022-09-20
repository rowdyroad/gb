package main

import (
	"net/http"


	"lesson3mod/pkg/test"
	testX "lesson3mod/pkg/x/test"
	gin "yandex.ru/module"
)


func main() {
	r := gin.Default()
	r.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"test":test.Test(), "testX": testX.Test()})
	})
	r.Run(":80")
}