package main

import (
	"fmt"
	"time"
)

var votes map[string]int = make(map[string]int)

func Democracy(in <-chan string, out chan<- string) {
	fmt.Println("Democracy initialised")

	timer := time.Tick(time.Second)
	for {
		select {
		case <-timer:
			max_v := 0
			max_k := ""

			for k, v := range votes {
				if v > max_v {
					max_v = v
					max_k = k
				}
			}

			if max_v > 0 {
				fmt.Println(max_k, "(", max_v, "votes)")
				out <- max_k
			}

			votes = make(map[string]int) // reset votes

		case vote := <-in:
			if _, ok := votes[vote]; !ok {
				votes[vote] = 1
			} else {
				votes[vote] += 1
			}
		}
	}
}
