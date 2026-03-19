package lab1

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

type Leaderboard struct {
	cityWealth    map[string]int64
	wealthCities  map[int64]map[string]bool
	sortedWealths []int64
}

func NewLeaderboard() *Leaderboard {
	return &Leaderboard{
		cityWealth:    make(map[string]int64),
		wealthCities:  make(map[int64]map[string]bool),
		sortedWealths: make([]int64, 0),
	}
}

func (l *Leaderboard) UpdateCity(city string, delta int64) {
	oldW := l.cityWealth[city]
	newW := oldW + delta

	if oldW > 0 {
		delete(l.wealthCities[oldW], city)
		if len(l.wealthCities[oldW]) == 0 {
			delete(l.wealthCities, oldW)
			l.removeSorted(oldW)
		}
	}

	l.cityWealth[city] = newW

	if newW > 0 {
		if l.wealthCities[newW] == nil {
			l.wealthCities[newW] = make(map[string]bool)
		}
		if len(l.wealthCities[newW]) == 0 {
			l.insertSorted(newW)
		}
		l.wealthCities[newW][city] = true
	}
}

func (l *Leaderboard) insertSorted(w int64) {
	idx := sort.Search(len(l.sortedWealths), func(i int) bool { return l.sortedWealths[i] >= w })
	l.sortedWealths = append(l.sortedWealths, 0)
	copy(l.sortedWealths[idx+1:], l.sortedWealths[idx:])
	l.sortedWealths[idx] = w
}

func (l *Leaderboard) removeSorted(w int64) {
	idx := sort.Search(len(l.sortedWealths), func(i int) bool { return l.sortedWealths[i] >= w })
	l.sortedWealths = append(l.sortedWealths[:idx], l.sortedWealths[idx+1:]...)
}

func (l *Leaderboard) GetLeader() (string, bool) {
	if len(l.sortedWealths) == 0 {
		return "", false
	}
	maxW := l.sortedWealths[len(l.sortedWealths)-1]
	cities := l.wealthCities[maxW]

	if len(cities) == 1 {
		for city := range cities {
			return city, true
		}
	}
	return "", false
}

type Person struct {
	city   string
	wealth int64
}

type Move struct {
	person string
	toCity string
}

func TwentytwoMain() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	readWord := func() string {
		scanner.Scan()
		return scanner.Text()
	}

	readInt := func() int {
		scanner.Scan()
		val, _ := strconv.Atoi(scanner.Text())
		return val
	}

	readInt64 := func() int64 {
		scanner.Scan()
		val, _ := strconv.ParseInt(scanner.Text(), 10, 64)
		return val
	}

	if !scanner.Scan() {
		return
	}
	n, _ := strconv.Atoi(scanner.Text())

	people := make(map[string]*Person)
	board := NewLeaderboard()

	for i := 0; i < n; i++ {
		name := readWord()
		city := readWord()
		wealth := readInt64()
		people[name] = &Person{city: city, wealth: wealth}
		board.UpdateCity(city, wealth)
	}

	m := readInt()
	k := readInt()

	movesByDay := make(map[int][]Move)
	var activeDays []int

	for i := 0; i < k; i++ {
		day := readInt()
		person := readWord()
		toCity := readWord()

		if len(movesByDay[day]) == 0 {
			activeDays = append(activeDays, day)
		}
		movesByDay[day] = append(movesByDay[day], Move{person: person, toCity: toCity})
	}

	sort.Ints(activeDays)

	cityWins := make(map[string]int)
	startDay := 1

	for _, day := range activeDays {
		if day >= startDay {
			daysElapsed := day - startDay + 1
			leader, unique := board.GetLeader()
			if unique {
				cityWins[leader] += daysElapsed
			}
		}

		for _, move := range movesByDay[day] {
			p := people[move.person]
			if p.city != move.toCity {
				board.UpdateCity(p.city, -p.wealth)
				board.UpdateCity(move.toCity, p.wealth)
				p.city = move.toCity
			}
		}

		startDay = day + 1
	}

	if m >= startDay {
		daysElapsed := m - startDay + 1
		leader, unique := board.GetLeader()
		if unique {
			cityWins[leader] += daysElapsed
		}
	}

	var winningCities []string
	for city, wins := range cityWins {
		if wins > 0 {
			winningCities = append(winningCities, city)
		}
	}

	sort.Strings(winningCities)

	writer := bufio.NewWriter(os.Stdout)
	for _, city := range winningCities {
		fmt.Fprintf(writer, "%s %d\n", city, cityWins[city])
	}
	writer.Flush()
}
