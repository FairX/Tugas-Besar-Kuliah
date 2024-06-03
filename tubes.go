package main

import (
	"fmt"
	"math/rand"
	"time"
)

const topPlay int = 10

type tiles [30]struct {
	n1, n2, totIdx, idxOut, balak int
}

type user [100]struct {
	status     string
	win, round int
	roundS     float64
}

type scoreboard [100]struct {
	statistics user
	name       string
	score      float64
}

func main() {
	var player scoreboard
	var stat user
	var n int
	mainMenu(&player, &stat, &n)
}

func mainMenu(player *scoreboard, stat *user, n *int) {
	var option string
	//ADD BACK AND RULES OPTIONS
	fmt.Println("==================================================")
	fmt.Println(" DOMINO SOLITAIRE ")
	fmt.Println(" 1 : Start Game")
	fmt.Println(" 2 : Tutorial ")
	fmt.Println(" 3 : Scoreboard")
	fmt.Println(" 9 : Exit Game")
	fmt.Print(" What's on your mind? ")
	fmt.Scanln(&option)
	fmt.Println()
	if option == "1" {
		fmt.Print(" Player's name: ")
		fmt.Scanln(&player[*n].name)
		fmt.Println()
		game(*n, &*player, &*stat)
		checkName(&*player, &*stat, &*n)
		*n++
	} else if option == "2" {
		rules()
	} else if option == "3" {
		if *n == 0 {
			fmt.Println(" There is no recorded gameplay yet!!")
			fmt.Println()
		} else {
			rankSort(&*player, &*stat, *n)
			rankList(*player, *stat, *n)
		}
	} else if option == "9" {
		fmt.Println(" Thank you for playing!!")
		fmt.Println("==================================================")
	} else {
		fmt.Println(" WRONG INPUT, please input again")
		fmt.Println()
	}
	if option != "9" {
		mainMenu(&*player, &*stat, &*n)
	}
}

func rules() {
	var x string
	fmt.Println(" Tutorial !!")
	fmt.Println(" 1. At the start of the game, you will be given 4 tiles.")
	fmt.Println(" 2. Each tile has 2 numbers, and these numbers will determine your next choice.")
	fmt.Println(" 3. You have 6 options to choose from based on the numbers on your tiles:")
	fmt.Println("   - number 1, you can replace the first tile.")
	fmt.Println("   - number 2, you can replace the second tile.")
	fmt.Println("   - number 3, you can replace the third tile.")
	fmt.Println("   - number 4, you can replace the fourth tile.")
	fmt.Println("   - number 0, you can finish the current round.")
	fmt.Println("   - number 9, you can choose to quit the game.")
	fmt.Println(" 4. You can replace your tiles up to two times per round.")
	fmt.Println(" 5. Every time you replace a tile, your score after winning will be deducted by 25 points.")
	fmt.Println(" 6. Each win will earn you 100 points, a draw will give you 75 points, and losing will result in 0 points.")
	fmt.Println()
	fmt.Println(" How to Win")
	fmt.Println(" 1. There are two ways to win the game:")
	fmt.Println("   - if the total sum of the numbers on your tiles is higher than that of your opponent's tiles.")
	fmt.Println("   - if you have a twin tile (two tiles with the same number), regardless of the total sum.")
	fmt.Println(" 2. If either you or your opponent has a twin tile, and the other player doesn't,")
	fmt.Println("    the player with the twin tile will be the winner.")
	fmt.Println(" 3. A draw can only occur if neither you nor your opponent has a twin tile.")
	fmt.Println(" 4. If both you and your opponent have a twin tile, the tile with the higher number will determine the winner.")
	fmt.Println()
	fmt.Print(" Press ENTER to continue.")
	fmt.Scanln(&x)
	fmt.Println()
}

func rankList(player scoreboard, stat user, n int) {
	var choice, choiceS, nouse string
	var choiceN int
	fmt.Printf(" Top %v players\n", topPlay)
	for i := 0; i < topPlay; i++ {
		if player[i].name != "" {
			fmt.Printf(" %d. Name: %s, Score: %.f, Win: %d, Total round: %d \n", i+1, player[i].name, player[i].score, stat[i].win, stat[i].round)
		}
	}
	fmt.Println()
	fmt.Print(" What stat are you looking for? (1 = name, 2 = rank, 9 = back) ")
	fmt.Scanln(&choice)
	fmt.Println()
	//ADD BACK TO SCOREBOARD OPTIONS
	if choice == "1" {
		fmt.Print(" Input the name that you are looking for: ")
		fmt.Scanln(&choiceS)
		fmt.Println()
		statList(player, stat, nameSearch(player, stat, n, choiceS))
		fmt.Println()
		fmt.Print(" Press ENTER to go back.")
		fmt.Scanln(&nouse)
		fmt.Println()
		rankList(player, stat, n)
	} else if choice == "2" {
		fmt.Print(" Input the rank that you are looking for: ")
		fmt.Scanln(&choiceN)
		fmt.Println()
		if choiceN > n {
			fmt.Println(" Player not found...")
		} else if choiceN > topPlay {
			fmt.Printf(" This player is not inside the top %d players\n", topPlay)
		} else if choiceN <= topPlay {
			statList(player, stat, choiceN-1)
		}
		fmt.Println()
		fmt.Print(" Press ENTER to go back.")
		fmt.Scanln(&nouse)
		fmt.Println()
		rankList(player, stat, n)
	} else if choice == "9" {
		//backtoabove
	} else {
		fmt.Println(" WRONG INPUT, please input again")
		fmt.Println()
		rankList(player, stat, n)
	}
}

