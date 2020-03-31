package rand

import (
	"fmt"
	"math/rand"
	"testing"
)

func Test_Rand(t *testing.T) {
	n := uint64(100)
	cntR := make([]int, n)
	cnt := make([]int, n)
	for i := 0; i < 1000; i++ {
		cntR[int(Rand()%n)] ++
		cnt[int(rand.Uint64() % n)] ++
	}

	for i := uint64(0); i < n; i++ {
		fmt.Println(i, cntR[i], cnt[i])
	}
}
