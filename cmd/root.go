package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	rootCmd = &cobra.Command{
		Use:   "vcd",
		Short: "A cli tool for VCD VMWare",
	}
)

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(InitConfig)

	viper.SetEnvPrefix("vcd")

	homeDir, err := os.UserHomeDir()

	if err != nil {
		log.Fatal("Error reading homedir")
	}

	rootCmd.PersistentFlags().String("config-file", homeDir+"/.config/vcd/config.yaml", "config file")
	viper.BindPFlag("config-file", rootCmd.PersistentFlags().Lookup("config-file"))
	viper.BindEnv("config-file")

	rootCmd.AddCommand(vmCmd)
	rootCmd.AddCommand(vdcCmd)

}

func InitConfig() {

	viper.SetConfigFile(viper.GetString("config-file"))

	viper.AutomaticEnv()

	err := viper.ReadInConfig()

	if err != nil {
		log.Fatal(fmt.Println("Error reading config file: ", viper.ConfigFileUsed()))
	}

}

func Output(results interface{}) {

	output, err := json.MarshalIndent(results, "", "\t")
	if err != nil {
		log.Fatal("Failed to ecnode results to json")
	}

	fmt.Println(string(output))

}
