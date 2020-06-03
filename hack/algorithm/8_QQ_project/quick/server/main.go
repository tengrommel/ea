package main

import (
	"fmt"
	"io"
	"net"
	"os"
	"strconv"
)

// 追加
func WriteFile(content []byte) {
	if len(content) != 0 {
		fp, err := os.OpenFile("sr_1.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0777)
		defer fp.Close()
		if err != nil {
			fmt.Println("open file err ", err)
		}
		_, err = fp.Write(content)
		if err != nil {
			fmt.Println("WriteFile ok")
		}
	}
}

// 判断文件是否存在，提取大小
func getFileStat() int64 {
	fileInfo, err := os.Stat("sr_1.txt")
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("文件不存在")
			return 0
		}
	}
	return fileInfo.Size() // 返回文件大小
}

func ServerConn(conn net.Conn) {
	defer conn.Close()
	for {
		var buf = make([]byte, 10)
		n, err := conn.Read(buf)
		if err != nil {
			if err == io.EOF {
				fmt.Println("server is EOF")
				return
			}
			fmt.Println("server read ", err)
		}
		fmt.Println("收到", string(buf[:n]))
		switch string(buf[:n]) {
		case "start-->":
			off := getFileStat()
			stringOff := strconv.FormatInt(off, 10)
			_, err := conn.Write([]byte(stringOff))
			if err != nil {
				fmt.Println("写入", err)
			}
			continue
		case "<--end":
			fmt.Println("文件写入ok")
			return
		}
		WriteFile(buf[:n])
	}
}

func main() {
	sr, err := net.Listen("tcp", "127.0.0.1:8848")
	if err != nil {
		fmt.Println("main err ", err)
	}
	defer sr.Close()
	conn, err := sr.Accept()
	if err != nil {
		fmt.Println("accept err ", err)
	}
	ServerConn(conn)
}
