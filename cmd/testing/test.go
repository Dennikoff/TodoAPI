package main

type B struct {
	i int
}

func main() {
	b := &B{
		i: 1,
	}
	sum(&b.i)
	println(b.i)
}

func sum(a interface{}) {
	*a.(*int) = 10
}
