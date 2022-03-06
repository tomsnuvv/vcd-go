package cmd

import (
	"github.com/spf13/cobra"
)

var vmCmd = &cobra.Command{
	Use:   "vm",
	Short: "Print the version number of Hugo",
}

func init() {

	vmCmd.AddCommand(vmListCmd)

}
