package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

var question = `The Elf finishes helping with the tent and sneaks back over to you. "Anyway, the second column says how the round needs to end: X means you need to lose, 
Y means you need to end the round in a draw, and Z means you need to win. Good luck!"

The total score is still calculated in the same way, but now you need to figure out what shape to choose so the round ends as indicated. The example above now goes like this:

In the first round, your opponent will choose Rock (A), and you need the round to end in a draw (Y), so you also choose Rock. This gives you a score of 1 + 3 = 4.
In the second round, your opponent will choose Paper (B), and you choose Rock so you lose (X) with a score of 1 + 0 = 1.
In the third round, you will defeat your opponent's Scissors with Rock for a score of 1 + 6 = 7.
Now that you're correctly decrypting the ultra top secret strategy guide, you would get a total score of 12.

Following the Elf's instructions for the second column, what would your total score be if everything goes exactly according to your strategy guide?`

var WinMap = map[string]string{
	"A": "Y",
	"B": "Z",
	"C": "X",
}

var LoseMap = map[string]string{
	"A": "Z",
	"B": "X",
	"C": "Y",
}

var DrawMap = map[string]string{
	"A": "X",
	"B": "Y",
	"C": "Z",
}

var winScore = 6
var drawScore = 3
var rockScore = 1
var papeScore = 2
var scissorsScore = 3

var scoreMap = map[string]int{
	"A X": rockScore + drawScore,
	"A Y": papeScore + winScore,
	"A Z": scissorsScore,
	"B X": rockScore,
	"B Y": papeScore + drawScore,
	"B Z": scissorsScore + winScore,
	"C X": rockScore + winScore,
	"C Y": papeScore,
	"C Z": scissorsScore + drawScore,
}

func main() {
	input, err := ioutil.ReadFile("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	plays := strings.Split(string(input), "\n")
	score := 0

	for _, round := range plays {
		play := strings.Split(round, " ")
		if play[1] == "X" {
			play[1] = LoseMap[play[0]]
		} else if play[1] == "Z" {
			play[1] = WinMap[play[0]]
		} else if play[1] == "Y" {
			play[1] = DrawMap[play[0]]
		}
		value := strings.Join(play, " ")

		score = score + scoreMap[value]
	}

	fmt.Println(score)
}
