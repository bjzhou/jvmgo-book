package main

import (
	"flag"
	"fmt"
)

func main() {
	cmd := parseCmd()
	if cmd.versionFlag {
		fmt.Println("version 0.0.1")
	} else if cmd.versionFlag || cmd.class == "" {
		flag.PrintDefaults()
	} else {
		startJVM(cmd)
	}
}

func startJVM(cmd *Cmd) {
	fmt.Printf("classpath: %s class: %s args: %v\n", cmd.cpOption, cmd.class, cmd.args)
}
