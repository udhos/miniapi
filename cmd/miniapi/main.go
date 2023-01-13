/*
This is the main package for the miniapi service.
*/
package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"runtime"

	"github.com/udhos/miniapi/env"
)

const version = "0.0.1"

func getVersion(me string) string {
	return fmt.Sprintf("%s version=%s runtime=%s GOOS=%s GOARCH=%s GOMAXPROCS=%d",
		me, version, runtime.Version(), runtime.GOOS, runtime.GOARCH, runtime.GOMAXPROCS(0))
}

func main() {

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
	path := env.String("ROUTE", "/v1/world")
	health := env.String("HEALTH", "/health")

	mux := http.NewServeMux()
	server := &http.Server{
		Addr:    addr,
		Handler: mux,
	}

	const root = "/"

	register(mux, addr, root, handlerRoot)
	register(mux, addr, path, handlerPath)
	register(mux, addr, health, handlerHealth)

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

func response(w http.ResponseWriter, r *http.Request, status int, message string) {
	reply := fmt.Sprintf(`{"message":"%s","status":"%d","path":"%s","method":"%s"}`,
		message, status, r.RequestURI, r.Method)
	httpJSON(w, reply, status)
}

func handlerRoot(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s %s %s - 404 not found", r.RemoteAddr, r.Method, r.RequestURI)
	response(w, r, http.StatusNotFound, "not found")
}

func handlerPath(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s %s %s - 200 ok", r.RemoteAddr, r.Method, r.RequestURI)
	response(w, r, http.StatusOK, "ok")
}

func handlerHealth(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s %s %s - 200 health ok", r.RemoteAddr, r.Method, r.RequestURI)
	response(w, r, http.StatusOK, "health ok")
}
