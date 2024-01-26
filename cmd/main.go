package main

import (
	"fmt"
)

var blocks = []map[string]bool{
	{
		"school": true,
		"gym":    false,
		"store":  false,
	},
	{
		"school": false,
		"gym":    true,
		"store":  false,
	},
	{
		"school": true,
		"gym":    true,
		"store":  false,
	},
	{
		"school": true,
		"gym":    false,
		"store":  true,
	},
	{
		"school": true,
		"gym":    false,
		"store":  false,
	}}

func searchInfostarctual(name string, blocks []map[string]bool, index int) int {
	length := len(blocks)
	var rightResult, leftResult = length, length
	for i := index; i < length; i++ {
		if blocks[i][name] == true {
			rightResult = i - index
			break
		}
	}
	for i := index; i > 0; i-- {
		if blocks[i][name] == true {
			leftResult = index - i
			break
		}
	}
	if rightResult > leftResult {
		return leftResult
	}
	return rightResult
}

func main() {
	result := []int{}
	for index, kvartal := range blocks {
		infostractual := []string{"school", "store", "gym"}
		maxDistance := 0
		for _, nameOfIfostarctula := range infostractual {
			value := kvartal[nameOfIfostarctula]
			var resultDistance int
			if value == false {
				resultDistance = searchInfostarctual(nameOfIfostarctula, blocks, index)
			}
			if maxDistance < resultDistance {
				maxDistance = resultDistance
			}
		}
		result = append(result, maxDistance)
	}
	var min, resultIndex int = len(result), 0
	for index, resultDistance := range result {
		if min > resultDistance {
			min = resultDistance
			resultIndex = index
		}
	}
	fmt.Printf("Result index kvartal %d \n", resultIndex)
}