func nameSearch(player scoreboard, stat user, n int, name string) int {
	for i := 0; i < n; i++ {
		if name == player[i].name {
			return i
		}
	}
	return -1
}

func rankSort(player *scoreboard, stat *user, n int) {
	var i, j int
	var temp scoreboard
	var tempS user
	i = 1
	for i < n {
		j = i
		temp[i] = player[j]
		tempS[i] = stat[j]
		for j > 0 && temp[i].score > player[j-1].score {
			player[j] = player[j-1]
			stat[j] = stat[j-1]
			j--
		}
		player[j] = temp[i]
		stat[j] = tempS[i]
		i++
	}
}

func statList(player scoreboard, stat user, n int) {
	if n == -1 {
		fmt.Println(" Player not found..")
	} else {
		fmt.Printf(" %s's statistics\n", player[n].name)
		for j := 0; j < stat[n].round; j++ {
			fmt.Printf(" round %d: %s, score: %.f\n", j+1, player[n].statistics[j].status, player[n].statistics[j].roundS)
		}
	}
}

func checkName(P *scoreboard, stat *user, n *int) {
	for i := 0; i < *n; i++ {
		if P[i].name == P[*n].name {
			if P[*n].score > P[i].score {
				P[i] = P[*n]
				stat[i] = stat[*n]
				P[*n] = P[*n+1]
				stat[*n] = stat[*n+1]
				*n--
			} else {
				P[*n] = P[*n+1]
				stat[*n] = stat[*n+1]
				*n--
			}
		}
	}
}

func game(match int, player *scoreboard, stat *user) {
	var tile tiles
	var n, option int
	tileList(&tile, &n)
	for option != 9 {
		fmt.Println(" Dealing ... ")
		playTab(&tile, &option, match, &*stat, &*player)
		if option == 9 {
			fmt.Printf(" Your total score is %.f \n", player[match].score)
			fmt.Printf(" With %d Win, and %d round.\n", stat[match].win, stat[match].round)
			fmt.Println()
		}
	}
}

func check(list tiles, num *int) {
	for i := 0; i < list[0].totIdx; i++ {
		for *num == list[i].idxOut {
			*num = rand.Intn(28)
			i = 0
		}
	}
}

func tileInsert(list *tiles, t1, t2 *int, num *int) {
	list[list[0].totIdx].idxOut = *num
	list[0].totIdx++
	*t1 = list[*num].n1
	*t2 = list[*num].n2
}

func randNum(list *tiles, t1, t2, rand1 *int) {
	rand.Seed(time.Now().UnixNano())
	*rand1 = rand.Intn(28)
	check(*list, &*rand1)
	tileInsert(list, &*t1, &*t2, &*rand1)
}

func playTab(list *tiles, option *int, match int, stat *user, player *scoreboard) {
	list[0].totIdx = 0
	list[0].idxOut = 0
	var t1, t2, t3, t4, t5, t6, t7, t8 int
	var idx int
	randNum(&*list, &t1, &t2, &idx)
	randNum(&*list, &t3, &t4, &idx)
	randNum(&*list, &t5, &t6, &idx)
	randNum(&*list, &t7, &t8, &idx)
	fmt.Println(" Your tiles: ")
	fmt.Printf(" (%d,%d) (%d,%d) (%d,%d) (%d,%d)\n", t1, t2, t3, t4, t5, t6, t7, t8)
	fmt.Println(" Replace X-th tile, X: 1-4 | Done: 0 | Exit: 9")
	fmt.Print(" Decision? ")
	fmt.Scanln(&*option)
	fmt.Println()
	pilihan(&*list, &*option, t1, t2, t3, t4, t5, t6, t7, t8, match, &*stat, &*player)
}

