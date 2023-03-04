/*
Copyright (c) 2023 John Dewey <john@dewey.ws>
*/

package cmd

import (
	"os"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/retr0h/goltrane/pkg/scale"
	"github.com/spf13/cobra"
)

// scaleListCmd represents the scaleList command
var scaleListCmd = &cobra.Command{
	Use:   "list",
	Short: "list known Scales",
	Long: `list the names of all the known scale-building rules
`,
	Run: func(cmd *cobra.Command, args []string) {
		s := scale.NewScale()
		scales := s.GetScalesByName()

		t := table.NewWriter()
		t.SetOutputMirror(os.Stdout)
		t.SetTitle("List Scales")
		t.AppendHeader(table.Row{"Name"})
		for _, scale := range scales {
			t.AppendRows([]table.Row{
				{scale},
			})
		}
		t.AppendSeparator()
		style := table.StyleColoredDark
		t.SetStyle(style)
		t.Render()
	},
}

func init() {
	scaleCmd.AddCommand(scaleListCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// scaleListCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// scaleListCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
