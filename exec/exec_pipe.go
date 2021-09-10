package main

import (
	"bytes"
	"io"
	"log"
	"os"
	"os/exec"
)

func main() {
	c1 := exec.Command("cat", "exec/input.txt")
	c2 := exec.Command("wc")

	reader, writer := io.Pipe()
	c1.Stdout = writer
	c2.Stdin = reader

	defer func() {
		if err := writer.Close(); err != nil {
			log.Println(err)
		}
		if err := reader.Close(); err != nil {
			log.Println(err)
		}
	}()

	var buf, e1, e2 bytes.Buffer
	c2.Stdout = &buf
	c1.Stderr = &e1
	c2.Stderr = &e2

	if err := c1.Start(); err != nil {
		log.Fatalf("failed to start c1, reason: %s", err)
	}
	if err := c2.Start(); err != nil {
		log.Fatalf("failed to start c2, reason: %s", err)
	}

	if err := c1.Wait(); err != nil {
		log.Fatalf("failed waiting for c1, reason: %s, stderr: %s", err, e1.String())
	}
	if err := writer.Close(); err != nil {
		log.Fatalf("failed closing writer, reason: %s, stderr: %s", err, e2.String())
	}

	if err := c2.Wait(); err != nil {
		log.Fatalf("failed waiting for c2, reason: %s", err)
	}
	if err := reader.Close(); err != nil {
		log.Fatalf("failed closing reader, reason: %s", err)
	}

	if _, err := io.Copy(os.Stdout, &buf); err != nil {
		log.Fatalf("failed to copy output, reason: %s", err)
	}
}
