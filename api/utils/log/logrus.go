package log

import (
	"fmt"
	"peckergo/api/config"
	"runtime"

	"github.com/multiplay/go-slack/chat"
	"github.com/multiplay/go-slack/lrhook"

	log "github.com/sirupsen/logrus"
)

const (
	slackHook = "https://hooks.slack.com/services"
)

var (
	// Hook lrhook
	Hook    *lrhook.Hook
	logSync *log.Logger
)

// Init log 初始化
func Init() {
	level := log.Level(config.GetInt("log.logLevel"))
	log.SetLevel(level)
	log.SetFormatter(&log.TextFormatter{
		ForceColors:     true,
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
	})
	log.Warnf("logLevel: [%+v]", level)

	if config.GetBool("slack.enabled") {
		key := config.GetString("slack.key")
		cfg := lrhook.Config{
			MinLevel: log.ErrorLevel,
			Message: chat.Message{
				Channel:   config.GetString("slack.channel"),
				IconEmoji: ":ghost:",
			},
			Async: true,
		}
		Hook = lrhook.New(cfg, fmt.Sprintf("%s/%s", slackHook, key))
		log.AddHook(Hook)

		cfg.Async = false
		logSync = log.New()
		logSync.AddHook(lrhook.New(cfg, fmt.Sprintf("%s/%s", slackHook, key)))
	}
}

// PanicRecover 崩溃日志输出和上传  在入口 func main() 第一行 调用 defer logutils.PanicRecover()
func PanicRecover() {
	if e := recover(); e != nil {
		var buf [4096]byte
		n := runtime.Stack(buf[:], false)
		logSync.Error(e, "\n", string(buf[:n]))
	}
}
