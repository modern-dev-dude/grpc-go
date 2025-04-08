Working project that starts up 3 servers a rendering engine, random number simulating another remote procedure, and a client that orchestrates the web server for the content.

To demo run the following 3 commands in three seperate terminals 
Start web client
`go run ./cmd/client/main.go`
Start random number server
`go run ./cmd/random-number/main.go`
Start rendering engine
`go run ./cmd/rendering-engine/main.go`

navigate to http://localhost:8080 and see the output! 

If it is broke, it wasnt my fault, it worked locally ;) haha
