package main

import (
	"ea/hack/algorithm/7_distributed_map_reduce/SlaverServer/SlaveServer/ByteCode"
	"fmt"
	"net"
	"os"
	"strconv"
	"time"
)

func CheckError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
}

func SendArray(arr []int, conn net.Conn) {
	myStringStart := ByteCode.IntToBytes(0)
	myStringStart = append(myStringStart, ByteCode.IntToBytes(0)...)
	myStringStart = append(myStringStart, ByteCode.IntToBytes(0)...)
	conn.Write(myStringStart)

	myArr := []int{1, 9, 2, 8, 3, 7, 6, 4, 5, 0}
	for i := 0; i < len(myArr); i++ {
		myBytesData := ByteCode.IntToBytes(0)
		myBytesData = append(myBytesData, ByteCode.IntToBytes(1)...)
		myBytesData = append(myBytesData, ByteCode.IntToBytes(myArr[i])...)
		conn.Write(myBytesData)
	}

	// 结束
	myBytesEnd := ByteCode.IntToBytes(0)
	myBytesEnd = append(myBytesEnd, ByteCode.IntToBytes(1)...)
	conn.Write(myBytesEnd)
}

// 接收结果
func ServerMsg(conn net.Conn) <-chan int {
	out := make(chan int, 1024)
	defer conn.Close()
	arr := make([]int, 0)
	for {
		buf := make([]byte, 24)
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("服务器关闭")
			return nil
		}
		if n == 24 {
			data1 := ByteCode.BytesToInt(buf[:8])
			data2 := ByteCode.BytesToInt(buf[8:16])
			data3 := ByteCode.BytesToInt(buf[16:])
			fmt.Println(data1, data2, data3)
			if data1 == 0 && data2 == 0 && data3 == 0 {
				arr = make([]int, 0, 0)
				fmt.Println("开始")
			}
			if data1 == 0 && data2 == 1 {
				fmt.Println("收到", data3)
				arr = append(arr, data3)
			}
			if data1 == 0 && data2 == 0 && (data3 == 1 || data3 == 2) {
				fmt.Println("收到数组", arr)
				for i := 0; i < len(arr); i++ {
					out <- arr[i]
				}
				close(out)
				return out
			}
		}
	}
}

// 归并结果
func Merge(in1, in2 <-chan int) <-chan int {
	out := make(chan int, 1024)
	go func() {
		v1, ok1 := <-in1
		v2, ok2 := <-in2
		for ok1 || ok2 { // 可以取出的话一直取
			if !ok2 || (ok1 && v1 <= v2) { // 轮番取出数据
				out <- v1
				v1, ok1 = <-in1
			} else {
				out <- v2
				v2, ok2 = <-in2
			}
		}
		close(out)
		fmt.Println("归并排序完成")
	}()
	return out
}

func main() {
	arrList := [][]int{{1, 100, 2, 93}, {1000, 89, 299, 199}}
	sortResult := make([]<-chan int, 0)
	for i := 0; i < 2; i++ {
		tcpAddr, err := net.ResolveTCPAddr("tcp4", "127.0.0.1:"+strconv.Itoa(8888+i))
		CheckError(err)
		conn, err := net.DialTCP("tcp", nil, tcpAddr)
		CheckError(err)
		// 发送数组
		SendArray(arrList[i], conn)
		sortResult = append(sortResult, ServerMsg(conn))
		// 接收结果
		// 归并
	}
	fmt.Println(sortResult)
	fmt.Println(sortResult[0])
	fmt.Println(sortResult[1])
	last := Merge(sortResult[0], sortResult[1])
	for v := range last {
		fmt.Printf("%d ", v)
	}
	time.Sleep(time.Second % 100)
}
