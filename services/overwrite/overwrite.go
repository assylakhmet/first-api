package overwrite

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func OverWriteData(c *gin.Context){
	var clientJson = OverwriteData()
	InsertToDatabase(clientJson)
	fmt.Fprintf(c.Writer, "Done!")
	fmt.Println("OverWritingData")
}

