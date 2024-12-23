package module3

func UniquePathsWithObstacles(obstacleGrid [][]int) int {
	m, n := len(obstacleGrid), len(obstacleGrid[0])

	if obstacleGrid[0][0] == 1 {
		return 0
	}

	hasObstacle := false
	for i := range m {
		if hasObstacle || obstacleGrid[i][0] == 1 {
			obstacleGrid[i][0] = 0
			hasObstacle = true
		} else {
			obstacleGrid[i][0] = 1
		}
	}
	hasObstacle = false
	for i := range n {
		if i == 0 {
			continue
		}
		if hasObstacle || obstacleGrid[0][i] == 1 {
			obstacleGrid[0][i] = 0
			hasObstacle = true
		} else {
			obstacleGrid[0][i] = 1
		}
	}

	i := 1
	for i < m {
		j := 1
		for j < n {
			if obstacleGrid[i][j] == 1 {
				obstacleGrid[i][j] = 0
			} else {
				obstacleGrid[i][j] = obstacleGrid[i][j-1] + obstacleGrid[i-1][j]
			}
			j += 1
		}
		i += 1
	}

	return obstacleGrid[m-1][n-1]
}
