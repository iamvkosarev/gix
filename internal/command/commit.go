package command

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func Commit(args []string) ([]string, error) {
	commitTime := time.Now()
	timeSecondsMap := map[string]int{
		"mS": -1,
		"mM": -60,
		"mH": -60 * 60,
		"mD": -60 * 60 * 24,
		"pS": 1,
		"pM": 60,
		"pH": 60 * 60,
		"pD": 60 * 60 * 24,
	}
	normalArgs := make([]string, 0)
	for i := 0; i < len(args); i++ {
		arg := args[i]
		normalArg := true
		for key, seconds := range timeSecondsMap {
			if arg == "-"+key {
				if i < len(args)-1 {
					v, err := strconv.Atoi(args[i+1])
					if err != nil {
						return nil, fmt.Errorf("invalid value for time key \"%s\", expected int", key)
					}
					commitTime = commitTime.Add(time.Duration(v*seconds) * time.Second)
					i++
					normalArg = false
					break
				} else {

					return nil, fmt.Errorf("null input after time key \"%s\"", key)
				}
			}
		}
		if normalArg {
			if !strings.HasPrefix(arg, "-") && i != 0 {
				normalArgs = append(normalArgs, fmt.Sprintf("%s", arg))
			} else {
				normalArgs = append(normalArgs, arg)
			}
		}
	}

	const gitCommiterDate = "GIT_COMMITTER_DATE"
	const gitAuthorDate = "GIT_AUTHOR_DATE"
	err := os.Setenv(gitCommiterDate, commitTime.Format(time.RFC3339))
	if err != nil {
		return nil, fmt.Errorf("error setting %s: %s\n", gitCommiterDate, err)
	}
	err = os.Setenv(gitAuthorDate, commitTime.Format(time.RFC3339))
	if err != nil {
		return nil, fmt.Errorf("error setting %s: %s\n", gitAuthorDate, err)
	}
	fmt.Printf("Commit at: %s\n", commitTime.Format("2006-01-02 15:04:05"))
	return normalArgs, nil
}
