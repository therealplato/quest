package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

var list = &cobra.Command{
	Use:  "list",
	Long: `quest list`,
	Run:  listHandler,
}

func listHandler(cmd *cobra.Command, params []string) {
	fmt.Println(`Quests:
	1 visit a friend
	2 brush teeth
	find a nearby hero
	find a far-away hero
	find a zero risk quest
	passive quest achievements - visit a city
	Complete a multiparty quest
		- LFG flag (looking for group irl)
	Fill a role on your Organic Farm
		- Post Quest "Work on a farm"
		- Accomplish above quest
	Work On an Organic Farm
	define: democracy, autocracy, voluntarism
	find information you seek
	share information with someone who doesn't know they seek it
	achieve enlightenment
	communicate a scientific truth
`)
}
