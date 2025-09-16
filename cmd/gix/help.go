package main

import (
	"fmt"
	"github.com/fatih/color"
)

func help() {
	cmdColor := color.New(color.FgHiMagenta)
	cmdInTextColor := color.New(color.FgCyan)
	argColor := color.New(color.FgYellow)
	formatCmd := cmdColor.SprintFunc()
	formatCmdInText := cmdInTextColor.SprintFunc()
	formatArg := argColor.SprintFunc()

	color.New(color.BgHiMagenta).Printf("gix - a bit powered version of Git")
	fmt.Printf("\n\n")
	color.Unset()
	printBackCommandHelp(formatCmd, formatCmdInText)

	printCommitCommandHelp(formatCmd, formatCmdInText)
	printCommitCommandArgsHelp(formatArg)

	printSaveCommandHelp(formatCmd, formatCmdInText)
	printCommitCommandArgsHelp(formatArg)
}

func printBackCommandHelp(formatCmd func(a ...interface{}) string, formatCmdInText func(a ...interface{}) string) {
	fmt.Printf(
		"%s - reset last commit and print the command to undo it.\nBased on %s command.\n\n", formatCmd("back"),
		formatCmdInText("git reset --soft HEAD~1"),
	)
}

func printCommitCommandHelp(formatCmd func(a ...interface{}) string, formatCmdInText func(a ...interface{}) string) {
	fmt.Printf(
		"%s - classic %s command, but with addition flags to modify time:\n", formatCmd("commit"),
		formatCmdInText("git commit"),
	)
}

func printCommitCommandArgsHelp(formatArg func(a ...interface{}) string) {
	argsDiscMap := map[string]string{
		"pD": "plus days",
		"pH": "plus hours",
		"pM": "plus minutes",
		"pS": "plus seconds",
		"mD": "minus days",
		"mH": "minus hours",
		"mM": "minus minutes",
		"mS": "minus seconds",
	}
	order := []string{
		"pD",
		"pH",
		"pM",
		"pS",
		"mD",
		"mH",
		"mM",
		"mS",
	}
	for _, arg := range order {
		fmt.Printf("  -%s - %s\n", formatArg(arg), argsDiscMap[arg])
	}
	fmt.Printf("\n")
}

func printSaveCommandHelp(formatCmd func(a ...interface{}) string, formatCmdInText func(a ...interface{}) string) {
	fmt.Printf(
		"%s - same %s command, but with additional saving step of all changes.\nBased on %s command.\n",
		formatCmd("save"),
		formatCmd("commit"),
		formatCmdInText("git add ."),
	)
}
