package cmd

import (
	"fmt"
	"log"

	"github.com/spencerStephan/term-tutor/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type SetupConfig struct {
	Email    string
	Interval string
	Database string
}

var setupConfig SetupConfig

var setupCmd = &cobra.Command{
	Use:   "setup",
	Short: "Set up the term-tutor application",
	Long:  "Set up the term-tutor application by configuring the database paths, email address, and other settings",
	Run: func(cmd *cobra.Command, args []string) {
		if err := viper.BindPFlags(cmd.Flags()); err != nil {
			log.Fatalf("Error binding flags: %s", err)
		}

		// Load the setup-specific configuration values from flags
		setupConfig.Email = viper.GetString("email")
		setupConfig.Interval = viper.GetString("interval")
		setupConfig.Database = viper.GetString("database")

		fmt.Printf("Setup configuration: Email=%s, Interval=%s, Database=%s\n", setupConfig.Email, setupConfig.Interval, setupConfig.Database)
		config.OverrideDatabasePaths(setupConfig.Database)
	},
}

func init() {
	setupCmd.Flags().StringVar(&setupConfig.Email, "email", "", "Email address")
	setupCmd.Flags().StringVar(&setupConfig.Interval, "interval", "", "Interval duration. Valid values are 'daily', 'weekly', 'monthly'")
	setupCmd.Flags().StringVar(&setupConfig.Database, "database", "", "Database path. Example: /path/to/database.db")

	rootCmd.AddCommand(setupCmd)
}
