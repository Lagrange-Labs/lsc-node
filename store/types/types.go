package types

import "fmt"

var (
	ErrBlockNotFound = fmt.Errorf("block not found")
	ErrBatchNotFound = fmt.Errorf("batch not found")
	ErrNodeNotFound  = fmt.Errorf("node not found")
)
