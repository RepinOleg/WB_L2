package pattern

/*
Стратегия — это поведенческий паттерн проектирования,
который определяет семейство схожих алгоритмов и помещает каждый из них в собственный класс,
после чего алгоритмы можно взаимозаменять прямо во время исполнения программы.

Паттерн Стратегия предлагает определить семейство схожих алгоритмов,
которые часто изменяются или расширяются, и вынести их в собственные классы, называемые стратегиями.

Вместо того, чтобы изначальный класс сам выполнял тот или иной алгоритм, он будет играть роль контекста,
ссылаясь на одну из стратегий и делегируя ей выполнение работы.
Чтобы сменить алгоритм, будет достаточно подставить в контекст другой объект-стратегию.

*/

type StrategySort interface {
	Sort([]int)
}

type BubbleSort struct {
}

func (s *BubbleSort) Sort(a []int) {
	size := len(a)
	if size < 2 {
		return
	}
	for i := 0; i < size; i++ {
		for j := size - 1; j >= i+1; j-- {
			if a[j] < a[j-1] {
				a[j], a[j-1] = a[j-1], a[j]
			}
		}
	}
}

type InsertionSort struct {
}

func (s *InsertionSort) Sort(a []int) {
	size := len(a)
	if size < 2 {
		return
	}
	for i := 1; i < size; i++ {
		var j int
		var buff = a[i]
		for j = i - 1; j >= 0; j-- {
			if a[j] < buff {
				break
			}
			a[j+1] = a[j]
		}
		a[j+1] = buff
	}
}

type Context struct {
	strategy StrategySort
}

func (c *Context) Algorithm(a StrategySort) {
	c.strategy = a
}

func (c *Context) Sort(s []int) {
	c.strategy.Sort(s)
}

//func main() {
//	data1 := []int{8, 2, 6, 7, 1, 3, 9, 5, 4}
//	data2 := []int{8, 2, 6, 7, 1, 3, 9, 5, 4}
//
//	ctx := new(Context)
//
//	ctx.Algorithm(&BubbleSort{})
//
//	ctx.Sort(data1)
//
//	ctx.Algorithm(&InsertionSort{})
//
//	ctx.Sort(data2)
//}
