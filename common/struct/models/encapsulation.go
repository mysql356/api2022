package models

func (s *Student) Set(fn string, ln string) {
	s.firstName = fn
	s.lastName = ln
}

func (s *Student) Get() (name string) {
	name = s.firstName + " " + s.lastName
	return
}
