package communitymodules

import (
	"os"
	"sort"
	"strings"

	"github.com/olekukonko/tablewriter"
)

type RowConverter func(row) []string
type TableInfo struct {
	Header       []string
	RowConverter RowConverter
}

var (
	CollectiveTableInfo = TableInfo{
		Header:       []string{"NAME", "REPOSITORY", "VERSION INSTALLED", "MANAGED"},
		RowConverter: func(r row) []string { return []string{r.Name, r.Repository, r.Version, r.Managed} },
	}
	InstalledTableInfo = TableInfo{
		Header:       []string{"NAME", "VERSION"},
		RowConverter: func(r row) []string { return []string{r.Name, r.Version} },
	}
	ManagedTableInfo = TableInfo{
		Header:       []string{"NAME"},
		RowConverter: func(r row) []string { return []string{r.Name} },
	}
	CatalogTableInfo = TableInfo{
		Header:       []string{"NAME", "REPOSITORY", "LATEST VERSION"},
		RowConverter: func(r row) []string { return []string{r.Name, r.Repository, r.LatestVersion} },
	}
)

func RenderModules(raw bool, moduleMap moduleMap, tableInfo TableInfo) {
	renderTable(
		raw,
		convertModuleMapToTable(moduleMap, tableInfo.RowConverter),
		tableInfo.Header)
}

func convertModuleMapToTable(moduleMap moduleMap, rowConverter RowConverter) [][]string {
	var moduleNames []string
	for key := range moduleMap {
		moduleNames = append(moduleNames, key)
	}
	sort.Strings(moduleNames)
	var result [][]string
	for _, key := range moduleNames {
		result = append(result, rowConverter(moduleMap[key]))
	}
	return result
}

// renderTable renders the table with the provided headers
func renderTable(raw bool, modulesData [][]string, headers []string) {
	if raw {
		for _, row := range modulesData {
			println(strings.Join(row, "\t"))
		}
	} else {
		var table [][]string
		table = append(table, modulesData...)

		twTable := setTable(table)
		twTable.SetHeader(headers)
		twTable.Render()
	}
}

// setTable sets the table settings for the tablewriter
func setTable(inTable [][]string) *tablewriter.Table {
	table := tablewriter.NewWriter(os.Stdout)
	table.AppendBulk(inTable)
	table.SetRowLine(true)
	table.SetAlignment(tablewriter.ALIGN_CENTER)
	table.SetColumnAlignment([]int{tablewriter.ALIGN_CENTER, tablewriter.ALIGN_LEFT, tablewriter.ALIGN_LEFT, tablewriter.ALIGN_LEFT})
	table.SetBorder(false)
	return table
}
