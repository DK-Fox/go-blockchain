/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

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

var from string
var to string
var msg string

// sendCmd represents the send command
var sendCmd = &cobra.Command{
	Use:   "send",
	Short: "Send a message from a node to another",
	Long:  `Send a message from a node to another`,
	Run: func(cmd *cobra.Command, args []string) {
		//trigger a trasaction
		cli.send(from, to, msg, nodeId)
		fmt.Println("send called from ", from, " to", to, "message ", msg)
	},
}

func init() {
	rootCmd.AddCommand(sendCmd)
	sendCmd.Flags().StringVarP(&from, "from", "f", "", "sender node address")
	sendCmd.Flags().StringVarP(&to, "to", "t", "", "receiver node address")
	sendCmd.Flags().StringVarP(&msg, "message", "m", "", "message")
}
