/*
Copyright © 2021 Paul Ebose <paulebose@gmail.com>

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program. If not, see <http://www.gnu.org/licenses/>.
*/
package cmd

import (
	"github.com/spf13/cobra"
)

// newRun executes when `new` command is run.
func newRun(cmd *cobra.Command, args []string) {
	initNoteFlags(cmd)
	cmd.Flags().Set("open", "true")
	noteRun(cmd, args)
}

// newCmd represents the `new` command
var newCmd = &cobra.Command{
	Use:   "new",
	Short: `Alias for "journal create note --open"`,
	Run:   newRun,
}

func init() {
	rootCmd.AddCommand(newCmd)
}
