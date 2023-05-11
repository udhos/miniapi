/*
This is the main package for the miniapi service.
*/
package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/udhos/miniapi/env"
)

const version = "1.0.5"

func getVersion(me string) string {
	return fmt.Sprintf("%s version=%s runtime=%s GOOS=%s GOARCH=%s GOMAXPROCS=%d",
		me, version, runtime.Version(), runtime.GOOS, runtime.GOARCH, runtime.GOMAXPROCS(0))
}

type config struct {
	paramList []string
}

func main() {

	app := config{}

	var showVersion bool
	flag.BoolVar(&showVersion, "version", showVersion, "show version")
	flag.Parse()

	me := filepath.Base(os.Args[0])

	{
		v := getVersion(me)
		if showVersion {
			fmt.Print(v)
			fmt.Println()
			return
		}
		log.Print(v)
	}

	addr := env.String("ADDR", ":8080")
	path := env.String("ROUTE", "/v1/hello;/v1/world;;")
	health := env.String("HEALTH", "/health")
	params := env.String("PARAMS", "param1;param2")

	pathList := strings.FieldsFunc(path, func(r rune) bool { return r == ';' })
	app.paramList = strings.FieldsFunc(params, func(r rune) bool { return r == ';' })

	mux := http.NewServeMux()
	server := &http.Server{
		Addr:    addr,
		Handler: mux,
	}

	const root = "/"

	register(mux, addr, root, func(w http.ResponseWriter, r *http.Request) { handlerRoot(&app, w, r) })
	register(mux, addr, health, func(w http.ResponseWriter, r *http.Request) { handlerHealth(&app, w, r) })

	for _, p := range pathList {
		register(mux, addr, p, func(w http.ResponseWriter, r *http.Request) { handlerPath(&app, w, r) })
	}

	go listenAndServe(server, addr)

	<-chan struct{}(nil)
}

func register(mux *http.ServeMux, addr, path string, handler http.HandlerFunc) {
	mux.HandleFunc(path, handler)
	log.Printf("registered on port %s path %s", addr, path)
}

func listenAndServe(s *http.Server, addr string) {
	log.Printf("listening on port %s", addr)
	err := s.ListenAndServe()
	log.Printf("listening on port %s: %v", addr, err)
}

// httpJSON replies to the request with the specified error message and HTTP code.
// It does not otherwise end the request; the caller should ensure no further
// writes are done to w.
// The error message should be JSON.
func httpJSON(w http.ResponseWriter, error string, code int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(code)
	fmt.Fprintln(w, error)
}

func toJSON(v interface{}) string {
	b, err := json.Marshal(v)
	if err != nil {
		log.Printf("toJSON: %v", err)
	}
	return string(b)
}

type responseBody struct {
	Request        responseRequest `json:"request"`
	Message        string          `json:"message"`
	Status         int             `json:"status"`
	ServerHostname string          `json:"server_hostname"`
	ServerVersion  string          `json:"server_version"`
}

type responseRequest struct {
	Headers   http.Header       `json:"headers"`
	Method    string            `json:"method"`
	URI       string            `json:"uri"`
	Host      string            `json:"host"`
	Body      string            `json:"body"`
	FormQuery url.Values        `json:"form_query"`
	FormPost  url.Values        `json:"form_post"`
	Params    map[string]string `json:"parameters"`
}

func response(app *config, w http.ResponseWriter, r *http.Request, status int, message string) {
	const me = "response"

	hostname, errHost := os.Hostname()
	if errHost != nil {
		log.Printf("%s hostname error: %v", me, errHost)
	}

	// take a copy of the body
	reqBody, errRead := io.ReadAll(r.Body)
	if errRead != nil {
		log.Printf("%s: body read error: %v", me, errRead)
	}
	r.Body = io.NopCloser(bytes.NewBuffer(reqBody)) // restore it

	errForm := r.ParseForm()
	if errForm != nil {
		log.Printf("%s: form error: %v", me, errForm)
	}

	errMultipart := r.ParseMultipartForm(32 << 20)
	if errMultipart != nil {
		log.Printf("%s: form multipart error: %v", me, errMultipart)
	}

	params := map[string]string{}

	for _, p := range app.paramList {
		params[p] = r.FormValue(p)
	}

	reply := responseBody{
		Request: responseRequest{
			Headers:   r.Header,
			Method:    r.Method,
			URI:       r.RequestURI,
			Host:      r.Host,
			Body:      string(reqBody),
			FormQuery: r.Form,
			FormPost:  r.PostForm,
			Params:    params,
		},
		Message:        message,
		Status:         status,
		ServerHostname: hostname,
		ServerVersion:  version,
	}

	body := toJSON(reply)

	httpJSON(w, body, status)
}

func handlerRoot(app *config, w http.ResponseWriter, r *http.Request) {
	log.Printf("%s %s %s - 404 not found", r.RemoteAddr, r.Method, r.RequestURI)
	response(app, w, r, http.StatusNotFound, "not found")
}

func handlerPath(app *config, w http.ResponseWriter, r *http.Request) {
	log.Printf("%s %s %s - 200 ok", r.RemoteAddr, r.Method, r.RequestURI)
	response(app, w, r, http.StatusOK, "ok")
}

func handlerHealth(app *config, w http.ResponseWriter, r *http.Request) {
	log.Printf("%s %s %s - 200 health ok", r.RemoteAddr, r.Method, r.RequestURI)
	response(app, w, r, http.StatusOK, "health ok")
}
