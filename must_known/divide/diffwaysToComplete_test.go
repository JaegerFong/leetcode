package divide

import (
	"testing"
)

func Test_diffWaysToCompute(t *testing.T) {
	exp := "2*3-4*5"
	res := diffWaysToCompute(exp)
	t.Log(res)
}
