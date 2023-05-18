package types

import "fmt"

var (
	ErrBlockNotFound = fmt.Errorf("block not found")
	ErrNodeNotFound  = fmt.Errorf("node not found")
)
