package main

import "net"

type User struct {
	Name string
	Addr string
	C    chan string
	conn net.Conn
}

func NewUser(conn net.Conn) *User {
	strAddr := conn.RemoteAddr().String()

	user := &User{
		Name: strAddr,
		Addr: strAddr,
		C:    make(chan string),
		conn: conn,
	}
	go user.ListenMessage()
	return user
}

func (r *User) ListenMessage() {
	for {
		msg := <-r.C
		r.conn.Write([]byte(msg + "\n"))
	}
}
