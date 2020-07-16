package main

import (
	"fmt"
	"math/rand"
	"time"
)

func roll() int {
	rand.Seed(time.Now().UnixNano())
	result := rand.Intn(101)
	time.Sleep(1 * time.Nanosecond)
	return result
}

func dice1(x int) int {
	bankroll := 2000
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
	bankroll := 2000
	bet := 1
	for i := 0; i < x; i++ {
		result := roll()
		if result >= 51 {
			bankroll += bet
			bet = 1
		} else {
			bankroll -= bet
			bet *= 2
			reroll := roll()
			if reroll >= 51 {
				bankroll += bet
				bet = 1
			} else {
				bankroll -= bet
				bet = 1
			}
		}
	}
	return bankroll
}

func dice3(x int) int {
	bankroll := 2000
	bet := 1
	for i := 0; i < x; i++ {
		result := roll()
		if result >= 51 {
			bankroll += bet
			bet = 1
		} else {
			bankroll -= bet
			bet *= 2
			reroll := roll()
			if reroll >= 51 {
				bankroll += bet
				bet = 1
			} else {
				bankroll -= bet
				bet *= 2
				reroll2 := roll()
				if reroll2 >= 51 {
					bankroll += bet
					bet = 1
				} else {
					bankroll -= bet
					bet = 1
				}
			}
		}
	}
	return bankroll
}

func dice4(x int) int {
	bankroll := 2000
	bet := 1
	for i := 0; i < x; i++ {
		result := roll()
		if result >= 51 {
			bankroll += bet
			bet = 1
		} else {
			bankroll -= bet
			bet *= 2
			reroll := roll()
			if reroll >= 51 {
				bankroll += bet
				bet = 1
			} else {
				bankroll -= bet
				bet *= 2
				reroll2 := roll()
				if reroll2 >= 51 {
					bankroll += bet
					bet = 1
				} else {
					bankroll -= bet
					bet *= 2
					reroll3 := roll()
					if reroll3 >= 51 {
						bankroll += bet
						bet = 1
					} else {
						bankroll -= bet
						bet = 1
					}
				}
			}
		}
	}
	return bankroll
}

func dice5(x int) int {
	bankroll := 2000
	bet := 1
	for i := 0; i < x; i++ {
		result := roll()
		if result >= 51 {
			bankroll += bet
			bet = 1
		} else {
			bankroll -= bet
			bet *= 2
			reroll := roll()
			if reroll >= 51 {
				bankroll += bet
				bet = 1
			} else {
				bankroll -= bet
				bet *= 2
				reroll2 := roll()
				if reroll2 >= 51 {
					bankroll += bet
					bet = 1
				} else {
					bankroll -= bet
					bet *= 2
					reroll3 := roll()
					if reroll3 >= 51 {
						bankroll += bet
						bet = 1
					} else {
						bankroll -= bet
						bet *= 2
						reroll4 := roll()
						if reroll4 >= 51 {
							bankroll += bet
							bet = 1
						} else {
							bankroll -= bet
							bet = 1
						}
					}
				}
			}
		}
	}
	return bankroll
}

func main() {
	x := dice1(100000)
	y := dice2(100000)
	z := dice3(100000)
	a := dice4(100000)
	b := dice5(100000)
	fmt.Println("using a 1 step Martingale we get")
	fmt.Println(x)
	fmt.Println("using a 2 step Martingale we get")
	fmt.Println(y)
	fmt.Println("using a 3 step Martingale we get")
	fmt.Println(z)
	fmt.Println("using a 4 step Martingale we get")
	fmt.Println(a)
	fmt.Println("using a 5 step Martingale we get")
	fmt.Println(b)
	/*	for i := 0; i < 1000; i++ {
			fmt.Println(roll())
		}
	*/
}

//pana la 51 inclusiv ai 52 de variante
// de la 52 pana la 100 ai 49 de variante
