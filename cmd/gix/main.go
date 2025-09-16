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
	preparedArgs, err := prepareCommand(args)
	if err != nil {
		fmt.Printf("%s %v\n", color.New(color.FgRed).Sprintf("Error:"), err)
		return
	}
	runAndPrint(preparedArgs)
}

func prepareCommand(args []string) ([]string, error) {
	switch args[0] {
	case "back":
		return command.Back()
	case "commit":
		return command.Commit(args)
	}
	return args, nil
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

	printRunningCommand(args)
	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
	}
}

func printRunningCommand(args []string) {
	changeColor := color.New(color.FgGreen).SprintFunc()
	for i := 1; i < len(args); i++ {
		arg := args[i]
		if arg[:1] == "-" && arg[:2] != "--" && i+1 < len(args) {
			args[i+1] = fmt.Sprintf("\"%s\"", args[i+1])
			i++
		}
	}
	fmt.Printf("Running: %s", changeColor(fmt.Sprintf("git %s\n", strings.Join(args, " "))))
}
