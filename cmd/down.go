// Prototype (☺) 2017 Manuel Ramirez Lopez <ramz@zhaw.ch>
// Copyright (©) 2017 Zürcher Hochschule für Angewandte Wissenschaften
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"os/exec"
	"github.com/spf13/cobra"
)

// downCmd represents the down command
var downCmd = &cobra.Command{
	Use:   "down",
	Short: "Delete a project in your old cluster",
	Long: `Delete a project in your old cluster`,
	Run: func(cmd *cobra.Command, args []string) {
		down(cmd, args)
	},
}


func down(cmd *cobra.Command, args []string) {
	getAllValue()
	PathTemplate = PathTemplate + "/" + ProjectFrom
	loginCluster(ClusterFrom, UsernameFrom, PasswordFrom)
	changeProject(ProjectFrom)
	fmt.Println("before")
	fmt.Println(ObjectsOc)
	ObjectsOc = getTypeObjects(ObjectsOc)
	fmt.Println("after")
	fmt.Println(ObjectsOc)
	for _, typeObject := range ObjectsOc {
		fullPath := PathTemplate + "/" + typeObject
		delete(fullPath)
	}
}

func delete(path string){
	CmdDelete := exec.Command("oc", "delete", "-f",  path)
	CmdDeleteOut, err := CmdDelete.Output()
	if CmdDeleteOut != nil {
		fmt.Println(string(CmdDeleteOut))
	}
	if err != nil {
		fmt.Println("Error deleting " + path)

	}
}

func init() {
	RootCmd.AddCommand(downCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// downCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// downCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
