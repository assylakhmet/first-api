package addition

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"first-api/configuration"
	m "first-api/models"
	s "first-api/services"
	"strconv"
)

func GetClient(c *gin.Context){
	c.Writer.Header().Set("Content-Type", "application/json")
	s.GetDataFromDataBase()
	for _, item := range m.Clients {
		if strconv.Itoa(item.ClientId) ==  c.Params.ByName("id") {
			json.NewEncoder(c.Writer).Encode(item)
			return
		}
	}
	json.NewEncoder(c.Writer).Encode(&m.Client{})
}


func Truncate(c *gin.Context){
	db, err := configuration.SqlConfig()
	if err != nil{
		fmt.Println(err.Error())
	}
	db.Query("truncate clients")
	m.Clients=nil
	fmt.Fprintf(c.Writer,"Done!")
	fmt.Println("truncate done;")
}
