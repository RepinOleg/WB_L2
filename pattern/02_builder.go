package pattern

/*
Паттерн Builder относится к порождающим паттернам уровня объекта.

Паттерн Builder определяет процесс поэтапного построения сложного продукта.
После того как будет построена последняя его часть, продукт можно использовать.
*/

type Builder interface {
	MakeHeader(str string)
	MakeBody(str string)
	MakeFooter(str string)
}

type Director struct {
	builder Builder
}

func (d *Director) Construct() {
	d.builder.MakeHeader("Header")
	d.builder.MakeBody("Body")
	d.builder.MakeFooter("Footer")
}

type ConcreteBuilder struct {
	product *Product
}

type Product struct {
	content string
}

func (p *Product) Show() string {
	return p.content
}

func (b *ConcreteBuilder) MakeHeader(str string) {
	b.product.content += "<header>" + str + "</header>"
}

func (b *ConcreteBuilder) MakeBody(str string) {
	b.product.content += "<article>" + str + "</article>"
}

func (b *ConcreteBuilder) MakeFooter(str string) {
	b.product.content += "<footer>" + str + "</footer>"
}

//func main() {
//	product := new(Product)
//
//	director := Director{&ConcreteBuilder{product: product}}
//	director.Construct()
//
//	result := product.Show()
//	fmt.Println(result)
//}
