package main

import (
	"bufio"
	"ea/hack/algorithm/7_distributed_map_reduce/SlaverServer/SlaveServer/ByteCode"
	"ea/hack/algorithm/7_distributed_map_reduce/SlaverServer/SlaveServer/sort"
	"fmt"
	"io"
	"math/rand"
	"net"
	"os"
	"runtime"
	"runtime/debug"
	"strconv"
	"strings"
	"time"
)

type Pass struct {
	PassWord string
	times    int
}

func Server(conn net.Conn) {
	if conn == nil {
		fmt.Println("无效连接")
		return
	}
	// 接收数据，处理
	arr := make([]interface{}, 0) // 泛用
	filePath := "."
	fmt.Println(filePath)
	myType := 0
	for {
		// 等待，接收信息
		buf := make([]byte, 24)
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("服务器关闭")
			return
		}
		if n == 24 {
			data1 := ByteCode.BytesToInt(buf[:8])   // 取出第一个数
			data2 := ByteCode.BytesToInt(buf[8:16]) // 取出第二个数
			data3 := ByteCode.BytesToInt(buf[16:])  // 取出第三个数

			if data1 == 0 {
				// 内存模式
				if data2 == 0 && data3 == 0 {
					// 开始
					arr = make([]interface{}, 0, 0)
				}
				if data2 == 1 {
					// 整数
					arr = append(arr, data3)
					myType = 1
				} else if data2 == 2 {
					// 实数
					arr = append(arr, ByteCode.ByteToFloat64(buf[16:]))
					myType = 2
				} else if data2 == 3 {
					// 字符串
					stringByte := make([]byte, data3, data3)
					length, _ := conn.Read(stringByte)
					if length == data3 {
						arr = append(arr, string(stringByte))
					}
					myType = 3
				} else if data2 == 4 {
					// 结构体
					buf1 := make([]byte, 8)
					length, _ := conn.Read(buf)
					if length == 8 {
						data4 := ByteCode.BytesToInt(buf1)
						stringByte := make([]byte, data4, data4)
						length, _ := conn.Read(stringByte)
						if length == data4 {
							tmpPass := Pass{string(stringByte), data3}
							arr = append(arr, tmpPass)
						}
					}
					myType = 4
				}
				if data2 == 0 && data3 == 1 {
					// 结束 从小到大
					fmt.Println("收到数组", arr)
					myData := new(sort.QuickSortData)
					myData.Data = arr
					myData.IsIncreasing = true
					QuickSortDataByType(myData, myType)
					myData.QuickSort()
					fmt.Println("最终数组", arr)
					arr = nil
					runtime.GC()
					debug.FreeOSMemory() //释放内存
					SendResult(myData, conn, myType, true)
				}
				if data2 == 0 && data3 == 2 {
					// 结束 从大到小
					fmt.Println("收到数组", arr)
					myData := new(sort.QuickSortData)
					myData.IsIncreasing = false
					QuickSortDataByType(myData, myType)
					myData.QuickSort()
					fmt.Println("最终数组", arr)
					arr = nil
					runtime.GC()
					debug.FreeOSMemory() //释放内存
					SendResult(myData, conn, myType, false)
				}
			} else if data1 == 1 {
				// 硬盘模式
			AAA:
				rand.Seed(time.Now().UnixNano())
				fileNameNumber := rand.Int() % 10000
				fileStorePath := filePath + strconv.Itoa(fileNameNumber) + ".txt"
				if _, err := os.Stat(fileStorePath); err == nil {
					goto AAA
				}
				fmt.Println(fileStorePath)
				saveFile, _ := os.Create(fileStorePath)
				var save *bufio.Writer
				if data2 == 0 && data3 == 0 {
					// 开始
					save = bufio.NewWriter(saveFile)
				}
				if data2 == 1 {
					// 整数
					//arr = append(arr, data3)
					fmt.Fprintln(save, strconv.Itoa(data3))
					myType = 1
				} else if data2 == 2 {
					// 实数
					floatNum := ByteCode.ByteToFloat64(buf[16:])
					fmt.Fprintln(save, strconv.FormatFloat(floatNum, 'f', 6, 64))
					myType = 2
				} else if data3 == 3 {
					stringByte := make([]byte, data3, data3)
					length, _ := conn.Read(stringByte)
					if length == data3 {
						//arr = append(arr, string(stringByte))
						fmt.Fprintln(save, string(stringByte))
					}
					myType = 3
				} else if data2 == 4 {
					// 结构体
					buf1 := make([]byte, 8)
					length, _ := conn.Read(buf)
					if length == 8 {
						data4 := ByteCode.BytesToInt(buf1)
						stringByte := make([]byte, data4, data4)
						length, _ := conn.Read(stringByte)
						if length == data4 {
							fmt.Fprintln(save, string(stringByte)+" # "+strconv.Itoa(data3))
						}
					}
					myType = 4
				}
				if data2 == 0 && data3 == 1 {
					// 结束，从大到小
					save.Flush()
					saveFile.Close()
					// 类型划分
					// 读取文件
					// 排序
					// 写入文件
					// 读取文件传输
					DiskFileSortAndSend(fileStorePath, myType, true, conn)
				}
				if data2 == 0 && data3 == 2 {
					// 结束，从小到大
					save.Flush()
					saveFile.Close()
					DiskFileSortAndSend(fileStorePath, myType, false, conn)
				}
			}
		}
	}
}

