package jk_util

var isAsc = true

/*Sort 采用归并排序,在原数组上进行排序，true为升序，false为降序*/
func Sort(arr *[]int, isAsc_ bool) int {
	isAsc = isAsc_
	l := len(*arr)
	var temp = make([]int, l)
	return sortGroup(arr, &temp, 0, l-1)
}
func sortGroup(arr, temp *[]int, left, right int) int {
	if left < right {
		mid := (left + right) / 2
		a := sortGroup(arr, temp, left, mid)
		b := sortGroup(arr, temp, mid+1, right)
		c := merge(arr, temp, left, mid, right)
		//fmt.Println(a,b,c)
		return a + b + c
	} else {
		return 0
	}
}
func merge(arr, temp *[]int, left, mid, right int) (cnt int) {
	i := left
	j := mid + 1
	t := 0
	if (isAsc && (*arr)[mid] <= (*arr)[j]) || (!isAsc && (*arr)[mid] >= (*arr)[j]) {
		return
	}
	for ; i <= mid && j <= right; t++ {
		if (isAsc && (*arr)[i] <= (*arr)[j]) || (!isAsc && (*arr)[i] >= (*arr)[j]) {
			(*temp)[t] = (*arr)[i]
			i++
		} else {
			(*temp)[t] = (*arr)[j]
			j++
			cnt += mid + 1 - i
		}
	}
	for ; i <= mid; i++ {
		(*temp)[t] = (*arr)[i]
		t++
	}
	for ; j <= right; j++ {
		(*temp)[t] = (*arr)[j]
		t++
	}
	for t = 0; t+left <= right; t++ {
		(*arr)[t+left] = (*temp)[t]
	}
	return
}

/*Sort 采用归并排序,返回排序后的数组，原数组不变，true为升序，false为降序*/
func Sort2(arr *[]int, isAsc_ bool) ([]int, int) {
	isAsc = isAsc_
	return sortGroup2(arr)
}
func sortGroup2(arr *[]int) ([]int, int) {
	length := len(*arr)
	if length <= 1 {
		return *arr, 0
	}
	num := length / 2
	arrLeft := (*arr)[:num]
	arrRight := (*arr)[num:]
	left, a := sortGroup2(&arrLeft)
	right, b := sortGroup2(&arrRight)
	result, c := merge2(&left, &right)
	return result, a + b + c
}
func merge2(left, right *[]int) (result []int, cnt int) {
	l, r := 0, 0
	lenLeft, lenRight := len(*left), len(*right)
	if (isAsc && (*left)[lenLeft-1] <= (*right)[r]) || (!isAsc && (*left)[lenLeft-1] >= (*right)[r]) {
		result = append(result, (*left)[l:]...)
		result = append(result, (*right)[r:]...)
		return
	}
	for l < lenLeft && r < lenRight {
		if (isAsc && (*left)[l] <= (*right)[r]) || (!isAsc && (*left)[l] >= (*right)[r]) {
			result = append(result, (*left)[l])
			l++
		} else {
			result = append(result, (*right)[r])
			r++
			cnt += lenLeft - l
		}
	}
	result = append(result, (*left)[l:]...)
	result = append(result, (*right)[r:]...)
	return
}
