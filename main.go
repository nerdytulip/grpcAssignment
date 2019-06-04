package main

import (
	"grpcAssignment/client"
	"grpcAssignment/server"
)

func main() {
	go func() {
		server.Start("50050")
	}()
	//time.Sleep(2 * time.Second)
	client.Start("50050")
}
