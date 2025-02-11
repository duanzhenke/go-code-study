package main

import (
	"fmt"
	"net"
	"strconv"
)

type server struct {
	Ip   string
	Port int
}

// 创建一个server接口
func (receiver server) NewServer(ip string, port int) *server {
	server := &server{
		Ip:   ip,
		Port: port,
	}
	return server
}

func (s *server) Start() {
	//socket listen
	listen, err := net.Listen("tcp", s.Ip+":"+strconv.Itoa(s.Port))
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

func (s *server) Handler(conn net.Conn) {
	// 当前链接的逻辑
	fmt.Println("链接建立成功了")
}
