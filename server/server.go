package server

import (
	"bytes"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
)

const (
	PipeName = `\\.\pipe\pipe-exec`
)

type ExecImpl struct {
}

func (c *ExecImpl) Exec(stream Exec_ExecServer) error {
	var (
		blob   []byte
		args   []string
		stdout bytes.Buffer
		stderr bytes.Buffer
	)
	for {
		c, err := stream.Recv()
		if err == io.EOF {
			log.Printf("Transfer of %d bytes successful", len(blob))
			break
		}
		if err != nil {
			log.Printf("err receiving %#v", err)
		}
		blob = append(blob, c.Chunk...)
		args = c.Args
	}

	tmpfile, err := ioutil.TempFile("", "sudo.*.exe")
	if err != nil {
		log.Fatalf("err creating tempfile %#v", err)
		return err
	}
	defer os.Remove(tmpfile.Name())

	if _, err := tmpfile.Write(blob); err != nil {
		tmpfile.Close()
		log.Fatalf("err writing %#v", err)
		return err
	}
	if err := tmpfile.Close(); err != nil {
		log.Fatalf("err closing tmpfile %#v", err)
		return err
	}

	cmd := exec.Command(tmpfile.Name(), args...)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	if err := cmd.Run(); err != nil {
		log.Fatalf("err running %#v", err)
	}

	if err := stream.SendAndClose(&Response{
		Stdout: stdout.Bytes(),
		Stderr: stderr.Bytes(),
	}); err != nil {
		log.Printf("err closing stream %#v", err)
		return err
	}
	return nil
}
