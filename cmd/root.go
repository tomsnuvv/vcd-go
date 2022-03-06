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
	cfgFile string
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
	cfgFile := viper.GetString("config-file")

	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".cobra" (without extension).
		viper.AddConfigPath(home + "/.config/vcd")
		viper.SetConfigType("yaml")
		viper.SetConfigName("config.yaml")
	}

	viper.AutomaticEnv()

	err := viper.ReadInConfig()

	if err != nil {
		log.Fatal(fmt.Println("Error reading config file: ", viper.ConfigFileUsed()))
	}

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file: ", viper.ConfigFileUsed())
	}
}

func Output(results interface{}) {

	output, err := json.MarshalIndent(results, "", "\t")
	if err != nil {
		log.Fatal("Failed to ecnode results to json")
	}

	fmt.Println(string(output))

}
