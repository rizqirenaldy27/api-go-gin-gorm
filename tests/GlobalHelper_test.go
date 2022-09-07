package tests

import (
	"testing"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/rizqirenaldy/api-go-gin-gorm/helpers"
	"gopkg.in/go-playground/assert.v1"
)

func TestLengthofArray(t *testing.T) {
	var friend = []string{"john", "andy", "grace", "joe"}

	result := helpers.LengthofArray(friend)
	assert.Equal(t, 4, result)
}

func TestContains(t *testing.T) {
	var friend = []string{"john", "andy", "grace", "joe"}
	friendSelect := "john"

	result := helpers.Contains(friend, friendSelect)
	assert.Equal(t, true, result)
}

func BenchmarkLengthofArray(b *testing.B) {
	var friend = []string{"john", "andy", "grace", "joe"}

	for i := 0; i < b.N; i++ {
		helpers.LengthofArray(friend)
	}
}

func BenchmarkContains(b *testing.B) {
	var friend = []string{"john", "andy", "grace", "joe"}
	friendSelect := "john"

	for i := 0; i < b.N; i++ {
		helpers.Contains(friend, friendSelect)
	}
}
