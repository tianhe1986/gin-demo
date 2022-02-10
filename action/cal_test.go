package action

import "testing"

func TestAdd(t *testing.T) {
	t.Run("pos", func(t *testing.T) {
		if ans := Add(1, 2); ans != 3 {
			t.Errorf("1 + 2 expected be 3, but %d got", ans)
		}

	})
	t.Run("neg", func(t *testing.T) {
		if ans := Add(2, -3); ans != -1 {
			t.Errorf("2 + -3 expected be -1, but %d got", ans)
		}
	})
}

func TestMul(t *testing.T) {
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
