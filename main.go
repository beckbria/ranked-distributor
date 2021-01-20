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
	items := []string{
		"Aardwolf - Early Bird Special",
		"Alesmith - Speedway Stout (Mokasida Coffee)",
		"Allagash - Interlude (2016)",
		"Angry Chair - When in Doubt",
		"Avery - Beast (2013)",
		"Bruery - 6 Geese a Laying",
		"Crooked Line - Labyrinth",
		"Cycle - Trademark Dispute (Brown/Vanilla)",
		"Cycle - Trademark Dispute (Green/Hazelnut)",
		"Cycle - Trademark Dispute (Yellow/Coffee+Cinnamon)",
		"Dark Horse - Plead the 5th",
		"de Garde - Anian",
		"de Garde - Apricot",
		"de Garde - Bluest",
		"de Garde - Bluest",
		"de Garde - Bu Weisse",
		"de Garde - Dark Harvest",
		"de Garde - Eponyme",
		"de Garde - Frais",
		"de Garde - Framboise",
		"de Garde - Framboise",
		"de Garde - Framboise",
		"de Garde - Framboise",
		"de Garde - Grand Blanc",
		"de Garde - Kriek",
		"de Garde - Kriek",
		"de Garde - Kriek",
		"de Garde - Kriek Premiere",
		"de Garde - L'Hiver Melange #2",
		"de Garde - Lee Kriek",
		"de Garde - Lee Kriek",
		"de Garde - Long Tooth",
		"de Garde - Petit Kriek",
		"de Garde - Purple",
		"de Garde - Purple Kriek",
		"de Garde - The Road",
		"Deschutes - Black Butte XXVI",
		"Deschutes - Mirror Mirror 2014",
		"Evil Twin - Imperial Biscotti Break",
		"Forager - Funky Dangerfield",
		"Forest & Main - Marius",
		"Fort George - Matroyshka",
		"Free State - Old Backus",
		"Fremont - Abominable (Can)",
		"Grassroots - Brother Soigne",
		"Grassroots - Brother Soigne",
		"Hair of the Dog - Doggie Claws 2013",
		"Haw River - No Holdsies Holdsie Uno",
		"Hill Farmstead - Elaborative",
		"Hill Farmstead - Florence",
		"Hill Farmstead - Nordic Saison",
		"Holy Mountain - Marionberry Table",
		"Holy Mountain - Temple of Heaven",
		"Homebrew Lambic",
		"Hoof Hearted - Skittley Bittley Bop",
		"Jackie O's - Funky South Paw",
		"Kane - Fall Saints",
		"Kane - Object Permanence",
		"Leelenau - Michilimackinad",
		"Lost Abbey - 15th Anniversary",
		"McKenzie - What Supp??",
		"Midnight Sun - Berserker",
		"Natty Greene's American Sour 2014",
		"Olde Hickory - Dread God",
		"Oso Tradeship (Blended)",
		"Oso Tradeship (Blood Orange)",
		"Pipeworks - Close Encounter",
		"Rare Barrel - Home Sour Home",
		"Russian River - Consecration",
		"Russian River - Supplication",
		"Sante Adairius/Cycle - Tandem",
		"Side Project - Pulling Nails #6",
		"Ska - Taster's Choice",
		"Surly - Darkness",
		"Surly - Darkness (Rye Barrels)",
		"Tired Hands - Blourison",
		"Tired Hands - Parageusia",
		"Tired Hands - Permashore",
		"Tired Hands - Whatever Nevermind",
		"Toppling Goliath - Kaiju Clash",
		"Toppling Goliath - Norseman's Wrath",
		"Treehoues - Simple Life",
		"Westbrook - 6th Anniversary (Bourbon Barrel)",
		"Westbrook - 6th Anniversary (Cabernet Barrel)",
		"Westbrook/Evil Twin - Mini Growler",
		"Wynwood - Pop's Porter",
	}
	users := []string{"BB", "CD", "MG" "RW", "CM"}
	prefs := [][]int{
		// BB
		{79, 51, 50, 65, 64, 49, 48, 57, 23, 17, 28, 74, 45, 71, 8, 7, 67, 30, 82, 81, 83, 52, 33, 11, 76, 75, 70, 61, 53, 18, 58, 29, 31, 27, 12, 14, 34, 3, 21, 44, 9, 69, 68, 46, 73, 16, 15, 13, 35, 32, 22, 47, 0, 4, 19, 55, 59, 77, 40, 62, 63, 1, 2, 24, 37, 39, 20, 25, 42, 56, 78, 84, 26, 36, 10, 6, 5, 60, 41, 80, 66, 38, 54, 72, 85, 43},
		// CD
		{74, 63, 7, 8, 73, 6, 82, 42, 59, 10, 47, 71, 3, 2, 83, 57, 40, 58, 41, 67, 80, 52, 16, 23, 17, 18, 31, 35, 33, 34, 65, 64, 85, 51, 11, 12, 15, 79, 9, 48, 49, 50, 56, 78, 4, 45, 39, 54, 75, 44, 55, 60, 14, 77, 62, 66, 70, 72, 76, 81, 0, 84, 38, 5, 68, 46, 37, 36, 69, 61, 1, 13, 29, 27, 25, 32, 20, 28, 24, 26, 30, 21, 19, 22, 53, 43},
		// MG
		{4, 27, 74, 20, 3, 57, 39, 79, 84, 0, 69, 14, 9, 1, 6, 60, 80, 73, 21, 13, 78, 22, 24, 61, 18, 23, 19, 5, 2, 31, 32, 85, 33, 10, 25, 40, 26, 46, 54, 41, 42, 49, 44, 34, 47, 29, 72, 50, 45, 48, 15, 55, 56, 59, 35, 17, 7, 30, 28, 62, 38, 63, 66, 58, 67, 68, 70, 75, 81, 76, 71, 77, 82, 12, 83, 16, 8, 37, 36, 65, 64, 51, 52, 43, 11, 53},
		// RW
		{17, 33, 30, 21, 34, 57, 26, 14, 56, 37, 7, 10, 47, 22, 60, 82, 61, 68, 59, 45, 3, 16, 42, 27, 46, 65, 50, 40, 41, 63, 39, 58, 52, 69, 55, 8, 75, 71, 32, 77, 13, 78, 51, 11, 23, 31, 35, 49, 53, 62, 67, 29, 19, 20, 24, 25, 12, 4, 15, 18, 36, 43, 64, 70, 83, 66, 2, 6, 44, 48, 76, 81, 5, 28, 74, 80, 84, 85, 73, 0, 1, 9, 38, 72, 79, 54},
		// CM
		{12, 14, 17, 34, 21, 73, 47, 56, 50, 78, 27, 74, 52, 71, 16, 42, 67, 33, 3, 32, 2, 44, 28, 39, 72, 31, 35, 48, 49, 51, 40, 57, 45, 58, 60, 76, 63, 70, 18, 11, 15, 59, 64, 75, 68, 69, 77, 65, 23, 61, 66, 62, 6, 81, 79, 83, 80, 54, 7, 82, 13, 29, 20, 22, 19, 30, 24, 26, 25, 41, 4, 8, 9, 10, 37, 1, 84, 46, 5, 0, 38, 53, 36, 85, 55, 43},
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
