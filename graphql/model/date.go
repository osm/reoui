package model

import (
	"fmt"
	"io"
	"regexp"
)

var re = regexp.MustCompile(`^\d{4}-(02-(0[1-9]|[12][0-9])|(0[469]|11)-(0[1-9]|[12][0-9]|30)|(0[13578]|1[02])-(0[1-9]|[12][0-9]|3[01]))$`)

type Date string

func (d *Date) UnmarshalGQL(v interface{}) error {
	ds, ok := v.(string)
	if !ok {
		return fmt.Errorf("Date must be a string")
	}

	if !re.Match([]byte(ds)) {
		return fmt.Errorf("Date is not valid, expects format of YYYY-MM-DD")
	}

	*d = Date(ds)
	return nil
}

func (d Date) MarshalGQL(w io.Writer) {
	w.Write([]byte(d))
}
