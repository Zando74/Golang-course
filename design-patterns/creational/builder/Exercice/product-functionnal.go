package main

import "fmt"

type ProductFunc struct {
	name, category, description string
	price                       float64
}

type ProductFuncMod func(*ProductFunc)

type ProductFuncBuilder struct {
	actions []ProductFuncMod
}

func (pfb *ProductFuncBuilder) WithName(name string) *ProductFuncBuilder {
	pfb.actions = append(pfb.actions, func(pf *ProductFunc) { pf.name = name })
	return pfb
}

func (pfb *ProductFuncBuilder) Build() *ProductFunc {
	productFunc := ProductFunc{}

	for _, action := range pfb.actions {
		action(&productFunc)
	}

	return &productFunc
}

// Add rest later
func (pfb *ProductFuncBuilder) WithCategory(category string) *ProductFuncBuilder {
	pfb.actions = append(pfb.actions, func(pf *ProductFunc) { pf.category = category })
	return pfb
}

func (pfb *ProductFuncBuilder) WithDescription(description string) *ProductFuncBuilder {
	pfb.actions = append(pfb.actions, func(pf *ProductFunc) { pf.description = description })
	return pfb
}

func (pfb *ProductFuncBuilder) WithPrice(price float64) *ProductFuncBuilder {
	pfb.actions = append(pfb.actions, func(pf *ProductFunc) { pf.price = price })
	return pfb
}

func MainFunc() {
	productFuncBuilder := ProductFuncBuilder{}

	productFunc := productFuncBuilder.
		WithName("apple").
		WithCategory("fruit").
		WithDescription("a delicious functionnal fruit").
		WithPrice(2).
		Build()

	fmt.Printf(" %s - %s - %s - %f \n", productFunc.name, productFunc.description, productFunc.category, productFunc.price)
}
