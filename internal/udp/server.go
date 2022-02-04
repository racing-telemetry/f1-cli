package udp

import "net"

type Server struct {
	conn net.Conn // UDP Connection
}

func Dial(addr *net.UDPAddr) (*Server, error) {
	conn, err := net.Dial("udp", addr.String())
	if err != nil {
		return nil, err
	}

	return &Server{conn: conn}, nil
}

func (s *Server) WriteSocket(buf []byte) error {
	_, err := s.conn.Write(buf)
	if err != nil {
		return err
	}

	return nil
}

func (s *Server) Close() error {
	return s.conn.Close()
}
