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
	"fmt"
	"git-user-switch/gituser"
	"git-user-switch/profile"
	"os"

	"github.com/spf13/cobra"
)

var targetGlobal = false
var targetSystem = false
var targetLocal = false

// setCmd represents the set command
var setCmd = &cobra.Command{
	Use:   "set",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		var gu gituser.GitUser
		var err error
		if targetGlobal {
			gu, err = gituser.New(gituser.TargetScopeGlobal)
			if err != nil {
				os.Exit(1)
			}
		} else if targetLocal {
			gu, err = gituser.New(gituser.TargetScopeLocal)
			if err != nil {
				os.Exit(1)
			}
		} else if targetSystem {
			gu, err = gituser.New(gituser.TargetScopeSystem)
			if err != nil {
				os.Exit(1)
			}
		} else {
			//とりあえずデフォルトはグローバルかな。頻繁に設定すると思うし。
			//内心自動で現在有効なスコープを特定して設定するの面倒くさい
			gu, err = gituser.New(gituser.TargetScopeGlobal)
			if err != nil {
				os.Exit(1)
			}
		}
		c := profile.Profiles{}
		if err := c.Load(); err != nil {
			fmt.Printf("error : %s\n", err)
			os.Exit(1)
		}
		for _, p := range c {
			if p.NickName == args[0] {
				gu.Name = p.Name
				gu.Email = p.Email
				gu.InsertUsernameTarget = p.InsertUsernameTarget
				if err := gu.SetConfig(); err != nil {
					fmt.Printf("error : %s\n", err)
					os.Exit(1)
				}
			}
		}
	},
	ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		ps := profile.Profiles{}
		ps.Load()

		nicknames := []string{}
		for _, p := range ps {
			nicknames = append(nicknames, p.NickName)
		}
		return nicknames, cobra.ShellCompDirectiveNoFileComp | cobra.ShellCompDirectiveNoSpace
	},
}

func init() {
	setCmd.Flags().BoolVarP(&targetGlobal, "global", "g", false, "")
	setCmd.Flags().BoolVarP(&targetSystem, "system", "s", false, "")
	setCmd.Flags().BoolVarP(&targetLocal, "local", "l", false, "")
	rootCmd.AddCommand(setCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// setCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// setCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