func pilihan(list *tiles, option *int, t1, t2, t3, t4, t5, t6, t7, t8, match int, stat *user, player *scoreboard) {
	var t9, t10, t11, t12, t13, t14, t15, t16 int
	var repCount, idx int
	randNum(&*list, &t9, &t10, &idx)
	randNum(&*list, &t11, &t12, &idx)
	randNum(&*list, &t13, &t14, &idx)
	randNum(&*list, &t15, &t16, &idx)
	for repCount < 2 && *option != 0 && *option != 9 {
		if *option >= 1 && *option <= 4 {
			repCount++
			if *option == 1 {
				randNum(&*list, &t1, &t2, &idx)
				list[0].totIdx--
				list[0].idxOut = idx
			} else if *option == 2 {
				randNum(&*list, &t3, &t4, &idx)
				list[0].totIdx--
				list[1].idxOut = idx
			} else if *option == 3 {
				randNum(&*list, &t5, &t6, &idx)
				list[0].totIdx--
				list[2].idxOut = idx
			} else if *option == 4 {
				randNum(&*list, &t7, &t8, &idx)
				list[0].totIdx--
				list[3].idxOut = idx
			}
			player[match].statistics[stat[match].round].roundS -= 25
		} else {
			fmt.Println(" WRONG INPUT, please input again")
			fmt.Println()
		}
		fmt.Println(" Your tiles: ")
		fmt.Printf(" (%d,%d) (%d,%d) (%d,%d) (%d,%d)\n", t1, t2, t3, t4, t5, t6, t7, t8)
		fmt.Println(" Replace X-th tile, X: 1-4 | Done: 0 | Exit: 9")
		fmt.Print(" Decision? ")
		fmt.Scanln(&*option)
		fmt.Println()
	}
	if *option == 0 || repCount == 2 {
		fmt.Println("----------------------------------------")
		fmt.Printf("|Dealer tiles:                         |\n")
		fmt.Printf("|(%d,%d) (%d,%d) (%d,%d) (%d,%d)               |\n", t9, t10, t11, t12, t13, t14, t15, t16)
		fmt.Println("|                                      |")
		winner(*list, match, &*stat, &*player)
	}
}

func balak(list *tiles, a, b int) int {
	var c int
	for i := a; i < b; i++ {
		if list[list[i].idxOut].n1 == list[list[i].idxOut].n2 {
			c++
			list[i].balak = list[i].idxOut
		}
	}
	return c
}

func scoring(stat *user, player *scoreboard, match int, condition string) {
	var score float64
	player[match].statistics[stat[match].round].status = condition
	stat[match].round++
	if condition == "Win" {
		fmt.Printf("|You Won!!                             |\n")
		stat[match].win++
		player[match].statistics[stat[match].round-1].roundS += 100
		score = player[match].statistics[stat[match].round-1].roundS
		player[match].score += score
	} else if condition == "Lose" {
		fmt.Printf("|You Lost!!                            |\n")
		player[match].statistics[stat[match].round-1].roundS = 0
	} else if condition == "Draw" {
		fmt.Printf("|Draw                                  |\n")
		player[match].statistics[stat[match].round-1].roundS += 75
		score = player[match].statistics[stat[match].round-1].roundS
		player[match].score += score
	}
	fmt.Printf("|This round score: %.f\n", score)
	fmt.Printf("|Total score: %.f\n", player[match].score)
	fmt.Println("----------------------------------------")
}

func winner(list tiles, match int, stat *user, player *scoreboard) {
	var playerScore, dealerScore, bsr, j int
	var condition string
	if balak(&list, 0, 8) == 0 {
		for i := 0; i < list[0].totIdx-4; i++ {
			playerScore += list[list[i].idxOut].n1 + list[list[i].idxOut].n2
		}
		for i := 4; i < list[0].totIdx; i++ {
			dealerScore += list[list[i].idxOut].n1 + list[list[i].idxOut].n2
		}
		if playerScore > dealerScore {
			condition = "Win"
		} else if playerScore < dealerScore {
			condition = "Lose"
		} else {
			condition = "Draw"
		}
	} else {
		if balak(&list, 0, 4) == 0 {
			condition = "Lose"
		} else if balak(&list, 4, 8) == 0 {
			condition = "Win"
		} else {
			balak(&list, 0, 8)
			for i := 0; i < 8; i++ {
				if list[i].balak > bsr {
					bsr = list[i].balak
					j = i
				}
			}
			if j >= 0 && j < 4 {
				condition = "Win"

			} else if j >= 4 && j < 8 {
				condition = "Lose"

			}
		}
	}
	scoring(&*stat, &*player, match, condition)
}

func tileList(list *tiles, n *int) {
	var t1, t2 int
	for t1 <= 6 && t2 <= 6 {
		for t1 != t2 {
			list[*n].n1 = t1
			list[*n].n2 = t2
			t2++
			*n++
		}
		if t1 == t2 {
			list[*n].n1 = t1
			list[*n].n2 = t2
			t1++
			t2 = 0
			*n++
		}
	}
}
