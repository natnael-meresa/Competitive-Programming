package main

func countingValleys(steps int32, path string) int32 {
	level := 0
	underSealevel := false
	count := 0
	for _, step := range path {

		if string(step) == "D" {
			level--
		} else if string(step) == "U" {
			level++
		}

		if level < 0 && !underSealevel {
			underSealevel = true
			count++
		}
		if level >= 0 && underSealevel {
			underSealevel = false
		}
	}

	return int32(count)

}
