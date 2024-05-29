package pkg

import "strconv"

type Str string

func (s Str) String() string {
	return string(s)
}

func (s Str) Int() (int, error) {
	return strconv.Atoi(string(s))
}

func (s Str) OnlyInt() int {
	res, _ := s.Int()
	return res
}

func (s Str) Uint32() (uint32, error) {
	res, err := s.Int()
	return uint32(res), err
}

func (s Str) OnlyUint32() uint32 {
	res, _ := s.Uint32()
	return res
}
