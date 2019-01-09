package exer10

func FibCon(n uint, cutoff uint, nc chan uint) {
	nc <- Fib(n, cutoff)
}

func Fib(n uint, cutoff uint) uint {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	}

	var fn1 uint
	var fn2 uint
	if n > cutoff {
		fn1c := make(chan uint)
		fn2c := make(chan uint)
		go FibCon(n-1, cutoff, fn1c)
		go FibCon(n-2, cutoff, fn2c)
		fn1 = <-fn1c
		fn2 = <-fn2c
	} else {
		fn1 = Fib(n-1, cutoff)
		fn2 = Fib(n-2, cutoff)
	}
	return fn1 + fn2
}
