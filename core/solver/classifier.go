package solver

type Classifier interface  {
	State() string
	TrainOnText([]string, []string) error
	Solve() (string, error)
}