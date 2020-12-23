package services

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)


const charset = "abcdefghijklmnopqrstuvwxyz" +
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ"

var seededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

func StringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func New(c *gin.Context){
	c.Writer.Header().Set("Content-Type", "application/json")
	var data string
	file, err := os.Create("hello.txt")

	if err != nil{
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()
	n1, err:=strconv.Atoi(c.Params.ByName("n"))
	fmt.Println(n1)
	data ="[\n"
	//var client []read.Client
	//file.WriteString(text)
	for i:=0;i<n1;i++{
		data += "\t{\n\t\"id\": " + strconv.Itoa(i+1)+", \n\t\"name\": \""+ StringWithCharset(5, charset)+"\"\n\t},"
		//client[i].ClientId = i
		//client[i].Name=StringWithCharset(5, charset)
	}

	//a, _ :=json.Marshal(client)
	data = strings.TrimSuffix(data, ",")
	data+="\n]"
	file.WriteString(data)


	fmt.Fprintf(c.Writer,data)
	fmt.Println("Done.")
}

