package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func calculateY(n int64, dy0 int64) int64 {
	if n == 0 {
		return 0
	}
	answer := calculateY(n-1, dy0) + calculateDY(n-1, dy0)
	return answer
}

func calculateDY(n int64, dy0 int64) int64 {
	if n == 0 {
		return dy0
	}
	dyminus1 := calculateDY(n-1, dy0)
	return dyminus1 - 1
}

func calculateX(n int64, dx0 int64) int64 {
	if n == 0 {
		return 0
	}
	answer := calculateX(n-1, dx0) + calculateDX(n-1, dx0)
	return answer
}

func calculateDX(n int64, dx0 int64) int64 {
	if n == 0 {
		return dx0
	}
	dxminus1 := calculateDX(n-1, dx0)
	if dxminus1 > 0 {
		return dxminus1 - 1
	} else if dxminus1 < 0 {
		return dxminus1 + 1
	}
	return 0
}

func main() {

	data, err := os.ReadFile("day17.txt")
	// data, err := os.ReadFile("day17_ex1.txt")
	if err != nil {
		log.Fatal(err)
	}

	target := strings.Split(strings.TrimSpace(string(data)), " ")

	fmt.Println(target)

	r, _ := regexp.Compile("([a-z]+)=(-?[0-9]+)..(-?[0-9]+)")

	x_range := r.FindAllStringSubmatch(string(data), -1)[0]
	x_min, _ := strconv.ParseInt(x_range[2], 0, 64)
	x_max, _ := strconv.ParseInt(x_range[3], 0, 64)

	y_range := r.FindAllStringSubmatch(string(data), -1)[1]
	y_min, _ := strconv.ParseInt(y_range[3], 0, 64)
	y_max, _ := strconv.ParseInt(y_range[2], 0, 64)

	x_vel := int64(x_max)
	x_prev := int64(0)

	n_list_for_x := make(map[int64][][]int64, 0)
	n_list_for_x_infinate := make(map[int64][][]int64, 0)
	for x_vel > 0 {
		changed := true
		n_term := int64(0)
		for n := int64(1); calculateX(n, x_vel) <= x_max && changed; n++ {
			if x_prev == calculateX(n, x_vel) {
				changed = false
			} else {
				x_prev = calculateX(n, x_vel)
			}
			n_term = n
			x_term := calculateX(n_term, x_vel)
			if x_term >= x_min && x_term <= x_max {
				if !changed {
					n_list_for_x_infinate[n_term] = append(n_list_for_x_infinate[n_term], []int64{x_term, x_vel})
					// fmt.Println(n_term, x_vel, x(n_term, x_vel), "no change")
				} else {
					n_list_for_x[n_term] = append(n_list_for_x[n_term], []int64{x_term, x_vel})
					// fmt.Println(n_term, x_vel, x(n_term, x_vel))
				}
			}
		}
		x_vel--
	}

	y_vel := -int64(y_max)
	n_list_for_y := make(map[int64][][]int64, 0)

	for y_vel > int64(y_max)-1 {
		n_term := int64(0)
		for n := int64(1); calculateY(n, y_vel) >= y_max; n++ {
			n_term = n
			y_term := calculateY(n_term, y_vel)
			if y_term <= y_min && y_term >= y_max {
				n_list_for_y[n_term] = append(n_list_for_y[n_term], []int64{y_term, y_vel})
			}
		}
		y_vel--
	}
	valid_speeds := make(map[[2]int64]bool, 0)
	for yk, v := range n_list_for_y {
		for _, yvel := range v {
			for _, xvel := range n_list_for_x[yk] {
				valid_speeds[[2]int64{xvel[1], yvel[1]}] = true
			}
			for xk, xval := range n_list_for_x_infinate {
				if xk <= yk {
					for _, xvel := range xval {
						valid_speeds[[2]int64{xvel[1], yvel[1]}] = true
					}
				}
			}
		}
	}
	max := int64(0)
	for speed := range valid_speeds {
		if speed[1] > max {
			max = speed[1]
		}

	}
	last_y := int64(0)
	for n := int64(1); calculateY(n, max) > last_y; n++ {
		last_y = calculateY(n, max)
	}
	fmt.Println(last_y, len(valid_speeds))
	// n := int64(0)
	// for n < 1000 {
	// 	n++
	// 	if y(n, y_vel) < y_max {
	// 		fmt.Println(y_max)
	// 		break
	// 	} else {
	// 		y_max = y(n, y_vel)
	// 	}
	// }

}
