package main

import (
	"flag"
	"fmt"
	"github.com/spf13/viper"
	"log"
	"net/http"
	"os"
	"path"
	"strconv"
)

var confFile, ip, logPath, logLevelC string
var port int
var logFile os.File

var InfoLog, WarnLog, DebugLog, ErrorLog *log.Logger

var logLevel = []string{"debug", "info", "warning", "error"}

func Find(s []string, x string) int {
	for i, n := range s {
		if x == n {
			return i
		}
	}
	return 0
}

func init() {
	flag.StringVar(&confFile, "config", "", "configuration file")
	log.SetPrefix("Log:")
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Lshortfile)

	flag.Parse()
	viper.SetConfigFile(confFile)
	if err := viper.ReadInConfig(); err != nil {

		log.Fatal("Ошибка чтения конфигурационного файла")
	}
	ip = viper.GetString("http_listen.ip")
	port = viper.GetInt("http_listen.port")
	logPath = viper.GetString("log_file")
	logLevelC = viper.GetString("log_level")

	_, err := os.Stat(logPath)
	if os.IsNotExist(err) {
		if errDir := os.MkdirAll(logPath, 0755); errDir != nil {
			log.Fatal("Ошибка при создании директории для логирования", errDir)
		}
	}
	logFile, err := os.OpenFile(path.Join(logPath, "info.log"), os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal("Ошибка при создании файла для логирования", err)
	}

	logLevelI := Find(logLevel, logLevelC)

	println(logLevelI)

	if logLevelI == 0 {
		DebugLog = log.New(logFile, "Info", log.Ldate|log.Lmicroseconds|log.Lshortfile)
	} else {
		DebugLog = log.New(os.Stdout, "Info", log.Ldate|log.Lmicroseconds|log.Lshortfile)
	}
	if logLevelI <= 1 {
		InfoLog = log.New(logFile, "Info", log.Ldate|log.Lmicroseconds|log.Lshortfile)
	} else {
		InfoLog = log.New(os.Stdout, "Info", log.Ldate|log.Lmicroseconds|log.Lshortfile)
	}
	if logLevelI <= 2 {
		WarnLog = log.New(logFile, "Info", log.Ldate|log.Lmicroseconds|log.Lshortfile)
	} else {
		WarnLog = log.New(os.Stdout, "Info", log.Ldate|log.Lmicroseconds|log.Lshortfile)
	}
	if logLevelI <= 3 {
		ErrorLog = log.New(logFile, "Info", log.Ldate|log.Lmicroseconds|log.Lshortfile)
	} else {
		ErrorLog = log.New(os.Stdout, "Info", log.Ldate|log.Lmicroseconds|log.Lshortfile)
	}

}

func main() {

	println("Запущено ...", os.Args[0])

	defer logFile.Close()

	log.Println(ip, port, logPath, logLevel)
	InfoLog.Println("Инфо левел", logLevelC)
	DebugLog.Println("Дебаг левел", logLevelC)
	WarnLog.Println("Варнинг левел", logLevelC)
	ErrorLog.Println("Error левел", logLevelC)
	// println(ip, port, logPath, logLevel )

	http.HandleFunc("/", func(responseWriter http.ResponseWriter, r *http.Request) {
		InfoLog.Printf("%v", r)
		fmt.Fprintf(responseWriter, "Hello world\n")
	})
	err := http.ListenAndServe(":"+strconv.Itoa(port), nil)
	println(err, ":"+strconv.Itoa(port))

}
