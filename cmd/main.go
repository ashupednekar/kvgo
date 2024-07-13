package main

import "github.com/ashupednekar/kvgo/cmd/server"

func main() {
	s := server.NewServer(":3000")
	s.Start()
}
