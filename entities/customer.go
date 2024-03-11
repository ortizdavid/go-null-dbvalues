package entities


type Customer struct {
	Id		int
	Name 	string
	Age		*int
	Gender	*string
	Height	*float32
}
