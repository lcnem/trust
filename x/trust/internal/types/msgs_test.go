package types

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestValidateTopicID(t *testing.T) {
	require.Equal(t, validateTopicID("a"), true)
	require.Equal(t, validateTopicID("ab"), true)
	require.Equal(t, validateTopicID("a-"), true)
	require.Equal(t, validateTopicID("-a"), false)
	require.Equal(t, validateTopicID("0"), false)
}
