package protocol

import (
	"fmt"
	"bytes"
	"io/ioutil"
	//"encoding/json"
	"net/http"
	"testing"
)

type SendData struct {
	username string
	password string 
}

func TestStartPostServer(t *testing.T){

	data := &SendData{}
	data.username = "xulang"
	data.password = "123"

	// var db []byte

	// db, _ = json.Marshal(data)
		
	da := "username=" + data.username + "&" + "password=" + data.password
	body := bytes.NewBuffer([]byte(da))


    resp, err := http.Post("http://192.168.9.222:9090/kity",
        "application/x-www-form-urlencoded", body)

    if err != nil {
        fmt.Println(err)
    }
 
    defer resp.Body.Close()
    res, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        // handle error
    }
 
    fmt.Println(string(res))
}