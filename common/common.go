package common

const A = "123"

func DifferenceSetOfIntSlice(A []int, B []int) []int {
	result := []int{}
	for _, value := range A {
		if !ExistDataInIntSlice(value, B) {
			result = append(result, value)
		}
	}
	return result
}
func ExistDataInIntSlice(data int, slice []int) bool {
	tmpMap := map[int]bool{}
	for _, MemberOfSlice := range slice {
		tmpMap[MemberOfSlice] = true
	}
	return tmpMap[data] == true
}
