package main

//// A utility function that returns
//// true if there is a subset of
//// arr[] with sum equal to given sum
//func isSubsetSum(arr []int, n int, sum int) bool {
//	// Base Cases
//	if sum == 0 {
//		return true
//	}
//	if n == 0 && sum != 0 {
//		return false
//	}
//
//	// If last element is greater than sum, then
//	// ignore it
//	if arr[n-1] > sum {
//		return isSubsetSum(arr, n-1, sum)
//	}
//
//	/* else, check if sum can be obtained by any of
//	the following
//	(a) including the last element
//	(b) excluding the last element*/
//
//	return isSubsetSum(arr, n-1, sum) || isSubsetSum(arr, n-1, sum-arr[n-1])
//}
//
//// Returns true if arr[] can be partitioned in two
//// subsets of equal sum, otherwise false
//func findPartion(arr []int, n int) bool {
//	// Calculate sum of the elements in array
//	sum := 0
//	for i := 0; i < n; i++ {
//		sum += arr[i]
//	}
//	// If sum is odd, there cannot be two subsets
//	// with equal sum
//	if sum%2 != 0 {
//		return false
//	}
//
//	// Find if there is subset with sum equal to
//	// half of total sum
//	return isSubsetSum(arr, n, sum/2)
//}

//func main() {
//	arr := []int{3, 1, 5, 9, 12}
//	n := len(arr)
//
//	// Function call
//	if findPartion(arr, n) == true {
//		fmt.Println("Can be divided into two subsets of equal sum")
//	} else {
//		fmt.Println("Can not be divided into two subsets of equal sum")
//	}
//}
