package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"sort"
)

// TODO: clean up the code
func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var T int

	if scanner.Scan() {
		input := scanner.Text()
		T = getInt(input)
	}

	for t := 0; t < T; t++ {
		var N int
		if scanner.Scan() {
			input := scanner.Text()
			N = getInt(input)
		}

		activities := make([]activity, N)
		for n := 0; n < N; n++ {
			if scanner.Scan() {
				input := strings.Split(scanner.Text(), " ")
				activities[n] = activity{startTime: getInt(input[0]), endTime: getInt(input[1]), index: n}
			}
		}

		sort.Sort(byStartTime(activities))
		endC, endJ := 0, 0
		var outputStr string
		for i := range activities {
			if activities[i].startTime >= endC {
				activities[i].assignee = 'C'
				endC = activities[i].endTime
			} else if activities[i].startTime >= endJ {
				activities[i].assignee = 'J'
				endJ = activities[i].endTime
			} else {
				outputStr = "IMPOSSIBLE"
				break
			}
		}

		if outputStr == "" {
			sort.Sort(byIndex(activities))
			builder := make([]byte, len(activities))
			for i := range activities {
				builder[i] = activities[i].assignee
			}
			outputStr = string(builder)
		}

		fmt.Printf("Case #%d: %s\n", t+1, outputStr)

	}
}

func getInt(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		panic(err.Error())
	}
	return n
}

type activity struct {
	startTime, endTime int
	assignee byte
	index int
}

type byIndex []activity

func (a byIndex) Len() int           { return len(a) }
func (a byIndex) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byIndex) Less(i, j int) bool { return a[i].index < a[j].index }

type byStartTime []activity

func (a byStartTime) Len() int           { return len(a) }
func (a byStartTime) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byStartTime) Less(i, j int) bool { return a[i].startTime < a[j].startTime }
