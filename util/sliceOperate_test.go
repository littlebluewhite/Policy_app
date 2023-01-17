package util

import (
	"github.com/stretchr/testify/require"
	"log"
	"testing"
)

func TestSliceRemove(t *testing.T) {
	t.Run("test1", func(t *testing.T) {
		a := []int{1, 2, 3, 4}
		b := SliceRemove(a, []int{1, 3})
		log.Println(b)
		require.ElementsMatch(t, b, []int{1, 3})
	})

	t.Run("test2", func(t *testing.T) {
		a := []string{"a", "b", "c", "d", "e", "f"}
		b := SliceRemove(a, []int{1, 2, 4})
		require.ElementsMatch(t, b, []string{"a", "d", "f"})
	})
}
