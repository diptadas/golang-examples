package exec

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"testing"

	shell "github.com/codeskyblue/go-sh"
)

func TestGoExecPipe(t *testing.T) {
	c1 := exec.Command("cat", "abc")
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
		t.Fatalf("failed to start c1, reason: %s", err)
	}
	if err := c2.Start(); err != nil {
		t.Fatalf("failed to start c2, reason: %s", err)
	}

	if err := c1.Wait(); err != nil {
		t.Fatalf("failed waiting for c1, reason: %s, stderr: %s", err, e1.String())
	}
	if err := writer.Close(); err != nil {
		t.Fatalf("failed closing writer, reason: %s, stderr: %s", err, e2.String())
	}

	if err := c2.Wait(); err != nil {
		t.Fatalf("failed waiting for c2, reason: %s", err)
	}
	if err := reader.Close(); err != nil {
		t.Fatalf("failed closing reader, reason: %s", err)
	}

	if _, err := io.Copy(os.Stdout, &buf); err != nil {
		t.Fatalf("failed to copy output, reason: %s", err)
	}
}

func TestGoShPipe(t *testing.T) {
	sh := shell.NewSession()
	sh.PipeFail = true
	sh.PipeStdErrors = true

	sh.Command("cat", "abc")
	sh.Command("wc")
	out, err := sh.Output()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(string(out))
}
