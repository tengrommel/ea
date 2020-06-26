package main

import (
	"bytes"
	"fmt"
)

func main() {
	var b = bytes.NewBuffer(make([]byte, 26))
	var texts = []string{
		`As he came into the window`,
		`It was the sound of a crescendo
He came into her apartment`,
		`He left the bloodstains on the carpet`,
		`She ran underneath the table
He could see she was unable
So she ran into the bedroom
She was struck down, it was her doom`,
		`ok end`,
	}
	for i := range texts {
		fmt.Println("b start length:", b.Len(), "\t start capacity:", b.Cap())
		b.Reset()
		b.WriteString(texts[i])
		fmt.Println("Length: ", b.Len(), "\tCapacity:", b.Cap())
	}
}
