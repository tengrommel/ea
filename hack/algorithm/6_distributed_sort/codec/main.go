package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"math"
)

//00
//12
//13
//19
//01
// 整数类型 字节
// 实数类型 字节

// 0 00 字符串
// 1 4 calc
// 1 5 cmdstr
// 0 0 1

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

func ByteToFloat32(bs []byte) float32 {
	bits := binary.LittleEndian.Uint32(bs)
	return math.Float32frombits(bits)
}

func ByteToFloat64(bs []byte) float64 {
	bits := binary.LittleEndian.Uint64(bs)
	return math.Float64frombits(bits)
}

func Float32ToByte(data float32) []byte {
	bits := math.Float32bits(data)
	bytes := make([]byte, 4)
	binary.LittleEndian.PutUint32(bytes, bits) // 数据填充
	return bytes
}

func Float64ToByte(data float64) []byte {
	bits := math.Float64bits(data)
	bytes := make([]byte, 4)
	binary.LittleEndian.PutUint64(bytes, bits)
	return bytes
}

func main() {
	// 字符串转字节
	fmt.Println(IntToBytes(31))
	fmt.Println(BytesToInt(IntToBytes(31)))
	fmt.Println(Float32ToByte(121.12334134))
	fmt.Println(ByteToFloat32(Float32ToByte(231.2433243)))
}
