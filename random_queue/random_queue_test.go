package random_queue

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Testほげ(t *testing.T) {
	assert.Equal(t, 1, 1)
}

// Queueのインターフェース
// add
// remove
// 償却実行時間は一定であること

func TestAddができる(t *testing.T) {

	sut := RandomQueue[int]{}
	sut.Add(1000)
	assert.Equal(t, 1, sut.Size())
}

func TestSize確認ができる(t *testing.T) {

	sut := RandomQueue[int]{
		a: []int{1, 2, 3},
	}
	assert.Equal(t, 3, sut.Size())
}
