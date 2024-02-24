package piscine

import "fmt"

func QuadB(x, y int) {
	if x>0 && y>0 {
		for i := 1; i <= y; i++ {
			for j := 1; j <= x; j++ {
				if i == 1 {
					if j == 1 {
						fmt.Print("/")
					} else if j == x {
						fmt.Print("\\")
					} else {
						fmt.Print("*")
					}
				} else if i == y {
					if j == 1 {
						fmt.Print("\\")
					} else if j == x {
						fmt.Print("/")
					} else {
						fmt.Print("*")
					}
				} else {
					if j == 1 || j == x {
						fmt.Print("*")
					} else {
						fmt.Print(" ")
					}
				}
			}
			fmt.Println()
		}
	}else{
		fmt.Print("")
	}
}
