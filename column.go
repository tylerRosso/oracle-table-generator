package main

import (
	"fmt"
	"strings"
)

type column struct {
	comment    string
	definition string
	name       string
}

type columns []column

func (c *column) commentScript(tableName string) string {
	return fmt.Sprintf("COMMENT ON COLUMN %s.%s IS '%s'\n/\n\n", tableName, c.name, c.comment)
}

func (c *column) script() string {
	return fmt.Sprintf("%s %s, ", c.name, c.definition)
}

func (cs *columns) commentsScript(tableName string) string {
	sb := strings.Builder{}

	for _, c := range *cs {
		sb.WriteString(c.commentScript(tableName))
	}

	return sb.String()
}

func (cs *columns) newColumn(tokens []string) {
	*cs = append(*cs, column{
		comment:    tokens[3],
		definition: tokens[2],
		name:       tokens[1],
	})
}

func (cs *columns) script() string {
	sb := strings.Builder{}

	for _, c := range *cs {
		sb.WriteString(c.script())
	}

	return strings.TrimSuffix(sb.String(), ", ")
}
