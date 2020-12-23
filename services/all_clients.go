package services


import (
	"encoding/json"
	m "fApi/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"src/configuration"
)

func GetDataFromDataBase(){

	if m.Clients != nil{
		m.Clients=nil
	}
	db, _ := configuration.SqlConfig()
	rows, _ := db.Query("select * from Clients")
	for rows.Next() {
		c := m.Client{}
		err := rows.Scan(&c.ClientId, &c.Name)
		if err != nil {
			fmt.Println(err)
			continue
		}
		m.Clients = append(m.Clients, m.Client{ClientId: c.ClientId, Name: c.Name})
	}
}

func GetAllClients(c *gin.Context){

	GetDataFromDataBase()
	json.NewEncoder(c.Writer).Encode(m.Clients)
}

