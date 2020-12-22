package overwrite

import (
	"fApi/configuration"
	m "fApi/models"
	"encoding/json"
	"fmt"
)


func InsertToDatabase(clientJson string){
	//парсинг данных
	var client []m.Client
	json.Unmarshal([]byte(clientJson), &client)
	//подключение к базе данных
	db, _ := configuration.SqlConfig()
	defer db.Close()
	//запись данных в базу данных
	for i := range client{
		_, err := db.Exec("insert into Clients (clientid, Name) values ($1, $2)", client[i].ClientId,client[i].Name)
		if err != nil{
			fmt.Println(err.Error())
		}
	}
}
