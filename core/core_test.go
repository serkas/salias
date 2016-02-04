package core


import (
	"testing"
	"salias/core/solver"
)

func TestNewClf(t *testing.T) {

	cls := solver.NewNaiveClassifier([]string{"a"})
	if(cls.State() != "empty"){
		t.Error("Expected Classifire in `empty` state, got", cls.State())
	}
}

func TestNewClfFromBuilder(t *testing.T) {
	cls := solver.NewNaiveClassifier([]string{"a"})
	if(cls.State() != "empty"){
		t.Error("Expected Classifire in `empty` state, got", cls.State())
	}
}

func TestBasicClassify(t *testing.T) {
	trainTask := []string{"a b", "b c", "a c", "d e", "d d e f"}
	trainSolution := []string{"1", "1", "1", "0", "0"}

	cls := solver.NewNaiveClassifier([]string{"a"})

	_ = cls.TrainOnText(trainTask, trainSolution)

	answer, _ :=cls.Solve("a b c")

	if(answer != "1"){
		t.Error("Classifire fails on easy task")
	}
}
