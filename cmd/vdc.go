package cmd

import (
	"github.com/spf13/cobra"
)

type Fields struct {
	Name string
}

var vdcCmd = &cobra.Command{
	Use:   "vdc",
	Short: "Command to manage vdc's (Virtual Data Centers)",
}

func init() {

	vdcCmd.AddCommand(vdcListCmd)

}
