package main

import (
	"first-api/configuration"
	"first-api/services"
	"fmt"
	logger "git.homebank.kz/HomeBank/elastic-logger.v2"
	"github.com/gin-gonic/gin"
	"net/http"

	_ "github.com/lib/pq"
)


func main(){

	_cfg := configuration.NewConfiguration("config.json")
	_log := logger.NewLogger("")
	_handler := services.NewHandler(_cfg, _log)


	r := gin.New()
	r.GET("/", index)
	r.GET("/overwrite",  _handler.OverwriteData)
	r.GET("/clients", _handler.GetAllClients)
	r.GET("/clients/id/:id", _handler.GetClient)
	r.GET("/truncate", _handler.Truncate)
	r.GET("/new-file/:n", services.New)
	http.ListenAndServe(":10010", r)
}

func index(c *gin.Context){
	fmt.Fprintf(c.Writer,`<a href="/overwrite">Переписать данные в с файла в базу</a>`)
	fmt.Fprintf(c.Writer,`<br/><a href="/clients">Вернуть все данные которые есть у нас в базе по клиентам</a>`)
	fmt.Fprintf(c.Writer,`<br/><a href="/truncate">Удалить все данные по клиентам</a>`)
}
