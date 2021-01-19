package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
)

// Given ranked preferences for multiple users, allocate items in accordance with their
// preferences.  A random initial order is chosen for the customers.  Each round the
// order is reversed (the user who selected last selects first in the next round) in
// a snake draft.

const (
	resultsFileName = "out.txt"
	logFileName     = "out.log"
)

func main() {
	users := []string{"BB", "CD", "MG", "RW"}
	items := []string{
		"Dirty Bucket",
		"Hand of Glory",
		"Fruitcake",
		"Abyss",
		"Miller",
		"Oberon",
		"Pliny",
		"Sucaba",
		"The Lily",
		"Centaur",
		"11th Anniversary",
		"M-43",
		"Utopias",
		"Candy Mountain",
		"Ultimate",
		"Monkee",
		"XXIV",
		"Black Butte",
		"Roasted Rye",
		"Strawberry",
		"Peanut Butter",
	}
	prefs := [][]int{
		{0, 10, 6, 12, 13, 5, 11, 19, 7, 8, 16, 1, 20, 14, 3, 17, 15, 18, 4, 2, 9}, // BB
		{0, 1, 10, 11, 18, 7, 20, 15, 19, 2, 5, 4, 9, 12, 3, 8, 6, 14, 16, 13, 17}, // CD
		{0, 6, 16, 12, 1, 13, 15, 3, 5, 7, 4, 2, 14, 9, 17, 11, 10, 19, 8, 20, 18}, // MG
		{0, 6, 7, 8, 4, 9, 11, 19, 17, 15, 20, 16, 3, 1, 5, 12, 14, 18, 10, 2, 13}, // RW
	}

	// Confirm that all preference lists are the same length and contain the proper values
	if len(prefs) != len(users) {
		log.Fatalf("Found %d users but %d preferences", len(prefs), len(users))
	}
	for i, p := range prefs {
		if len(p) != len(items) {
			log.Fatalf("prefs[%d] contains %d elements, but there are %d items", i, len(p), len(items))
		}
		if !validPrefs(p) {
			log.Fatalf("prefs[%d] does not contain the numbers from 0-%d", i, len(p)-1)
		}
	}

	outFile, err := os.Create(resultsFileName)
	check(err)
	defer outFile.Close()
	out := bufio.NewWriter(outFile)

	logFile, err := os.Create(logFileName)
	check(err)
	defer logFile.Close()
	logOut := bufio.NewWriter(logFile)

	// Generate the order
	order := makeOrder(len(users))
	logOut.WriteString("Order: ")
	for _, i := range order {
		logOut.WriteString(users[i])
		logOut.WriteString(" ")
	}
	logOut.WriteString("\n\n")

	// Store the items selected for each user
	selected := make([][]int, len(users))
	for i := range selected {
		selected[i] = make([]int, 0)
	}
	// Record which items are still available
	taken := make(map[int]bool)

	for iter := range items {
		u := order[iter%len(order)]
		pick, idx := pickItem(prefs[u], taken)
		if pick < 0 || idx < 0 {
			log.Fatalf("No items available at iteration %d", iter)
		}
		logOut.WriteString(fmt.Sprintf("%s selected %s (%d)\n", users[u], items[pick], pick))
		selected[u] = append(selected[u], pick)

		// We know that all items before idx have been selected, so prune the list
		prefs[u] = prefs[u][idx+1:]
	}
	logOut.Flush()

	// Write the final selections
	for u, name := range users {
		out.WriteString(fmt.Sprintf("%s:\n", name))
		for _, i := range selected[u] {
			out.WriteString(fmt.Sprintf("%s (%d)\n", items[i], i))
		}
		out.WriteString("\n")
	}
	out.Flush()
}

// order generates the repeating order in which picks are selected.  It randomly assigns the
// initial order and returns it combined with a reverse (snake) order.  Thus, a four-user order
// would return something like [0,3,2,1,1,2,3,0]
func makeOrder(size int) []int {
	o := make([]int, size*2)
	for i := 0; i < size; i++ {
		o[i] = i
	}
	rand.Shuffle(size, func(i, j int) { o[i], o[j] = o[j], o[i] })
	for i := size; i < len(o); i++ {
		o[i] = o[len(o)-(i+1)]
	}
	return o
}

// validPrefs indicates if this slice contains all of the distinct values between 0 and its length
func validPrefs(p []int) bool {
	seen := make(map[int]bool)
	for _, i := range p {
		if i < 0 || i >= len(p) || seen[i] {
			return false
		}
		seen[i] = true
	}
	return true
}

// pickItem selects the first item available from the user's preferences.  Returns the item and
// its index.  Returns -1, -1 if no items are still available
func pickItem(prefs []int, taken map[int]bool) (int, int) {
	for i, p := range prefs {
		if !taken[p] {
			taken[p] = true
			return p, i
		}
	}
	return -1, -1
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
