// Single-threaded server application using ZeroMQ
package main

import (
  zmq "github.com/pebbe/zmq4"
  "log"
)

// Number of workers to start
const max_workers int32 = 5

// Connection to bind 
const client_bind string = "tcp://*:5555"

func Reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

func main() {
  log.Println("Started")
  
  zmqs, _ := zmq.NewSocket(zmq.REP)
  defer zmqs.Close()
  zmqs.Bind(client_bind)

  for {
    log.Println("Waiting for message")
    data, _ := zmqs.RecvMessage(0)
    log.Printf("Received message: %s", data)
    reply := Reverse(data[0])
    zmqs.SendMessage(reply)
    log.Printf("Sent reply: %s", reply)
  }
}