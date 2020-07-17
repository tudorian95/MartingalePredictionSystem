package main

import (
	"fmt"
	"math/rand"
	"time"
)

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

func dice4(x int) int {
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
					bet *= 2
					bankroll -= bet
					reroll3 := roll()
					if reroll3 >= 51 {
						bankroll += 2 * bet
						bet = 1
					} else {
						bet = 1
					}
				}
			}
		}
	}
	return bankroll
}

func dice5(x int) int {
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
					bet *= 2
					bankroll -= bet
					reroll3 := roll()
					if reroll3 >= 51 {
						bankroll += 2 * bet
						bet = 1
					} else {
						bet *= 2
						bankroll -= bet
						reroll4 := roll()
						if reroll4 >= 51 {
							bankroll += 2 * bet
							bet = 1
						} else {
							bet = 1
						}
					}
				}
			}
		}
	}
	return bankroll
}

func dice6(x int) int {
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
					bet *= 2
					bankroll -= bet
					reroll3 := roll()
					if reroll3 >= 51 {
						bankroll += 2 * bet
						bet = 1
					} else {
						bet *= 2
						bankroll -= bet
						reroll4 := roll()
						if reroll4 >= 51 {
							bankroll += 2 * bet
							bet = 1
						} else {
							bet *= 2
							bankroll -= bet
							reroll5 := roll()
							if reroll5 >= 51 {
								bankroll += 2 * bet
								bet = 1
							} else {
								bet = 1
							}
						}
					}
				}
			}
		}
	}
	return bankroll
}

func main() {
	rand.Seed(time.Now().UnixNano())
	x := dice1(100000000)
	y := dice2(100000000)
	z := dice3(100000000)
	a := dice4(100000000)
	b := dice5(100000000)
	c := dice6(100000000)
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
	fmt.Println("using a 6 step Martingale we get")
	fmt.Println(c)
}
