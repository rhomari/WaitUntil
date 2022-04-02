package main

import (
	"fmt"
	"log"
	"net/http"
	"runtime"
	"sync"
	"time"
)

func main() {
	wg := &sync.WaitGroup{}
	wg.Add(2)
	go setupHttpServer()
	go makeTask(time.Now().Add(time.Second*10), wg)
	wg.Wait()

}
func setupHttpServer() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello from back-end you requested: %s\n", r.URL.Path)
	})
	log.Fatal(http.ListenAndServe(":2304", nil))
}
func makeTask(nexttime time.Time, waitgroup *sync.WaitGroup) {

	time.Sleep(time.Until(nexttime))
	log.Println("I'am  doing stuff without front-end user intervention")
	waitgroup.Done()
	waitgroup.Add(1)
	log.Println(runtime.NumGoroutine())
	makeTask(nexttime.Add(time.Second*10), waitgroup)
}
