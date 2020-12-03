package day3

// Traverse the map line by line and returns a true if a tree was encountered.
func traverseMap(line string, location int) bool {
	length := len(line)
	if string(line[location%length]) == "#" {
		return true
	}

	return false
}
