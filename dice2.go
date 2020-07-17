package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var m sync.Mutex

func roll() int {
	result := rand.Intn(101)
	return result
}

func dice1(x int) int {
	bankroll := 100
	bet := 1
	for i := 0; i < x; i++ {
		result := roll()
		if result >= 51 {
			bankroll += bet
		} else {
			bankroll -= bet
		}
	}
	return bankroll
}

func dice2(x int) int {
	bankroll := 100
	bet := 1
	for i := 0; i < x; i++ {
		result := roll()
		if result >= 51 {
			bankroll += bet
		} else {
			bankroll -= bet
			bet *= 2
			bankroll -= bet
			reroll := roll()
			if reroll >= 51 {
				bankroll += 2 * bet
				bet = 1
			} else {
				bet = 1
			}
		}
	}
	return bankroll
}

func dice3(x int) int {
	bankroll := 100
	bet := 1
	for i := 0; i < x; i++ {
		result := roll()
		if result >= 51 {
			bankroll += bet
		} else {
			bankroll -= bet
			bet *= 2
			bankroll -= bet
			reroll := roll()
			if reroll >= 51 {
				bankroll += 2 * bet
				bet = 1
			} else {
				bet *= 2
				bankroll -= bet
				reroll2 := roll()
				if reroll2 >= 51 {
					bankroll += 2 * bet
					bet = 1
				} else {
					bet = 1
				}
			}
		}
	}
	return bankroll
}

func main() {
	rand.Seed(time.Now().UnixNano())
	/*	x := dice1(100)
		y := dice2(100)
		z := dice3(100)
		fmt.Println("using a 1 step Martingale we get")
		fmt.Println(x)
		fmt.Println("using a 2 step Martingale we get")
		fmt.Println(y)
		fmt.Println("using a 3 step Martingale we get")
		fmt.Println(z)*/

	for i := 0; i < 15; i++ {
		fmt.Println(roll())
	}

}
