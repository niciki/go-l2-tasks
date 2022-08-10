package main

import (
	"fmt"
	"log"
)

/*
Порождающий шаблон проектирования, предоставляющий подклассам (дочерним классам) интерфейс
для создания экземпляров некоторого класса. В момент создания наследники могут определить,
какой класс создавать. Иными словами, данный шаблон делегирует создание объектов наследникам
родительского класса. Это позволяет использовать в коде программы не конкретные классы,
а манипулировать абстрактными объектами на более высоком уровне.

Структура:
- Product — продукт
	определяет интерфейс объектов, создаваемых абстрактным методом;
- ConcreteProduct — конкретный продукт
	реализует интерфейс Product;
- Creator — создатель
	объявляет фабричный метод, который возвращает объект типа Product. Может также содержать реализацию этого метода «по умолчанию»;
	может вызывать фабричный метод для создания объекта типа Product;
- ConcreteCreator — конкретный создатель
	переопределяет фабричный метод таким образом, чтобы он создавал и возвращал объект класса ConcreteProduct.

*/

type Product interface {
	Name() string
}

// concrete product 1
type Pineapple struct {
}

func (a *Pineapple) Name() string {
	return "pineapple"
}

// concrete product 2
type Apple struct {
}

func (a *Apple) Name() string {
	return "apple"
}

// concrete product 3
type Cherry struct {
}

func (c *Cherry) Name() string {
	return "cherry"
}

type Creator interface {
	CreateProduct(str string) Product
}

type ConcreteCreator struct {
}

func NewCreator() Creator {
	return &ConcreteCreator{}
}

func (c *ConcreteCreator) CreateProduct(str string) Product {
	var product Product
	switch str {
	case "apple":
		product = &Apple{}
	case "pineapple":
		product = &Pineapple{}
	case "cherry":
		product = &Cherry{}
	default:
		log.Fatal("unknown product")
	}
	return product
}

func main() {
	c := NewCreator()
	fmt.Println(c.CreateProduct("apple").Name())
	fmt.Println(c.CreateProduct("cherry").Name())
	fmt.Println(c.CreateProduct("pineapple").Name())
	fmt.Println(c.CreateProduct("strawberry").Name())
}
