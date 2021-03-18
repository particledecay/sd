package main

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestSd(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Sd Suite")
}
