package main

import "net"

type User struct {
	Name   string
	Addr   string
	C      chan string
	conn   net.Conn
	server *Server
}

//创建一个用户的API

func NewUser(conn net.Conn, server *Server) *User {
	userAddr := conn.RemoteAddr().String() //当前连接的ip地址。通过 remoteADDr获取
	user := &User{
		Name:   userAddr,
		Addr:   userAddr,
		C:      make(chan string),
		conn:   conn,
		server: server,
	}
	//启动监听当前user channel消息的goroutine
	go user.ListenMessage()
	return user
}

//用户的上线功能
func (this *User) Online() {
	//	用户上线，将用户加入到onlineMap中
	this.server.maplock.Lock()
	this.server.OnlineMpa[this.Name] = this
	this.server.maplock.Unlock()
	//	广播当前用户上线消息
	this.server.BroadCast(this, "已上线")
}

//用户的下线功能
func (this *User) Offline() {
	//	用户下线，将用户从onlineMap中删除
	this.server.maplock.Lock()
	delete(this.server.OnlineMpa, this.Name)
	this.server.maplock.Unlock()
	this.server.BroadCast(this, "已下线")

}

//用户处理消息的业务
func (this *User) DoMessage(msg string) {
	this.server.BroadCast(this, "msg")

}

//监听当前User channel的方法，一旦有消息，就直接发送给对端客户端
func (this *User) ListenMessage() {
	for {
		msg := <-this.C
		this.conn.Write([]byte(msg + "\n"))

	}
}
