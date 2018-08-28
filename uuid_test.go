package uuid_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/tkitsunai/uuid"
)

func TestUUID_NoHyphen(t *testing.T) {
	test, _ := uuid.NewUUID()
	if strings.Index(test.NoHyphen(), "-") > 0 {
		t.Fatal("fail ", strings.Index(test.NoHyphen(), "-"))
	}
	fmt.Println(test.String())
}

func TestUUID_WithHyphen(t *testing.T) {
	test, _ := uuid.NewUUID()
	if strings.Index(test.String(), "-") <= 0 {
		t.Fatal("fail ", strings.Index(test.String(), "-"))
	}
	fmt.Println(test.String())
}

func BenchmarkUUID_NoHyphen(b *testing.B) {
	for i := 0; i < b.N; i++ {
		test, _ := uuid.NewUUID()
		test.NoHyphen()
	}
}
