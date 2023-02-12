package execAPI

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
)

var exeName = "/bin/bash"
var args = ""
var cmd *exec.Cmd
var stdin io.WriteCloser

func checkExists() bool {
	_, err := exec.LookPath(exeName)
	return err == nil

}

func StartCommandLine() {
	if !checkExists() {
		return
	}
	cmd := exec.Command(exeName, args)
	//cmd := exec.Command("ping", "127.0.0.1")
	//cmd := exec.Command("env", "")

	env := os.Environ()
	cmd.Env = env

	var err error
	stdin, err = cmd.StdinPipe()
	if err != nil {
		fmt.Println(err)
		cmd = nil
		return
	}
	defer stdin.Close() // the doc says subProcess.Wait will close it, but I'm not sure, so I kept this line

	stdout, _ := cmd.StdoutPipe()
	cmd.Stderr = os.Stderr

	err = cmd.Start()
	if err != nil {
		fmt.Println("An error occured: ", err)
		cmd = nil
		return
	}

	buf := bufio.NewReader(stdout)
	for {
		line, _, err := buf.ReadLine()
		if err != nil {
			fmt.Println("An error on read occured: ", err)
			cmd = nil
			return
		}

		fmt.Println(string(line))
	}

	// cmd.Wait()
}

func RunCommand(comand string) {
	if cmd == nil {
		return
	}

	_, err := io.WriteString(stdin, comand+"\n")
	if err != nil {
		fmt.Println(err)
		return
	}

}
