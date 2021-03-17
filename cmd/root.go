package cmd

import (
	"errors"

	"github.com/particledecay/sd/conf"
	"github.com/particledecay/sd/internal"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var verbose bool

var rootCmd = &cobra.Command{
	Use:   "sd",
	Short: "sd executes commands in your scripts directory",
	Long: `sd allows you to add, remove, and execute scripts present in
			an organized directory on your computer`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		if verbose {
			zerolog.SetGlobalLevel(zerolog.DebugLevel)
		}
	},
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("an action is required")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 2 {
			log.Fatal().Msg("a category and script name is required")
		}

		scriptPath := conf.GetScriptPath(args[0], args[1], false)
		internal.RunScript(scriptPath, args[2:])
	},
}

func init() {
	// flags
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "display debug log messages")
}

// Execute combines all of the available command functions
func Execute() {
	rootCmd.AddCommand(editCmd)

	if err := rootCmd.Execute(); err != nil {
		log.Fatal().Msgf("error during execution: %v", err)
	}
}
