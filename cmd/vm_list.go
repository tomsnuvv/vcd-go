package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/vmware/go-vcloud-director/v2/types/v56"
)

func init() {

	vmCmd.Flags().String("name", "", "Name of the vm")
	vmCmd.Flags().Int("limit", 20, "Limit the amount of records")

}

var vmListCmd = &cobra.Command{
	Use:   "list",
	Short: "List the vm's",
	Run:   vmList,
}

func vmList(cmd *cobra.Command, args []string) {
	client := GetClient()

	org, err := client.GetOrgByName(viper.GetString("organisation"))
	if err != nil {
		log.Fatal(fmt.Println("Failed to find org: ", org))
	}

	vdc, err := org.GetVDCByName(viper.GetString("vdc"), true)
	if err != nil {
		log.Fatal(fmt.Println("Failed to find org: ", org))
	}

	results, err := vdc.QueryVmList(types.VmQueryFilterOnlyDeployed)
	if err != nil {
		log.Fatal(fmt.Println("Failed to find org: ", org))
	}

	fmt.Println(results)
}
