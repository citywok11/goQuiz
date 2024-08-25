package quiz

import (
	"bufio"
	"fmt"
	"goQuiz/internal/models"
	"os"
	"strconv"
	"strings"
)

func displayQuestion(q models.Question) {
	fmt.Println(q.Text)
	for i, option := range q.Options {
		fmt.Printf("%d. %s\n", i+1, option)
	}
}

func promptForAnswer(minOption, maxOption int) int {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Printf("Enter your answer (%d-%d): ", minOption, maxOption)
		scanner.Scan()
		input := strings.TrimSpace(scanner.Text())

		answer, err := strconv.Atoi(input)
		if err != nil || answer < minOption || answer > maxOption {
			fmt.Printf("Invalid input. Please enter a number between %d and %d.\n", minOption, maxOption)
			continue
		}
		return answer
	}
}