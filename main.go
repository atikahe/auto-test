package main

import (
	"fmt"
	"log"
	"os"
	"path"
	"time"

	"github.com/atikahe/auto-test/pkg/codex"
	"github.com/atikahe/auto-test/pkg/promptgen"
	"github.com/atikahe/auto-test/pkg/testfile"
	"github.com/spf13/cobra"
)

const (
	OpenAIAPIBaseURL string = "https://api.openai.com/v1"
)

var (
	filePath     string
	customPrompt string
	overwrite    bool
)

func loading() {
	for {
		fmt.Print(".")
		time.Sleep(500 * time.Millisecond)
		fmt.Print("\b \b")
		time.Sleep(500 * time.Millisecond)
	}
}

func main() {
	// Load env
	OpenAIAPIKey := os.Getenv("OPENAI_API_KEY")
	if OpenAIAPIKey == "" {
		log.Fatal("OPENAI_API_KEY is required.")
	}

	// Define flags
	cmd := &cobra.Command{
		Use:   "auto-test",
		Short: "An  example of using auto-test commands.",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) > 0 && (args[0] == "help" || args[0] == "--help" || args[0] == "-h") {
				cmd.Help()
				return
			}

			fmt.Print("Running auto-test")

		},
	}

	cmd.Flags().StringVarP(&filePath, "file", "f", "", "Path to file.")
	cmd.Flags().StringVarP(&customPrompt, "prompt", "p", "", "An additional custom prompt. It will be added to existing prompt by default.")
	cmd.Flags().BoolVarP(&overwrite, "override", "o", false, "Tells the program to overwrite existing prompt. False by default.")

	if err := cmd.Execute(); err != nil {
		log.Fatalf("unable to run command: %s", err)
	}

	if filePath == "" {
		log.Fatal("unable to run auto-test: file location required")
	}

	go loading()

	// Process prompt
	inputPrompt := promptgen.Customized{
		Text:      customPrompt,
		Overwrite: overwrite,
	}

	prompt, _ := promptgen.Build(filePath, inputPrompt)

	maxTokens := 1024
	if len(prompt) < maxTokens {
		maxTokens = len(prompt)
	}

	// Initiate codex
	cdx := codex.New(OpenAIAPIKey, OpenAIAPIBaseURL)
	response, err := cdx.CreateCompletion(&codex.CompletionArgs{
		Model:       cdx.Model,
		Prompt:      prompt,
		MaxTokens:   maxTokens,
		Temperature: 0,
	})
	if err != nil {
		log.Fatalf("%s", err)
	}

	// Generate test file
	generatedTestString := promptgen.AddPrefix(path.Ext(filePath), response.Choices[0].Text)
	fileName, err := testfile.Generate(filePath)
	if err != nil {
		log.Fatalf("failed to generate filename: %s", err)
	}

	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("error creating file: %s", err)
	}
	defer file.Close()

	_, err = file.WriteString(generatedTestString)
	if err != nil {
		log.Fatalf("error writing to file: %s", err)
	}

	fmt.Println()
	log.Println("Test file is successfully generated. Please make some necessary modifications if needed.")
}
