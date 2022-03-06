package cmd

import (
	"github.com/spf13/cobra"
)

var vmCmd = &cobra.Command{
	Use:   "vm",
	Short: "Command to manage vm's",
}

func init() {

	vmCmd.AddCommand(vmListCmd)

}
