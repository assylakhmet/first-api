package models

type Client struct{
	//модель для данных о клтентах
	ClientId int 	`json:"id"`
	Name string 	`json:"name"`
}

var Clients []Client
