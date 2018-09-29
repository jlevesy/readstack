package cmd

import (
	"fmt"
	"os"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"google.golang.org/grpc"

	"github.com/jlevesy/readstack/api"
)

var (
	cfgFile string

	host string
)

var rootCmd = &cobra.Command{
	Use:   "readstackctl",
	Short: "Manage and share your reading todo list",
}

// Execute is the entrypoint of readstackctl
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVarP(&host, "host", "H", "localhost:8080", "Address of the Readstack server")
}

func initClient() (*grpc.ClientConn, api.ItemClient, error) {
	conn, err := grpc.Dial(host, grpc.WithInsecure())

	if err != nil {
		return nil, nil, err
	}

	return conn, api.NewItemClient(conn), nil
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		viper.AddConfigPath(home)
		viper.SetConfigName(".readstackctl")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
