package exek

import (
	"fmt"
	"strings"
	"os/exec"
)

func Call(cmdTmpl string, args ...interface{}) string {
	if len(args) > 0 {
		cmdTmpl = fmt.Sprintf(cmdTmpl, args...)
	}
	commands := parseCmds(cmdTmpl)

	result := ""
	for _, str := range commands {
		o := exeCmd(str)
		if len(o) > 0 {
			result += strings.TrimRight(o, "\n")
		}
	}

	return result
}

func parseCmds(cmdString string) []string {
	return strings.Split(cmdString, "&&")
}

func exeCmd(cmd string) string {
	var out []byte
	var err error
	out, err = exec.Command("bash", "-c", cmd).Output()
	if err != nil {
		panic(err)
	}

	return (string)(out)
}
