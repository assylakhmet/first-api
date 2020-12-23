package main

import (
	"first-api/services"
	"first-api/services/addition"
	"first-api/services/overwrite"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"

	_ "github.com/lib/pq"
)


func main(){
	handleRequest()
}

func handleRequest(){
	r := gin.New()
	r.GET("/", index)
	r.GET("/overwrite", overwrite.OverwriteData)
	r.GET("/clients", services.GetAllClients)
	r.GET("/clients/id/:id", addition.GetClient)
	r.GET("/truncate", addition.Truncate)
	r.GET("/new-file/:n", addition.New)
	http.ListenAndServe(":10010", r)
}

func index(c *gin.Context){
	fmt.Fprintf(c.Writer,`<a href="/overwrite">Переписать данные в с файла в базу</a>`)
	fmt.Fprintf(c.Writer,`<br/><a href="/clients">Вернуть все данные которые есть у нас в базе по клиентам</a>`)
	fmt.Fprintf(c.Writer,`<br/><a href="/truncate">Удалить все данные по клиентам</a>`)
}
