/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bytes"
	"context"
	"encoding/binary"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/autrk/pluggable-calculator/internal/operation"
	extism "github.com/extism/go-sdk"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds two numbers",
	Long:  `Performs addition on two numbers.`,
	Args:  cobra.MatchAll(cobra.ExactArgs(2), cobra.OnlyValidArgs),
	Run:   runAddCommand,
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func runAddCommand(cmd *cobra.Command, args []string) {
	a, err := strconv.ParseFloat(args[0], 64)
	if err != nil {
		panic(err)
	}

	b, err := strconv.ParseFloat(args[1], 64)
	if err != nil {
		panic(err)
	}

	sum := operation.Add(a, b)
	fmt.Printf("Addition of %s is %.3f\n", args, sum)

	wasmFile := "./plugins/wasm/add.wasm"

	manifest := extism.Manifest{
		Wasm: []extism.Wasm{
			extism.WasmFile{
				Path: wasmFile,
			},
		},
	}

	ctx := context.Background()
	config := extism.PluginConfig{EnableWasi: true}
	plugin, err := extism.NewPlugin(ctx, manifest, config, []extism.HostFunction{})

	if err != nil {
		fmt.Printf("Failed to initialize plugin: %v\n", err)
		os.Exit(1)
	}

	// Floats in Bytes umwandeln
	var buf bytes.Buffer
	if err := binary.Write(&buf, binary.LittleEndian, a); err != nil {
		log.Fatal(err)
	}
	if err := binary.Write(&buf, binary.LittleEndian, b); err != nil {
		log.Fatal(err)
	}

	// Plugin ausführen
	exit, result, err := plugin.Call("add", buf.Bytes())
	if err != nil {
		fmt.Println(err)
		os.Exit(int(exit))
	}

	// Ergebnis (8 Bytes) zurück in float64 konvertieren
	if err := binary.Read(bytes.NewReader(result), binary.LittleEndian, &sum); err != nil {
		fmt.Println(err)
		os.Exit(int(exit))
	}

	fmt.Printf("Result: %.2f\n", sum) // ➔ Result: 10.00
}
