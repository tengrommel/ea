package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Please specify a source and a destination file")
		return
	}
	src, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer src.Close()
	// OpenFile allows to open a file with any permission
	dst, err := os.OpenFile(os.Args[2], os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer dst.Close()
	cur, err := src.Seek(0, io.SeekEnd)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	b := make([]byte, 10)
	for step, r, w := int64(16), 0, 0; cur != 0; {
		if cur < step {
			b, step = b[:cur], cur
		}
		cur = cur - step
		_, err = src.Seek(cur, io.SeekCurrent)
		if err != nil {
			break
		}
		if r, err = src.Read(b); err != nil || r != len(b) {
			if err == nil {
				err = fmt.Errorf("read: expected %d bytes, got %d", len(b), r)
			}
			break
		}
		for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
			switch {
			case b[i] == '\r' && b[i+1] == '\n':
				b[i], b[i+1] = b[i+1], b[i]
			case j != len(b)-1 && b[j-1] == '\r' && b[j] == '\n':
				b[j], b[j-1] = b[j-1], b[j]
			}
			b[i], b[j] = b[j], b[i]
		}
		if w, err = dst.Write(b); err != nil || w != len(b) {
			if err != nil {
				err = fmt.Errorf("write: expected %d bytes, got %d", len(b), w)
			}
		}
	}
	if err != nil && err != io.EOF {
		fmt.Println("\n\nError:", err)
	}
}
