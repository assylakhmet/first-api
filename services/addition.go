package services

import (
	"encoding/json"
	"first-api/configuration"
	m "first-api/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

func (handler *Handler) GetClient(c *gin.Context){
	c.Writer.Header().Set("Content-Type", "application/json")
	GetDataFromDataBase(handler)
	for _, item := range m.Clients {
		if strconv.Itoa(item.ClientId) ==  c.Params.ByName("id") {
			json.NewEncoder(c.Writer).Encode(item)
			return
		}
	}
	json.NewEncoder(c.Writer).Encode(&m.Client{})
}


func (handler *Handler) Truncate(c *gin.Context){
	db, err := configuration.SqlConfig(handler.Config)
	if err != nil{
		fmt.Println(err.Error())
	}
	db.Query("truncate clients")
	m.Clients=nil
	fmt.Fprintf(c.Writer,"Done!")
	fmt.Println("truncate done;")
}
