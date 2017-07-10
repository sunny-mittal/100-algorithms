package main

import "fmt"

func play(discsLeft, src, dst, aux int) {
	if discsLeft != 0 {
		play(discsLeft-1, src, aux, dst)
		fmt.Printf("%d => %d\n", src, dst)
		play(discsLeft-1, aux, dst, src)
	}
}

func main() {
	play(4, 1, 2, 3)
}
