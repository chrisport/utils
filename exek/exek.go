package exek

import (
	"fmt"
	"strings"
	"os/exec"
)

// CallToLines executes the specified command (formatted with args) and returns the result Split by new lines
func CallToLines(cmdTmpl string, args ...interface{}) []string {
	o := Call(cmdTmpl, args...)
	return strings.Split(o, "\n")
}

// Call executes the specified command (formatted with args) and returns the result as string
func Call(cmdTmpl string, args ...interface{}) string {
	if len(args) > 0 {
		cmdTmpl = fmt.Sprintf(cmdTmpl, args...)
	}
	commands := parseCmds(cmdTmpl)

	result := execCommands(commands)

	return result
}

func execCommands(commands []string) string {
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
		fmt.Println(err)
	}

	return (string)(out)
}
