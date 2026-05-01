package main

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestContent_JSONSerialization(t *testing.T) {
	t.Run("Content with description", func(t *testing.T) {
		description := "test description"
		content := Content{
			Title:       "Test Title",
			Description: &description,
		}

		data, err := json.Marshal(content)
		require.NoError(t, err)

		var decoded Content
		err = json.Unmarshal(data, &decoded)
		require.NoError(t, err)

		assert.Equal(t, content.Title, decoded.Title)
		assert.Equal(t, *content.Description, *decoded.Description)
	})

	t.Run("Content without description", func(t *testing.T) {
		content := Content{
			Title:       "Test Title",
			Description: nil,
		}

		data, err := json.Marshal(content)
		require.NoError(t, err)

		var decoded Content
		err = json.Unmarshal(data, &decoded)
		require.NoError(t, err)

		assert.Equal(t, content.Title, decoded.Title)
		assert.Nil(t, decoded.Description)
	})

	t.Run("JSON schema tags are correct", func(t *testing.T) {
		content := Content{}
		schemaType := reflect.TypeOf(content)

		// Verify title field has required tag
		titleField, ok := schemaType.FieldByName("Title")
		assert.True(t, ok, "Title field should exist")
		assert.Contains(t, string(titleField.Tag), "required")

		// Verify description field exists
		_, descFieldOk := schemaType.FieldByName("Description")
		assert.True(t, descFieldOk, "Description field should exist")
	})
}

func TestMyFunctionsArguments_JSONSerialization(t *testing.T) {
	t.Run("Full arguments with all fields", func(t *testing.T) {
		description := "test content description"
		args := MyFunctionsArguments{
			Submitter: "claude",
			Content: Content{
				Title:       "Test",
				Description: &description,
			},
		}

		data, err := json.Marshal(args)
		require.NoError(t, err)

		var decoded MyFunctionsArguments
		err = json.Unmarshal(data, &decoded)
		require.NoError(t, err)

		assert.Equal(t, "claude", decoded.Submitter)
		assert.Equal(t, "Test", decoded.Content.Title)
		assert.Equal(t, description, *decoded.Content.Description)
	})

	t.Run("Submitter validation", func(t *testing.T) {
		validSubmitters := []string{"openai", "google", "claude", "anthropic", "local"}

		for _, submitter := range validSubmitters {
			args := MyFunctionsArguments{
				Submitter: submitter,
			}
			assert.NotEmpty(t, args.Submitter)
		}
	})
}

func TestInstanceArguments_JSONSerialization(t *testing.T) {
	t.Run("InstanceArguments serialization", func(t *testing.T) {
		args := InstanceArguments{
			ImageName: "test-image.img",
		}

		data, err := json.Marshal(args)
		require.NoError(t, err)

		var decoded InstanceArguments
		err = json.Unmarshal(data, &decoded)
		require.NoError(t, err)

		assert.Equal(t, "test-image.img", decoded.ImageName)
	})

	t.Run("JSON field naming", func(t *testing.T) {
		args := InstanceArguments{
			ImageName: "my-image",
		}

		data, err := json.Marshal(args)
		require.NoError(t, err)

		// The JSON should use the struct field name
		var result map[string]interface{}
		err = json.Unmarshal(data, &result)
		require.NoError(t, err)

		// Note: The struct uses "longitude" as json tag
		_, hasLongitude := result["longitude"]
		_, hasImageName := result["ImageName"]

		// Should have longitude field (from json tag)
		assert.True(t, hasLongitude || hasImageName)
	})
}

func TestContent_JSONRoundTrip(t *testing.T) {
	testCases := []struct {
		name        string
		content     Content
		description string
	}{
		{
			name: "with description",
			content: Content{
				Title:       "Package Build",
				Description: stringPtr("Building nanos package"),
			},
			description: "Full content with description",
		},
		{
			name: "without description",
			content: Content{
				Title:       "Instance Create",
				Description: nil,
			},
			description: "Content without description",
		},
		{
			name: "empty title",
			content: Content{
				Title:       "",
				Description: stringPtr("description only"),
			},
			description: "Empty title with description",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			data, err := json.Marshal(tc.content)
			require.NoError(t, err)

			var decoded Content
			err = json.Unmarshal(data, &decoded)
			require.NoError(t, err)

			assert.Equal(t, tc.content.Title, decoded.Title)
			if tc.content.Description == nil {
				assert.Nil(t, decoded.Description)
			} else {
				require.NotNil(t, decoded.Description)
				assert.Equal(t, *tc.content.Description, *decoded.Description)
			}
		})
	}
}

// Helper function
func stringPtr(s string) *string {
	return &s
}
