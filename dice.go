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
	bankroll := 1000
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
	bankroll := 1000
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
	bankroll := 1000
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

func dicePerfected(x int) int {
	bankroll := 1000
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
					bet = 100
					bankroll -= bet
					reroll3 := roll()
					if reroll3 >= 51 {
						bankroll += bet
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

func dice4(x int) int {
	bankroll := 1000
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
	bankroll := 1000
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
	bankroll := 1000
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

func dice7(x int) int {
	bankroll := 1000
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
								bet *= 2
								bankroll -= bet
								reroll6 := roll()
								if reroll6 >= 51 {
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
	}
	return bankroll
}

func dice8(x int) int {
	bankroll := 1000
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
								bet *= 2
								bankroll -= bet
								reroll6 := roll()
								if reroll6 >= 51 {
									bankroll += 2 * bet
									bet = 1
								} else {
									bet *= 2
									bankroll -= bet
									reroll7 := roll()
									if reroll7 >= 51 {
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
		}
	}
	return bankroll
}

func dice9(x int) int {
	bankroll := 1000
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
								bet *= 2
								bankroll -= bet
								reroll6 := roll()
								if reroll6 >= 51 {
									bankroll += 2 * bet
									bet = 1
								} else {
									bet *= 2
									bankroll -= bet
									reroll7 := roll()
									if reroll7 >= 51 {
										bankroll += 2 * bet
										bet = 1
									} else {
										bet *= 2
										bankroll -= bet
										reroll8 := roll()
										if reroll8 >= 51 {
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
			}
		}
	}
	return bankroll
}

func dice10(x int) int {
	bankroll := 1000
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
								bet *= 2
								bankroll -= bet
								reroll6 := roll()
								if reroll6 >= 51 {
									bankroll += 2 * bet
									bet = 1
								} else {
									bet *= 2
									bankroll -= bet
									reroll7 := roll()
									if reroll7 >= 51 {
										bankroll += 2 * bet
										bet = 1
									} else {
										bet *= 2
										bankroll -= bet
										reroll8 := roll()
										if reroll8 >= 51 {
											bankroll += 2 * bet
											bet = 1
										} else {
											bet *= 2
											bankroll -= bet
											reroll9 := roll()
											if reroll9 >= 51 {
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
				}
			}
		}
	}
	return bankroll
}

func dice11(x int) int {
	bankroll := 1000
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
								bet *= 2
								bankroll -= bet
								reroll6 := roll()
								if reroll6 >= 51 {
									bankroll += 2 * bet
									bet = 1
								} else {
									bet *= 2
									bankroll -= bet
									reroll7 := roll()
									if reroll7 >= 51 {
										bankroll += 2 * bet
										bet = 1
									} else {
										bet *= 2
										bankroll -= bet
										reroll8 := roll()
										if reroll8 >= 51 {
											bankroll += 2 * bet
											bet = 1
										} else {
											bet *= 2
											bankroll -= bet
											reroll9 := roll()
											if reroll9 >= 51 {
												bankroll += 2 * bet
												bet = 1
											} else {
												bet *= 2
												bankroll -= bet
												reroll10 := roll()
												if reroll10 >= 51 {
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
					}
				}
			}
		}
	}
	return bankroll
}

func dice12(x int) int {
	bankroll := 1000
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
								bet *= 2
								bankroll -= bet
								reroll6 := roll()
								if reroll6 >= 51 {
									bankroll += 2 * bet
									bet = 1
								} else {
									bet *= 2
									bankroll -= bet
									reroll7 := roll()
									if reroll7 >= 51 {
										bankroll += 2 * bet
										bet = 1
									} else {
										bet *= 2
										bankroll -= bet
										reroll8 := roll()
										if reroll8 >= 51 {
											bankroll += 2 * bet
											bet = 1
										} else {
											bet *= 2
											bankroll -= bet
											reroll9 := roll()
											if reroll9 >= 51 {
												bankroll += 2 * bet
												bet = 1
											} else {
												bet *= 2
												bankroll -= bet
												reroll10 := roll()
												if reroll10 >= 51 {
													bankroll += 2 * bet
													bet = 1
												} else {
													bet *= 2
													bankroll -= bet
													reroll11 := roll()
													if reroll11 >= 51 {
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
						}
					}
				}
			}
		}
	}
	return bankroll
}

func dice13(x int) int {
	bankroll := 1000
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
								bet *= 2
								bankroll -= bet
								reroll6 := roll()
								if reroll6 >= 51 {
									bankroll += 2 * bet
									bet = 1
								} else {
									bet *= 2
									bankroll -= bet
									reroll7 := roll()
									if reroll7 >= 51 {
										bankroll += 2 * bet
										bet = 1
									} else {
										bet *= 2
										bankroll -= bet
										reroll8 := roll()
										if reroll8 >= 51 {
											bankroll += 2 * bet
											bet = 1
										} else {
											bet *= 2
											bankroll -= bet
											reroll9 := roll()
											if reroll9 >= 51 {
												bankroll += 2 * bet
												bet = 1
											} else {
												bet *= 2
												bankroll -= bet
												reroll10 := roll()
												if reroll10 >= 51 {
													bankroll += 2 * bet
													bet = 1
												} else {
													bet *= 2
													bankroll -= bet
													reroll11 := roll()
													if reroll11 >= 51 {
														bankroll += 2 * bet
														bet = 1
													} else {
														bet *= 2
														bankroll -= bet
														reroll12 := roll()
														if reroll12 >= 51 {
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
							}
						}
					}
				}
			}
		}
	}
	return bankroll
}

func dice14(x int) int {
	bankroll := 1000
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
								bet *= 2
								bankroll -= bet
								reroll6 := roll()
								if reroll6 >= 51 {
									bankroll += 2 * bet
									bet = 1
								} else {
									bet *= 2
									bankroll -= bet
									reroll7 := roll()
									if reroll7 >= 51 {
										bankroll += 2 * bet
										bet = 1
									} else {
										bet *= 2
										bankroll -= bet
										reroll8 := roll()
										if reroll8 >= 51 {
											bankroll += 2 * bet
											bet = 1
										} else {
											bet *= 2
											bankroll -= bet
											reroll9 := roll()
											if reroll9 >= 51 {
												bankroll += 2 * bet
												bet = 1
											} else {
												bet *= 2
												bankroll -= bet
												reroll10 := roll()
												if reroll10 >= 51 {
													bankroll += 2 * bet
													bet = 1
												} else {
													bet *= 2
													bankroll -= bet
													reroll11 := roll()
													if reroll11 >= 51 {
														bankroll += 2 * bet
														bet = 1
													} else {
														bet *= 2
														bankroll -= bet
														reroll12 := roll()
														if reroll12 >= 51 {
															bankroll += 2 * bet
															bet = 1
														} else {
															bet *= 2
															bankroll -= bet
															reroll13 := roll()
															if reroll13 >= 51 {
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
								}
							}
						}
					}
				}
			}
		}
	}
	return bankroll
}

func dice15(x int) int {
	bankroll := 1000
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
								bet *= 2
								bankroll -= bet
								reroll6 := roll()
								if reroll6 >= 51 {
									bankroll += 2 * bet
									bet = 1
								} else {
									bet *= 2
									bankroll -= bet
									reroll7 := roll()
									if reroll7 >= 51 {
										bankroll += 2 * bet
										bet = 1
									} else {
										bet *= 2
										bankroll -= bet
										reroll8 := roll()
										if reroll8 >= 51 {
											bankroll += 2 * bet
											bet = 1
										} else {
											bet *= 2
											bankroll -= bet
											reroll9 := roll()
											if reroll9 >= 51 {
												bankroll += 2 * bet
												bet = 1
											} else {
												bet *= 2
												bankroll -= bet
												reroll10 := roll()
												if reroll10 >= 51 {
													bankroll += 2 * bet
													bet = 1
												} else {
													bet *= 2
													bankroll -= bet
													reroll11 := roll()
													if reroll11 >= 51 {
														bankroll += 2 * bet
														bet = 1
													} else {
														bet *= 2
														bankroll -= bet
														reroll12 := roll()
														if reroll12 >= 51 {
															bankroll += 2 * bet
															bet = 1
														} else {
															bet *= 2
															bankroll -= bet
															reroll13 := roll()
															if reroll13 >= 51 {
																bankroll += 2 * bet
																bet = 1
															} else {
																bet *= 2
																bankroll -= bet
																reroll14 := roll()
																if reroll14 >= 51 {
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
									}
								}
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
	x := dice1(10000)
	y := dice2(10000)
	z := dice3(10000)
	a := dice4(10000)
	b := dice5(10000)
	c := dice6(10000)
	d := dice7(10000)
	e := dice8(10000)
	f := dice9(10000)
	g := dice10(10000)
	h := dice11(10000)
	i := dice12(10000)
	j := dice13(10000)
	k := dice14(10000)
	l := dice15(10000)
	m := dicePerfected(10000)
	fmt.Println("using a 1 step Martingale we get")
	fmt.Println(" ", x)
	fmt.Println("using a 2 step Martingale we get")
	fmt.Println(" ", y)
	fmt.Println("using a 3 step Martingale we get")
	fmt.Println(" ", z)
	fmt.Println("using a 4 step Martingale we get")
	fmt.Println(" ", a)
	fmt.Println("using a 5 step Martingale we get")
	fmt.Println(" ", b)
	fmt.Println("using a 6 step Martingale we get")
	fmt.Println(" ", c)
	fmt.Println("using a 7 step Martingale we get")
	fmt.Println(" ", d)
	fmt.Println("using a 8 step Martingale we get")
	fmt.Println(" ", e)
	fmt.Println("using a 9 step Martingale we get")
	fmt.Println(" ", f)
	fmt.Println("using a 10 step Martingale we get")
	fmt.Println(" ", g)
	fmt.Println("using a 11 step Martingale we get")
	fmt.Println(" ", h)
	fmt.Println("using a 12 step Martingale we get")
	fmt.Println(" ", i)
	fmt.Println("using a 13 step Martingale we get")
	fmt.Println(" ", j)
	fmt.Println("using a 14 step Martingale we get")
	fmt.Println(" ", k)
	fmt.Println("using a 15 step Martingale we get")
	fmt.Println(" ", l)
	fmt.Println("using dicePerfected Martingale we get")
	fmt.Println(" ", m)
}
