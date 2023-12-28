package main

import "math"

func coinChange(coins []int, amount int) int {
	var result = make([]int, amount+1)
	for i := range result {
		result[i] = math.MaxInt - 1
	}
	result[0] = 0
	for i := 1; i < amount+1; i++ {
		for j := range coins {
			if i >= coins[j] {
				if result[i] > result[i-coins[j]]+1 {
					result[i] = result[i-coins[j]] + 1
				}
			}
		}
	}
	if result[amount] == math.MaxInt-1 {
		result[amount] = -1
	}
	return result[amount]
}
func main() {

}
