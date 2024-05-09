package pattern

import (
	"strings"
)

/*
Паттерн Facade относится к структурным паттернам уровня объекта.

Паттерн Facade предоставляет высокоуровневый унифицированный интерфейс в виде набора имен методов к набору взаимосвязанных классов
или объектов некоторой подсистемы, что облегчает ее использование.

Разбиение сложной системы на подсистемы позволяет упростить процесс разработки,
а также помогает максимально снизить зависимости одной подсистемы от другой.
Однако использовать такие подсистемы становиться довольно сложно.
Один из способов решения этой проблемы является паттерн Facade.
Наша задача, сделать простой, единый интерфейс, через который можно было бы взаимодействовать с подсистемами.

В качестве примера можно привести интерфейс автомобиля.
Современные автомобили имеют унифицированный интерфейс для водителя, под которым скрывается сложная подсистема.
Благодаря применению навороченной электроники, делающей большую часть работы за водителя, тот может с лёгкостью управлять автомобилем,
не задумываясь, как там все работает.
*/

type Car struct {
	engine      *Engine
	suspension  *Suspension
	brakeSystem *BrakeSystem
}

func NewCar() *Car {
	return &Car{
		engine:      &Engine{},
		suspension:  &Suspension{},
		brakeSystem: &BrakeSystem{},
	}
}

func (c *Car) Drive() string {
	result := []string{
		c.engine.StartEngine(),
		c.suspension.CheckSuspension(),
		c.brakeSystem.CheckBrakeSystem(),
	}
	return strings.Join(result, "\n")
}

type Engine struct {
}

func (e *Engine) StartEngine() string {
	return "The engine started"
}

type Suspension struct {
}

func (s *Suspension) CheckSuspension() string {
	return "suspension is ok"
}

type BrakeSystem struct {
}

func (b *BrakeSystem) CheckBrakeSystem() string {
	return "brake system is ok"
}

//func main() {
//	myCar := NewCar()
//	result := myCar.Drive()
//	fmt.Println(result)
//}
