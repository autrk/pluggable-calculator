/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"fmt"
	"os"

	extism "github.com/extism/go-sdk"
	"github.com/spf13/cobra"
)

// countCmd represents the count command
var countCmd = &cobra.Command{
	Use:   "count",
	Short: "Counts the number of vowels in a given string",
	Long: `Counts the number of vowels (a, e, i, o, u) in the provided string using the "count_vowels"
function from the plugin. The result is returned as a JSON-encoded response.`,
	Run: count,
}

func init() {
	rootCmd.AddCommand(countCmd)
}

func count(cmd *cobra.Command, args []string) {
	fmt.Println("count called")

	manifest := extism.Manifest{
		Wasm: []extism.Wasm{
			extism.WasmUrl{
				Url: "https://github.com/extism/plugins/releases/latest/download/count_vowels.wasm",
			},
		},
	}

	ctx := context.Background()
	config := extism.PluginConfig{}
	plugin, err := extism.NewPlugin(ctx, manifest, config, []extism.HostFunction{})

	if err != nil {
		fmt.Printf("Failed to initialize plugin: %v\n", err)
		os.Exit(1)
	}

	data := []byte("Hello, World!")
	exit, out, err := plugin.Call("count_vowels", data)
	if err != nil {
		fmt.Println(err)
		os.Exit(int(exit))
	}

	response := string(out)
	fmt.Println(response)
}
