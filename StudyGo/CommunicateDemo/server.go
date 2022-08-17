package main

import (
	"fmt"
	"io"
	"net"
	"sync"
)

type Server struct {
	Ip   string
	Port int
	//在线用户的列表
	OnlineMpa map[string]*User
	maplock   sync.RWMutex
	Message   chan string
}

//创建一个server接口
func NewSever(ip string, port int) *Server {
	server := &Server{
		Ip:        ip,
		Port:      port,
		OnlineMpa: make(map[string]*User),
		Message:   make(chan string),
	}
	return server
}

//监听Message广播消息channel的goroutine，一旦有消息就发送给全部在线User
func (this *Server) ListenMessager() {
	for {
		msg := <-this.Message
		this.maplock.Lock()
		for _, cli := range this.OnlineMpa {
			cli.C <- msg
		}
		this.maplock.Unlock()
	}
}

//广播消息的方法
func (this *Server) BroadCast(user *User, msg string) {
	sendMsg := "[" + user.Addr + "]" + user.Name + ":" + msg
	this.Message <- sendMsg
}

func (this *Server) Handler(conn net.Conn) {
	//	...当前的业务链。
	//fmt.Println("连接建立成功")
	user := NewUser(conn, this)
	//	用户上线，将用户加入到onlineMap
	user.Online()
	//接收客户端发送的消息
	go func() {
		buf := make([]byte, 5096)
		for {
			n, err := conn.Read(buf)
			if n == 0 {
				user.Offline()
				return
			}
			if err != nil && err != io.EOF {
				fmt.Println("Conn Read err:", err)
			}
			//提取用户的消息（去除\n）
			msg := string(buf[:n-1])

			//	将得到的消息进行广播
			user.DoMessage(msg)
		}
	}()
	//	当前handler阻塞
	select {}

}

//启动服务器的接口
func (this *Server) Start() {
	//	socket Listen
	listener, err := net.Listen("tcp", fmt.Sprintf("%v:%v", this.Ip, this.Port))
	if err != nil {
		fmt.Println("net.Listen err", err)
		return
	}
	//	close listen socket
	defer listener.Close()

	//启动监听Message的Goroutine
	go this.ListenMessager()
	//	accept

	for {
		//	accept
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("listener accept err:", err)
			continue
		}
		//	do handler
		go this.Handler(conn)
	}

	//	do handler

}
