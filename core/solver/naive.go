package solver

import (
	"github.com/serkas/salias/core/objects"
	"math"
)

type NaiveClassifier struct {
	Info         map[string]string
	classes      []string
	priors       map[string]float64
	conditionals map[TermClass]float64
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
	cls.priors = priors(solved)
	cls.conditionals = conditionals(set, solved)
	return err
}

func (cls *NaiveClassifier) Solve(task string) (string, error) {
	var err error
	var result string
	bestScore := -math.MaxFloat64

	scores := cls.classScores(task)

	for class, score := range scores {
		println(class, score)
		if score > bestScore {
			result = class
			bestScore = score
		}
	}
	return result, err
}

func (cls *NaiveClassifier) classScores(task string) map[string]float64 {
	scores := map[string]float64{}

	_, terms := termCount([]string{task})

	for _, class := range cls.classes {
		termRelevance := 0.0

		for term, termCount := range terms {
			if cProb, ok := cls.conditionals[TermClass{term, class}]; ok {
				logProb := math.Log10(cProb)
				termRelevance = termRelevance + logProb*float64(termCount)
			}

		}
		scores[class] = math.Log10(cls.priors[class]) + termRelevance
	}
	return scores
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

func conditionals(set []string, solved []string) map[TermClass]float64 {
	prob := map[TermClass]float64{}

	vocab := vocabulary(set)
	vSize := len(vocab)

	classes := byClass(set, solved)

	for class, terms := range classes {
		countClass, inClassCounts := termCount(terms)

		for _, term := range vocab {
			termCount := int(inClassCounts[term])
			value := float64(termCount+1) / float64(countClass+vSize)
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

func byClass(trainTask []string, trainSolution []string) map[string][]string {
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
