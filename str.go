package flagvar

type StrVar struct {
	val      string
	assigned bool
}

func NewStrVar(defval string) *StrVar {
	v := new(StrVar)
	v.val = defval
	return v
}

func (dv *StrVar) Set(opt string) error {
	dv.assigned = true
	dv.val = opt
	return nil
}

func (dv *StrVar) String() string {
	return dv.val
}

func (dv *StrVar) Value() (string, bool) {
	return dv.val, dv.assigned
}

func (dv *StrVar) IsAssigned() bool {
	return dv.assigned
}
