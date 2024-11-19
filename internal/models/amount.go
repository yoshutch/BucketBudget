package models

type Amount struct {
	Cents int64
}

func NewAmount(dollars, cents int64) (Amount, error) {
	// validate stuff
	return Amount{
		Cents: cents,
	}, nil
}

func (a Amount) Add(other Amount) Amount {
	return Amount{
		Cents: a.Cents + other.Cents,
	}
}
