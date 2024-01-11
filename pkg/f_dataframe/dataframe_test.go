package dataframe

import (
	"fmt"
	"testing"
	datasource "tiny_planner/pkg/b_datasource"
	exprLogi "tiny_planner/pkg/c_expr_logical"
)

func TestLogicalPlan_DataFrame(t *testing.T) {
	ctx := NewContext()
	df := ctx.ReadCsv("/Users/arjunsunilkumar/GolandProjects/0learning/tiny_planner/test/data/aggregate_test_100.csv",
		datasource.CsvReadOptions{HasHeader: true})

	df = df.
		//Filter(exprLogi.Eq(exprLogi.Column{Name: "state"}, exprLogi.LiteralString{Val: "CO"})).
		Project([]exprLogi.LogicalExpr{
			exprLogi.Column{Name: "c1"},
			exprLogi.Column{Name: "c2"},
		})

	logicalPlan := df.LogicalPlan()
	fmt.Println(exprLogi.PrettyPrint(logicalPlan, 0))

	df.Show()
}