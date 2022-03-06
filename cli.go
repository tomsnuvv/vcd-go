package main

import (
	"fmt"
	"net/url"
	"os"

	"github.com/vmware/go-vcloud-director/v2/govcd"
)

type Config struct {
	User     string
	Token    string
	Org      string
	Href     string
	VDC      string
	Insecure bool
}

func (c *Config) Client() (*govcd.VCDClient, error) {
	u, err := url.ParseRequestURI(c.Href)
	if err != nil {
		return nil, fmt.Errorf("unable to pass url: %s", err)
	}

	vcdclient := govcd.NewVCDClient(*u, c.Insecure)
	err = vcdclient.SetToken(c.Org, govcd.AuthorizationHeader, c.Token)
	if err != nil {
		return nil, fmt.Errorf("unable to authenticate: %s", err)
	}
	return vcdclient, nil
}

func main() {
	if len(os.Args) < 6 {
		fmt.Println("Syntax: example user token org VCD_IP VDC ")
		os.Exit(1)
	}
	config := Config{
		User:     os.Args[1],
		Token:    os.Args[2],
		Org:      os.Args[3],
		Href:     fmt.Sprintf("https://%s/api", os.Args[4]),
		VDC:      os.Args[5],
		Insecure: true,
	}

	client, err := config.Client() // We now have a client
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	org, err := client.GetOrgByName(config.Org)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	vdc, err := org.GetVDCByName(config.VDC, false)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("Org URL: %s\n", org.Org.HREF)
	fmt.Printf("VDC URL: %s\n", vdc.Vdc.HREF)

}
