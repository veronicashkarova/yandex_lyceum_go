package main

type Student struct {
	name (string)
	solvedProblems (int)
	scoreForOneTask (float64)
	passingScore (float64)
}

func (p Student) IsExcellentStudent () bool {
	return (float64(p.solvedProblems) * p.scoreForOneTask) >= p.passingScore
}