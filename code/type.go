package main

// 别名
type PersonA Person
type Str string

// 定义结构体
type Person struct {
	Name     string
	RealName Str
	Age      int
}

// 需要强制实现接口的所有方法,来检查结构体Pa是否实现所有方法
var _ Personer = (*Pa)(nil)

// 定义接口，只有方法定义，接口的方法需被实现，golang不强制要求
type Personer interface {
	Run()
	//See()
}

type Pa struct {
}

type Pb struct {
}

//结构体实现接口的方案，可以认为实现了此接口
func (a *Pa) Run() {}
func (b *Pb) Run() {}

//返回的是一个接口，可以是结构体 也可以是指针
func NewPersonerPa() Personer {
	//retrun Pa{}
	return &Pa{}
}

//定义函数类型
//以下是定义一个函数类型handler
type handler func(name string) int

//针对这个函数类型可以再定义方法，如：
func (h handler) add(name string) int {
	return h(name) + 10
}

//参见golang option设计模式
type Option func(person *Person)

func WithName(name string) Option {
	return func(m *Person) {
		m.Name = name
	}
}
