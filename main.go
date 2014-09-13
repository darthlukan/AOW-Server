package main

import (
	"encoding/json"
	"fmt"
	"github.com/Radiobox/web_responders"
	"github.com/stretchr/goweb"
	"io/ioutil"
	"log"
	"math/rand"
	"net"
	"net/http"
	"os"
	"path"
	"strings"
	"time"
)

const (
	escape = "\x1b"
)

var (
	projectRoot string
	bookPath    string
	quotes      []string
	goPath      = os.Getenv("GOPATH")
	messages    = web_responders.NewMessageMap()
	book        = "aow.txt"
	cfgFile     = "config.json"
	config      *Config
)

type Config struct {
	AndroidKey string
	FFOSKey    string
	TizenKey   string
	WebKey     string
	DevKey     string
}

func OpenBook(filename string) []string {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	quotes := strings.Split(string(content), "\n\n")
	return quotes
}

func init() {

	rand.Seed(time.Now().UTC().UnixNano())

	if goPath == "" {
		projectRoot = "."
	} else {
		projectRoot = path.Join(goPath, "src", "github.com", "darthlukan", "AOW-Server")
	}

	cfg, err := os.Open(cfgFile)
	if err != nil {
		panic(err)
	}

	decoder := json.NewDecoder(cfg)
	config = &Config{}
	decoder.Decode(&config)

	bookPath = path.Join(projectRoot, book)
	quotes = OpenBook(bookPath)
}

func main() {
	log.Println("Starting server...")

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
