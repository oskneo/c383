package exer9

import (
	"math"
	"math/rand"
	"time"
)

type Chunk struct {
	x_sums  float64
	x2_sums float64
}

func RandomArray(length int, maxInt int) []int {
	// TODO: create a new random generator with a decent seed; create an array with length values from 0 to values-1.
	s := rand.NewSource(time.Now().UnixNano())
	rd := rand.New(s)
	slice := make([]int, length)
	for i := 0; i < length; i++ {
		slice[i] = rd.Intn(maxInt)
	}
	return slice
}

func MSgoroutine(arr []int, index int, chunkarray chan Chunk) {
	x := 0.0
	x2 := 0.0
	for i := 0; i < len(arr); i++ {
		n := float64(arr[i])
		x += n
		x2 += n * n
	}
	chunkarray <- Chunk{x, x2}

}

func MeanStddev(arr []int, chunks int) (mean, stddev float64) {
	if len(arr)%chunks != 0 {
		panic("You promised that chunks would divide slice size!")
	}
	// TODO: calculate the mean and population standard deviation of the array, breaking the array into chunks segments
	// and calculating on them in parallel.
	xx2 := make(chan Chunk, chunks)
	j := 0
	for i := 0; i < len(arr); i += len(arr) / chunks {
		go MSgoroutine(arr[i:(i+len(arr)/chunks)], j, xx2)
		j++
	}
	x := 0.0
	x2 := 0.0

	for i := 0; i < chunks; i++ {
		ck := <-xx2
		x += ck.x_sums
		x2 += ck.x2_sums
	}
	mean = x / float64(len(arr))
	stddev = math.Sqrt(x2/float64(len(arr)) - mean*mean)
	return mean, stddev

}
