package helpers

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReplaceCharacterAtIndexInString(t *testing.T) {

	t.Run("every other character", func(t *testing.T) {

		str := "abcdefghijklmnopqrstuvwxyz"
		indices := []int{1, 3, 5, 7, 9, 11, 13, 15, 17, 19, 21, 23, 25}

		expected := "a_c_e_g_i_k_m_o_q_s_u_w_y_"
		actual := ReplaceCharacterAtIndexInString(str, "_", indices)

		assert.Equal(t, expected, actual)
	})
}
