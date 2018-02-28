// Copyright 2016 Palantir Technologies, Inc.
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
	"github.com/palantir/godel/framework/pluginapi"
	"github.com/spf13/cobra"

	"github.com/palantir/go-extimport/extimport"
)

var (
	RootCmd = &cobra.Command{
		Use:   "extimport [flags] [packages]",
		Short: "checks whether project imports any external packages (packages that are not within the project or its vendor directories)",
		RunE: func(cmd *cobra.Command, args []string) error {
			return extimport.Run(projectDirFlagVal, args, listFlagVal, allFlagVal, cmd.OutOrStdout())
		},
	}

	projectDirFlagVal string
	listFlagVal       bool
	allFlagVal        bool
)

func init() {
	pluginapi.AddProjectDirPFlagPtr(RootCmd.Flags(), &projectDirFlagVal)
	RootCmd.Flags().BoolVar(&listFlagVal, "list", false, "print external dependencies one per line")
	RootCmd.Flags().BoolVar(&allFlagVal, "all", false, "list all external dependencies, including those multiple levels deep")
}
