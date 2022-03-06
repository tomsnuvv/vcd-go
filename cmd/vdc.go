package cmd

import (
	"github.com/spf13/cobra"
)

var vdcCmd = &cobra.Command{
	Use:   "vdc",
	Short: "Command to manage vdc's (Virtual Data Centers)",
}

func init() {

	vdcCmd.AddCommand(vdcListCmd)

}
