package ContainerWithMostWater

func maxArea(height []int) int {
	length := len(height)
	max := 0
	start := 0
	end := length - 1
	tempArea := 0
	minL := 0
	for start < end {
		if height[start] < height[end] {
			minL = height[start]
			tempArea = minL * (end - start)
			start++
		} else {
			minL = height[end]
			tempArea = minL * (end - start)
			end--
		}

		if tempArea > max {
			max = tempArea
		}
	}

	return max
}
