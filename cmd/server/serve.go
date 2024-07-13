package server

import (
	"log"
	"net"

	"github.com/ashupednekar/kvgo/internal/hashtable"
)

type Server struct {
	ListenAddr string
	t          hashtable.HashTable
	ln         net.Listener
	quitchan   chan struct{}
}

func NewServer(addr string) *Server {
	return &Server{
		ListenAddr: addr,
		t:          hashtable.NewHashTable(),
		quitchan:   make(chan struct{}),
	}
}

func (s *Server) Start() error {
	ln, err := net.Listen("tcp", s.ListenAddr)
	if err != nil {
		return err
	}
	s.ln = ln
	s.AcceptLoop()

	<-s.quitchan
	return nil
}

func (s *Server) AcceptLoop() {
	println("Accepting connections at ", s.ListenAddr)
	for {
		conn, err := s.ln.Accept()
		if err != nil {
			log.Fatal("error while accepting: ", err)
		}
		go s.HandleConn(conn)
	}
}
