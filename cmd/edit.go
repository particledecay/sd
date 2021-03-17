package cmd

import (
	"errors"

	"github.com/particledecay/sd/conf"
	"github.com/particledecay/sd/internal"
	"github.com/spf13/cobra"
)

var editCmd = &cobra.Command{
	Use:   "edit",
	Short: "Create or edit a script",
	Long:  `Launch $EDITOR to edit a new or existing script`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 2 {
			return errors.New("you must supply a category and script name")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		scriptPath := conf.GetScriptPath(args[0], args[1], true)

		internal.EditScript(scriptPath)
	},
}
