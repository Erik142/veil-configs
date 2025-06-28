package app

import (
	"fmt"
	"os"

	"github.com/Erik142/veil-configs/internal/server"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "server",
	Short: "Nebula Config Server",
	Long: `The Nebula Config Server provides Nebula configuration files
via gRPC to clients.`,
	Run: func(cmd *cobra.Command, args []string) {
		addr := viper.GetString("server.address")
		if addr == "" {
			addr = ":50051"
		}
		logrus.Infof("Starting server on %s", addr)
		if err := server.StartServer(addr); err != nil {
			logrus.Fatalf("Failed to start server: %v", err)
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here, will
	// be available to all subcommands and application globally.
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.server.yaml)")
	rootCmd.PersistentFlags().String("address", ":50051", "The address to listen on")
	viper.BindPFlag("server.address", rootCmd.PersistentFlags().Lookup("address"))

	// Cobra also supports local flags, which will only run when this command
	// is called directly, e.g.:
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".server" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".server")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
