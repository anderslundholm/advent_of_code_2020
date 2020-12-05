package day5

func translateBinarySpacePartitionerSegment(partitionSegment string) int8 {
	var pos int8 = 0
	maxBitShift := len(partitionSegment) - 1
	for i, partition := range partitionSegment {
		// int32 "B" OR "R"
		if partition == 66 || partition == 82 {
			// Binary left shift from 1 when in the "higher half" of rows
			// starting with 6 or 2 = 64 or 4, then inclusive OR to count up the value
			pos |= 1 << (maxBitShift - i)
		}
	}
	return pos
}

func parseBoardingpass(boardingPass string) int {
	row := translateBinarySpacePartitionerSegment(boardingPass[:7])
	col := translateBinarySpacePartitionerSegment(boardingPass[7:])
	seatID := int(row)*8 + int(col)

	return seatID

}
