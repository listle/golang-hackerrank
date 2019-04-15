// Complete the countingValleys function below.
func countingValleys(n int32, s string) int32 {

	var altitude, countvalleys int32
	altitude, countvalleys = 0, 0

	for _, c := range s {
		i := string(c)

		if i == "U" {
			altitude++
		} else {
			if altitude == 0 {
				countvalleys++
			}
			altitude--
		}

	}
	return countvalleys
}
