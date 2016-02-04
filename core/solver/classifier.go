package solver
import (
	"fmt"
	"salias/core/objects"
)


type Classifier interface  {
	State() string
	TrainOnText([]string, []string) error
	Solve() (string, error)
}

type NaiveClassifier struct {
	Info map[string]string
	classes []string
}

func (cls *NaiveClassifier) State() string {
	stat, ok := cls.Info["state"]
	if ok {
		return stat
	}
	return "empty"
}

func (cls *NaiveClassifier) TrainOnText(set []string, solved []string) error {
	var err error

	//priorClass := priors(solved)

	// Tokenize doc
	for _, str := range set {
		doc := objects.MakeDoc(str)
		fmt.Println(doc)
	}
	return err
}

func (cls *NaiveClassifier) Solve(task string) (string, error) {
	var err error

	return "0", err
}



func priors(solved []string) map[string]float32  {
	// Compute class (prior) probabilities
	classCounts := map[string]int{}
	for _, y := range solved {
		classCounts[y] = classCounts[y] + 1
	}
	priorClass := map[string]float32{}
	for c, count := range classCounts {
		priorClass[c] = float32(count)/float32(len(solved))
	}
	return priorClass
}

func termCount(set []string) (int, map[string]int) {
	counts := map[string]int{}
	var total int

	for _, sample := range set {
		doc := objects.MakeDoc(sample)
		docCounts := doc.CountTokens()

		for t, value := range docCounts {
			counts[t] += value
			total += value
		}
	}
	return total, counts
}
