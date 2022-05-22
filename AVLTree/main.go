package main

func main() {

	b := &AVLTree{}

	b.Add(5)
	b.Add(-5)
	b.Add(6)
	b.Add(55)
	b.Add(56)
	b.Add(57)
	b.Add(58)
	b.Add(59)
	b.Add(60)
	b.Add(61)
	b.Add(-1)
	b.Add(-10)
	b.Remove(55)
	Print2D(b)
	b.Remove(-5)
	b.Remove(60)

	//Print2D(b)
}
