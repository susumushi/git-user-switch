/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

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
	"errors"
	"git-user-switch/profile"
	"strings"

	"github.com/spf13/cobra"
)

// defineCmd represents the define command
var defineCmd = &cobra.Command{
	Use:   "define",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		c := profile.Profiles{}
		c.Load(config)
		//if err := c.Load(); err != nil {
		//fmt.Printf("error : %s\n", err)
		//os.Exit(1)
		//}
		if len(args) < 3 {
			return errors.New("profile values must be specified")
		}
		targetURLs := strings.Split(args[3], `,`)
		if err := c.Set(profile.Profile{
			Name:                 args[0],
			Email:                args[1],
			NickName:             args[2],
			InsertUsernameTarget: targetURLs,
		}); err != nil {
			return err
		}
		if err := c.Save(config); err != nil {
			return err
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(defineCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// defineCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// defineCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
