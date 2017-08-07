package main

import (
	"fmt"
	"math/rand"
	"time"
)
func main() {
	jokes := []string{
		"what did the tree wear to the pool? swimming trunks!",
		"when nothing goes right, go left!",
		"I don't need a hair stylist, my pillow gives me a new hairstyle evevry morning",
		"my computer once beat me at chess. but it was no match for me at kick boxing",
		"my room is like the bermuda triangle, stuff goes in and is never seen again",
		"life always offers you a second chance. it's called tommorow.",
		"my bed is a magical place where I suddenly remember everything I forgot to do",
		"the future is shaped by your dreams, so at night stop wasting time and go to sleep",
		"they say don't try this at home, so I am coming over to your house to try it.",
		"lazy people fact #39841297423, you were too lazy to read that number.",
		"what did the cat get on his test? A purrfect score.",
		"I would like to live like a poor man, only with lots of money.",
		"I wish my wallet came with free refills",
		"I did not trip and fall. I attacked the floor and I think I am winning.",
		"A cop pulled me over and told me 'papers', so I said 'scissors, I win!' and drove off",
		"Isn't it funny how red white and blue represent freedom, unless they are flashing behind you?",
		"Friends are like a four leaf clover hard to find and lucky to have",
		"for the best seat in the house you have to have a dog to help you find it",
	}
	rand.Seed(time.Now().Unix())
	quote := jokes[rand.Intn(len(jokes))]
	fmt.Println(quote)
}