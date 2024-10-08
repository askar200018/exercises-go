package problem1

func isChangeEnough(changes [4]int, total float32) bool {
	current := float32(0)

	current += float32(changes[0]) * 0.25
	current += float32(changes[1]) * 0.10
	current += float32(changes[2]) * 0.05
	current += float32(changes[3]) * 0.01

	return current >= total
}
