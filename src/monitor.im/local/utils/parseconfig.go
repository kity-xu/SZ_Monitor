package utils

import (
	"encoding/xml"
	"io/ioutil"
	"monitor.im/config"
	l4g "github.com/alecthomas/log4go"
)

/** 
	enconding/xml:
	func Marshal(v interface{}) ([]byte, error)
 	func Unmarshal(data []byte, v interface{}) error
 	上面两个函数的文档（注释）很长。从文档知道，Marshal是将v代表的数据转为XML格式（生成XML）；而Unmarshal刚好相反，是解析XML，同时将解析的结果保存在v中。

 	func MarshalIndent(v interface{}, prefix, indent string) ([]byte, error)
	 这个函数和Marshal的不同是，每个XML元素会增加前缀和缩进

**/

const (
	path_appconfig = config.GO_CONFIGS_APP
)

type ParseResult struct {
	Version  string         `xml:"Version,attr"`
	Application []Application	`xml:"Application"`
}

type Application struct {
	Record 		[]Record  	`xml:"Record"`
	Category	string		`xml:"Category,attr"`
	Caption		string		`xml:"Caption,attr"`
}
type Record struct {
	Command		[]Command	`xml:"Command"`
	Name		string		`xml:"Name,attr"`
	Caption		string		`xml:"Caption,attr"`
}
type Command struct {
	Cmd 		string		`xml:"Cmd,attr"`
	Caption 	string		`xml:"Caption,attr"`
}

func ParseXml(path string) (ParseResult, error) {
	var result ParseResult
	content, err := ioutil.ReadFile(path)
	if err != nil {
		l4g.Error(err)
		return result, err
	}

	e := xml.Unmarshal(content, &result)
	if e != nil {
		l4g.Error(e)
		return result, err
	}

	l4g.Info(result)
	return result, nil
}

func GetAppName(path string) ([]string, error) {
	str, err := ParseXml(path)
	if err != nil {
		return nil, err
	}
	var ss		[]string
	var n		int
	var apps 	Application
	for n, apps = range str.Application {
		ss = append(ss, apps.Category)
	}

	l4g.Info("application total: %d,  that respectively are: %s\n",n+1, ss)
	return ss, nil
}