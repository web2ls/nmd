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
	"io/ioutil"
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

		isDeepFlag, _ := cmd.Flags().GetBool("deep")

		if isDeepFlag {
			deleteMultipleTargetFolder()
		} else {
			deleteSingleTargetFolder()
		}
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
	startCmd.Flags().BoolP("deep", "d", false, "delete target folder from all folders inside on one deep level")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func deleteSingleTargetFolder() {
	err := os.RemoveAll("node_modules")
	checkError(err)

	fmt.Println("Folder has been deleted")
}

func deleteMultipleTargetFolder() {
	dirContent, err := ioutil.ReadDir(".")
	checkError(err)

	var directoryList []string
	for _, item := range dirContent {
		if item.IsDir() {
			directoryList = append(directoryList, item.Name())
		}
	}

	for _, checkingDir := range directoryList {
		err = os.Chdir(checkingDir)
		checkError(err)

		err = os.RemoveAll("node_modules")
		checkError(err)

		err = os.Chdir("./../")
		checkError(err)
	}

	fmt.Println("All target folders inside has been deleted")
}
