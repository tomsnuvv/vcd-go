package cmd

import (
	"fmt"
	"log"
	"net/url"
	"os"

	"github.com/spf13/viper"
	"github.com/vmware/go-vcloud-director/v2/govcd"
)

type Config struct {
	User     string
	Token    string
	Org      string
	Url      string
	VDC      string
	Insecure bool
}

func (c *Config) Client() (*govcd.VCDClient, error) {
	u, err := url.ParseRequestURI(c.Url)
	if err != nil {
		return nil, fmt.Errorf("unable to pass url: %s", err)
	}

	vcdclient := govcd.NewVCDClient(*u, c.Insecure)
	err = vcdclient.SetToken(c.Org, govcd.ApiTokenHeader, c.Token)
	if err != nil {
		return nil, fmt.Errorf("unable to authenticate: %s", err)
	}
	return vcdclient, nil
}

func GetClient() *govcd.VCDClient {

	config := Config{
		User:     viper.GetString("user"),
		Token:    viper.GetString("token"),
		Org:      viper.GetString("organisation"),
		Url:      viper.GetString("url"),
		VDC:      viper.GetString("vdc"),
		Insecure: true,
	}

	client, err := config.Client() // We now have a client
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return client
}

func GetOrg(client *govcd.VCDClient) *govcd.Org {
	org, err := client.GetOrgByName(viper.GetString("organisation"))
	if err != nil {
		log.Fatal(fmt.Println("Failed to find org: ", org))
	}

	return org
}

func GetOrgAndVDC(client *govcd.VCDClient) (*govcd.Org, *govcd.Vdc) {
	org, err := client.GetOrgByName(viper.GetString("organisation"))
	if err != nil {
		log.Fatal(fmt.Println("Failed to find org: ", org))
	}

	vdc, err := org.GetVDCByName(viper.GetString("vdc"), true)
	if err != nil {
		log.Fatal(fmt.Println("Failed to find org: ", org))
	}

	return org, vdc
}
