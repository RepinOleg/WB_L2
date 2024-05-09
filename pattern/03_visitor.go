package pattern

/*
Паттерн Visitor относится к поведенческим паттернам уровня объекта.

Паттерн Visitor позволяет обойти набор элементов (объектов) с разнородными интерфейсами,
а также позволяет добавить новый метод в класс объекта, при этом, не изменяя сам класс этого объекта.
*/

type Visitor interface {
	VisitSushiBar(s *SushiBar) string
	VisitPizzeria(pi *Pizzeria) string
	VisitBurgerBar(b *BurgerBar) string
}

type Place interface {
	Accept(v Visitor) string
}

type People struct {
}

func (p *People) VisitSushiBar(s *SushiBar) string {
	return s.BuySushi()
}

// VisitPizzeria implements visit to Pizzeria.
func (p *People) VisitPizzeria(pi *Pizzeria) string {
	return pi.BuyPizza()
}

// VisitBurgerBar implements visit to BurgerBar.
func (p *People) VisitBurgerBar(b *BurgerBar) string {
	return b.BuyBurger()
}

type City struct {
	places []Place
}

func (c *City) Add(p Place) {
	c.places = append(c.places, p)
}

func (c *City) Accept(v Visitor) string {
	var result string
	for _, p := range c.places {
		result += p.Accept(v)
	}
	return result
}

type SushiBar struct {
}

func (s *SushiBar) Accept(v Visitor) string {
	return v.VisitSushiBar(s)
}

func (s *SushiBar) BuySushi() string {
	return "Buy sushi..."
}

type Pizzeria struct {
}

func (p *Pizzeria) Accept(v Visitor) string {
	return v.VisitPizzeria(p)
}

func (p *Pizzeria) BuyPizza() string {
	return "Buy pizza..."
}

type BurgerBar struct {
}

func (b *BurgerBar) Accept(v Visitor) string {
	return v.VisitBurgerBar(b)
}

func (b *BurgerBar) BuyBurger() string {
	return "Buy burger..."
}
