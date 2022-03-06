package cmd

import (
	"fmt"
	"log"

	"github.com/jinzhu/copier"
	"github.com/spf13/cobra"
	"github.com/vmware/go-vcloud-director/v2/types/v56"
)

func init() {

	vmListCmd.Flags().Int("limit", 20, "Limit the amount of records")

}

var vmListCmd = &cobra.Command{
	Use:   "list",
	Short: "List the vm's",
	Run:   vmList,
}

type VMStruct struct {
	Name        string
	GuestOS     string
	Cpus        int
	MemoryMB    int
	Status      string
	NetworkName string
	IpAddress   string
}

func vmList(cmd *cobra.Command, args []string) {
	client := GetClient()

	org, vdc := GetOrgAndVDC(client)

	results, err := vdc.QueryVmList(types.VmQueryFilterOnlyDeployed)
	if err != nil {
		log.Fatal(fmt.Println("Failed to find org: ", org))
	}

	vms := []VMStruct{}

	copier.CopyWithOption(&vms, &results, copier.Option{IgnoreEmpty: true, DeepCopy: true})

	Output(vms)
}
