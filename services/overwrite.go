package services

import (
	"bufio"
	"encoding/json"
	m "fApi/models"
	"first-api/configuration"
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
)

func  (handler *Handler)  OverwriteData(c *gin.Context){
	var clientJson = ReadData("hello.txt")
	InsertToDatabase(handler, clientJson)
	fmt.Fprintf(c.Writer, "Done!")
	fmt.Println("OverWritingData")
}


func ReadData(path string) string{
	//прочитать данные и записать все в строку
	file, err :=os.Open(path)
	if err!=nil{
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()
	r := bufio.NewReader(file)
	s, e := Readln(r)
	var a string
	for e == nil {
		a=a+s
		s,e = Readln(r)
	}
	return a
}


func InsertToDatabase(handler *Handler,clientJson string){
	//парсинг данных
	var client []m.Client
	json.Unmarshal([]byte(clientJson), &client)
	//подключение к базе данных
	db, _ := configuration.SqlConfig(handler.Config)
	defer db.Close()
	//запись данных в базу данных
	for i := range client{
		_, err := db.Exec("insert into Clients (clientid, Name) values ($1, $2)", client[i].ClientId,client[i].Name)
		if err != nil{
			fmt.Println(err.Error())
		}
	}
}


func Readln(r *bufio.Reader) (string, error) {
	var (isPrefix bool = true
		err error = nil
		line, ln []byte
	)
	for isPrefix && err == nil {
		line, isPrefix, err = r.ReadLine()
		ln = append(ln, line...)
	}
	return string(ln),err
}

