package lib

import "errors"

func (gl *lvaVM) GetError() error {
	errstr := gl.GetGlobal("Error").String()
	if errstr == "nil" {
		return nil
	}
	return errors.New(errstr)
}
