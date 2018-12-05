package main

import (
	"io/ioutil"
	"strings"
	"time"
	"fmt"
	"strconv"
	"sort"
)

type GuardLog struct {
	time time.Time
	note string
}

var layout = "2006-01-02 15:04"
/*
	Only part 1
 */
func main() {
	file, _ := ioutil.ReadFile("4th/input.txt")
	fileSlice := strings.Split(string(file), "\r\n")

	guardLogs := []GuardLog{}


	fmt.Println("Startin loop at line 32")
	for _, line := range fileSlice {
		line =strings.TrimPrefix(line, "[")
		splittedLine := strings.SplitAfter(line, "]")
		timestamp := strings.TrimSuffix(splittedLine[0], "]")
		note := strings.TrimSpace(splittedLine[1])
		time, err := time.Parse(layout, timestamp)
		if err != nil {
			fmt.Println("error:", err)
		}
		//fmt.Println(time.Minute())
		guardLog := GuardLog{time, note}
		guardLogs = append(guardLogs, guardLog)
	}
	//Sorterer alle linjene i rett rekkefølge
	for i := 0; i < len(guardLogs); i++ {
		for j := 0; j < len(guardLogs)-1; j++ {
			if guardLogs[j].time.After(guardLogs[j+1].time) {
				switched := guardLogs[j+1]
				guardLogs[j+1] = guardLogs[j]
				guardLogs[j] = switched

			}
		}
	}

	//Registrerer hver vakt med de minuttene han har sovet på
	var guardId int
	guardsAsleep := make(map[int][]int)
	for i := 0; i < len(guardLogs); i++{
		note := guardLogs[i].note
		if strings.HasPrefix(note, "G") {
			guardId,_ = strconv.Atoi(strings.TrimSpace(strings.SplitAfter(strings.SplitAfter(guardLogs[i].note, "#")[1]," ")[0]))
		} else if strings.HasPrefix(note, "f") {
			startSleep := guardLogs[i].time.Minute()
			endSleep := guardLogs[i+1].time.Minute()
			for i := startSleep; i < endSleep; i++ {
				guardsAsleep[guardId] = append(guardsAsleep[guardId], i)
			}
		}
	}

	//Finner den vakten som har sovet mest på jobben
	biggestSleeper := 0
	idOfBiggest := 0
	for i, e := range guardsAsleep {
		if len(e) > biggestSleeper {
			biggestSleeper = len(e)
			idOfBiggest = i
		}
	}

	//Finner de minuttene vakten har sovet på
	minutesSleptOn :=[]int{}
	for  i := 0; i < 60; i++ {
		for j := 0; j < len(guardsAsleep[idOfBiggest]); j++ {
			if guardsAsleep[idOfBiggest][j] == i {
				minutesSleptOn = append(minutesSleptOn, guardsAsleep[idOfBiggest][j])
				j = len(guardsAsleep[idOfBiggest])
			}
		}
	}
	//Finner ut hvor mange ganger vakten har sovet på hvert minutt
	minutesMap := map[int]int{}
	for i := 0; i < len(minutesSleptOn); i++ {
		for j := 0; j < len(guardsAsleep[idOfBiggest]); j++ {
			if guardsAsleep[idOfBiggest][j] == minutesSleptOn[i] {
				minutesMap[minutesSleptOn[i]]++
			}
		}
	}

	//Legger til frekvensene av hvert minutt i ei liste og sorterer lista
	listOfFrequncies := []int{}
	for _, v := range minutesMap {
		listOfFrequncies = append(listOfFrequncies, v)
	}
	sort.Ints(listOfFrequncies)

	//Finner det minuttet som har høyest frekvens
	desiredMinute := 0
	for k, v := range minutesMap {
		if v == listOfFrequncies[len(listOfFrequncies)-1] {
			desiredMinute = k
		}
	}
	fmt.Println("THE ID: ", idOfBiggest, "* THE SAFEST MINUTE: ", desiredMinute, "= ", idOfBiggest*desiredMinute)
}
