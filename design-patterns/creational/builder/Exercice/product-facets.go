package main

import (
	"fmt"
)

type Producer struct {
	name string
	city string
}

type Attribute struct {
	name, category, description string
	price                       float64
}

type Product struct {
	attribute *Attribute
	producer  *Producer
}

type ProductBuilder struct {
	product *Product
}

type ProductProducerBuilder struct {
	ProductBuilder
}

type ProductAttributeBuilder struct {
	ProductBuilder
}

func NewProductBuilder() *ProductBuilder {
	return &ProductBuilder{&Product{attribute: &Attribute{}, producer: &Producer{}}}
}

func (it *ProductBuilder) Attribute() *ProductAttributeBuilder {
	return &ProductAttributeBuilder{*it}
}

func (p *ProductAttributeBuilder) WithName(name string) *ProductAttributeBuilder {
	p.product.attribute.name = name
	return p
}

func (p *ProductAttributeBuilder) WithPrice(price float64) *ProductAttributeBuilder {
	p.product.attribute.price = price
	return p
}

func (p *ProductAttributeBuilder) WithCategory(category string) *ProductAttributeBuilder {
	p.product.attribute.category = category
	return p
}

func (p *ProductAttributeBuilder) WithDescription(description string) *ProductAttributeBuilder {
	p.product.attribute.description = description
	return p
}

func (it *ProductBuilder) Producer() *ProductProducerBuilder {
	return &ProductProducerBuilder{*it}
}

func (ppb *ProductProducerBuilder) WithName(name string) *ProductProducerBuilder {
	ppb.product.producer.name = name
	return ppb
}

func (ppb *ProductProducerBuilder) From(city string) *ProductProducerBuilder {
	ppb.product.producer.city = city
	return ppb
}

func (p *ProductBuilder) Build() *Product {
	return p.product
}

func MainFacets() {
	productBuilder := NewProductBuilder()

	product := productBuilder.
		Attribute().
		WithName("apple").
		WithCategory("fruit").
		WithDescription("a delicious fruit").
		WithPrice(1).
		Producer().
		WithName("denis").
		From("Chicago").
		Build()

	fmt.Printf(" %s - %s - %s - %f - %s - %s \n", product.attribute.name, product.attribute.description, product.attribute.category, product.attribute.price, product.producer.city, product.producer.name)
}
