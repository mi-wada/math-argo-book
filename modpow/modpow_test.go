package modpow_test

import (
	"fmt"
	"testing"

	"github.com/mi-wada/math-argo-book/modpow"
)

func TestModpow(t *testing.T) {
	t.Parallel()
	for _, tt := range []struct {
		a, b, mod, want int
	}{
		{2, 2, 5, 4},
		{2, 2, 4, 0},
	} {
		t.Run(fmt.Sprintf("Modpow(%v, %v, %v) = %v", tt.a, tt.b, tt.mod, tt.want), func(t *testing.T) {
			t.Parallel()
			got := modpow.Modpow(tt.a, tt.b, tt.mod)
			if got != tt.want {
				t.Errorf(
					"Modpow(%v, %v, %v) = %v, want %v",
					tt.a, tt.b, tt.mod, got, tt.want,
				)
			}
		})
	}
}
