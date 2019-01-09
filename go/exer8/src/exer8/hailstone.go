package exer8

// TODO: your Hailstone, HailstoneSequenceAppend, HailstoneSequenceAllocate functions

func Hailstone(n uint) uint {
	if n%2 == 0 {
		return n / 2
	} else {
		return 3*n + 1
	}

}

func HailstoneSequenceAppend(n uint) []uint {
	var h1 []uint
	h1 = append(h1, n)
	for n != 1 {
		n = Hailstone(n)
		h1 = append(h1, n)
	}

	return h1
}

func HailstoneSequenceAllocate(n uint) []uint {
	length := 1
	n2 := n
	for n2 != 1 {
		n2 = Hailstone(n2)
		length++
	}
	slice := make([]uint, length)
	for i := 0; i < length; i++ {
		slice[i] = n
		n = Hailstone(n)
	}
	return slice
}
