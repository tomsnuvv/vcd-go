package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

func init() {

	placementPolicyCmd.Flags().Int("limit", 20, "Limit the amount of records")

}

var placementPolicyListCmd = &cobra.Command{
	Use:   "list",
	Short: "List the placement policies",
	Run:   placementPolicyList,
}

type PlacementPolicyStruct struct {
	Name        string
	GuestOS     string
	Cpus        int
	MemoryMB    int
	Status      string
	NetworkName string
	IpAddress   string
}

func placementPolicyList(cmd *cobra.Command, args []string) {
	client := GetClient()

	org, _ := GetOrgAndVDC(client)

	results, err := org.GetAllVdcComputePolicies(nil)
	if err != nil {
		log.Fatal(fmt.Println("Failed to find org: ", org))
	}

	// placementPolicies := []PlacementPolicyStruct{}

	// copier.CopyWithOption(&vms, &results, copier.Option{IgnoreEmpty: true, DeepCopy: true})

	Output(results)
}
