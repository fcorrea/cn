/*
 * Ceph Nano (C) 2018 Red Hat, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

/*
 * Below main package has canonical imports for 'go get' and 'go build'
 * to work with all other clones of github.com/ceph/cn repository. For
 * more information refer https://golang.org/doc/go1.4#canonicalimports
 */

package cmd

import (
	"fmt"

	"github.com/apcera/termtables"
	"github.com/spf13/cobra"
)

var (
	cmdImages = &cobra.Command{
		Use:   "images [command] [arg]",
		Short: "Interact with images",
		Args:  cobra.NoArgs,
	}
)

func init() {
	cmdImages.AddCommand(
		cliImagesList(),
	)
}

// cliFlavorsList is the Cobra CLI call
func cliImagesList() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "ls",
		Short: "Print the list of images",
		Args:  cobra.NoArgs,
		Run:   listImages,
	}
	return cmd
}

func listImages(cmd *cobra.Command, args []string) {
	table := termtables.CreateTable()
	table.AddHeaders("NAME", "IMAGE_NAME")
	for image := range getItemsFromGroup(IMAGES) {
		table.AddRow(image, getImageName(image))
	}
	fmt.Println(table.Render())
}
