// 48. 旋转图像
// 反转法
// 90度旋转 = 水平反转 + 主对角线反转(左上角到右下角)
func rotate(matrix [][]int)  {
    n := len(matrix)
    // 1 水平反转 [i, j] i / 2
    for i := 0; i < n / 2; i++ {
        for j := 0; j < n; j++ {
            matrix[i][j], matrix[n-1-i][j] = matrix[n-1-i][j], matrix[i][j]
        }
    }
    // 2. 主对角线反转
    for i := 0; i < n; i++ {
        for j := 0; j < i; j++ {
            matrix[i][j], matrix[j][i] = matrix[j][i],matrix[i][j]
        }
    }
}

// 180度旋转 = 水平反转 + 垂直反转
// 270度旋转 = 垂直旋转 + 主对角线反转
// 副对角线反转(左下角到右上角)
