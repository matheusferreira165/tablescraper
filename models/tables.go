package models

type TableData struct {
	Headers []string
	Rows    []string
}

type Table struct {
	Table []TableData
}
