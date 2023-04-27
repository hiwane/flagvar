package flagvar

import (
	"fmt"
)

type ChoiceVar struct {
	val      string
	cand     []string
	assigned bool
}

func NewChoiceVar(defval string, cand []string) *ChoiceVar {
	v := new(ChoiceVar)
	v.val = defval
	v.cand = cand
	return v
}

func (cv *ChoiceVar) isValid(s string) bool {
	for _, c := range cv.cand {
		if c == s {
			return true
		}
	}
	return false
}

func (cv *ChoiceVar) Set(opt string) error {
	if !cv.isValid(opt) {
		return fmt.Errorf("invalid choice: choose from %v", cv.cand)
	}
	cv.assigned = true
	cv.val = opt
	return nil
}

func (cv *ChoiceVar) String() string {
	return cv.val
}

func (cv *ChoiceVar) Value() (string, bool) {
	return cv.val, cv.assigned
}

func (cv *ChoiceVar) IsAssigned() bool {
	return cv.assigned
}
