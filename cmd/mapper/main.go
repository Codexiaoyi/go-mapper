package main

import (
	"log"

	"github.com/Codexiaoyi/go-mapper/internal/mapper"
	"github.com/spf13/cobra"
)

var rootCommand = &cobra.Command{
	Use:     "mapper",
	Short:   "mapper is a tool of map struct.",
	Long:    "mapper is a tool of map struct.",
	Version: "v2.0.0",
	Run:     mapper.Run,
}

func init() {
	rootCommand.AddCommand(mapper.GenCmd)
}

func main() {
	if err := rootCommand.Execute(); err != nil {
		log.Fatal(err)
	}
}
