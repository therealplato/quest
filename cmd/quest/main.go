package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/therealplato/quest"
)

func main() {
	parseCommand(os.Args...)
}

func parseCommand(params ...string) {
	cmdlen := len(params)
	if cmdlen == 0 {
		teeth := quest.Q{T: "brush teeth"}
		teeth.Achieve()
		fmt.Println(teeth)
		return
	}
	command := params[1]
	i, err := strconv.Atoi(command)
	if err != nil {
		// do things for quest i
		_ = i
		fmt.Printf("quest %d: \n", i)
		return
	}
	switch command {
	case "add":
		// add quest
		if len(params) == 2 {
			fmt.Println("add what?")
			return
		}
		fmt.Printf("adding quest %s\n", params[2:])
	default:
		fmt.Printf("idk %s\n", command)
	}

}
