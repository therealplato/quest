package main

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/therealplato/quest"
)

var root = &cobra.Command{
	Use:   "quest",
	Short: "your quest log",
	Long:  `how many quests could a quest log hold if the holder quested quests?`,
	Run:   oneQuest,
}

func main() {
	root.AddCommand(add)
	root.AddCommand(list)
	root.Execute()
}

func oneQuest(cmd *cobra.Command, params []string) {
	// $ quest 1 done
	// $ quest 1 active?
	// $ quest 1 lfg
	// $ quest 1 details
	// $ quest 999 details
	// $ quest 1 restart/reset
	cmdlen := len(params)
	if cmdlen == 0 {
		teeth := quest.Q{T: "brush teeth", Status: quest.PROGRESS}
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
}
