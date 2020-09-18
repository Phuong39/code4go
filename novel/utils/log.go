/*
* 作者：刘时明
* 时间: 2019/9/30-11:14
* 作用：
 */
package utils

import (
	"fmt"
	log "github.com/jeanphorn/log4go"
	"os"
	"runtime"
	"strings"
)

// InitLogPath 初始化日志文件目录
func InitLogPath(svcName string) {
	arrName := strings.Split(svcName, ".")
	if len(arrName) > 1 {
		svcName = arrName[len(arrName)-1]
	}
	logPath := "/data/log/" + svcName
	exist, err := PathExists(logPath)
	if err != nil {
		panic(fmt.Sprintf("InitLogPath error: %s", err.Error()))
	}
	if !exist {
		if strings.ToLower(runtime.GOOS) == "windows" { //增加windows支持
			logPath = strings.Replace(logPath, "/data/", "", -1)
		}
		err := os.MkdirAll(logPath, os.ModePerm)
		if err != nil {
			panic(fmt.Sprintf("InitLogPath os.Mkdir falied: %s", err.Error()))
		}
	}
	// log
	processFetch := strings.Split(svcName, ".")
	if len(processFetch) == 0 {
		panic("invalid svcName")
	}
	logFileName := fmt.Sprintf("%s/ddlive_%s_api_d.log", logPath, processFetch[len(processFetch)-1])
	log.AddFilter(LogPlugin, log.DEBUG, log.NewFileLogWriter(logFileName, true, true))
}
