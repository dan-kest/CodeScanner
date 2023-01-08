package null

import (
	"database/sql"
	"database/sql/driver"
)

// Extention of sql.NullString for more generic usage.
type String struct {
	Val   string
	valid bool
}

// Create new String instance with value.
func NewString(str string) String {
	return String{
		Val:   str,
		valid: true,
	}
}

// Set value, also set to not null.
func (i *String) Set(str string) {
	i.Val = str
	i.valid = true
}

// Return true if null.
func (i *String) IsNull() bool {
	return !i.valid
}

// Return true if not null.
func (i *String) IsNotNull() bool {
	return i.valid
}

// String implements the Stringer interface and returns the internal string
// so you can use a String in a fmt.Println statement for example
func (i *String) String() string {
	if !i.valid {
		return ""
	}

	return i.Val
}

// Scan implements the Scanner interface for SQL read.
func (i *String) Scan(value interface{}) error {
	if value == nil {
		i.Val = ""
		i.valid = false
		return nil
	}

	var ns sql.NullString

	if err := ns.Scan(value); err != nil {
		i.Val = ""
		i.valid = false
		return err
	}

	i.Val = ns.String
	i.valid = ns.Valid
	return nil
}

// Value implements the Valuer interface for SQL write.
func (i String) Value() (driver.Value, error) {
	if !i.valid {
		return nil, nil
	}

	return i.Val, nil
}
