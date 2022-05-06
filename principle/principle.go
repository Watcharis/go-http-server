package principle

type Police struct {
	Student
}

type Student struct{}

func (p Police) TestOne() int {
	return 1
}

func (s Student) TestOne() int {
	return 1
}
