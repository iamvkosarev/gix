package command

import (
	"fmt"
	"os/exec"
)

func Save(args []string) ([]string, error) {
	cmd := exec.Command(
		"git", []string{
			"add",
			".",
		}...,
	)
	err := cmd.Run()
	if err != nil {
		return nil, fmt.Errorf("failed to run \"git add\" command: %v", err)
	}
	for i, arg := range args {
		if arg == "save" {
			args[i] = "commit"
		}
	}

	args, err = Commit(args)
	if err != nil {
		return nil, fmt.Errorf("failed to commit changes: %v", err)
	}
	return args, nil
}
