package solver

import (
	"salias/core/objects"
)

type NaiveClassifier struct {
	Info    map[string]string
	classes []string
}

type TermClass struct {
	term  string
	class string
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
	//priorP := priors(solved)
	//classes := byClass(set, solved)
	return err
}


func (cls *NaiveClassifier) Solve(task string) (string, error) {
	var err error
	return "0", err
}


func priors(solved []string) map[string]float64 {
	// Compute class (prior) probabilities
	classCounts := map[string]int{}
	for _, y := range solved {
		classCounts[y] = classCounts[y] + 1
	}
	priorClass := map[string]float64{}
	for c, count := range classCounts {
		priorClass[c] = float64(count) / float64(len(solved))
	}
	return priorClass
}

func conditionals(set []string, solved []string) (map[TermClass]float64) {
	prob := map[TermClass]float64{}

	vocab := vocabulary(set)
	vSize := len(vocab)

	classes := byClass(set, solved)

	for class, terms := range classes {
		countClass, inClassCounts := termCount(terms)

		for _, term := range vocab {
			termCount := int(inClassCounts[term])
			value := float64(termCount + 1) / float64(countClass + vSize)
			prob[TermClass{term, class}] = value
		}
	}
	
	return prob
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

func byClass(trainTask  []string, trainSolution []string) (map[string][]string) {
	classes := map[string][]string{}

	for i, taskStr := range trainTask {
		cls := trainSolution[i]
		classes[cls] = append(classes[cls], taskStr)
	}
	return classes
}

func vocabulary(set []string) []string {
	vocab := []string{}
	_, allByTerm := termCount(set)
	for term, _ := range allByTerm {
		vocab = append(vocab, term)
	}
	return vocab
}
