package requests

import (
	_ "encoding/json"
	"fmt"
	_ "math"
	"net/http"
	_ "strconv"
	_"strings"

	_ "github.com/tidwall/gjson"
)


func GetExtraInfo(APIurl string, urls []string) bool{
	myClient := &http.Client{}
	res, _ := myClient.Get(APIurl)

	defer res.Body.Close()
	test := "https://api.csgofloat.com/?url="
	for _, value := range urls{
		test = APIurl + value
		fmt.Println(test)
		test = "https://api.csgofloat.com/?url="
	}

	return true
}