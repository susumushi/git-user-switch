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

	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,

	RunE: func(cmd *cobra.Command, args []string) error {
		c := profile.Profiles{}
		if err := c.Load(config); err != nil {
			return err
		}

		if len(args) < 1 {
			return errors.New("delete target nickname must be specified")
		}

		notMatched := true
		for _, p := range c {
			if p.NickName == args[0] {
				c.Delete(p)
				notMatched = false
			}
		}
		if notMatched {
			return errors.New("delete target nickname not found")
		}
		if err := c.Save(config); err != nil {
			return err
		}
		return nil
	},
	ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		ps := profile.Profiles{}
		ps.Load(config)

		nicknames := []string{}
		for _, p := range ps {
			nicknames = append(nicknames, p.NickName)
		}
		return nicknames, cobra.ShellCompDirectiveNoFileComp | cobra.ShellCompDirectiveNoSpace
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
