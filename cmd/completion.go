/*
Copyright 2020 The KubeSphere Authors.

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
)

// CompletionOptions is the option of completion command
type CompletionOptions struct {
	Type string
}

// ShellTypes contains all types of shell
var ShellTypes = []string{
	"zsh", "bash", "powerShell",
}

var completionOptions CompletionOptions

func init() {
	rootCmd.AddCommand(completionCmd)

	flags := completionCmd.Flags()
	flags.StringVarP(&completionOptions.Type, "type", "t", "",
		fmt.Sprintf("Generate different types of shell which are %v", ShellTypes))

	err := completionCmd.RegisterFlagCompletionFunc("type", func(cmd *cobra.Command, args []string, toComplete string) (
		i []string, directive cobra.ShellCompDirective) {
		return ShellTypes, cobra.ShellCompDirectiveDefault
	})
	if err != nil {
		completionCmd.PrintErrf("register flag type for sub-command doc failed %#v\n", err)
	}
}

var completionCmd = &cobra.Command{
	Use:   "completion",
	Short: "Generate shell completion scripts",
	Long: `Generate shell completion scripts
Normally you don't need to do more extra work to have this feature if you've installed kk by brew`,
	Example: `Installing bash completion on macOS using homebrew
If running Bash 3.2 included with macOS
brew install bash-completion
or, if running Bash 4.1+
brew install bash-completion@2
You may need to add the completion to your completion directory by the following command
kk completion > $(brew --prefix)/etc/bash_completion.d/kk
If you get trouble, please visit https://github.com/jenkins-zh/jenkins-cli/issues/83.

Load kk completion code for bash into the current shell
source <(kk completion --type bash)

In order to have good experience on zsh completion, ohmyzsh is a good choice.
Please install ohmyzsh by the following command
sh -c "$(curl -fsSL https://raw.githubusercontent.com/ohmyzsh/ohmyzsh/master/tools/install.sh)"
Get more details about onmyzsh from https://github.com/ohmyzsh/ohmyzsh

Load the kk completion code for zsh[1] into the current shell
source <(kk completion --type zsh)
Set the kk completion code for zsh[1] to autoload on startup
kk completion --type zsh > "${fpath[1]}/_kk"`,
	RunE: func(cmd *cobra.Command, _ []string) (err error) {
		shellType := completionOptions.Type
		switch shellType {
		case "zsh":
			err = rootCmd.GenZshCompletion(cmd.OutOrStdout())
		case "powerShell":
			err = rootCmd.GenPowerShellCompletion(cmd.OutOrStdout())
		case "bash":
			err = rootCmd.GenBashCompletion(cmd.OutOrStdout())
		case "":
			err = cmd.Help()
		default:
			err = fmt.Errorf("unknown shell type %s", shellType)
		}
		return
	},
}
