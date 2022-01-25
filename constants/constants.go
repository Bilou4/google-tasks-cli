package constants

import (
	"github.com/jedib0t/go-pretty/v6/table"
	"google.golang.org/api/tasks/v1"
)

var (
	Srv   *tasks.Service
	Table table.Writer
)
