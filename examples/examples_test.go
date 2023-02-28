package examples_test

import (
	"testing"

	"github.com/singlestore-labs/terraform-provider-singlestore/examples"
	"github.com/stretchr/testify/require"
)

func TestEmbedsExamples(t *testing.T) {
	require.NotEmpty(t, examples.Provider)
	require.NotEmpty(t, examples.Regions)
}