package main

import "github.com/p886/5-million-steps/cmd"

func main() {
	cmd.RootCmd.AddCommand(cmd.ServerCmd)
	cmd.Execute()
}
