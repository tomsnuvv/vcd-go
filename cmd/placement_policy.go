package cmd

import (
	"github.com/spf13/cobra"
)

var placementPolicyCmd = &cobra.Command{
	Use:   "placement-policy",
	Short: "Command to manage placement policies",
}

func init() {

	placementPolicyCmd.AddCommand(placementPolicyListCmd)

}
