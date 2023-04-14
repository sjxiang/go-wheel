package orm


// 这种叫做标记接口
// expr 这个方法并不具有实际意义和实际用处
type Expression interface {
	expr()
}
