package dataflow

import (
	"testing"

	"github.com/stackql/any-sdk/pkg/dto"
	"github.com/stackql/stackql/internal/stackql/taxonomy"
)

func TestUpsertStandardDataFlowVertex(t *testing.T) {
	testdc := NewStandardDataFlowCollection(dto.NewDataFlowCfg(10, 10))
	annotation := taxonomy.NewAnnotationCtxSplitMap("test", "test", "test", "test", "test")
}
