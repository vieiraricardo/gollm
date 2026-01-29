// Package providers implements LLM provider interfaces and implementations.
package providers

// ZhipuProvider implements the Provider interface for Z.ai's GLM API.
// Z.ai provides an Anthropic-compatible API for their GLM models (including GLM-4.7),
// allowing us to reuse the AnthropicProvider implementation with a different endpoint.
type ZhipuProvider struct {
	*AnthropicProvider
}

// NewZhipuProvider creates a new Zhipu provider instance.
// It uses the Anthropic provider implementation since Z.ai's API is Anthropic-compatible.
//
// Parameters:
//   - apiKey: Z.ai API key for authentication
//   - model: The model to use (e.g., "glm-4.7", "glm-4")
//   - extraHeaders: Additional HTTP headers for requests
//
// Returns:
//   - A configured ZhipuProvider instance
func NewZhipuProvider(apiKey, model string, extraHeaders map[string]string) Provider {
	// Create the base Anthropic provider
	anthropicProvider := NewAnthropicProvider(apiKey, model, extraHeaders)

	// Wrap it in ZhipuProvider to override the endpoint
	return &ZhipuProvider{
		AnthropicProvider: anthropicProvider.(*AnthropicProvider),
	}
}

// Name returns "zhipu" as the provider identifier.
func (p *ZhipuProvider) Name() string {
	return "zhipu"
}

// Endpoint returns the Z.ai API endpoint URL.
// Z.ai provides an Anthropic-compatible API at https://api.z.ai/api/anthropic
func (p *ZhipuProvider) Endpoint() string {
	return "https://api.z.ai/api/anthropic/v1/messages"
}
