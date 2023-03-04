package cmd

/*
Copyright (c) 2023 John Dewey <john@dewey.ws>
*/

import (
	"os"

	// "github.com/retr0h/goltrane/pkg/migrate"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	directory string
	verbose   bool
	log       *logrus.Logger
)

// rootCmd represents the base command when called without any subcommands.
var rootCmd = &cobra.Command{
	Use:   "goltrane",
	Short: "A music theory library/CLI.",
	Long: `A music theory library/CLI written in [Go][].
`,

	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		log = &logrus.Logger{
			Out:       os.Stderr,
			Formatter: &logrus.JSONFormatter{PrettyPrint: true},
			Hooks:     make(logrus.LevelHooks),
			Level:     logrus.InfoLevel,
		}

		// m.Logger.WithError(err).Fatal("failed to get db")
		// log.WithFields(logrus.Fields{
		// 	"animal": "walrus",
		// }).Info("A walrus appears")
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
}
