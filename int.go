package flagvar

import (
	"fmt"
	"strconv"
)

type IntVar struct {
	min      int
	max      int
	val      int
	assigned bool
}

func NewIntVar(defval, min, max int) *IntVar {
	v := new(IntVar)
	v.min = min
	v.max = max
	v.val = defval
	return v
}

func (dv *IntVar) Set(opt string) error {
	v, err := strconv.Atoi(opt)
	if err != nil {
		return err
	}
	if v < dv.min || v > dv.max {
		return fmt.Errorf("out of range")
	}

	dv.assigned = true
	dv.val = v
	return nil
}

func (dv *IntVar) String() string {
	return fmt.Sprintf("%d", dv.val)
}

func (dv *IntVar) Value() (int, bool) {
	return dv.val, dv.assigned
}
