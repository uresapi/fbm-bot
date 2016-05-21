package main_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestFbmBot(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "FbmBot Suite")
}
