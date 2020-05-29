package main

import (
	"bufio"
	"ea/hack/algorithm/7_distributed_map_reduce/DistributedServer/SplitFile"
	"ea/hack/algorithm/7_distributed_map_reduce/SlaverServer/SlaveServer/ByteCode"
	"fmt"
	"io"
	"net"
	"os"
	"strconv"
	"strings"
)

func CheckError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
}

// 发送数据
func SendFile(sendFilePath string, conn net.Conn, dataType int, isMem bool, isIncreasing bool) {
	if isMem {
		// 控制slaverServer
		myStringStart := ByteCode.IntToBytes(0)
		myStringStart = append(myStringStart, ByteCode.IntToBytes(0)...)
		myStringStart = append(myStringStart, ByteCode.IntToBytes(0)...)
		conn.Write(myStringStart)

		// 读取文件并发送
		fi, err := os.Open(sendFilePath)
		if err != nil {
			fmt.Println(err)
			return
		}
		br := bufio.NewReader(fi)
		for {
			line, _, err := br.ReadLine()
			if err == io.EOF {
				break
			}
			if dataType == 1 {
				data, _ := strconv.Atoi(string(line))
				dataByte := ByteCode.IntToBytes(0)
				dataByte = append(myStringStart, ByteCode.IntToBytes(1)...)
				dataByte = append(myStringStart, ByteCode.IntToBytes(data)...)
				conn.Write(dataByte)
			} else if dataType == 2 {
				data, _ := strconv.ParseFloat(string(line), 64)
				dataByte := ByteCode.IntToBytes(0)
				dataByte = append(myStringStart, ByteCode.IntToBytes(2)...)
				dataByte = append(myStringStart, ByteCode.Float64ToByte(data)...)
				conn.Write(dataByte)
			} else if dataType == 3 {
				data := string(line)
				dataByte := ByteCode.IntToBytes(0)
				dataByte = append(myStringStart, ByteCode.IntToBytes(3)...)
				dataByte = append(myStringStart, data...)
				conn.Write(dataByte)
			} else if dataType == 4 {
				data := string(line)
				dataList := strings.Split(data, " # ")
				times, _ := strconv.Atoi(dataList[1])
				password := dataList[0]
				dataByte := ByteCode.IntToBytes(0)
				dataByte = append(myStringStart, ByteCode.IntToBytes(4)...)
				dataByte = append(myStringStart, ByteCode.IntToBytes(times)...)
				dataByte = append(myStringStart, ByteCode.IntToBytes(len(password))...)
				dataByte = append(myStringStart, []byte(password)...)
				conn.Write(dataByte)
			}
		}

		fi.Close()
		if isIncreasing {
			// 结束
			myBytesEnd := ByteCode.IntToBytes(0)
			myBytesEnd = append(myBytesEnd, ByteCode.IntToBytes(0)...)
			myBytesEnd = append(myBytesEnd, ByteCode.IntToBytes(1)...)
			conn.Write(myBytesEnd)
		} else {
			// 结束
			myBytesEnd := ByteCode.IntToBytes(0)
			myBytesEnd = append(myBytesEnd, ByteCode.IntToBytes(0)...)
			myBytesEnd = append(myBytesEnd, ByteCode.IntToBytes(2)...)
			conn.Write(myBytesEnd)
		}

	} else {
		// 控制slaverServer
		myStringStart := ByteCode.IntToBytes(1)
		myStringStart = append(myStringStart, ByteCode.IntToBytes(0)...)
		myStringStart = append(myStringStart, ByteCode.IntToBytes(0)...)
		conn.Write(myStringStart)

		// 读取文件并发送
		fi, err := os.Open(sendFilePath)
		if err != nil {
			fmt.Println(err)
			return
		}
		br := bufio.NewReader(fi)
		for {
			line, _, err := br.ReadLine()
			if err == io.EOF {
				break
			}
			if dataType == 1 {
				data, _ := strconv.Atoi(string(line))
				dataByte := ByteCode.IntToBytes(1)
				dataByte = append(myStringStart, ByteCode.IntToBytes(1)...)
				dataByte = append(myStringStart, ByteCode.IntToBytes(data)...)
				conn.Write(dataByte)
			} else if dataType == 2 {
				data, _ := strconv.ParseFloat(string(line), 64)
				dataByte := ByteCode.IntToBytes(1)
				dataByte = append(myStringStart, ByteCode.IntToBytes(2)...)
				dataByte = append(myStringStart, ByteCode.Float64ToByte(data)...)
				conn.Write(dataByte)
			} else if dataType == 3 {
				data := string(line)
				dataByte := ByteCode.IntToBytes(1)
				dataByte = append(myStringStart, ByteCode.IntToBytes(3)...)
				dataByte = append(myStringStart, data...)
				conn.Write(dataByte)
			} else if dataType == 4 {
				data := string(line)
				dataList := strings.Split(data, " # ")
				times, _ := strconv.Atoi(dataList[1])
				password := dataList[0]
				dataByte := ByteCode.IntToBytes(1)
				dataByte = append(myStringStart, ByteCode.IntToBytes(4)...)
				dataByte = append(myStringStart, ByteCode.IntToBytes(times)...)
				dataByte = append(myStringStart, ByteCode.IntToBytes(len(password))...)
				dataByte = append(myStringStart, []byte(password)...)
				conn.Write(dataByte)
			}
		}

		fi.Close()

		if isIncreasing {
			// 结束
			myBytesEnd := ByteCode.IntToBytes(1)
			myBytesEnd = append(myBytesEnd, ByteCode.IntToBytes(0)...)
			myBytesEnd = append(myBytesEnd, ByteCode.IntToBytes(1)...)
			conn.Write(myBytesEnd)
		} else {
			// 结束
			myBytesEnd := ByteCode.IntToBytes(1)
			myBytesEnd = append(myBytesEnd, ByteCode.IntToBytes(0)...)
			myBytesEnd = append(myBytesEnd, ByteCode.IntToBytes(2)...)
			conn.Write(myBytesEnd)
		}

	}
}

// 接收数据
func RecvFile(saveFilePath string, conn net.Conn, dataType int) {
	saveFile, _ := os.Create(saveFilePath)
	var save *bufio.Writer
	save = bufio.NewWriter(saveFile)
	save.Flush()
	saveFile.Close()
}

//
func main() {
	// 输入文件路径
	path1 := "file1"
	// 结果文件夹
	//path2 := "file2"

	pathResult := "tmp"

	ipList := []string{"127.0.0.1:8848", "127.0.0.1:8883"}

	// 切割
	N := len(ipList)
	saveList := SplitFile.SplitFile(path1, pathResult, N)
	fmt.Println(saveList)
	connList := make([]net.Conn, 0)

	for i := 0; i < len(ipList); i++ {
		tcpAddr, err := net.ResolveTCPAddr("tcp4", ipList[i])
		CheckError(err)
		conn, err := net.DialTCP("tcp", nil, tcpAddr)
		CheckError(err)
		connList = append(connList, conn)
	}
	// 发送
	for i := 0; i < len(connList); i++ {
		go func() {
			SendFile(saveList[i], connList[i], 3, true, true)
		}()
	}
	// 接收
	for i := 0; i < len(connList); i++ {
		savePath := strings.Replace(saveList[i], ".txt", "sort.txt", -1)
		go func() {
			RecvFile(savePath, connList[i], 3)
		}()
	}

	// 归并

}
