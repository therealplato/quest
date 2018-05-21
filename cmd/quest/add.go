package main

import (
	"fmt"

	"github.com/spf13/cobra"
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
