// Copyright © 2016 Mario Vejlupek <mario@vejlupek.cz>
//
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

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/spf13/cobra"
	"golang.org/x/net/context"
	"os"
	"io/ioutil"
	"strings"
)

// cleanupCmd represents the cleanup command
var cleanupCmd = &cobra.Command{
	Use:   "cleanup",
	Short: "Cleanup unused containers",
	Long:  `Check for lock files and try to compare them to running containers, remove missing`,
	Run: func(cmd *cobra.Command, args []string) {

		checkLockFiles( listContainers() )

	},
}

func init() {
	RootCmd.AddCommand(cleanupCmd)
}

func listContainers() []string {

	defaultHeaders := map[string]string{"User-Agent": "engine-api-cli-1.0"}
	cli, err := client.NewClient("unix:///var/run/docker.sock", "v1.22", nil, defaultHeaders)

	checkError( err )

	options := types.ContainerListOptions{}
	containers, err := cli.ContainerList(context.Background(), options)

	checkError( err )

	var containersId = make( []string, 99 )

	for index, c := range containers {

		containersId[index] = c.ID

	}

	return containersId

}

func checkLockFiles( containers []string ) {

	//TODO: use ENV
	files, _ := ioutil.ReadDir("/opt/docker-shared/agentlock")

	for _, f := range files {

		//TODO: use ENV
		lockFilePath := "/opt/docker-shared/agentlock/" + f.Name()

		data, err := ioutil.ReadFile( lockFilePath )

		checkError( err )

		if stringInSlice( string(data), containers ) {

			fmt.Printf("matching container %s for file %s [KEEPING]\n", string(data), lockFilePath )

		} else {

			fmt.Printf("abadoned file %s for container %s [REMOVING]\n", lockFilePath, string(data) )

			var err = os.Remove( lockFilePath )
			checkError( err )

		}

	}
}

func checkError(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(0)
	}
}

func stringInSlice(str string, list []string) bool {
	for _, v := range list {

		if strings.TrimSpace( v ) == strings.TrimSpace( str ) {
			return true
		}
	}
	return false
}