package v1

import (
	"app/helpers"

	//"bytes"
	//"encoding/json"
	//"errors"
	"fmt"
	//"reflect"
	"strconv"
	//"time"

	"github.com/gin-gonic/gin"
)

func FindTerminalByID(c *gin.Context) {

	id := c.Param("id")

	terminalId , _ := strconv.Atoi(id)

	authorization := c.GetHeader("Authorization")

	statusCode, result := CurlFindTerminalByID(terminalId, authorization)

	c.JSON(statusCode, result)
}

func CurlFindTerminalByID(id int, authorization string) (statusCode int, result map[string]interface{})  {

	headers := map[string]string {
		"Authorization": authorization,
	}

	url := fmt.Sprintf("http://api-merchant.eggsmartpos.local/v1/terminals/%d", id)

	statusCode, result = helpers.HttpClient("GET", url, nil, headers, "")

	return statusCode,  result
}
