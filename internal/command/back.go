package command

import (
	"bytes"
	"fmt"
	"github.com/fatih/color"
	"os"
	"os/exec"
	"strings"
)

func Back() ([]string, error) {
	cmd := exec.Command(
		"git", []string{
			"log",
			"--format=%s@@%b@@%cD",
			"-n",
			"1",
			"-i",
		}...,
	)

	s := ""
	buf := bytes.NewBufferString(s)
	cmd.Stdin = os.Stdin
	cmd.Stdout = buf
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		return nil, err
	}
	outRawArgs := strings.Split(buf.String(), "@@")
	if len(outRawArgs) != 3 {
		return nil, fmt.Errorf("usage: not correct out args: %s", s)
	}
	// TODO: Add time return
	outArgs := []string{
		"-m",
		outRawArgs[0],
	}
	if outRawArgs[1] != "" {
		outArgs = append(
			outArgs,
			"-m",
			outRawArgs[1],
		)
	}
	changeColor := color.New(color.FgHiMagenta).SprintFunc()
	fmt.Printf("Revert: %s\n", changeColor(fmt.Sprintf("gix Commit %s", strings.Join(outArgs, " "))))
	return []string{
		"reset",
		"--soft",
		"HEAD~1",
	}, nil
}
