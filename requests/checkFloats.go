package requests

import (
	_ "encoding/json"
	"fmt"
	_ "io"
	"io/ioutil"
	_ "math"
	"net/http"
	_ "strconv"
	_"strings"
	_ "time"

	_ "github.com/tidwall/gjson"
)


func GetExtraInfo(url string) bool{
	myClient := &http.Client{}
	res, _ := myClient.Get(url)

	data, _ := ioutil.ReadFile("items.json")

	fmt.Println(string(data))

	defer res.Body.Close()

	return true
}