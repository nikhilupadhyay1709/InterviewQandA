package main

import "fmt"

func main() {
	var t1, t2, t3, tn int
	t1 = 5
	for i := 1; i <= t1; i++ {

		for j := 1; j <= t1-i; j++ {
			fmt.Print("  ")
			t3++
		}
		for {
			if t3 <= t1-1 {
				fmt.Printf(" %d", i+t2)
				t3++
			} else {
				tn++
				fmt.Printf(" %d", (i + t2 - 2*tn))
			}
			t2++

			if t2 == 2*i-1 {
				break
			}

		}
		t3 = 0
		tn = 0
		t2 = 0
		fmt.Println("")
	}

}
