package protocol

import (
	"time"
	"strings"
	"monitor.im/local/utils"
	l4g "github.com/alecthomas/log4go"
	"os/exec"
	"monitor.im/config"
)

type MonitorCode struct {

}

const (
		py_dirpath_app = config.PY_DIRPATH_APP
	)


func shellCommand(name string) string {
	cmd := exec.Command("python", py_dirpath_app + name)
	if out, err := cmd.CombinedOutput(); err == nil {
		//l4g.Debug(string(out))		//out return value of call *.py
		return string(out)
	}else{
		l4g.Debug("running shellCommand error ....")
		return ""
	}
}

func start(name string){
	for ;; {
		info := shellCommand(name)
		if strings.EqualFold(info,""){
			l4g.Debug("shellCommand return value is null...")
			return
		}
		l4g.Info(info)			//最终的数据
		time.Sleep(100000)
	}	
}

func getFilesPY() ([]string, error) {
	/**
	*	arr 与 Category 对比，如果apps有arr没有，则说明少了脚本文件或是多了无效的配置 则忽略该配置 
	*	如果arr有 Category 没有，则以配置文件为主（即 取他们的交集）
	*
	**/
	arr, err := utils.GetfilesBywalkdir(py_dirpath_app, "pY")					//遍历python/application 目录，查看已有的脚本
	if err != nil {
		return nil, err
	}
	l4g.Debug(arr)

	Category, err := utils.GetAppName(config.GO_CONFIGS_APP + "appconfigs.xml")		//读取配置文件查看已有的脚本
	if err != nil {
		return nil, err
	}
	l4g.Debug(Category)

	var apps []string

	for _, v := range Category {
		 for _, v2 := range arr {
		 	if !strings.EqualFold(v, v2) {
		 		apps = append(apps, v2+".py")
		 		break
		 	}else{
		 		continue
		 	}
		 }
	}

	l4g.Debug("apps :", apps)
	return apps, nil
}

func (this *MonitorCode)StartMonitor() {
	apps, err := getFilesPY()
	if err != nil || apps == nil{
		return 
	}

	// i := 0
	// for ; i<len(apps); i++ {
	// 	go start(apps[i])
	// }
	start(apps[0])
}