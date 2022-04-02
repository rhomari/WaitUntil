package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

func main() {
	wg := &sync.WaitGroup{}
	wg.Add(2)
	go setupHttpServer()
	go func() {
		startTime, _ := time.Parse(time.RFC3339, "2022-04-02T11:55:00Z") //scheduling a task at the given date and time
		time.Sleep(time.Until(startTime))
		log.Println("I'am  doing stuff widthout front-end user intervention")
		wg.Done()
		//you can add a new schedule here for the next time
	}()

	wg.Wait()

}
func setupHttpServer() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello from back-end you requested: %s\n", r.URL.Path)
	})
	log.Fatal(http.ListenAndServe(":2304", nil))
}
