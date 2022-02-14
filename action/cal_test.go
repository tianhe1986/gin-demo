package action

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	t.Run("pos", func(t *testing.T) {
		assert.Equal(t, 3, Add(1, 2))
	})
	t.Run("neg", func(t *testing.T) {
		assert.Equal(t, -1, Add(2, -3))
	})
}

func TestSub(t *testing.T) {
	t.Run("pos", func(t *testing.T) {
		assert.Equal(t, 13, Sub(52, 39))
	})
	t.Run("neg", func(t *testing.T) {
		assert.Equal(t, -154, Sub(-12, 142))
	})
}

func TestMul(t *testing.T) {
	// 使用t.Errorf报错
	t.Run("pos", func(t *testing.T) {
		if ans := Mul(2, 3); ans != 6 {
			t.Errorf("2 * 3 expected be 6, but %d got", ans)
		}
	})
	t.Run("neg", func(t *testing.T) {
		if ans := Mul(2, -3); ans != -6 {
			t.Errorf("2 * -3 expected be -6, but %d got", ans)
		}
	})
}

func TestDiv(t *testing.T) {
	t.Run("err", func(t *testing.T) {
		ans, err := Div(6, 0)
		assert.NotNil(t, err)
		assert.Equal(t, 0, ans)
		assert.Equal(t, "math: dividend can not be zero", err.Error())
	})
	t.Run("pos", func(t *testing.T) {
		ans, err := Div(12, 4)
		assert.Nil(t, err)
		assert.Equal(t, 3, ans)
	})
	t.Run("neg", func(t *testing.T) {
		ans, err := Div(-32, 4)
		assert.Nil(t, err)
		assert.Equal(t, -8, ans)
	})
}

func TestSum(t *testing.T) {
	nums := []int{2, 4, -5}
	ans := Sum(nums)
	assert.Equal(t, 1, ans)

	nums = []int{}
	ans = Sum(nums)
	assert.Equal(t, 0, ans)
}
