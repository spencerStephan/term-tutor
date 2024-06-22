package cmd

import (
	"fmt"
	"os"

	"github.com/spencerStephan/term-tutor/config"

	"github.com/spf13/cobra"
)

var (
	cfgFile string

	rootCmd = &cobra.Command{
		Use:   "tutor",
		Short: "Term-tutor is a terminal-based application that challenges you as you learn.",
		Long:  "Challenge yourself with flashcards, set reminders, and plan your lessons, all within the terminal. Read the full documentation at term-tutor.io",
		Run: func(cmd *cobra.Command, args []string) {
			dbPaths := config.RootCfg.Database
			fmt.Printf("Database paths: Windows=%s, Mac=%s, Linux=%s\n", dbPaths.Windows, dbPaths.Mac, dbPaths.Linux)
		},
	}
)

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fscanln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "Set a custom config file")
}

func initConfig() {
	config.InitRootConfig(cfgFile)
}
