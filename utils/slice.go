package utils

// DeleteVal nums: old slice, val: the value to delete, return: new slice
func DeleteVal(nums []string, val string) []string {
	// create a new slice
	var newNums []string
	for _, v := range nums {
		// if the value is not equal to val, append it to the new slice
		if v != val {
			newNums = append(newNums, v)
		}
	}
	return newNums
}
