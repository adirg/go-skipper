// Copyright © 2017 NAME HERE <EMAIL ADDRESS>
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
	"context"
	"log"
	"path/filepath"
	"strings"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/spf13/cobra"
)

// buildCmd represents the build command
var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "Build a container",

	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here
		files, err := filepath.Glob("Dockerfile.*")
		if err != nil {
			log.Fatal(err)
		}

		var buildOptions = types.ImageBuildOptions{
			Remove: true,
		}
		images := make([]string, len(files))
		cli, err := client.NewEnvClient()
		for i, file := range files {
			file_parts := strings.SplitN(file, ".", 2)
			images[i] = file_parts[1]

			buildOptions.Dockerfile = file
			_, err := cli.ImageBuild(context.Background(), nil, buildOptions)
			if err != nil {
				panic(err)
			}

		}

		if err != nil {
			panic(err)
		}
	},
}

func init() {
	RootCmd.AddCommand(buildCmd)
}
