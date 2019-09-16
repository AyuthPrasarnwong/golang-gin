package v1

import (
	"app/controllers"
	"app/helpers"
	"app/models"
	//"bytes"
	"encoding/json"
	//"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	//"reflect"
	"strconv"
	"strings"
	//"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

type VerifyOtp models.VerifyOtp
type Order models.Order
type User models.User
type Meta controllers.Meta
type Pagination controllers.Pagination
type Links controllers.Links

type CheckoutParams struct {
	VerifyOtp VerifyOtp `json:"verify_otp"`
	Order    Order     `json:"order"`
}

func CheckoutOrder(c *gin.Context) {
	//var params CheckoutParams
	//
	//c.BindJSON(&params)
	//
	//in := make(chan int)
	//
	//outlet := make(chan map[string]interface{})
	////verify := make(chan map[string]interface{})
	//user := make(chan map[string]interface{})
	//
	////body := map[string]string{
	////	"otp_code":  params.VerifyOtp.OtpCode,
	////	"otp_ref":   params.VerifyOtp.OtpRef,
	////	"auth_code": params.VerifyOtp.AuthCode,
	////	"mobile":    params.VerifyOtp.Mobile,
	////}
	//
	//go FindOutletByID(in, outlet)
	////go CurlVerifyOtp(in, verify, body)
	//go CheckUser(in, user)
	//
	//in <- params.Order.OutletId
	//
	//c.Header("Content-Type", "application/json")
	//
	//userInfo := <-user
	//
	//data := userInfo["data"]
	//
	//userData := User{}
	//
	//b, err := json.Marshal(data)
	//
	//if err != nil {
	//
	//}
	//
	//json.Unmarshal([]byte(b), &userData)
	//
	//fmt.Println("userData", userData)
	//
	//loc, err := time.LoadLocation("Asia/Bangkok")
	//
	//if err != nil {
	//	panic(err)
	//}
	//
	//now := time.Now().In(loc)
	//dateTimeLayout := "2006-01-02 15:04:05"
	//
	//params.Order.PaymentStatus = "waiting"
	//params.Order.OrderStatus = "waiting"
	//params.Order.CustomerId = userData.Id
	//params.Order.PaymentChannelReference = params.VerifyOtp.Mobile
	//params.Order.ClientId = 3
	//params.Order.OrderCreatedAt = now.Format(dateTimeLayout)
	//params.Order.CustomerType = "member"
	//
	//order := make(chan map[string]interface{})
	//
	//statusCode := make(chan int)
	//
	//orderJson, _ := json.Marshal(&params.Order)
	//
	//bodyOrder := bytes.NewBuffer(orderJson)
	//
	//go CreateOrder(in, order, bodyOrder, statusCode)
	//
	//result := map[string]interface{}{
	//	"outlet": <-outlet,
	//	"order":  <-order,
	//	"user":   data,
	//}

	result := map[string]interface{}{
		"order":  "test",
	}

	c.JSON(http.StatusOK, result)
}

func CreateOrder(in <-chan int, out chan<- map[string]interface{}, body io.Reader, statusCode chan<- int) {

	//fmt.Println("in", in)
	//
	//url := "http://api-order.eggsmartpos.local/v2/orders"
	//
	//code, result := helpers.HttpClient("POST", url, body)
	//
	//out <- result
	//
	//statusCode <- code
}

func FindOrderByID(c *gin.Context) {

	id := c.Param("id")

	orderId , _ := strconv.Atoi(id)

	authorization := c.GetHeader("Authorization")

	statusCode, result := CurlFindOrderByID(orderId, authorization)

	c.JSON(statusCode, result)
}

func CurlFindOrderByID(id int, authorization string) (statusCode int, result map[string]interface{})  {

	fmt.Println("id", id)

	headers := map[string]string {
		"Authorization": authorization,
	}

	fmt.Println("headers", headers)

	url := fmt.Sprintf("http://api-order.eggsmartpos.local/v2/orders/%d", id)

	statusCode, result = helpers.HttpClient("GET", url, nil, headers, "")

	return statusCode,  result
}

func GetAllOrders(c *gin.Context) {

	authorization := c.GetHeader("Authorization")

	query := c.Request.URL.Query().Encode()

	statusCode, result := CurlGetAllOrders(query, authorization)

	meta := new(Meta)

	metaData, err := json.Marshal(result["meta"])

	if err != nil {

	}

	json.Unmarshal([]byte(metaData), &meta)

	error := godotenv.Load()

	if error != nil {
		log.Fatal("Error loading .env file")
	}

	protocol := "https://"

	if os.Getenv("APP_ENV") == "local" {
		protocol = "http://"
	}

	if meta.Pagination.Links.Next != "" {
		query := strings.Split(meta.Pagination.Links.Next, "?")

		url := fmt.Sprintf("%s%s%s?%s", protocol, c.Request.Host, c.Request.URL.Path,  query[1])
		meta.Pagination.Links.Next = url
	}

	if meta.Pagination.Links.Previous != "" {
		query := strings.Split(meta.Pagination.Links.Previous, "?")
		url := fmt.Sprintf("%s%s%s?%s", protocol, c.Request.Host, c.Request.URL.Path,  query[1])
		meta.Pagination.Links.Previous = url
	}

	if meta.Pagination.Links.Next == "" && meta.Pagination.Links.Previous == "" {

	}
	result["meta"] = &meta

	//fmt.Println("meta", meta.Pagination.Links)
	//
	//fmt.Println("result", result["meta"])

	c.JSON(statusCode, result)
}

func CurlGetAllOrders(query string, authorization string) (statusCode int, result map[string]interface{})  {

	fmt.Println("query", query)

	headers := map[string]string {
		"Authorization": authorization,
	}

	fmt.Println("headers", headers)

	url := "http://api-order.eggsmartpos.local/v2/orders"

	statusCode, result = helpers.HttpClient("GET", url, nil, headers, query)

	return statusCode,  result
}


