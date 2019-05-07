
package cmd

import (
	"fmt"
	"os"
	"strings"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/yulibaozi/kubectl-switch/server"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "kubectl-switch",
	Short: "switch is kubectl plugin that switch freely between multiple clusters",
	Long:  `Kubernetes multi-cluster command line management tool.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func cluster(err error) bool {
	return strings.Contains(err.Error(), "unknown command")
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		if !cluster(err) {
			if !rootCmd.SilenceErrors {
				rootCmd.Println("Error:", err.Error())
				rootCmd.Printf("Run '%v --help' for usage.\n", rootCmd.CommandPath())
			}
			os.Exit(1)
		}
		clusterName := os.Args[1]
		clusterNames := server.GetClusterNames()
		if !clusterNames[clusterName] {
			rootCmd.Println("Error:", err.Error())
			rootCmd.Printf("Run '%v --help' for usage.\n", rootCmd.CommandPath())
			os.Exit(1)
		}
		cmdShi := &server.CmdShim{
			SubCmd: clusterName,
			Args:   os.Args[2:],
			Run:    server.Exec,
		}
		cmdShi.Run(cmdShi)
		return
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.kubectl-switch.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".kubectl-switch" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".kubectl-switch")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	// if err := viper.ReadInConfig(); err == nil {
	// 	fmt.Println("Using config file:", viper.ConfigFileUsed())
	// }
}
