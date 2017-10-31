package presto

import (
	"fmt"
	"sync"
)

type Data [][]interface{}

func (this Data) StringSlices() [][]string {
	data := make([][]string, len(this))
	var wg sync.WaitGroup
	for idx := range this {
		wg.Add(1)
		go func(index int) {
			defer wg.Done()
			row := make([]string, 0, len(this[idx]))
			for _, field := range this[idx] {
				if field == nil {
					row = append(row, "")
					continue
				}
				row = append(row, fmt.Sprint(field))
			}
			data[index] = row
		}(idx)
	}
	wg.Wait()
	return data
}
