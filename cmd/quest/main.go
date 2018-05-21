package main

import "github.com/spf13/cobra"

var root = &cobra.Command{
	Use:   "quest",
	Short: "quest logs quests",
	Long:  `how many quests could a quest log hold if the holder quested quests?`,
	Run:   all,
}

func main() {
	root.AddCommand(add)
	root.Execute()
}
