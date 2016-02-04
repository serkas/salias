package core


import (
	"testing"
	"salias/core/classifier"
)

func TestNewClassifier(t *testing.T) {

	cls := new(classifier.Classifier)
	if(cls.State() != "empty"){
		t.Error("Expected classifier in `empty` state, got", cls.State())
	}
}

func TestNewClassifierFromBuilder(t *testing.T) {

	cls := classifier.NewClassifier()
	if(cls.State() != "empty"){
		t.Error("Expected classifier in `empty` state, got", cls.State())
	}
}

