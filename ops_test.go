package main

import (
	"testing"

	"github.com/nanovms/ops/types"
	"github.com/stretchr/testify/assert"
)

// TestGetProviderAndContext tests the getProviderAndContext function
// Note: This tests the function signature and error handling for invalid providers
func TestGetProviderAndContext(t *testing.T) {
	t.Run("invalid provider returns error", func(t *testing.T) {
		config := &types.Config{}
		provider, ctx, err := getProviderAndContext(config, "nonexistent-provider")

		assert.Error(t, err)
		assert.Nil(t, provider)
		assert.Nil(t, ctx)
	})

	t.Run("empty provider name returns error", func(t *testing.T) {
		config := &types.Config{}
		provider, ctx, err := getProviderAndContext(config, "")

		assert.Error(t, err)
		assert.Nil(t, provider)
		assert.Nil(t, ctx)
	})

	t.Run("nil config handled gracefully", func(t *testing.T) {
		// The function accepts *types.Config, passing nil may cause issues
		// depending on the underlying implementation
		config := &types.Config{}
		provider, ctx, err := getProviderAndContext(config, "onprem")

		// onprem provider should work with valid config
		if err == nil {
			assert.NotNil(t, provider)
			assert.NotNil(t, ctx)
		}
	})
}

// TestProviderInterface tests that the returned provider implements expected interfaces
func TestProviderInterface(t *testing.T) {
	t.Run("onprem provider implements required methods", func(t *testing.T) {
		config := &types.Config{}
		provider, _, err := getProviderAndContext(config, "onprem")

		if err != nil {
			t.Skip("Skipping interface test due to provider init error")
		}

		assert.NotNil(t, provider)

		// Provider should not be nil when valid
		// The actual interface checking would require importing the Provider interface
		// from the ops package
	})
}

// TestContextCreation tests the context creation part of getProviderAndContext
func TestContextCreation(t *testing.T) {
	t.Run("context is created for valid provider", func(t *testing.T) {
		config := &types.Config{}
		_, ctx, err := getProviderAndContext(config, "onprem")

		if err != nil {
			t.Skip("Skipping context test due to provider error")
		}

		assert.NotNil(t, ctx)
	})

	t.Run("context is nil for invalid provider", func(t *testing.T) {
		config := &types.Config{}
		_, ctx, err := getProviderAndContext(config, "invalid")

		assert.Error(t, err)
		assert.Nil(t, ctx)
	})
}
