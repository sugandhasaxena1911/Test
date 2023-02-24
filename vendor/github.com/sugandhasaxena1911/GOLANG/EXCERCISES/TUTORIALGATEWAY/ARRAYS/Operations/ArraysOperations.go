// Package Operations provides all the operations that we can do in Arrays
package Operations

import "fmt"

// Add2DArray Add the correspoding elements of both arrays and return the new array
func Add2DArray(arr1, arr2 [][]int) [][]int {
	arr3 := [][]int{}
	for i := 0; i < len(arr1); i++ {
		temp := []int{}
		for j := 0; j < len(arr1[i]); j++ {
			temp = append(temp, arr1[i][j]+arr2[i][j])
		}
		arr3 = append(arr3, temp)
	}
	return arr3
}

// Print2DArray Prints the 2d array
func Print2DArray(arr1 [][]int) {
	for i, _ := range arr1 {
		for _, vv := range arr1[i] {
			fmt.Print(vv, " ")

		}
		fmt.Println(" ")

	}
}

// CountDuplicates Counts he duplicates in the slice
func CountDuplicates(sl []int) int {

	m1 := make(map[int]int, len(sl))
	for _, v := range sl {
		m1[v] = v
	}
	return len(sl) - len(m1)
}

// EnterMatrix Enter 2d array of mentioned rows and columns
func EnterMatrix(n, m int) [][]int {
	var num int

	arr1 := [][]int{}
	for x := 0; x < n; x++ {
		arr2 := []int{}
		for y := 0; y < m; y++ {
			fmt.Print("Enter number")
			fmt.Scan(&num)
			arr2 = append(arr2, num)
		}
		arr1 = append(arr1, arr2)
	}

	return arr1

}

// SumOfEachColumnRow sum of each column and row of 2d array
func SumOfEachColumnRow(arr1 [][]int) {
	for i, _ := range arr1 {
		sumRow := 0
		for _, vv := range arr1[i] {
			sumRow = sumRow + vv
		}
		fmt.Println("Sum of Row", i, "is : ", sumRow)
	}

	for y := 0; y < len(arr1[0]); y++ {
		sumColumn := 0
		for x := 0; x < len(arr1); x++ {
			sumColumn = sumColumn + arr1[x][y]
		}
		fmt.Println("Sum of column", y, "is : ", sumColumn)

	}
}

// IsSymmetric finds if the 2d arrya is symmetric
func IsSymmetric(arr1 [][]int) bool {
	var b bool
	b = true
	arr2 := TransposeMatrix(arr1)
	fmt.Println(arr1)
	fmt.Println("Transpose of matrix ")
	fmt.Println(arr2)
	r1 := len(arr1)
	c1 := len(arr1[0])

	if r1 != c1 {
		b = false
		return b
	}
	for x := 0; x < r1; x++ {
		for y := 0; y < c1; y++ {
			if arr1[x][y] != arr2[x][y] {
				b = false
				return b
			}
		}
	}

	return b

}

// TransposeMatrix Transpose the matrix ie row to columns and vice versa
func TransposeMatrix(arr1 [][]int) [][]int {
	// calculate rows and columns of new matrix
	rows := len(arr1[0])
	cols := len(arr1)

	arr2 := [][]int{}
	for r := 0; r < rows; r++ { //3
		arr3 := []int{}
		for c := 0; c < cols; c++ { //2
			arr3 = append(arr3, arr1[c][r])
		}
		arr2 = append(arr2, arr3)
	}

	return arr2
}

// CheckIdentity Finds if the 2d array is identity or not
func CheckIdentity(arr1 [][]int) bool {
	//var identity bool
	identity := true
	//size of square matrix
	n := len(arr1)
	for x := 0; x < n; x++ {

		for y := 0; y < n; y++ {
			if y == x {
				if arr1[x][y] != 1 {
					identity = false
					break
				}
			} else {
				if arr1[x][y] != 0 {
					identity = false
					break
				}
			}

		}

	}
	return identity

}

// InterchangeDiagnols Interchange the diagnols of the 2d array and returns new array
func InterchangeDiagnols(arr1 [][]int) [][]int {
	//diag1 : 0,0  1,1   2,2 ..... diag2 : 0,n-1   1,n-1 ....
	n := len(arr1)

	for x := 0; x < n; x++ {

		arr1[x][x], arr1[x][n-x-1] = arr1[x][n-x-1], arr1[x][x]
	}

	return arr1

}

// LowerTri Prints lower triangle of array
func LowerTri(arr1 [][]int) [][]int {
	n := len(arr1)
	for x := 0; x < n; x++ {
		for y := 0; y < n; y++ {
			if y > x {
				arr1[x][y] = 0
			}
		}
	}

	return arr1
}

// UpperTri prints upper triangle of arrya
func UpperTri(arr1 [][]int) [][]int {
	n := len(arr1)
	for x := 0; x < n; x++ {
		for y := 0; y < n; y++ {
			if y < x {
				arr1[x][y] = 0
			}
		}
	}

	return arr1
}

// MinMax Print min ax of slice
func MinMax(sl []int) (int, int) {
	min := sl[0]
	max := sl[0]
	for _, v := range sl {
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}

	}
	return min, max
}

// ReverseArray reverse the slice
func ReverseArray(sl []int) []int {
	m := len(sl)
	for i, _ := range sl {
		if i >= m/2 {
			break
		}
		sl[i], sl[m-i-1] = sl[m-i-1], sl[i]

	}

	return sl
}
