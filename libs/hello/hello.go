package hello

import (
	"fmt"

	"github.com/google/uuid"
)

func Greet(audience string) string {
	_ = uuid.New()
	return fmt.Sprintf("Hello man, %s!", audience)
}
