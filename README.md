# 0mq-multithreaded

Example code for the blog post at http://www.aviancarrier.co.uk/blog/2015/04/05/multithreaded-servers-with-zeromq-and-go.html.

You will need Go installed. 

Try running with;

    go run singlethread_server.go

or 

    go run multithread_server.go

and the client application with

    go run client.go localhost "Hello World"
