package server

import (
	"fmt"
	"log"
	"net"
	"strings"
)

func (s *Server) HandleConn(conn net.Conn) {
	defer conn.Close()
	buf := make([]byte, 2048)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			log.Fatal("error handling conn: ", err)
			continue
		}
		msg := buf[:n]
		println("msg: ", string(msg))

		// ch := s.GetChan(conn.RemoteAddr().String())

		//go func() {
		//	for s := range ch {
		//		fmt.Println("to send: ", s)
		//		fmt.Fprint(conn, s)
		//	}
		//}()

		l := strings.Split(string(msg), ":")
		switch {
		case strings.HasPrefix(string(msg), "set:"):
			HandleSet(*s, l, conn)
		case strings.HasPrefix(string(msg), "get:"):
			HandleGet(*s, l, conn)
		case strings.HasPrefix(string(msg), "ttl:"):
			HandleTTL(*s, l, conn)
		default:
			fmt.Fprintf(conn, "invalid payload received, skipping")
		}

	}
}

func HandleSet(s Server, l []string, conn net.Conn) {
	fmt.Println("handling set for ", l)
	// TODO: deserialize
	s.t.Set(l[1], l[2])
	fmt.Fprintf(conn, "OK\n")
}

func HandleGet(s Server, l []string, conn net.Conn) {
	fmt.Println("handling get for ", l)
	v, err := s.t.Get(l[1])
	if err != nil {
		fmt.Fprintf(conn, "not found")
	}
	fmt.Fprintf(conn, "%s\n", v)
}

func HandleTTL(s Server, l []string, conn net.Conn) {
	fmt.Println("handling ttl for ", l)
}
