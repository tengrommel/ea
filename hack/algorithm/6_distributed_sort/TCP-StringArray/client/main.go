package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"net"
	"time"
)

func IntToBytes(n int) []byte {
	data := int64(n)                                 // 数据类型的转换
	byteBuffer := bytes.NewBuffer([]byte{})          // 字节集合
	binary.Write(byteBuffer, binary.BigEndian, data) // 按照二进制写入字节
	return byteBuffer.Bytes()
}

func BytesToInt(bs []byte) int {
	byteBuffer := bytes.NewBuffer(bs) // 根据二进制写入二进制结合
	var data int64
	binary.Read(byteBuffer, binary.BigEndian, &data)
	return int(data)
}

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8848")
	defer conn.Close()
	if err != nil {
		fmt.Println("客户端连接失败")
		return
	}
	time.Sleep(1 * time.Second)
	// 等待给服务器输入信息
	go func() {
		// 开始
		myBytesStart := IntToBytes(0)
		myBytesStart = append(myBytesStart, IntToBytes(0)...)
		conn.Write(myBytesStart)

		myArray := []string{"1", "9", "2", "8", "3", "7", "6", "4", "5", "0"}
		for i := 0; i < len(myArray); i++ {
			myBytesData := IntToBytes(3)
			myBytesData = append(myBytesData, IntToBytes(len(myArray[i]))...)
			conn.Write(myBytesData)
			conn.Write([]byte(myArray[i]))
		}
		// 结束
		myBytesEnd := IntToBytes(0)
		myBytesEnd = append(myBytesEnd, IntToBytes(1)...)
		conn.Write(myBytesEnd)
	}()

	arr := make([]string, 0)
	for {
		// 等待，接收信息
		buf := make([]byte, 16)
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("服务器关闭")
			return
		}
		if n == 16 {
			data1 := BytesToInt(buf[:len(buf)/2]) // 取出第一个数
			data2 := BytesToInt(buf[len(buf)/2:]) // 取出第二个数
			if data1 == 0 && data2 == 0 {
				// 开始
				arr = make([]string, 0, 0)
			}
			if data1 == 3 {
				// 接收数组
				stringByte := make([]byte, data2, data2)
				length, _ := conn.Read(stringByte)
				if length == data2 { // 校验长度
					arr = append(arr, string(stringByte))
				}
			}
			if data1 == 0 && data2 == 1 {
				// 结束
				fmt.Println("收到数组 arr:", arr)
				arr = nil
			}
		}
	}
}
