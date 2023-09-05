package database

import (
	"fmt"
	"testing"

	"github.com/ortelius/scec-commons/model"
)

func TestEmptyJson(_ *testing.T) {

	fmt.Printf("let compverdetail=JSON.parse('%s');\n", EmptyJSON(model.NewComponentVersionDetails()))
	fmt.Printf("let applications=JSON.parse('%s');\n", EmptyJSON(model.NewApplications()))
	fmt.Printf("let appver=JSON.parse('%s');\n", EmptyJSON(model.NewApplicationVersion()))
	fmt.Printf("let appverdetail=JSON.parse('%s');\n", EmptyJSON(model.NewApplicationVersionDetails()))
	fmt.Printf("let auditlog=JSON.parse('%s');\n", EmptyJSON(model.NewAuditLog()))
	fmt.Printf("let auditrec=JSON.parse('%s');\n", EmptyJSON(model.NewAuditRecord()))
	fmt.Printf("let compattrs=JSON.parse('%s');\n", EmptyJSON(model.NewCompAttrs()))
}
