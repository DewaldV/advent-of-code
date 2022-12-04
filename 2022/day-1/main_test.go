package main

import (
	"testing"
)

func TestMain(t *testing.T) {
	elfCalorieList := `1000
2000
3000

4000

5000
6000

7000
8000
9000

10000
`

	calories, err := mostCalories(elfCalorieList)
	if err != nil {
		t.Errorf("error converting calorie list, %s", err)
	}

	if calories != 24000 {
		t.Errorf("most calories incorrect, got: %d, expected: 24000", calories)
	}
}
