package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Create a new type of 'deck'
// which is a slice of strings.
type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	// _ is the index value which isn't necessary, so since we're discarding it's use. Use a "_"
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

func (d deck) print() {
	for i, card := range d {
		//In this case, we *are* interested in the index value, hence leaving it in.
		fmt.Println(i, card)
	}
}

//This function takes in two values, returns two decks.
func deal(d deck, handSize int) (deck, deck) {
	//Return value 1 is a subset of deck starting at 0, ending at handsize
	//Return value 2 is subset of deck starting at handSize to end.
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	//Concatenate strings, seperate with ","
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		//Option #1 - log the error and return a call to newDeck()

		//Option #2 - log the error and entirely quit the program.
		fmt.Println("Error:", err)
		os.Exit(1)

	}

	s := strings.Split(string(bs), ",")
	return deck(s)
}

// func (d deck) shuffle() {
// 	for i := range d {
// 		newPosition = rand.Intn(len(d) - 1)
// 		d[i], d[newPosition] = d[newPosition], d[i]
// 	}
// }

func (d deck) shuffle() {
	//Pick a randomized source, this time being the time.
	source := rand.NewSource(time.Now().UnixNano())
	//r is now the randomizer.
	r := rand.New(source)

	//For each card in the deck.
	for i := range d {
		//pick a random Integer between 0 and length of deck - 1.
		newPosition := r.Intn(len(d) - 1)
		//Switch current position of i with random card
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
