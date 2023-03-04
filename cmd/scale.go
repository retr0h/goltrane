/*
Copyright (c) 2023 John Dewey <john@dewey.ws>
*/

package cmd

import (
	"github.com/spf13/cobra"
)

// scaleCmd represents the scale command
var scaleCmd = &cobra.Command{
	Use:   "scale",
	Short: "scale sub-command",
	Long: `The scale sub-command parent
`,
	Run: func(cmd *cobra.Command, args []string) {
	},
}

// go run main.go scale list
// go run main.go scale note C --mode aug
// music-theory scale "C aug"

func init() {
	rootCmd.AddCommand(scaleCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// scaleCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// scaleCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
