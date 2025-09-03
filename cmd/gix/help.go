package main

import (
	"fmt"
	"github.com/fatih/color"
)

func help() {
	cmdColor := color.New(color.FgHiMagenta)
	cmdInTextColor := color.New(color.FgCyan)
	argColor := color.New(color.FgYellow)
	printCmd := cmdColor.PrintfFunc()
	formatCmdInText := cmdInTextColor.SprintFunc()
	formatArg := argColor.SprintFunc()

	color.New(color.BgHiMagenta).Printf("gix - like X-Men but about Git")
	fmt.Printf("\n\n")
	color.Unset()
	printCmd("back")
	fmt.Printf(
		" - reset the commit and print the command to undo it. Based on %s command.\n\n",
		formatCmdInText("git reset --soft HEAD~1"),
	)
	printCmd("commit")
	fmt.Printf(" - classic %s command, but with addition flags to modify time:\n", formatCmdInText("git commit"))
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
	for arg, disc := range argsDiscMap {
		fmt.Printf("  -%s - %s\n", formatArg(arg), disc)
	}
}
