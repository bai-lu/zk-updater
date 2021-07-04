package node

import (
	"bytes"
	"os"
	"os/exec"
	"testing"
)

func TestCmd(t *testing.T) {
	params := []string{
		"t.sh",
	}
	cmd := exec.Command("/bin/bash", params...)
	cmd.Env = os.Environ()
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		t.Errorf("执行命令失败 %v", err)
	}
	// err = cmd.Wait()
	// if err != nil {
	// 	t.Errorf("执行命令失败 %v", err)
	// }

	t.Log(cmd.Args)
	t.Log(stdout.String())
	t.Log(stderr.String())

}
