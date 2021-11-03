package main

import(
	"fmt"
	"os"
	"time"
	"log"
	"net/http"
)

var PAGE_PATH string = "index.html" //The page that will be used to "bait" 
var LOG_FILE_PATH string = "logs.txt" //The file containing the logs
var PAGE []byte = getFileContent(PAGE_PATH)

func getFileContent(path string) []byte {
	data, err := os.ReadFile(path)
	if err != nil {
		fmt.Printf(err.Error())
		os.Exit(1)
	}
	return data
}


func logIP(addr string) {
	f, err := os.OpenFile("logs.txt",os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()
	if _, err := f.WriteString("[*] "+time.Now().Format("2006-01-02 15:04:05")+" : "+addr+"\n"); err != nil {
		log.Println(err)
	}
}


func rootHandler(writer http.ResponseWriter, request *http.Request) {
	writer.Write(PAGE)
	addr := request.RemoteAddr
	logIP(addr)
}

func main() {
	fmt.Println("Server has started")
	http.HandleFunc("/", rootHandler)
	http.ListenAndServe(":8080",nil)
}   
