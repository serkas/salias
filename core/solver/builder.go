package solver


func NewNaiveClassifier(class []string) *NaiveClassifier {
	clf := NaiveClassifier{classes: class}
	return &clf
}