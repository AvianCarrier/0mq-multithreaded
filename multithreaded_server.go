// Multithreaded server application using ZeroMQ
package main

import (
  zmq "github.com/pebbe/zmq4"
  "log"
)

// Number of workers to start
const max_workers int32 = 5

// Connection to bind 
const client_bind string = "tcp://*:5555"
const worker_bind string = "inproc://workers"

func main() {
  
  log.Println("Starting workers")
  for i := 0; i != 5; i = i + 1 {
    log.Printf("Starting worker %d", i)
    go worker(i)
  }

  clients, _ := zmq.NewSocket(zmq.ROUTER)
  defer clients.Close() 
  clients.Bind(client_bind)
  log.Println("Client listener started")

  workers, _ := zmq.NewSocket(zmq.DEALER)
  defer workers.Close()
  workers.Bind(worker_bind)
  log.Println("Worker listener started")

  log.Println("Starting proxy")
  err := zmq.Proxy(clients, workers, nil)
  log.Fatalf("Proxy interrupted:", err)
}

func Reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

func worker(worker_id int) {
  log.Printf("Worker %d: Started", worker_id)
  
  zmqs, _ := zmq.NewSocket(zmq.REP)
  defer zmqs.Close()
  zmqs.Connect(worker_bind)

  for {
    log.Printf("Worker %d: Waiting for message", worker_id)
    data, _ := zmqs.RecvMessage(0)
    log.Printf("Worker %d: Received message: %s", worker_id, data)
    reply := Reverse(data[0])
    zmqs.SendMessage(reply)
    log.Printf("Worker %d: Sent reply: %s", worker_id, reply)
  }
}
