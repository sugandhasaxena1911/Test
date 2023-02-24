// Package sugandha this is my package sugandha
package hello

import (
	"fmt"
	"github.com/sugandhasaxena1911/GOLANG/EXCERCISES/TUTORIALGATEWAY/ARRAYS/Operations"
	"github.com/sugandhasaxena1911/GOLANG/EXCERCISES/TUTORIALGATEWAY/NUMBERS/MathOperations"
)

func Suggu() {
	fmt.Println("hello sugandha ")
	sl := []int{1, 2, 5, 6, 2, 8, 5, 9, 1}
	dupCnt := Operations.CountDuplicates(sl)

	fmt.Println("the number of duplicates were ", dupCnt)
	n := 567
	m := MathOperations.NoOfDigits(n)
	fmt.Println(m)

}
