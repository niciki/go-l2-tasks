package main

import "fmt"

/*
Класс ConcreteComponent — класс, в который с помощью шаблона Декоратор добавляется новая
функциональность. Абстрактный класс Component определяет интерфейс
для использования всех этих классов

Требуется для реализации:

1. Базовый интерфейс Component который предоставляет интерфейс для класса декоратора и компонента;
2. Класс Decorator1, реализующий интерфейс Component и перезагружающий все методы Component
2. Класс Decorator2, реализующий интерфейс Component и перезагружающий все методы Component,
также к ним добавляется функционал;
3. Класс ConcreteComponent реализующий интерфейс Component и который будет обернут декоратором.
*/
type Component interface {
	Operation() string
}

type ConcretComponent struct {
}

func (c ConcretComponent) Operation() string {
	return "ConcretComponent"
}

type Decorator1 struct {
	component Component
}

func (d Decorator1) Operation() string {
	return "Decorator1 " + d.component.Operation()
}

type Decorator2 struct {
	component Component
}

func (d *Decorator2) Operation() string {
	return "Decorator2 " + d.component.Operation()
}

func (d *Decorator2) AddedBehavior() string {
	return "Added Behavior: " + d.Operation()
}

func main() {
	d1 := Decorator1{component: ConcretComponent{}}
	fmt.Println(d1.Operation())
	d2 := Decorator2{component: Decorator1{component: ConcretComponent{}}}
	fmt.Println(d2.AddedBehavior())
}
