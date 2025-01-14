package util_test

import (
	"testing"

	"github.com/singlestore-labs/singlestore-go/management"
	"github.com/singlestore-labs/terraform-provider-singlestoredb/internal/provider/util"
	"github.com/stretchr/testify/require"
)

func TestDerefInt(t *testing.T) {
	var i *int
	result := util.Deref(i)
	require.Equal(t, 0, result)

	value := 10
	i = &value
	result = util.Deref(i)
	require.Equal(t, value, result)
}

func TestDerefSlice(t *testing.T) {
	var s *[]int
	result := util.Deref(s)
	require.Equal(t, []int(nil), result)

	value := []int{1, 2, 3}
	s = &value

	result = util.Deref(s)
	require.Equal(t, value, result)
}

func TestJoin(t *testing.T) {
	result := util.Join([]management.WorkspaceState{management.WorkspaceStateACTIVE, management.WorkspaceStateSUSPENDED}, ", ")
	require.Equal(t, result, "ACTIVE, SUSPENDED")
}
