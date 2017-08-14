package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
)

func stream(reader io.Reader, writer io.Writer) error {
	bufReader := bufio.NewReader(reader)
	for {
		data, err := bufReader.ReadByte()
		if err != nil {
			return err
		}
		_, err = writer.Write([]byte{40, data, 41})
		if err != nil {
			return err
		}
	}
}

func main() {
	// disable input buffering
	exec.Command("stty", "-F", "/dev/tty", "cbreak", "min", "1").Run()

	err := stream(os.Stdin, os.Stdout)
	if err != nil {
		fmt.Println(err)
	}
}