func SendResult(myData *sort.QuickSortData, conn net.Conn, myType int, isIncreasing bool) {
	if myType == 1 {
		// 整数
		myStringStart := ByteCode.IntToBytes(0)
		myStringStart = append(myStringStart, ByteCode.IntToBytes(0)...)
		myStringStart = append(myStringStart, ByteCode.IntToBytes(0)...)
		conn.Write(myStringStart)
		for i := 0; i < len(myData.Data); i++ {
			conn.Write(ByteCode.IntToBytes(0))
			myBytesData := ByteCode.IntToBytes(1)
			conn.Write(myBytesData)
			conn.Write(ByteCode.IntToBytes(myData.Data[i].(int)))
		}
		// 结束
		myBytesEnd := ByteCode.IntToBytes(0)
		myBytesEnd = append(myBytesEnd, ByteCode.IntToBytes(0)...)
		if isIncreasing {
			myBytesEnd = append(myBytesEnd, ByteCode.IntToBytes(1)...)
		} else {
			myBytesEnd = append(myBytesEnd, ByteCode.IntToBytes(2)...)
		}
		conn.Write(myBytesEnd)
	} else if myType == 2 {
		// 实数
		myStringStart := ByteCode.IntToBytes(0)
		myStringStart = append(myStringStart, ByteCode.IntToBytes(0)...)
		myStringStart = append(myStringStart, ByteCode.IntToBytes(0)...)
		conn.Write(myStringStart)
		for i := 0; i < len(myData.Data); i++ {
			conn.Write(ByteCode.IntToBytes(0))
			myBytesData := ByteCode.IntToBytes(2)
			conn.Write(myBytesData)
			conn.Write(ByteCode.Float64ToByte(myData.Data[i].(float64)))
		}
		// 结束
		myBytesEnd := ByteCode.IntToBytes(0)
		myBytesEnd = append(myBytesEnd, ByteCode.IntToBytes(0)...)
		if isIncreasing {
			myBytesEnd = append(myBytesEnd, ByteCode.IntToBytes(1)...)
		} else {
			myBytesEnd = append(myBytesEnd, ByteCode.IntToBytes(2)...)
		}
		conn.Write(myBytesEnd)
	} else if myType == 3 {
		// 字符串
		myStringStart := ByteCode.IntToBytes(0)
		myStringStart = append(myStringStart, ByteCode.IntToBytes(0)...)
		myStringStart = append(myStringStart, ByteCode.IntToBytes(0)...)
		conn.Write(myStringStart)

		for i := 0; i < len(myData.Data); i++ {
			conn.Write(ByteCode.IntToBytes(0))
			myBytesData := ByteCode.IntToBytes(3)
			myBytesData = append(myBytesData, ByteCode.IntToBytes(len(myData.Data[i].(string)))...)
			conn.Write(myBytesData)
			conn.Write([]byte(myData.Data[i].(string)))
		}
		// 结束
		myBytesEnd := ByteCode.IntToBytes(0)
		myBytesEnd = append(myBytesEnd, ByteCode.IntToBytes(0)...)
		if isIncreasing {
			myBytesEnd = append(myBytesEnd, ByteCode.IntToBytes(1)...)
		} else {
			myBytesEnd = append(myBytesEnd, ByteCode.IntToBytes(2)...)
		}
		conn.Write(myBytesEnd)
	} else if myType == 4 {
		// 结构体
		// 字符串
		myStringStart := ByteCode.IntToBytes(0)
		myStringStart = append(myStringStart, ByteCode.IntToBytes(0)...)
		myStringStart = append(myStringStart, ByteCode.IntToBytes(0)...)
		conn.Write(myStringStart)
		for i := 0; i < len(myData.Data); i++ {
			conn.Write(ByteCode.IntToBytes(0))
			myBytesData := ByteCode.IntToBytes(4)
			myBytesData = append(myBytesData, ByteCode.IntToBytes(myData.Data[i].(Pass).times)...)
			myBytesData = append(myBytesData, ByteCode.IntToBytes(len(myData.Data[i].(Pass).PassWord))...)
			conn.Write(myBytesData)
			conn.Write([]byte(myData.Data[i].(Pass).PassWord))
		}
		// 结束
		myBytesEnd := ByteCode.IntToBytes(0)
		myBytesEnd = append(myBytesEnd, ByteCode.IntToBytes(0)...)
		if isIncreasing {
			myBytesEnd = append(myBytesEnd, ByteCode.IntToBytes(1)...)
		} else {
			myBytesEnd = append(myBytesEnd, ByteCode.IntToBytes(2)...)
		}
		conn.Write(myBytesEnd)

	}
}

