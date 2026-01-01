package print

import (
	"os"

	"github.com/olekukonko/tablewriter"
	"github.com/olekukonko/tablewriter/renderer"
)

// Table is Struct for printable table
type Table struct {
	// Table Header as []string{"id","name","value" .....}
	Header []string
	// Table Body is content of table.
	Body [][]string
}

// Print function prints the table itself.
func (t *Table) Print(Caption string) {
	table := tablewriter.NewTable(os.Stdout,
		tablewriter.WithHeader(t.Header),
		tablewriter.WithRenderer(renderer.NewMarkdown()),
	)
	table.Bulk(t.Body)
	table.Render()
}
