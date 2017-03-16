package twig_test

import (
	"testing"

	"github.com/mattrbianchi/twig"
)

func TestCanLog(t *testing.T) {
	twig.Info("stuff")
}
