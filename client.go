// Single threaded client application using ZeroMQ
package main

import (
  zmq "github.com/pebbe/zmq4"
  "os"
  "fmt"
  "log"
)

// Connection to bind 
const client_bind string = "tcp://%s:5555"

func main() {

  // Check our inputs
  if len(os.Args) != 3 {
    log.Fatalf("Requires 2 arguments: <hostname> <string>")
  }

  conn := fmt.Sprintf(client_bind, os.Args[1])
 
  server, err := zmq.NewSocket(zmq.REQ)
  if err != nil {
    log.Fatalf("Connection failed: %s", err)
  }

  defer server.Close() 
  server.Connect(conn)

  log.Printf("Connected to %s", conn)

  server.SendMessage(os.Args[2])
  data, _ := server.RecvMessage(0)
  
  log.Printf("Received from server: %s", data)
}
