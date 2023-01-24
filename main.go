package main

import "github.com/roblesoft/bookd/pkg/server"

func main() {
	server := &server.Server{Port: ":3000"}
	server.Start()
}
