/*
Copyright © 2019 Alex Slobodyansky <web2ls@yandex.ru>

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
	"os"

	"github.com/spf13/cobra"
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Basic command for delete node_modules folder in the current directory",
	Long:  `Find and delete node_modules folder in the current directory`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Deletting node_modules folder...")

		err := os.RemoveAll("node_modules")

		if err != nil {
			fmt.Println("folder not found or error occurs on deleting")
		} else {
			fmt.Println("Folder has been deleted")
		}
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
}
