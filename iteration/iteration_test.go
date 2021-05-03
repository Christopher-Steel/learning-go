package iteration

import "fmt"
import "testing"

func TestRepeat(t *testing.T) {
	actual := Repeat("a", 7)
	expected := "aaaaaaa"

	if actual != expected {
		t.Errorf("expected %q but got %q", expected, actual)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}

func ExampleRepeat() {
	fmt.Println(Repeat("daddy", 3))
	// Output: daddydaddydaddy
}