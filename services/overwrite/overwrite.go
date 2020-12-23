package overwrite

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func OverwriteData(c *gin.Context){
	var clientJson = ReadData("hello.txt")
	InsertToDatabase(clientJson)
	fmt.Fprintf(c.Writer, "Done!")
	fmt.Println("OverWritingData")
}

