package helpers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEnvWithDefault(t *testing.T) {

	t.Run("success get value", func(t *testing.T) {
		key := "test key"
		want := "test value"
		t.Setenv(key, want)

		got := EnvWithDefault(key, "default")
		assert.Equal(t, got, want)

	})
	t.Run("success get default", func(t *testing.T) {
		key := "nonexistent value"
		want := "default"

		got := EnvWithDefault(key, "default")
		assert.Equal(t, got, want)

	})

}
