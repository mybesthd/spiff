package dynaml

import (
	"strconv"

	"github.com/cloudfoundry-incubator/spiff/yaml"
)

type IntegerExpr struct {
	Value int
}

func (e IntegerExpr) Evaluate(Binding) (yaml.Node, bool) {
	return node(e.Value), true
}

func (e IntegerExpr) String() string {
	return strconv.Itoa(e.Value)
}
