package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func myfibonacci() func(int) int {
	var fib[10] int
	return func(x int) int {
		if x==0 {
			fib[x] = x
			return fib[x]
		} else if x==1 {
			fib[x] = x
			return fib[x]
		} else {
			fib[x] = fib[x-2] + fib[x-1]
			return fib[x]
		}
	}
}

// fibonacci is a function that returns
// a function that returns an int.
func smartfibonacci() func() int {
    a, b := 1, 0
    return func() int {
        a, b = b, a+b
        return a
    }
}

func main() {
	//f1 := myfibonacci()
	f2 := smartfibonacci()
	for i := 0; i < 10; i++ {
		//fmt.Println(f1(i))
		fmt.Println(f2())
	}
}

