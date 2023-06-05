package leetcode

func twoSum(nums []int, target int) []int {
	// irritating over both slices
	// checking for the target

	// 0(n2) --> brute force solution
	for i, first := range nums {
		for j, second := range nums {
			if first+second == target && i != j {
				return []int{i, j}
			}
		}
	}
	// just return nothin
	return nil
}

// func twoSum(nums []int, target int) []int {
// 	m := make(map[int]int)
// 	for k, v := range nums {
// 		if idx, ok := m[target-v]; ok {
// 			return []int{idx, k}
// 		}
// 		m[v] = k
// 	}
// 	return nil
// }
