package main

/*

Exercise Prompt: You have been tasked with implementing a complex object creation process using the Builder design pattern.
The object you need to create is a Product with multiple attributes such as name, price, category, and description.

Your goal is to implement three different variations of the Builder pattern: Builder Facets, Functional Builder, and Builder Parameter.

Builder Facets:

Implement a ProductBuilder struct that allows you to set different attributes of the Product using separate methods for each attribute.
The ProductBuilder should have methods like WithName(name string), WithPrice(price float64), WithCategory(category string), and WithDescription(description string) to set the corresponding attributes.
Finally, implement a Build method that returns the fully constructed Product object.
Functional Builder:

Implement a ProductBuilder function that returns a closure.
The closure should have methods like WithName(name string), WithPrice(price float64), WithCategory(category string), and WithDescription(description string) to set the corresponding attributes.
The closure should also have a Build method that returns the fully constructed Product object.
Builder Parameter:

Implement a ProductBuilder struct that takes all the attributes as parameters in its constructor.
The ProductBuilder should have a Build method that returns the fully constructed Product object.
Your task is to implement the three variations of the Builder pattern as described above and demonstrate their usage by creating a Product object with different attributes using each technique.

Note: You can create a simple Product struct with the required attributes and a String method to print its details.

*/

func main() {
	MainFacets()
	MainFunc()
	MainParam()
}
