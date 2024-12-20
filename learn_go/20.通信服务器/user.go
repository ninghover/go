package main

import (
	"net"
	"strings"
)

type User struct {
	Name   string
	Addr   string
	C      chan string
	conn   net.Conn
	server *Server
}

func NewUser(conn net.Conn, server *Server) *User {
	userAddr := conn.RemoteAddr().String()
	user := &User{
		Name:   userAddr,
		Addr:   userAddr,
		C:      make(chan string),
		conn:   conn,
		server: server,
	}
	go user.ListenMessage()
	return user
}

// 用户上线
func (user *User) Online() {
	user.server.mapLock.Lock()
	user.server.OnlineMap[user.Name] = user
	user.server.mapLock.Unlock()
	user.server.BroadCast(user, "上线了")
}

// 用户下线
func (user *User) Offline() {
	user.server.mapLock.Lock()
	delete(user.server.OnlineMap, user.Name)
	user.server.mapLock.Unlock()
	user.server.BroadCast(user, "下线了")
}

func (user *User) SendMsg(msg string) {
	user.conn.Write([]byte(msg))
}

// 用户消息广播
func (user *User) DoMessage(msg string) {
	if msg == "who" { // 查看所有在线用户
		user.server.mapLock.Lock()
		for _, u := range user.server.OnlineMap {
			sendMsg := "[" + u.Addr + "]" + u.Name + ":在线...\n"
			user.SendMsg(sendMsg)
		}
		user.server.mapLock.Unlock()
	} else if len(msg) >= 7 && msg[:7] == "rename " { // 更新用户名
		newName := strings.Split(msg, " ")[1]
		_, ok := user.server.OnlineMap[newName]
		if ok {
			user.SendMsg("新用户名已存在！\n")
		} else {
			user.server.mapLock.Lock()
			delete(user.server.OnlineMap, user.Name)
			user.Name = newName
			user.server.OnlineMap[newName] = user
			user.server.mapLock.Unlock()
			user.SendMsg("用户名更新成功！\n")
		}

	} else if len(msg) >= 3 && msg[:3] == "to " { // 私聊功能
		toName := strings.Split(msg, " ")[1]
		toMsg := strings.Split(msg, " ")[2]
		toUser, ok := user.server.OnlineMap[toName]
		if !ok {
			user.SendMsg(toName + "不在线！\n")
		} else {
			toUser.SendMsg(toMsg)
		}

	} else {
		user.server.BroadCast(user, msg)
	}

}

// 监听当前User的channel，一旦有消息就发送给对端客户端
func (u *User) ListenMessage() {
	for {
		msg := <-u.C
		u.conn.Write([]byte(msg + "\n"))
	}
}
