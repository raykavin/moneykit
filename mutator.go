package moneykit

type mutator struct {
	calc *calculator
}

var mutate = mutator{calc: &calculator{}}
