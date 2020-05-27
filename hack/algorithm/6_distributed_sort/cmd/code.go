package main

import (
	"fmt"
	"io/ioutil"
	"os/exec"
)

func GoCmd(cmdString string) string {
	cmd := exec.Command(cmdString)
	cmd.Run()
	return ""
}

func GoCmdWithResult(cmdString string) string {
	cmd := exec.Command(cmdString)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return "error1"
	}
	if err := cmd.Start(); err != nil {
		return "error2"
	}
	outByte, err := ioutil.ReadAll(stdout) // 读取所有的输出
	stdout.Close()
	if err := cmd.Wait(); err != nil {
		return "error3"
	}
	return string(outByte)
}

func main() {
	result := GoCmdWithResult("ifconfig")
	fmt.Println(result)
}
