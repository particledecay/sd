package build

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("PrintVersion", func() {
	It("Should print nothing if a version is not set", func() {
		Version = ""
		ver := PrintVersion(false)

		Expect(ver).To(BeEmpty())
	})

	It("Should print a short version if set and not verbose", func() {
		Version = "1.2.3"
		ver := PrintVersion(false)

		Expect(ver).To(Equal("v1.2.3"))
	})

	It("Should print a long version if set and verbose", func() {
		Version = "1.2.3"
		Commit = "abcdef1234"
		Date = "20200101"
		ver := PrintVersion(true)

		Expect(ver).To(ContainSubstring("Version:"))
		Expect(ver).To(ContainSubstring("v1.2.3"))
		Expect(ver).To(ContainSubstring("SHA:"))
		Expect(ver).To(ContainSubstring("abcdef1234"))
		Expect(ver).To(ContainSubstring("Built On:"))
		Expect(ver).To(ContainSubstring("20200101"))
	})
})
