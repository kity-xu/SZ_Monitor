package protocol

import (
	"encoding/json"
	"net/http"
	"time"
	"strings"
	"fmt"
	l4g "github.com/alecthomas/log4go"
)

const(
	port = "9090"
)

type ResultJsonBean struct{
 	Code    int         `json:"code"`  
    Data    interface{} `json:"data"`  
    Message string      `json:"message"`  
}
func NewResultJsonBean() *ResultJ sonBean {
	return &ResultJsonBean{}
}

type HttpServer struct {

}

func NewHttpServer() *HttpServer {
	return &HttpServer{}
}

func (ps *HttpServer)StartPostServer() error{
	l4g.Info("start http post server ...")

	http.HandleFunc("/kity", post)	//服务器+主机，来确定唯一服务。shell: hostname

	if err := http.ListenAndServe("192.168.9.222:"+port, nil); err != nil {
		l4g.Error("Listen localhost:%d error... The error is:%s",port, err.Error())
		return err
	}
	return nil
}

func post(w http.ResponseWriter, req *http.Request){
	l4g.Info("post request...")

	time.Sleep(time.Second * 1)

	req.ParseForm()

	arg1, f1 := req.Form["username"]
	arg2, f2 := req.Form["password"]

	if !(f1 && f2){
		l4g.Info("post request 参数错误...")
		fmt.Fprintf(w, "参数错误%s---%s",arg1,arg2)
		return
	}

	result := NewResultJsonBean()

	name := arg1[0]
	password := arg2[0]

	if strings.EqualFold(name, "xulang") && strings.EqualFold(password, "123"){
		l4g.Info("post success...")

		result.Code = 200
		result.Message = "认证成功"
		//result.Data = {"Hello":"World"}
	}else{
		l4g.Info("post faild...")

		result.Code = 101
		result.Message = "认证失败"
	}

	bytes, _ := json.Marshal(result)
	fmt.Fprint(w, string(bytes))
}