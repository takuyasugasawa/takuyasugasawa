package utils

import (
	"io"
	"log"
	"os"
)

func LoggingSettings(logFile string) {
	// 読み書き、作成、追記
	logfile, err := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln(err)
	}

	// ログの書き込み先を指定（標準とlogfile）
	multiLogFile := io.MultiWriter(os.Stdout, logfile)
	// ログのフォーマットを指定
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	//
	log.SetOutput(multiLogFile)
}
