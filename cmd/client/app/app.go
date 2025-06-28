package app

import (
	"fmt"
	"os"

	"github.com/Erik142/veil-configs/internal/client"
	pb "github.com/Erik142/veil-configs/pkg/proto"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "client",
	Short: "Nebula Config Client",
	Long: `The Nebula Config Client fetches Nebula configuration files
from a gRPC server.`,
	Run: func(cmd *cobra.Command, args []string) {
		serverAddr := viper.GetString("client.server_address")
		clientID := viper.GetString("client.client_id")
		outputFile := viper.GetString("client.output_file")

		if clientID == "" {
			logrus.Fatal("Client ID is required. Use --client-id or set client.client_id in config.")
		}

		if outputFile == "" {
			outputFile = fmt.Sprintf("nebula_config_%s.yaml", clientID)
		}

		logrus.Infof("Fetching config for client %s from %s", clientID, serverAddr)

		conn, err := grpc.Dial(serverAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			logrus.Fatalf("did not connect: %v", err)
		}
		defer conn.Close()

		c := pb.NewNebulaConfigServiceClient(conn)

		if err := client.Run(c, clientID, outputFile); err != nil {
			logrus.Fatalf("Failed to fetch and save config: %v", err)
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
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.client.yaml)")
	rootCmd.PersistentFlags().String("server-address", "localhost:50051", "The address of the Nebula config server")
	viper.BindPFlag("client.server_address", rootCmd.PersistentFlags().Lookup("server-address"))

	rootCmd.PersistentFlags().String("client-id", "", "The client ID to request configuration for")
	viper.BindPFlag("client.client_id", rootCmd.PersistentFlags().Lookup("client-id"))

	rootCmd.PersistentFlags().String("output-file", "", "The output file name for the Nebula config")
	viper.BindPFlag("client.output_file", rootCmd.PersistentFlags().Lookup("output-file"))

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

		// Search config in home directory with name ".client" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".client")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}