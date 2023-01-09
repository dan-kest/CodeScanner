package tests

import (
	"errors"
	"fmt"
)

var (
	ErrNoResult = errors.New("result expected: got no result")
)

func CompareField(got interface{}, want interface{}, name interface{}) error {
	prefix := ""
	if name != nil {
		prefix = fmt.Sprintf("%v: ", name)
	}

	if got != want {
		return fmt.Errorf("%sgot '%v', want '%v'", prefix, got, want)
	}

	return nil
}

func CompareError(wantErr bool, got error, want error) error {
	if wantErr {
		if got == nil {
			return fmt.Errorf("error expected: got no error")
		}
		if got.Error() != want.Error() {
			gotMsg := got.Error()
			wantMsg := want.Error()

			return fmt.Errorf("error mismatch: got '%s', want '%s'", gotMsg, wantMsg)
		}
	} else if got != nil {
		gotMsg := got.Error()

		return fmt.Errorf("error unexpected: got '%s'", gotMsg)
	}

	return nil
}
