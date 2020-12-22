package services


import (
"encoding/json"
"fApi/configuration"
m "fApi/models"
"fmt"
"github.com/gin-gonic/gin"
)

func GetDataFromDataBase(c *gin.Context){
	c.Writer.Header().Set("Content-Type", "application/json")
	db, _ := configuration.SqlConfig()
	rows, _ := db.Query("select * from Clients")
	for rows.Next(){
		c := m.Client{}
		err := rows.Scan(&c.ClientId, &c.Name)
		if err != nil{
			fmt.Println(err)
			continue
		}
		m.Clients = append(m.Clients, m.Client{ClientId: c.ClientId,Name: c.Name})
	}
	json.NewEncoder(c.Writer).Encode(m.Clients)
}

