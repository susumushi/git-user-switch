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
	"git-user-switch/gituser"
	"git-user-switch/profile"

	"github.com/spf13/cobra"
)

var targetGlobal = false
var targetSystem = false
var targetLocal = false

// setCmd represents the set command
var setCmd = &cobra.Command{
	Use:   "set [profile nickname]",
	Short: "Set defined profile to git config",
	Long: `Set defined profile to git config

Set sub command is setting git user by defined user profile,
You must specify profile by nickname, you can set config target
switching by scope flags.`,
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		var gu gituser.GitUser
		if targetGlobal {
			gu, err = gituser.New(gituser.TargetScopeGlobal)
			if err != nil {
				return err
			}
		} else if targetLocal {
			gu, err = gituser.New(gituser.TargetScopeLocal)
			if err != nil {
				return err
			}
		} else if targetSystem {
			gu, err = gituser.New(gituser.TargetScopeSystem)
			if err != nil {
				return err
			}
		} else {
			//とりあえずデフォルトはグローバルかな。頻繁に設定すると思うし。
			//内心自動で現在有効なスコープを特定して設定するの面倒くさい
			gu, err = gituser.New(gituser.TargetScopeGlobal)
			if err != nil {
				return err
			}
		}
		c := profile.Profiles{}
		if err := c.Load(); err != nil {
			return err
		}
		profileIsNotMatched := true
		for _, p := range c {
			if p.NickName == args[0] {
				profileIsNotMatched = false
				gu.Name = p.Name
				gu.Email = p.Email
				gu.InsertUsernameTarget = p.InsertUsernameTarget
				if err := gu.SetConfig(); err != nil {
					return err
				}
			}
		}
		if profileIsNotMatched {
			return errors.New("profiles is not matched git user information")
		}
		return nil
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
	Args: cobra.ExactValidArgs(1),
}

func init() {
	setCmd.Flags().BoolVarP(&targetGlobal, "global", "g", false, "set to global scope (ex. ~/.gitconfig)")
	setCmd.Flags().BoolVarP(&targetSystem, "system", "s", false, "set to system scope (ex. /etc/gitconfig)")
	setCmd.Flags().BoolVarP(&targetLocal, "local", "l", false, "set to local scope (ex. ${YOUR_GIT_REPOSITORY}/.git/config)")
	rootCmd.AddCommand(setCmd)

}
