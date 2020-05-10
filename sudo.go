package jk_util

import "strconv"

func SudoCalculation(sd *[][]int){
	var zeroList [][2]int = make([][2]int, 0)
	for i := range (*sd) {
		for j := range (*sd)[i]{
			if (*sd)[i][j]==0 {
				zeroList = append(zeroList, [2]int{i,j})
			}
		}
	}
	var minMap = make(map[string]int)
	for a,lens:=0,len(zeroList);a<lens;a++ {
		i,j:= zeroList[a][0], zeroList[a][1]
		var x=1
		for x=1;x<10;x++ {
			iStr, jStr :=strconv.Itoa(i),strconv.Itoa(j)
			min := minMap[iStr+","+jStr]
			if check(sd,i,j,x,min){
				(*sd)[i][j] = x
				minMap[iStr+","+jStr] = x
				break
			}
		}
		if x==10 {
			for a_:=a;a_<lens;a_++ {
				i_,j_ := zeroList[a_][0], zeroList[a_][1]
				(*sd)[i_][j_] = 0;
				iStr, jStr :=strconv.Itoa(i_),strconv.Itoa(j_)
				minMap[iStr+","+jStr]=0
			}
			a=a-2;
		}
	}
}
func check(sd *[][]int,i int,j int,x int,min int) bool{
	if x<min {
		return false
	}
	for j_ := range (*sd)[i] {
		if (*sd)[i][j_] == x {
			return false
		}
	}
	for i_ := range (*sd) {
		if (*sd)[i_][j] == x {
			return false
		}
	}
	a,b := i/3,j/3
	for m:=a*3;m<3+a*3;m++ {
		for n:=b*3;n<3+b*3;n++ {
			if (*sd)[m][n]==x {
				return false
			}
		}
	}
	return true
}
