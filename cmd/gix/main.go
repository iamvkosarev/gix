package main

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/iamvkosarev/git-plus/internal/command"
	"os"
	"os/exec"
	"strings"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 || strings.Contains(args[0], "help") {
		help()
		return
	}
	runAndPrint(prepareCommand(args))
}

func prepareCommand(args []string) []string {
	switch args[0] {
	case "back":
		return command.Back()
	case "commit":
		return command.Commit(args)
	}
	return args
}

func runAndPrint(args []string) {
	if len(args) == 0 {
		fmt.Println("Null args in input")
		return
	}
	cmd := exec.Command("git", args...)

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	changeColor := color.New(color.FgGreen).SprintFunc()
	fmt.Printf("Running: %s", changeColor(fmt.Sprintf("git %s\n", strings.Join(args, " "))))
	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
	}
}
