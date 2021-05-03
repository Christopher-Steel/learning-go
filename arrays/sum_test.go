package arrays

import "reflect"
import "testing"

func TestSum(t *testing.T) {
	t.Run("variable size collection", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		actual := Sum(numbers)
		expected := 6
	
		if actual != expected {
			t.Errorf("got %d but expected %d", actual, expected)
		}
	})
}

func TestSumAll(t *testing.T) {
	actual := SumAll([]int{1, 2}, []int{3, 4, 5})
	expected := []int{3, 12}

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("got %v but expected %v", actual, expected)
	}
}