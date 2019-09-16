package v1


import (
	//"bytes"
	//"encoding/json"
	//"errors"
	"fmt"
	//"io"
	//"log	"
	//"net/http"
	//"reflect"
	//"time"
)

func CheckUser(in <-chan int, out chan<- map[string]interface{}) {

	fmt.Println("in", in)

	url := "http://api-account.eggsmartpos.local/api/v1/members/profile"

	out <- MakeRequest("GET", url, nil)
}