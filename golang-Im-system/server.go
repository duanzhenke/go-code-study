package main

import (
	"fmt"
	"net"
	"strconv"
)

type Server struct {
	Ip   string
	Port int
}

// 创建一个server接口
func NewServer(ip string, port int) *Server {
	server := &Server{
		Ip:   ip,
		Port: port,
	}
	return server
}

func (s *Server) Start() {
	//socket listen
	addStr := s.Ip + ":" + strconv.Itoa(s.Port)
	listen, err := net.Listen("tcp", addStr)
	if err != nil {
		fmt.Println("listen err:", err)
		return
	}

	// close socket listen
	defer listen.Close()
	//accept
	for {
		listenConn, err := listen.Accept()
		if err != nil {
			fmt.Println("Accept err:", err)
			continue
		}
		//do handler
		go s.Handler(listenConn)
	}

}

func (s *Server) Handler(conn net.Conn) {
	// 当前链接的逻辑
	fmt.Println("链接建立成功了")
}
