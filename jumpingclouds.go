import "fmt"

// Complete the jumpingOnClouds function below.
func jumpingOnClouds(c []int32) int32 {

	var numJumps int32 = 0
	currentIndex := 0
	maxJump := 0

	for {
		fmt.Println(currentIndex)
		if currentIndex == (len(c) - 1) {
			break
		}
		currentIndex++
		numJumps++

		//try max jump
		maxJump = currentIndex + 1
		if maxJump < len(c) {
			if c[maxJump] != 1 {
				currentIndex = maxJump
			}
		}
	}

	return numJumps
}