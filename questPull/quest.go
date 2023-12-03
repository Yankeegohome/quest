package questPull

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func QuestPull(filename string) ([]quest, error) {
	if fObj, err := os.Open(filename); err == nil {
		csvR := csv.NewReader(fObj)
		if cLines, err1 := csvR.ReadAll(); err1 == nil {
			return ParseQuestions(cLines), nil
		}
		return nil, fmt.Errorf("error reading data %s in scv - %s", err.Error(), filename)
	} else {
		return nil, fmt.Errorf("error opennig csv file - %s ", filename, err.Error())
	}
}

func ParseQuestions(lines [][]string) []quest {
	r := make([]quest, len(lines))
	for i := 0; i < len(lines); i++ {
		r[i] = quest{Q: lines[i][0], A: lines[i][1]}
	}
	log.Println(r[0])
	return r
}

type quest struct {
	Q string
	A string
}
