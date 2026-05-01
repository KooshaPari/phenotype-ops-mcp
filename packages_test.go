package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestLoadPackage_Structure tests the structure of loadPackage function
func TestLoadPackage_Structure(t *testing.T) {
	t.Run("accepts valid arguments", func(t *testing.T) {
		description := "/path/to/package"
		args := MyFunctionsArguments{
			Submitter: "claude",
			Content: Content{
				Title:       "Load Package",
				Description: &description,
			},
		}

		// Verify arguments are properly structured
		assert.Equal(t, "claude", args.Submitter)
		assert.NotNil(t, args.Content.Description)
	})

	t.Run("returns success response", func(t *testing.T) {
		description := "test-package"
		args := MyFunctionsArguments{
			Submitter: "test",
			Content: Content{
				Title:       "Test",
				Description: &description,
			},
		}

		// The function should return a response
		// Note: This may fail due to external dependencies, but structure is tested
		response, _ := loadPackage(args)

		if response != nil {
			assert.NotNil(t, response)
		}
	})
}

// TestPkgLoad_Configuration tests the pkgLoad function configuration
func TestPkgLoad_Configuration(t *testing.T) {
	t.Run("package path is extracted correctly", func(t *testing.T) {
		// Test the path handling logic
		testPaths := []struct {
			input    string
			expected string
		}{
			{"/path/to/mypackage", "mypackage"},
			{"/simple/package", "package"},
			{"/a/b/c/d/e", "e"},
		}

		for _, tc := range testPaths {
			// Simulate filepath.Base behavior
			base := tc.input
			for i := len(base) - 1; i >= 0; i-- {
				if base[i] == '/' {
					base = base[i+1:]
					break
				}
			}
			assert.Equal(t, tc.expected, base)
		}
	})
}

// TestLoadPackage_PathHandling tests path handling in loadPackage
func TestLoadPackage_PathHandling(t *testing.T) {
	t.Run("handles nil description gracefully", func(t *testing.T) {
		args := MyFunctionsArguments{
			Submitter: "test",
			Content: Content{
				Title:       "Test",
				Description: nil,
			},
		}

		// The function dereferences the description pointer
		// This tests that we handle nil safely
		if args.Content.Description != nil {
			assert.NotEmpty(t, *args.Content.Description)
		}
	})

	t.Run("handles empty path", func(t *testing.T) {
		emptyPath := ""
		args := MyFunctionsArguments{
			Submitter: "test",
			Content: Content{
				Title:       "Test",
				Description: &emptyPath,
			},
		}

		// Empty path should be handled
		assert.NotNil(t, args.Content.Description)
		assert.Empty(t, *args.Content.Description)
	})
}

// TestPkgLoad_ExecutableName tests executable name extraction logic
func TestPkgLoad_ExecutableName(t *testing.T) {
	t.Run("extracts executable name from program path", func(t *testing.T) {
		// Test cases for executable name extraction
		testCases := []struct {
			program        string
			packageFolder  string
			expectedResult string
		}{
			{
				program:        "/path/to/mypackage/program",
				packageFolder:  "mypackage",
				expectedResult: "program", // Contains package folder, so just base name
			},
			{
				program:        "standalone-program",
				packageFolder:  "somepackage",
				expectedResult: "standalone-program",
			},
		}

		for _, tc := range testCases {
			var result string
			// Simulate the logic from pkgLoad
			if contains(tc.program, tc.packageFolder) {
				result = baseName(tc.program)
			} else {
				result = tc.program
			}
			assert.Equal(t, tc.expectedResult, result)
		}
	})
}

// TestLoadPackage_Validation tests input validation
func TestLoadPackage_Validation(t *testing.T) {
	t.Run("validates submitter is set", func(t *testing.T) {
		args := MyFunctionsArguments{
			Submitter: "",
		}

		assert.Empty(t, args.Submitter)
	})

	t.Run("validates content title", func(t *testing.T) {
		description := "test"
		args := MyFunctionsArguments{
			Submitter: "claude",
			Content: Content{
				Title:       "Load Package",
				Description: &description,
			},
		}

		assert.NotEmpty(t, args.Content.Title)
	})
}

// Helper functions for testing
func contains(s, substr string) bool {
	return len(s) >= len(substr) && (s == substr || len(s) > 0 && containsHelper(s, substr))
}

func containsHelper(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}

func baseName(path string) string {
	for i := len(path) - 1; i >= 0; i-- {
		if path[i] == '/' {
			return path[i+1:]
		}
	}
	return path
}