func QuickSortDataByType(myData *sort.QuickSortData, myType int) {
	if myType == 1 {
		myData.MyFunc = func(data1, data2 interface{}, isIncreasing bool) bool {
			if isIncreasing {
				return data1.(int) < data2.(int)
			} else {
				return data1.(int) > data2.(int)
			}
		}
	} else if myType == 2 {
		myData.MyFunc = func(data1, data2 interface{}, isIncreasing bool) bool {
			if isIncreasing {
				return data1.(float64) < data2.(float64)
			} else {
				return data1.(float64) > data2.(float64)
			}
		}
	} else if myType == 3 {
		myData.MyFunc = func(data1, data2 interface{}, isIncreasing bool) bool {
			if isIncreasing {
				return data1.(string) < data2.(string)
			} else {
				return data1.(string) > data2.(string)
			}
		}
	} else if myType == 4 {
		myData.MyFunc = func(data1, data2 interface{}, isIncreasing bool) bool {
			if isIncreasing {
				return data1.(Pass).times < data2.(Pass).times
			} else {
				return data1.(Pass).times > data2.(Pass).times
			}
		}
	}
}

func DiskFileSortAndSend(filePath string, myType int, isIncreasing bool, conn net.Conn) {
	arr := make([]interface{}, 0)
	fi, err := os.Open(filePath)
	defer fi.Close()
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
		if myType == 1 {
			data, _ := strconv.Atoi(string(line))
			arr = append(arr, data)

		} else if myType == 2 {
			data, _ := strconv.ParseFloat(string(line), 64)
			arr = append(arr, data)
		} else if myType == 3 {
			data := string(line)
			arr = append(arr, data)
		} else if myType == 4 {
			data := string(line)
			dataList := strings.Split(data, " # ")
			times, _ := strconv.Atoi(dataList[1])
			arr = append(arr, Pass{dataList[0], times})
		}
	}
	// 排序
	myData := new(sort.QuickSortData)
	myData.IsIncreasing = true
	myData.Data = arr
	QuickSortDataByType(myData, myType)
	myData.QuickSort()
	fmt.Println("最终数组", arr)

	newPath := strings.Replace(filePath, ".txt", "sort.txt", -1)
	saveFile, _ := os.Create(newPath)
	save := bufio.NewWriter(saveFile)
	for i := 0; i < len(myData.Data); i++ {
		if myType == 1 {
			fmt.Fprintln(save, strconv.Itoa(myData.Data[i].(int)))
		} else if myType == 2 {
			fmt.Fprintln(save, strconv.FormatFloat(myData.Data[i].(float64), 'f', 6, 64))
		} else if myType == 3 {
			fmt.Fprintln(save, myData.Data[i].(string))
		} else if myType == 4 {
			fmt.Fprintln(save, myData.Data[i].(Pass).PassWord+" # "+strconv.Itoa(myData.Data[i].(Pass).times))
		}
	}
	save.Flush()
	saveFile.Close()

	// 释放内存
	arr = nil
	runtime.GC()
	debug.FreeOSMemory()

	myStringStart := ByteCode.IntToBytes(1)
	myStringStart = append(myStringStart, ByteCode.IntToBytes(0)...)
	myStringStart = append(myStringStart, ByteCode.IntToBytes(0)...)
	conn.Write(myStringStart)

	fileSort, err := os.Open(newPath)
	if err != nil {
		fmt.Println(err)
		return
	}
	brSort := bufio.NewReader(fileSort)
	for {
		line, _, err := brSort.ReadLine()
		if err == io.EOF {
			break
		}
		if myType == 1 {
			data, _ := strconv.Atoi(string(line))
			//arr = append(arr, data)
			for i := 0; i < len(myData.Data); i++ {
				conn.Write(ByteCode.IntToBytes(1))
				conn.Write(ByteCode.IntToBytes(1))
				conn.Write(ByteCode.IntToBytes(data))
			}
		} else if myType == 2 {
			data, _ := strconv.ParseFloat(string(line), 64)
			//arr = append(arr, data)
			for i := 0; i < len(myData.Data); i++ {
				conn.Write(ByteCode.IntToBytes(1))
				conn.Write(ByteCode.IntToBytes(2))
				conn.Write(ByteCode.Float64ToByte(data))
			}
		} else if myType == 3 {
			conn.Write(ByteCode.IntToBytes(1))
			conn.Write(ByteCode.IntToBytes(3))
			conn.Write(ByteCode.IntToBytes(len(string(line))))
			conn.Write(line)
		} else if myType == 4 {
			data := string(line)
			dataList := strings.Split(data, " # ")
			password := dataList[0]
			times, _ := strconv.Atoi(dataList[1])
			conn.Write(ByteCode.IntToBytes(1))
			conn.Write(ByteCode.IntToBytes(4))
			conn.Write(ByteCode.IntToBytes(times))
			conn.Write(ByteCode.IntToBytes(len(password)))
			conn.Write([]byte(password))
		}
	}
	fileSort.Close()

	myBytesEnd := ByteCode.IntToBytes(0)
	myBytesEnd = append(myBytesEnd, ByteCode.IntToBytes(0)...)
	if isIncreasing {
		myBytesEnd = append(myBytesEnd, ByteCode.IntToBytes(1)...)
	} else {
		myBytesEnd = append(myBytesEnd, ByteCode.IntToBytes(2)...)
	}
	conn.Write(myBytesEnd)
}

func main() {
	server, err := net.Listen("tcp", "127.0.0.1:8848")
	defer server.Close()
	if err != nil {
		fmt.Println("服务器开启失败")
		return
	}
	fmt.Println("开启服务器成功")
	for {
		conn, err := server.Accept()
		if err != nil {
			fmt.Println("链接出错")
		}
		go Server(conn) // 并发处理
	}
}
