// Copyright 息 2016 Shigeyuki Fujishima <shigeyuki.fujishima@gmail.com>
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
	"os"
	"strconv"
	"strings"

	log "github.com/Sirupsen/logrus"
	"github.com/spf13/cobra"
	"golang.org/x/text/width"
)

var Zenkaku bool

func convert2int(args []string) ([]int, error) {
	ret := make([]int, len(args))
	for i, n := range args {
		x, err := strconv.Atoi(n)
		if err != nil {
			return []int{}, err
		}
		ret[i] = x
	}

	return ret, nil
}

func convert2zenkaku(s string) string {
	return width.Widen.String(s)
}

func add(nums []int, x int) int {
	log.Debugf("nums: %v, x: %d", nums, x)
	if len(nums) == 0 {
		return x
	}
	x = nums[0] + x
	if len(nums) == 1 {
		return x
	}
	nums = nums[1:]
	return add(nums, x)
}

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "just adding 2 arguments",
	Long:  `just adding 2 arguments.`,
	Run: func(cmd *cobra.Command, args []string) {
		if Verbose {
			log.SetLevel(log.DebugLevel)
			log.Debug("Verbose mode")
		}

		log.Debugf("Args: %v", args)
		if len(args) < 2 {
			log.Errorf("Need 2 params. You passed %d args", len(args))
			os.Exit(1)
		}

		nums, err := convert2int(args)
		if err != nil {
			log.Errorf("Failed to convert arguments into number: %v", err.Error())
			os.Exit(1)
		}

		result := add(nums, 0)
		resultMessage := fmt.Sprintf("%s = %d", strings.Join(args, " + "), result)
		if Zenkaku {
			log.Debug("Convert a message to Zenkaku")
			resultMessage = convert2zenkaku(resultMessage)
		}
		log.Infof(">>> %s", resultMessage)
	},
}

func init() {
	RootCmd.AddCommand(addCmd)

	addCmd.Flags().BoolVarP(&Zenkaku, "zenkaku", "z", false, "Output the result as Zenkaku")

}
