package orm


type op string

const (
	opEQ = "="
	opLT = "<"
	opGT = ">"
)

type Column struct {
	name string
}

func C(name string) Column {
	return Column{name: name}
}


func (c Column) Eq(val any) Predicate {
	return Predicate{
		column: c,
		op: opEQ,
		arg: val,
	}
}


// 谓词
type Predicate struct {
	column Column
	op     op
	arg    any
}



