package presto

import (
	"fmt"
	"testing"
)

func TestStringSlices(t *testing.T) {
	data := Data{
		{370125442440519680, 1508315181, 1508114292.000000},
		{370125442440519680, 1508315181, 1508114292.000000},
		{370125442440519680, 1508315181, 1508114292.000000},
		{370125442440519680, 1508315181, 1508114292.000000},
		{370125442440519680, 1508315181, 1508114292.000000},
		{370125442440519680, 1508315181, 1508114292.000000},
		{370125442440519680, 1508315181, 1508114292.000000},
		{370125442440519680, 1508315181, 1508114292.000000},
	}

	slices := data.StringSlices()
	for idx := range slices {
		fmt.Println(slices[idx])
	}
	t.Error("########")
}