package solver

import (
	"testing"
	"math"
)


func TestVocabulary(t *testing.T)  {
	task := []string{"a", "a", "a"}
	v := vocabulary(task)

	if (v[0] != "a") {
		t.Error("Expected `a` as first vocabulary entry, got", v[0])
	}

	if (len(v) != 1) {
		t.Error("Expected vocabulary of size 1, got", len(v))
	}

	task2 := []string{"a", "b", "c", "d", "a", "b"}
	v2 := vocabulary(task2)

	if (len(v2) != 4) {
		t.Error("Expected vocabulary of size 4, got", len(v2))
	}

}
func TestPrior(t *testing.T) {
	trainSolution := []string{"1", "1", "1", "1", "0"}
	pr := priors(trainSolution)

	if (pr["0"] != 0.2 || pr["1"] != 0.8) {
		t.Error("Expected map[1:0.8 0:0.2] for,", trainSolution, ", got", pr)
	}
}

func TestTermCounts(t *testing.T) {
	trainTask := []string{"a b", "b c", "a c", "d e", "d d e f"}

	total, counts := termCount(trainTask)

	if (counts["a"] != 2) {
		t.Error("Expected 2 occurences of `a`, got", counts["a"])
	}
	if (counts["d"] != 3) {
		t.Error("Expected 3 occurences of `d`, got", counts["d"])
	}

	if (total != 12) {
		t.Error("Expected 12 terms, got", total)
	}
}

func TestByClass(t *testing.T) {
	task := []string{"a", "b", "c"}
	solution := []string{"0", "1", "0"}

	classes := byClass(task, solution)

	if (classes["0"][0] != "a") {
		t.Error("First doc in class `0` expcted `a`, got", classes)
	}
}


func TestConditionals(t *testing.T) {
	task := []string{"a b", "b c c", "a a b"}
	solution := []string{"0", "1", "0"}

	prob := conditionals(task, solution)
	a0Prob := prob[TermClass{"a", "0"}]
	a1Prob := prob[TermClass{"a", "1"}]
	b0Prob := prob[TermClass{"b", "0"}]

	if (math.Abs(a0Prob - float64(0.5)) > 0.05) {
		t.Error("Conditional probability for class `0` and term `a` expected ~0.5, got", a0Prob)
	}

	if (math.Abs(a1Prob - float64(0.1666)) > 0.05) {
		t.Error("Conditional probability for class `1` and term `a` expected ~0.1666, got", a1Prob)
	}

	if (math.Abs(float64(b0Prob - 0.375)) > 0.05) {
		t.Error("Conditional probability for class `1` and term `a` expected ~0.375, got", b0Prob)
	}
}


func TestClassScores(t *testing.T) {
	cls := NewNaiveClassifier([]string{"0", "1"})

	task := []string{"a b", "b c c", "a a b"}
	solution := []string{"0", "1", "0"}

	cls.TrainOnText(task, solution)

	test := "a b c"
	scores := cls.classScores(test)


	if (math.Abs(float64(scores["0"] - 0.015625)) > 0.01) {
		t.Error("Class `0` proportional probability wrong expected ~0.015625, got", scores["0"])
	}

	if (math.Abs(float64(scores["1"] - 0.00925)) > 0.01) {
		t.Error("Class `1` proportional probability wrong expected ~0.00925, got", scores["1"])
	}

}