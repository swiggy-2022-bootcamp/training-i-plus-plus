package main

import "fmt"

type PC struct {
	processor     string
	ram, ssd, hdd int
}

type PCBuilder struct {
	processor     string
	ram, ssd, hdd int
}

func (P *PCBuilder) WithProcessor(p string) *PCBuilder {
	P.processor = p
	return P
}

func (P *PCBuilder) WithRam(r int) *PCBuilder {
	P.ram = r
	return P
}

func (P *PCBuilder) WithSSD(s int) *PCBuilder {
	P.ssd = s
	return P
}

func (P *PCBuilder) WithHDD(h int) *PCBuilder {
	P.hdd = h
	return P
}

func (P *PCBuilder) Build() PC {
	return PC{
		P.processor,
		P.ram,
		P.ssd,
		P.hdd,
	}
}

func NewBuilder() *PCBuilder {
	return &PCBuilder{}
}

func main() {

	pc1 := NewBuilder().
		WithProcessor("i5").
		WithRam(8).
		WithSSD(256).
		Build()

	pc2 := NewBuilder().
		WithProcessor("i5").
		WithRam(8).
		WithHDD(1024).
		Build()

	pc3 := NewBuilder().
		WithProcessor("i5").
		WithRam(8).
		WithSSD(256).
		WithHDD(1024).
		Build()

	fmt.Printf("%+v\n", pc1)
	fmt.Printf("%+v\n", pc2)
	fmt.Printf("%+v\n", pc3)
}
