package pendulumcli

import (
	"context"
	"fmt"
)

// Generator define how to generate line
// Generate method return channels that send one line
type Generator interface {
	Generate(ctx context.Context, tickNum int) []<-chan fmt.Stringer
}
