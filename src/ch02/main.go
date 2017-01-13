package main

import (
	"ch02/classpath"
	"fmt"
	"strings"
)

func main() {
	cmd := parseCmd()
	if cmd.versionFlag {
		fmt.Println("version 0.0.1")
	} else if cmd.versionFlag || cmd.class == "" {
		printUsage()
	} else {
		startJVM(cmd)
	}
}

func startJVM(cmd *Cmd) {
	fmt.Printf("classpath: %s class: %s args: %v\n", cmd.cpOption, cmd.class, cmd.args)
	cp := classpath.Parse(cmd.XjreOption, cmd.cpOption)
	className := strings.Replace(cmd.class, ".", "/", -1)
	data, entry, err := cp.ReadClass(className)
	if err != nil {
		fmt.Printf("could not find or load mian class %s\n", cmd.class)
		return
	}
	fmt.Printf("entry: %s\n", entry.String())
	fmt.Printf("class data: %v\n", data)
}
