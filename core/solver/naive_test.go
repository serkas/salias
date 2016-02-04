package solver

import "testing"



func TestPrior(t *testing.T) {
	trainSolution := []string{"1", "1", "1", "1", "0"}
	pr := priors(trainSolution)

	if(pr["0"] != 0.2 || pr["1"] != 0.8){
		t.Error("Expected map[1:0.8 0:0.2] for,", trainSolution,", got", pr)
	}
}

func TestTermCounts(t *testing.T) {
	trainTask := []string{"a b", "b c", "a c", "d e", "d d e f"}

	total, counts := termCount(trainTask)

	if(counts["a"] != 2){
		t.Error("Expected 2 occurences of `a`, got", counts["a"])
	}
	if(counts["d"] != 3){
		t.Error("Expected 3 occurences of `d`, got", counts["d"])
	}

	if(total != 12){
		t.Error("Expected 12 terms, got", total)
	}
}