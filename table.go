package main

import (
	"fmt"
	"strings"
)

type table struct {
	columns     columns
	comment     string
	constraints constraints
	indexes     indexes
	name        string
}

type tables []table

func (t *table) createConstraintsScript() string {
	return t.constraints.script(t.name)
}

func (t *table) createIndexScript() string {
	return t.indexes.script(t.name)
}

func (t *table) commentsOnColumnsScript() string {
	return t.columns.commentsScript(t.name)
}

func (t *table) commentOnTableScript() string {
	return fmt.Sprintf("COMMENT ON TABLE %s IS '%s'\n/\n\n", t.name, t.comment)
}

func (t *table) createTableScript() string {
	return fmt.Sprintf("CREATE TABLE %s (%s)\n/\n\n", t.name, t.columns.script())
}

func (t *table) script() string {
	sb := strings.Builder{}

	sb.WriteString(t.createTableScript())
	sb.WriteString(t.commentOnTableScript())
	sb.WriteString(t.commentsOnColumnsScript())
	sb.WriteString(t.createConstraintsScript())
	sb.WriteString(t.createIndexScript())

	return sb.String()
}

func (ts *tables) newTable(tokens []string) {
	*ts = append(*ts, table{
		columns:     make(columns, 0, 10),
		comment:     tokens[1],
		constraints: make(constraints, 0, 5),
		name:        tokens[0],
	})
}

func (ts *tables) script() string {
	sb := strings.Builder{}

	for _, t := range *ts {
		sb.WriteString(t.script())
	}

	return sb.String()
}
