package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

func init() {

	vdcListCmd.Flags().Int("limit", 20, "Limit the amount of records")

}

var vdcListCmd = &cobra.Command{
	Use:   "list",
	Short: "List the vdc's",
	Run:   vdcList,
}

func vdcList(cmd *cobra.Command, args []string) {
	client := GetClient()

	org := GetOrg(client)

	results, err := org.QueryOrgVdcList()
	if err != nil {
		log.Fatal(fmt.Println("Failed to find vdc's"))
	}

	Output(results)

}
