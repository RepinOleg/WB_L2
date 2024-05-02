package main

import (
	"bufio"
	"fmt"
	"github.com/mitchellh/go-ps"
	"os"
	"strconv"
	"strings"
	"syscall"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("$ ")
		str, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
		exec(str)
	}
}

func exec(command string) {
	command = strings.TrimRight(command, "\n")
	args := strings.Split(command, " ")

	if args[0] == "\\quit" || args[0] == "\\q" {
		os.Exit(0)
	}

	switch args[0] {
	case "cd":
		cd(args)
	case "pwd":
		pwd()
	case "echo":
		echo(args)
	case "kill":
		kill(args)
	case "ps":
		processStatus()
	}
}

func processStatus() {
	sliceProc, _ := ps.Processes()
	for _, proc := range sliceProc {
		fmt.Printf("Process name: %v process id: %v\n", proc.Executable(), proc.Pid())
	}
}

func kill(args []string) {
	pid, err := strconv.Atoi(args[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, args[1], " argument must be job IDs")
		return
	}
	process, err := os.FindProcess(pid)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Process not found:", err)
		return
	}

	err = process.Signal(syscall.SIGKILL)
	if err != nil {
		fmt.Println("Failed to kill process:", err)
	}
}

func cd(args []string) {
	if len(args) < 2 {
		return
	}
	err := os.Chdir(args[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
}

func pwd() {
	currentPwd, err := os.Getwd()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	fmt.Println(currentPwd)
}

func echo(args []string) {
	if len(args) < 2 {
		fmt.Println()
		return
	}
	fmt.Println(args[1])
}
