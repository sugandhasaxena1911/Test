package MathOperations

import "fmt"

func factorial(n int) int {
	var result int
	if n <= 1 {
		return 1
	}

	result = n * factorial(n-1)
	return result

}

func NoOfDigits(num int) int {
	n := num
	dg := 0
	for {
		n = n / 10
		dg = dg + 1
		if n <= 0 {
			break
		}

	}

	return dg
}

func PerfectNums(num int) []int {
	perfect := []int{}

	for x := 1; x <= num; x++ {
		// check if x is perfect
		sl := divisors(x)
		fmt.Println(x, sl)
		sum := 0
		for _, v := range sl {
			sum = sum + v
		}
		if sum == x {
			perfect = append(perfect, x)
		}
	}

	return perfect

}

func divisors(num int) []int {
	sl := []int{}
	for x := 1; x < num; x++ {
		if num%x == 0 {
			sl = append(sl, x)
		}

	}
	return sl
}
func prime(n int) []int {
	sl := []int{}
	for x := 1; x <= n; x++ {
		check := 1
		//find if x is prime
		for y := 2; y < x; y++ {
			if x%y == 0 {
				check = 0
				break
			}
		}
		if check == 1 {
			sl = append(sl, x)
		}
	}
	return sl
}
func PrintNum(start, end int) {
	fmt.Println(start)
	if start == end {
		return
	}
	PrintNum(start+1, end)

}
