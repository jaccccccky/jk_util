package jk_util

import (
	"errors"
	"strconv"
)

/*SudoCalculation 计算并填写矩阵*/
func SudoCalculation(sd *[][]int) error {
	var zeroList [][2]int = make([][2]int, 0)
	for i := range *sd {
		for j := range (*sd)[i] {
			if (*sd)[i][j] == 0 {
				zeroList = append(zeroList, [2]int{i, j})
			} else {
				if !checkInit(sd, i, j, (*sd)[i][j]) {
					return errors.New("数独题目不合法！")
				}
			}
		}
	}
	var minMap = make(map[string]int)
	for a, lens := 0, len(zeroList); a < lens; a++ {
		i, j := zeroList[a][0], zeroList[a][1]
		var x = 1
		for x = 1; x < 10; x++ {
			iStr, jStr := strconv.Itoa(i), strconv.Itoa(j)
			min := minMap[iStr+","+jStr]
			if check(sd, i, j, x, min) {
				(*sd)[i][j] = x
				minMap[iStr+","+jStr] = x
				break
			}
		}
		if x == 10 {
			for a_ := a; a_ < lens; a_++ {
				i_, j_ := zeroList[a_][0], zeroList[a_][1]
				(*sd)[i_][j_] = 0
				iStr, jStr := strconv.Itoa(i_), strconv.Itoa(j_)
				minMap[iStr+","+jStr] = 0
			}
			a = a - 2
		}
	}
	return nil
}

/*check 检查单元格合法性*/
func checkInit(sd *[][]int, i int, j int, x int) bool {
	if x == 0 {
		return true
	}
	for j_ := range (*sd)[i] {
		if j_ != j {
			if (*sd)[i][j_] == x {
				return false
			}
		}
	}
	for i_ := range *sd {
		if i_ != i {
			if (*sd)[i_][j] == x {
				return false
			}
		}
	}
	a, b := i/3, j/3
	for m := a * 3; m < 3+a*3; m++ {
		for n := b * 3; n < 3+b*3; n++ {
			if m != i || n != j {
				if (*sd)[m][n] == x {
					return false
				}
			}
		}
	}
	return true
}

/*check 检查单元格填写的合法性*/
func check(sd *[][]int, i int, j int, x int, min int) bool {
	if x < min {
		return false
	}
	if (*sd)[i][j] == x {
		return false
	}
	return checkInit(sd, i, j, x)
}
