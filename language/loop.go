package main

import "fmt"

func main() {
	for {
		fmt.Println("loop 1")
		break
	}

	isRun := true
	for isRun {
		fmt.Println("loop 2")
		isRun = false
	}

	for i := 0; i < 2; i++ {
		if i == 1 {
			continue
		}
		fmt.Println("loop 3")
	}

	// slice
	sl := []int{1, 2, 3, 4}
	idx := 0
	for idx < len(sl) {
		fmt.Print(" ", idx, ":", sl[idx])
		idx++
	}
	fmt.Println(" - Slice")

	// or range
	for idx, val := range sl {
		fmt.Print(" ", idx, ":", val)
	}
	fmt.Println(" - Slice range")

	// map
	profile := map[int]string{1: "A", 2: "B", 4: "D"}
	for key := range profile {
		fmt.Println(key, ":", profile[key])
	}

	for _, val := range profile {
		fmt.Print(">", val)
	}
	fmt.Println()

	str := "Итерация по символам строки"
	for pos, char := range str {
		fmt.Printf("%d : %#U, ", pos, char)
	}
	fmt.Println()

	fmt.Println("Stop")
}
