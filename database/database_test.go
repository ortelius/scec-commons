package database

import (
	"fmt"
	"testing"

	"github.com/ortelius/scec-commons/model"
)

func TestEmptyJson(t *testing.T) {

	fmt.Printf("let compverdetail=JSON.parse('%s');\n", EmptyJson(model.NewComponentVersionDetails()))
	fmt.Printf("let applications=JSON.parse('%s');\n", EmptyJson(model.NewApplications()))
	fmt.Printf("let appver=JSON.parse('%s');\n", EmptyJson(model.NewApplicationVersion()))
	fmt.Printf("let appverdetail=JSON.parse('%s');\n", EmptyJson(model.NewApplicationVersionDetails()))
	fmt.Printf("let auditlog=JSON.parse('%s');\n", EmptyJson(model.NewAuditLog()))
	fmt.Printf("let auditrec=JSON.parse('%s');\n", EmptyJson(model.NewAuditRecord()))
	fmt.Printf("let compattrs=JSON.parse('%s');\n", EmptyJson(model.NewCompAttrs()))
}
