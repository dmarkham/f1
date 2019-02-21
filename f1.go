package f1

import (
	"math"
)

// F1 builds the F1 score.
// https://en.wikipedia.org/wiki/F1_score
type F1 struct {
	TruePositive float64
	Type1Error   float64 // false Positive
	TrueNegative float64
	Type2Error   float64 // false Negative
}

// New returns a empty F1 Struct
// this struct is *not* thread safe please protect it as you see fit
func New() *F1 {
	return &F1{}
}

// Score will let you weight the F1 Score by beta
// Two commonly used beta's are the 2.0 measure,
// which weighs recall higher than precision (by placing more emphasis on false negatives),
// and the 0.5 measure, which weighs recall lower than precision
// (by attenuating the influence of false negatives).
func (f *F1) Score(beta float64) float64 {
	return f.TruePositive * (1 + math.Pow(beta, 2)) /
		(((1 + math.Pow(beta, 2)) * f.TruePositive) + (math.Pow(beta, 2) * f.Type2Error) + f.Type1Error)
}

// this should match Score(1)
// this is here just for testing Score
func (f *F1) score2() float64 {
	return 2 * ((f.Precision() * f.Recall()) / (f.Precision() + f.Recall()))
}

// Precision f.TruePositive / (f.TruePositive + f.Type1Error)
// high precision means that an algorithm returned substantially
// more relevant results than irrelevant ones
func (f *F1) Precision() float64 {
	return f.TruePositive / (f.TruePositive + f.Type1Error)
}

//Recall f.TruePositive / (f.TruePositive + f.Type2Error)
// high recall means that an algorithm returned most of the relevant results.
func (f *F1) Recall() float64 {
	return f.TruePositive / (f.TruePositive + f.Type2Error)
}
