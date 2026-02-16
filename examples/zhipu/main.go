// Package main demonstrates the usage of Zhipu's GLM models through the gollm library.
// Zhipu provides an Anthropic-compatible API for their GLM models including GLM-4.7.
package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/teilomillet/gollm"
)

func main() {
	// Get API key from environment variable
	apiKey := os.Getenv("ZHIPU_API_KEY")
	if apiKey == "" {
		log.Fatal("ZHIPU_API_KEY environment variable is not set")
	}

	// Create LLM client with Zhipu provider
	llm, err := gollm.NewLLM(
		gollm.SetProvider("zhipu"),
		gollm.SetModel("glm-4.7"),
		gollm.SetAPIKey(apiKey),
		gollm.SetMaxTokens(500),
		gollm.SetTemperature(0.7),
	)
	if err != nil {
		log.Fatalf("Failed to create LLM client: %v", err)
	}

	ctx := context.Background()

	// Example 1: Simple text generation
	fmt.Println("=== Example 1: Simple Text Generation ===")
	prompt1 := llm.NewPrompt("Explique em português o que é computação quântica de forma simples.")

	response, err := llm.Generate(ctx, prompt1)
	if err != nil {
		log.Fatalf("Generate failed: %v", err)
	}

	fmt.Printf("\nPergunta: %s\n\n", prompt1)
	fmt.Printf("Resposta: %s\n\n", response)

	// Example 2: With system prompt
	fmt.Println("=== Example 2: With System Prompt ===")
	llm.SetSystemPrompt("Você é um assistente útil que responde de forma concisa e clara.", "")

	prompt2 := llm.NewPrompt("Quais são os principais benefícios da programação em Go?")

	response2, err := llm.Generate(ctx, prompt2)
	if err != nil {
		log.Fatalf("Generate with system prompt failed: %v", err)
	}

	fmt.Printf("\nPergunta: %s\n\n", prompt2)
	fmt.Printf("Resposta: %s\n\n", response2)

	// Example 3: Structured conversation
	fmt.Println("=== Example 3: Conversation ===")
	prompt3 := gollm.NewPrompt(
		"Liste 3 vantagens de usar modelos GLM da Zhipu",
		gollm.WithDirectives(
			"Seja conciso",
			"Use tópicos para a lista",
			"Explique cada vantagem em uma frase",
		),
	)

	response3, err := llm.Generate(ctx, prompt3)
	if err != nil {
		log.Fatalf("Generate with prompt failed: %v", err)
	}

	fmt.Printf("\nResposta:\n%s\n", response3)
}
