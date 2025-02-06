package main

import (
	"bytes"
	_ "embed"
	"fmt"
	"io"
	"math/rand"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

//go:embed magic.txt
var magicString string

func runCommand(command string, args ...string) (string, string, error) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer

	cmd := exec.Command(command, args...)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()

	return stdout.String(), stderr.String(), err
}

func main() {
	zname := fmt.Sprintf("__%d%d%d%d__.zip", rand.Intn(10), rand.Intn(10), rand.Intn(10), rand.Intn(10))

	exepath, err := os.Executable()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	exe, err := os.Open(exepath)
	defer exe.Close()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	content, err := io.ReadAll(exe)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	pos := strings.Index(string(content), magicString)
	if pos == -1 {
		panic("first magic string not found")
	}

	content = content[pos+len(magicString):]

	pos = strings.Index(string(content), magicString)
	if pos == -1 {
		fmt.Println("zip not found")
		os.Exit(1)
	}

	zcontent := content[pos+len(magicString):]

	output, err := os.OpenFile(zname, os.O_CREATE|os.O_WRONLY, os.ModePerm)
	defer output.Close()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	if _, err = output.Write(zcontent); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	if runtime.GOOS == "windows" {
		if _, se, err := runCommand("tar", "-xf", zname); err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		} else if se != "" {
			fmt.Println(se)
			os.Exit(1)
		}
	} else {
		if _, se, err := runCommand("unzip", zname); err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		} else if se != "" {
			fmt.Println(se)
			os.Exit(1)
		}
	}

	runCommand("rm", "-rf", zname, "__MACOSX", "rsc/.DS_Store")
}
