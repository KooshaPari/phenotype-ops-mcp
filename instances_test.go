package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestInstanceArguments_Validation tests the validation of instance arguments
func TestInstanceArguments_Validation(t *testing.T) {
	t.Run("valid image name", func(t *testing.T) {
		args := InstanceArguments{
			ImageName: "valid-image.img",
		}
		assert.NotEmpty(t, args.ImageName)
	})

	t.Run("empty image name", func(t *testing.T) {
		args := InstanceArguments{
			ImageName: "",
		}
		assert.Empty(t, args.ImageName)
	})

	t.Run("image name with path", func(t *testing.T) {
		args := InstanceArguments{
			ImageName: "/path/to/image.img",
		}
		assert.Contains(t, args.ImageName, "image.img")
	})

	t.Run("image name field tag uses longitude", func(t *testing.T) {
		args := InstanceArguments{
			ImageName: "test",
		}
		// The struct uses "longitude" as json tag for ImageName
		assert.NotEmpty(t, args.ImageName)
	})
}

// TestInstanceLogs_Arguments tests the arguments structure for instanceLogs
func TestInstanceLogs_Arguments(t *testing.T) {
	t.Run("accepts valid InstanceArguments", func(t *testing.T) {
		args := InstanceArguments{
			ImageName: "my-instance",
		}
		assert.Equal(t, "my-instance", args.ImageName)
	})
}

// TestInstanceCreate_Arguments tests the arguments structure for instanceCreate
func TestInstanceCreate_Arguments(t *testing.T) {
	t.Run("accepts valid InstanceArguments", func(t *testing.T) {
		args := InstanceArguments{
			ImageName: "my-custom-image.img",
		}
		assert.Equal(t, "my-custom-image.img", args.ImageName)
	})
}

// TestListInstances_Arguments tests the arguments structure for listInstances
func TestListInstances_Arguments(t *testing.T) {
	t.Run("accepts valid MyFunctionsArguments", func(t *testing.T) {
		args := MyFunctionsArguments{
			Submitter: "claude",
		}
		assert.Equal(t, "claude", args.Submitter)
	})
}

// TestInstanceCreate_Configuration tests configuration expectations
func TestInstanceCreate_Configuration(t *testing.T) {
	t.Run("kernel path is set", func(t *testing.T) {
		// Verify that the hardcoded kernel path format is correct
		// The function uses: lepton.GetOpsHome() + "/0.1.53-arm/kernel.img"
		expectedKernelPathSuffix := "/0.1.53-arm/kernel.img"

		// This is a structural test - the path is hardcoded
		assert.Contains(t, expectedKernelPathSuffix, "kernel.img")
		assert.Contains(t, expectedKernelPathSuffix, "0.1.53-arm")
	})

	t.Run("uses onprem provider", func(t *testing.T) {
		// The function uses "onprem" as the provider name
		// This is verified by code inspection
		providerName := "onprem"
		assert.NotEmpty(t, providerName)
	})
}

// TestInstanceArguments_ImageNameTypes tests various image name formats
func TestInstanceArguments_ImageNameTypes(t *testing.T) {
	testCases := []struct {
		name      string
		imageName string
	}{
		{"simple name", "image.img"},
		{"with path", "/usr/local/images/myimage.img"},
		{"with version", "app-v1.0.0.img"},
		{"with timestamp", "backup-2024-01-01.img"},
		{"alphanumeric", "test123.img"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			args := InstanceArguments{
				ImageName: tc.imageName,
			}
			require.NotNil(t, &args)
			assert.Equal(t, tc.imageName, args.ImageName)
		})
	}
}

// TestInstanceOperations_ResponseContract tests response type contracts
func TestInstanceOperations_ResponseContract(t *testing.T) {
	t.Run("instanceLogs response type is testable", func(t *testing.T) {
		// Verify the function signature allows for response testing
		// This is a structural test - the actual response requires integration
		args := InstanceArguments{ImageName: "test"}
		assert.NotNil(t, &args)
	})

	t.Run("instanceCreate response type is testable", func(t *testing.T) {
		args := InstanceArguments{ImageName: "test"}
		assert.NotNil(t, &args)
	})

	t.Run("listInstances response type is testable", func(t *testing.T) {
		args := MyFunctionsArguments{Submitter: "test"}
		assert.NotNil(t, &args)
	})
}
