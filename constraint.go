package main

import (
	"fmt"
	"strings"
)

type kind string

const (
	primary kind = "PRIMARY KEY"
	foreign kind = "FOREIGN KEY"
	unique  kind = "UNIQUE"
)

type constraint struct {
	columns            string
	kind               kind
	name               string
	deleteRule         string
	parentTable        string
	parentTableColumns string
}

type constraints []constraint

func newConstraint(kind kind, tokens []string) constraint {
	newConstraint := constraint{
		columns: tokens[3],
		kind:    kind,
		name:    tokens[2],
	}

	if kind == foreign {
		newConstraint.deleteRule = tokens[6]
		newConstraint.parentTable = tokens[4]
		newConstraint.parentTableColumns = tokens[5]
	}

	return newConstraint
}

func (c *constraint) script(tableName string) string {
	sb := strings.Builder{}

	sb.WriteString(fmt.Sprintf("ALTER TABLE %s ADD CONSTRAINT %s %s (%s)", tableName, c.name, c.kind, c.columns))

	if c.kind == foreign {
		sb.WriteString(fmt.Sprintf(" REFERENCES %s (%s)", c.parentTable, c.parentTableColumns))

		if c.deleteRule != "" {
			sb.WriteString(fmt.Sprintf(" ON DELETE %s", c.deleteRule))
		}
	}

	sb.WriteString("\n/\n\n")

	return sb.String()
}

func (c *constraints) newForeignConstraint(tokens []string) {
	*c = append(*c, newConstraint(foreign, tokens))
}

func (c *constraints) newPrimaryConstraint(tokens []string) {
	*c = append(*c, newConstraint(primary, tokens))
}

func (c *constraints) newUniqueConstraint(tokens []string) {
	*c = append(*c, newConstraint(unique, tokens))
}

func (c *constraints) script(tableName string) string {
	sb := strings.Builder{}

	for _, c := range *c {
		sb.WriteString(c.script(tableName))
	}

	return sb.String()
}
