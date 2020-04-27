package deamon

import (
	"fmt"
	"os"
	"os/exec"
)

func init() {
	if os.Args[len(os.Args)-1:][0] == "-d=false" {
		return
	}
	os.Args = append(os.Args, "-d=true")
	args := os.Args[1:]
	i := 0
	for ; i < len(args); i++ {
		if args[i] == "-d=true" {
			args[i] = "-d=false"
			break
		}
	}
	cmd := exec.Command(os.Args[0], args...)
	cmd.Start()
	fmt.Println("[PID]", cmd.Process.Pid)
	os.Exit(0)
}
