package lang

type Record struct {
	fields map[interface{}]interface{}
}

func NewRecord() *Record {
	r := new(Record)
	r.fields = make(map[interface{}]interface{})
	return r
}

func (r *Record) GetField(n string) (interface{}, bool) {
	val, ok := r.fields[n]
	return val, ok
}

func (r *Record) SetField(n string, val interface{}) {
	r.fields[n] = val
}
