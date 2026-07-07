package sqlutil

import (
	"fmt"
	"regexp"

	"github.com/Masterminds/squirrel"
)

var identifierPattern = regexp.MustCompile(`^[A-Za-z_][A-Za-z0-9_]*$`)

func StringListEqual(column string, values []string) (squirrel.Sqlizer, error) {
	if !identifierPattern.MatchString(column) {
		return nil, fmt.Errorf("invalid SQL identifier %q", column)
	}
	if len(values) == 0 {
		return squirrel.Expr("1 = 0"), nil
	}
	copied := append([]string(nil), values...)
	return squirrel.Eq{column: copied}, nil
}
