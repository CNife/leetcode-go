package asteroid_collision

func AsteroidCollision(asteroids []int) []int {
	var stack []int
	for _, asteroid := range asteroids {
		asteroidExists := true
		for asteroidExists && len(stack) > 0 && stack[len(stack)-1] > 0 && asteroid < 0 {
			last, absAsteroid := stack[len(stack)-1], -asteroid
			if last <= absAsteroid {
				stack = stack[:len(stack)-1]
			}
			if last >= absAsteroid {
				asteroidExists = false
			}
		}
		if asteroidExists {
			stack = append(stack, asteroid)
		}
	}
	return stack
}
