package main

import (
	"fmt"
	"time"

	"github.com/inancgumus/screen"
)

func main() {

	type placeholder [5]string
	zero := placeholder{
		"0 0 0",
		"0   0",
		"0   0",
		"0   0",
		"0 0 0",
	}

	one := placeholder{
		"0 0  ",
		"  0  ",
		"  0  ",
		"  0  ",
		"0 0 0",
	}

	two := placeholder{
		"0 0 0",
		"    0",
		"0 0 0",
		"0    ",
		"0 0 0",
	}

	three := placeholder{
		"0 0 0",
		"    0",
		"0 0 0",
		"    0",
		"0 0 0",
	}

	four := placeholder{
		"0   0",
		"0   0",
		"0 0 0",
		"    0",
		"    0",
	}

	five := placeholder{
		"0 0 0",
		"0    ",
		"0 0 0",
		"    0",
		"0 0 0",
	}

	six := placeholder{
		"0 0 0",
		"0    ",
		"0 0 0",
		"0   0",
		"0 0 0",
	}

	seven := placeholder{
		"0 0 0",
		"    0",
		"    0",
		"    0",
		"    0",
	}

	eight := placeholder{
		"0 0 0",
		"0   0",
		"0 0 0",
		"0   0",
		"0 0 0",
	}

	nine := placeholder{
		"0 0 0",
		"0   0",
		"0 0 0",
		"    0",
		"0 0 0",
	}

	colon := placeholder{
		"     ",
		"  0  ",
		"     ",
		"  0  ",
		"     ",
	}

	digits := [...]placeholder{
		zero, one, two, three, four, five, six, seven, eight, nine,
	}

	/*for _, digit := range digits {
		for _, line := range digit {
			fmt.Println(line)
		}
		fmt.Println()
	}


	fmt.Printf("hour: %d, min %d, sec: %d\n", hour, min, sec)
	for line := range digits[0] {
		for digit := range digits {
			fmt.Print(digits[digit][line], "  ")
		}
		fmt.Println()
	}
	*/
	screen.Clear()
	for {
		screen.MoveTopLeft()
		now := time.Now()
		hour, min, sec := now.Hour(), now.Minute(), now.Second()

		/*clock := [...]placeholder{
			digits[3], digits[2],
			digits[1], digits[0],
			digits[4], digits[5],
		}*/

		clock := [...]placeholder{
			digits[hour/10], digits[hour%10],
			colon,
			digits[min/10], digits[min%10],
			colon,
			digits[sec/10], digits[sec%10],
		}

		//	fmt.Printf("hour: %d, min %d, sec: %d\n", hour, min, sec)
		for line := range clock[0] {
			for index, digit := range clock {
				next := clock[index][line]
				if digit == colon && sec%2 == 0 {
					next = "     "
				}
				fmt.Print(next, "  ")
			}
			fmt.Println()
		}

		time.Sleep(time.Second)
	}
}
