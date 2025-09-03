package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Total Player: ")
	totalPlayer, err := inputTotalPlayer(scanner)
	if err != nil {
		fmt.Println("Error: Invalid input for total number of players:", err.Error())
		return
	}

	fmt.Print("List Score (Separate By Space): ")
	listScore, err := inputListScore(scanner, totalPlayer)
	if err != nil {
		fmt.Println("Error: Invalid score list input:", err.Error())
		return
	}

	fmt.Print("Total GITS Player: ")
	totalGitsPlayer, err := inputTotalPlayer(scanner)
	if err != nil {
		fmt.Println("Error: Invalid input for total GITS players:", err.Error())
		return
	}

	fmt.Print("List GITS Score (Separate By Space): ")
	listGitsScore, err := inputListScore(scanner, totalGitsPlayer)
	if err != nil {
		fmt.Println("Error: Invalid GITS score list input:", err.Error())
		return
	}

	gitsPlayerRanking, err := denseRanking(listScore, listGitsScore)
	if err != nil {
		fmt.Println("Error: Failed to calculate player ranking:", err.Error())
	}

	// Convert Gits Player Ranking To Array String For Print Output
	gitsPlayerRankingStr := make([]string, len(gitsPlayerRanking))
	for i, rank := range gitsPlayerRanking {
		gitsPlayerRankingStr[i] = strconv.Itoa(rank)
	}

	fmt.Println("GITS Player Ranking: ", strings.Join(gitsPlayerRankingStr, " "))
}

// Read Total Player From Scanner And Convert To Int
func inputTotalPlayer(scanner *bufio.Scanner) (int, error) {
	scanner.Scan()
	return strconv.Atoi(scanner.Text())
}

// Read List Score From Scanner And Convert To Array Int
func inputListScore(scanner *bufio.Scanner, totalPlayer int) ([]int, error) {
	scanner.Scan()
	listScoreStr := strings.Fields(scanner.Text())

	// List Score Lenght Must Be Exaxly Same With Total Player
	if len(listScoreStr) != totalPlayer {
		return nil, fmt.Errorf("input mismatch: expected %v scores, but received %v", totalPlayer, len(listScoreStr))
	}

	// Convert From Array String To Array Int
	listScore := make([]int, totalPlayer)
	for i, value := range listScoreStr {
		score, err := strconv.Atoi(value)
		if err != nil {
			return nil, fmt.Errorf("invalid score format at position %d: %v", i+1, value)
		}

		listScore[i] = score
	}

	return listScore, nil
}

// Calculate Dense Ranking From List Score And Gits Score
func denseRanking(listScore []int, gitsScore []int) ([]int, error) {
	sort.Slice(listScore, func(i, j int) bool {
		return listScore[i] > listScore[j]
	})

	// Build unique descending scores
	uniqueScore := []int{}
	for _, score := range listScore {
		if len(uniqueScore) == 0 || uniqueScore[len(uniqueScore)-1] != score {
			uniqueScore = append(uniqueScore, score)
		}
	}

	result := make([]int, 0, len(gitsScore))
	// Find Gits Score
	for _, gs := range gitsScore {
		idx := sort.Search(len(uniqueScore), func(i int) bool {
			return uniqueScore[i] <= gs
		})
		rank := idx + 1
		result = append(result, rank)
	}

	return result, nil
}
