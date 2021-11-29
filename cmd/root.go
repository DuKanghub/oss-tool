/*
Copyright © 2021 DuKang <dukang@dukanghub.com>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	"time"
)

var (
	cfgFile string
	platForm string
	endPoint string
	accessKeyId string
	accessKeySecret string
	bucketName string
	objectName string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "oss-tool",
	Short: "自用OSS工具",
	Long: `自用OSS工具, 
支持的OSS平台有:
- 阿里云
`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.oss-tool.yaml)")
	rootCmd.PersistentFlags().StringVarP(&platForm, "platform", "p", "ali", "OSS平台, 可选值有：ali")
	rootCmd.PersistentFlags().StringVarP(&endPoint, "endpoint", "e", "", "OSS地域，必传")
	rootCmd.PersistentFlags().StringVarP(&accessKeyId, "access_key_id", "k", "", "accessKeyId，必传")
	rootCmd.PersistentFlags().StringVarP(&accessKeySecret, "access_key_secret", "s", "", "accessKeySecret，必传")
	rootCmd.PersistentFlags().StringVarP(&bucketName, "bucket_name", "b", "", "用来存放上传文件的桶名字，必传")
	rootCmd.PersistentFlags().StringVarP(&objectName, "object_name", "o", fmt.Sprintf("%d%d%d", time.Now().Year(), time.Now().Month(), time.Now().Day()), "OSS文件夹名字，必传")
	// flag绑定
	cobra.CheckErr(viper.BindPFlag("platform", rootCmd.PersistentFlags().Lookup("platform")))
	cobra.CheckErr(viper.BindPFlag("endpoint", rootCmd.PersistentFlags().Lookup("endpoint")))
	cobra.CheckErr(viper.BindPFlag("access_key_id", rootCmd.PersistentFlags().Lookup("access_key_id")))
	cobra.CheckErr(viper.BindPFlag("access_key_secret", rootCmd.PersistentFlags().Lookup("access_key_secret")))
	cobra.CheckErr(viper.BindPFlag("bucket_name", rootCmd.PersistentFlags().Lookup("bucket_name")))
	cobra.CheckErr(viper.BindPFlag("object_name", rootCmd.PersistentFlags().Lookup("object_name")))
	// 设置必须flag
	//cobra.CheckErr(rootCmd.MarkFlagRequired(endPoint))
	//cobra.CheckErr(rootCmd.MarkFlagRequired(accessKeyId))
	//cobra.CheckErr(rootCmd.MarkFlagRequired(accessKeySecret))
	//cobra.CheckErr(rootCmd.MarkFlagRequired(bucketName))
	// 设置默认值
	viper.SetDefault("platform", "ali")
	viper.SetDefault("object_name", fmt.Sprintf("%d%d%d", time.Now().Year(), time.Now().Month(), time.Now().Day()))

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
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".oss-tool" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".oss-tool")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
