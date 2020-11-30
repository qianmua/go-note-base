package alg_leetcode

/**
 * @author
 * @date 2020/11/30 15:12
 * create by Goland
 * @version 1.0
 */ 
 
// 给定一个矩阵，然后找到所有含有 0 的地方，把该位置所在行所在列的元素全部变成 0。

func serZeros(matrix [][]int){
	row , col := len(matrix) , len(matrix)

	var rowZero = make([]bool , row )
	var colZero =  make([]bool , col )

	// mark
	for i := 0 ; i < row ; i++ {
		for j :=0 ; j < col  ; j ++ {
			if matrix[i][j] == 0  {
				rowZero[i] = true
				colZero[j] = true
			}
		}
	}
	// row
	for i := 0; i < row ;i++  {
		if rowZero[i] {
			setRow(matrix , i)
		}
	}

	//col
	for i := 0; i < col ;i++  {
		if colZero[i] {
			setRow(matrix , i)
		}
	}



}

func setRow(matrix [][]int , r int){
	for i := 0; i < len(matrix[r]) ;i++  {
		matrix[r][i] = 0
	}
}
func serCol(matrix [][]int , c int){
	for i := 0; i < len(matrix) ;i++  {
		matrix[i][c] = 0
	}
}