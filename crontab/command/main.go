package main

import (
	"context"
	"fmt"
	"os/exec"
	"time"
)

type result struct {
	err    error
	output []byte
}

func main() {
	var (
		ctx        context.Context
		cmd        *exec.Cmd
		cancelFunc context.CancelFunc
		resultChan chan *result
	)
	resultChan = make(chan *result, 1000)

	ctx, cancelFunc = context.WithCancel(context.TODO())

	go func() {
		var (
			output []byte
			err    error
		)
		cmd = exec.CommandContext(ctx, "/bin/bash", "-c", "sleep 5;ls -l")
		output, err = cmd.CombinedOutput()
		resultChan <- &result{err: err, output: output}
	}()
	time.Sleep(1 * time.Second)
	cancelFunc()
	res := <-resultChan
	fmt.Println(res.output, res.err)
}
