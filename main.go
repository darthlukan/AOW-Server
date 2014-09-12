package main

import (
	"fmt"
	"github.com/Radiobox/web_responders"
	"github.com/stretchr/goweb"
	"html/template"
	"log"
	"net"
	"net/http"
	"os"
	"path"
	"time"
)

var (
	projectRoot string
	templates   *template.Template
	goPath      = os.Getenv("GOPATH")
	messages    = web_responders.NewMessageMap()
)

func main() {
	log.Println("Starting server...")

	if goPath == "" {
		projectRoot = "."
	} else {
		projectRoot = path.Join(goPath, "src", "github.com", "darthlukan", "AOW-Server")
	}
	// Move to site server
	templates = template.Must(template.ParseGlob(projectRoot + "/html/*"))
	//

	// API Endpoints
	goweb.Map("/ping", pingHandler)
	goweb.Map("/getquote/{apiKey}", AowQuote)
	// End API Endpoints

	address := ":3000"
	if port := os.Getenv("PORT"); port != "" {
		address = ":" + port
	}
	server := &http.Server{
		Addr:           address,
		Handler:        &LoggedHandler{goweb.DefaultHttpHandler()},
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	listener, listenErr := net.Listen("tcp", address)
	if listenErr != nil {
		log.Panicf("Could not listen for TCP on %s: %s", address, listenErr)
	}
	log.Println("Server loaded, check localhost" + address)
	server.Serve(listener)
}

func colorize(r *http.Request) string {
	format := fmt.Sprintf("%s[94m %s %s[92m%s: %s[91m%s, %s[0m", escape, r.Proto, escape, r.Method, escape, r.URL, escape)
	return format
}

type LoggedHandler struct {
	baseHandler http.Handler
}

func (handler *LoggedHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	go log.Printf(colorize(r))
	handler.baseHandler.ServeHTTP(w, r)
}
