package main

import "fmt"

type taxSystem interface {
	caliculateTax() int
}

type indianTax struct {
	taxper int
	income int
}

func (i *indianTax) caliculateTax() int {
	fmt.Println(" indian caliculateTax", i)
	tax := i.income * i.taxper / 100 // 100*10/100 = 10
	return tax
}

type singaTax struct {
	taxper int
	income int
}

func (i *singaTax) caliculateTax() int {
	fmt.Println(" singapure caliculateTax", i)
	tax := i.income * i.taxper / 100 // 100*5/100 = 5
	return tax
}

func Tax() {
	// Alice
	indianTax := &indianTax{
		taxper: 10,
		income: 100,
	}

	singaTax := &singaTax{
		taxper: 5,
		income: 100,
	}

	taxSystems := []taxSystem{indianTax, singaTax}
	total := claculateTotalTax(taxSystems)
	fmt.Println("Total  tax will be", total)
}

func claculateTotalTax(taxSystems []taxSystem) int {
	fmt.Println("claculateTotalTax", taxSystems)
	totalTax := 0
	// ind, sig
	for _, t := range taxSystems {
		totalTax += t.caliculateTax() // totalTax = totalTax+ t.caliculateTax()
	}
	return totalTax
}
