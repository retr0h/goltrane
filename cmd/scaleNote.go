/*
Copyright (c) 2023 John Dewey <john@dewey.ws>
*/

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// scaleNoteCmd represents the scaleNote command
var scaleNoteCmd = &cobra.Command{
	Use:   "note",
	Short: "build a Scale",
	Long: `calculate the note pitch classes for a specified Scale
`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("scaleNote called")
	},
}

func init() {
	scaleCmd.AddCommand(scaleNoteCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// scaleNoteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// scaleNoteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
