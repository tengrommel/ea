package main

import (
	"fmt"
	"io"
	"net"
	"os"
	"strconv"
	"time"
)

// 接收
func clientRead(conn net.Conn) int {
	buf := make([]byte, 5)
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println("client read err", err)
	}
	off, err := strconv.Atoi(string(buf[:n]))
	return off
}

// 发送
func clientWrite(conn net.Conn, data []byte) {
	_, err := conn.Write(data)
	if err != nil {
		fmt.Println("client write err", err)
	}
	fmt.Println("写入数据成功", string(data))
}

func clientConn(conn net.Conn) {
	defer conn.Close() // 关闭链接
	clientWrite(conn, []byte("start -->"))
	off := clientRead(conn) // 读取数据
	fp, err := os.OpenFile("1.txt", os.O_RDONLY, 0777)
	if err != nil {
		fmt.Println("open err", err)
	}
	defer fp.Close()
	_, err = fp.Seek(int64(off), 0)
	if err != nil {
		fmt.Println("seek err", err)
	}
	for {
		data := make([]byte, 10)
		n, err := fp.Read(data)
		if err != nil {
			if err == io.EOF {
				// 标记文件结束
				time.Sleep(time.Second)
				clientWrite(conn, []byte("<-- end"))
				fmt.Println("send all ok")
				break
			}
		}
		clientWrite(conn, data[:n])
	}
}

func main() {
	conn, err := net.DialTimeout("tcp", "127.0.0.1:8848", time.Second*10)
	if err != nil {
		fmt.Println("err", err)
		panic(err)
	}
	clientConn(conn)
}
