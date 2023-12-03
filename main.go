package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"quest/questPull"
	"strconv"
	"sync"
	"time"
)

func main() {
	wg := sync.WaitGroup{}
	fname := flag.String("f", "questPull/quest.csv", "path for csv file")
	timer := flag.Int("t", 30, "timer of the quiz")
	flag.Parse()
	questions, err := questPull.QuestPull(*fname)
	if err != nil {
		exit(fmt.Sprintf("500"))
	}
	ansC := make(chan string)
	correctAns := 0
	tObj := time.NewTimer(time.Duration(*timer) * time.Second)
	wg.Add(4)
questionsLoop:
	for i, quest := range questions {
		var answer string
		fmt.Printf("Question %d: %s=", i+1, quest.Q)
		go func() {
			defer wg.Done()
			fmt.Sprintf(strconv.Itoa(i))
			fmt.Scanf("%s", &answer)
			ansC <- answer

		}()
		select {
		case <-tObj.C:
			log.Println("ok")
			break questionsLoop
		case iAns := <-ansC:
			if iAns == quest.A {
				correctAns++
				fmt.Println("Dalabeb kogda ti v perviy raz suda zahodish")
			}
			if i == len(questions) {
				close(ansC)
			}
			continue
		}
	}
	fmt.Printf("Your result is %d, out of %d\n", correctAns, len(questions))
	fmt.Printf("Press enter to exit")
	<-ansC
	wg.Wait()
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
