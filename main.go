package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/websocket"
)

//func rootHandler(w http.ResponseWriter, req *http.Request) {
//	index.Execute(w, nil)
//	return
//}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type DataPoint struct {
	X time.Time
	Y int
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	// pagebytes, _ := fmt.Fprintf(w, p.Body)
	var isUpgrade bool = strings.Join(r.Header["Upgrade"], "") == "websocket"
	fmt.Println(isUpgrade)
	if isUpgrade {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			fmt.Println("could not upgrade")
			log.Fatal(err)
		}

		for {
			var x time.Time = time.Now()
			var y int = rand.Intn(100)
			err = conn.WriteJSON(DataPoint{X: x, Y: y})
			if err != nil {
				fmt.Println("could not write message")
				log.Fatal(err)
			}
			time.Sleep(100 * time.Millisecond)
		}
	} else {
		pageBytes, _ := ioutil.ReadFile("index.html")
		fmt.Fprintf(w, "%s", pageBytes)
	}
}

func main() {
	http.HandleFunc("/", rootHandler)
	http.ListenAndServe(":80", nil)

}
