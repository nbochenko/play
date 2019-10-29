package main

import (
	_ "github.com/Propertyfinder/pf-go-common/v4/pkg/log"
	"github.com/sirupsen/logrus"
)

func main() {
	array := []int{9, 2, 1, 2, 3, 4, 5}
	logrus.Println(array)

	res := reverseArray(array)
	logrus.Println(res)

	array = []int{9, 2, 1, 2, 3, 4, 5}
	res2 := findXMinInArray(array, 3)
	logrus.Println(res2)
}

func reverseArray(array []int) []int {
	reverseIndex := len(array)
	var temp int
	for k, v := range array {
		if len(array)/2 < k {
			logrus.Println(k)
			break
		}
		reverseIndex--
		temp = array[reverseIndex]
		array[reverseIndex] = v
		array[k] = temp
	}

	return array
}

func findXMinInArray(array []int, x int) []int {
	selected := make([]int, x)

	for index := 0; index < x; index++ {
		min, key := findMin(array)
		selected[index] = min
		copy(array[key:], array[key+1:])
		array = array[:len(array)-1]
	}

	return selected
}

func findMin(array []int) (int, int) {
	var min int = array[0]
	var key int
	for k, v := range array {
		if v < min {
			min = v
			key = k
		}
	}

	return min, key
}
