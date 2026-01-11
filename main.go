package main

import (
	"fmt"

	"github.com/tanq16/ai-context/cmd"
	"github.com/tanq16/ai-context/config"
)

func main() {
	if m := config.DefaultModel(); m != "" {
		fmt.Printf("Default AI model: %s\n", m)
	}
	cmd.Execute()
}
