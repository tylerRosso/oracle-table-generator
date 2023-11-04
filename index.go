package main

import (
	"fmt"
	"strings"
)

type index struct {
	columns string
	name    string
}

type indexes []index

func (i *index) script(tableName string) string {
	return fmt.Sprintf("CREATE INDEX %s ON %s (%s)\n/\n\n", i.name, tableName, i.columns)
}

func (i *indexes) script(tableName string) string {
	sb := strings.Builder{}

	for _, i := range *i {
		sb.WriteString(i.script(tableName))
	}

	return sb.String()
}

func (i *indexes) newIndex(tokens []string) {
	*i = append(*i, index{
		columns: tokens[3],
		name:    tokens[2],
	})
}
