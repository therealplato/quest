package main

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/therealplato/quest"
)

var add = &cobra.Command{
	Use:  "add",
	Long: `quest add buy hello kitty merch`,
	Run:  addHandler,
}

func addHandler(cmd *cobra.Command, params []string) {
	// it started with quest add
	if len(params) < 1 {
		fmt.Println("add what?")
		return
	}
	fmt.Printf("adding quest %s\n", params[:])
}

func all(cmd *cobra.Command, params []string) {
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
	default:
		fmt.Printf("idk %s\n", command)
	}

}
