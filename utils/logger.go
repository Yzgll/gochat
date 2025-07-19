package utils

import (
	"fmt"
	"os"
	"runtime"
	"strings"

	"github.com/sirupsen/logrus"
)

var Log = logrus.New()

func InitLogger() {
	Log.SetOutput(os.Stdout)
	Log.SetReportCaller(true) // ✅ 开启调用者信息（文件+行号）
	Log.Info("日志系统启动")
	Log.SetFormatter(&logrus.TextFormatter{
		ForceColors:     true,                  // 彩色输出
		FullTimestamp:   true,                  // 显示完整时间
		TimestampFormat: "2006-01-02 15:04:05", // 自定义时间格式
		CallerPrettyfier: func(f *runtime.Frame) (string, string) {
			relPath := strings.Replace(f.File, "\\", "/", -1)
			projectRoot := "gochat/"
			idx := strings.Index(relPath, projectRoot)
			if idx != -1 {
				relPath = relPath[idx:]
			}
			return "", relPath + ":" + fmt.Sprintf("%d", f.Line)
		},
	})
}
