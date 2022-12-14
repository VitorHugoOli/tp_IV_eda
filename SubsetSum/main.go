package main

import "fmt"

func isSubsetSum(arr []int, n int, sum int, p1 *[]int, p2 *[]int) (bool, *[]int, *[]int) {

	if sum == 0 {
		return true, p1, p2
	}

	if n == 0 && sum != 0 {
		return false, p1, p2
	}

	if arr[n-1] > sum {
		return isSubsetSum(arr, n-1, sum, p1, p2)
	}

	*p1 = append(*p1, arr[n-1])
	include, p1Include, p2Include := isSubsetSum(arr, n-1, sum-arr[n-1], p1, p2)
	if include {
		return true, p1Include, p2Include
	}

	*p2 = append(*p2, arr[n-1])
	exclude, _, p2Exclude := isSubsetSum(arr, n-1, sum, p1, p2)
	if exclude {
		return true, p1, p2Exclude
	}

	return false, p1, p2
}

func findPartion(arr []int) (bool, *[]int, *[]int) {
	sum := 0

	//Caculando a soma dos elementos do array
	for _, v := range arr {
		sum += v
	}

	// Se a soma for impar, nao ha dois subconjuntos com soma igual
	if sum%2 != 0 {
		return false, nil, nil
	}

	return isSubsetSum(arr, len(arr), sum/2, &[]int{}, &[]int{})
}

func main() {
	arr := []int{3, 1, 5, 9, 12}

	// Function call
	if b, r1, r2 := findPartion(arr); b {
		fmt.Println("Can be divided into two subsets of equal sum", r1, r2)
	} else {
		fmt.Println("Can not be divided into two subsets of equal sum")
	}
}
