package main


type C1 struct {
	C SList
}

type T1 interface {
	Run()
	Start()
	Stop()
}

type S1 struct {}

type SList []*T1

func (s1 *S1) Run() {

}

func (s1 *S1) Start() {

}

func (s1 *S1) Stop() {

}

func main() {
	s1 := &S1{}

	c1 := &C1{}

	Slist := &SList{}

	*Slist = append(*Slist, s1)
}
