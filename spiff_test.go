package main

import (
	"io/ioutil"
	"os"
	"os/exec"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/vito/cmdtest"
	. "github.com/vito/cmdtest/matchers"
)

var _ = Describe("Running spiff", func() {
	spiff, err := cmdtest.Build("github.com/cloudfoundry-incubator/spiff")
	if err != nil {
		panic(err)
	}

	Describe("merge", func() {
		var merge *cmdtest.Session

		Context("when given a bad file path", func() {
			BeforeEach(func() {
				merge, err = cmdtest.Start(exec.Command(spiff, "merge", "foo.yml"))
				Expect(err).NotTo(HaveOccurred())
			})

			It("says file not found", func() {
				Expect(merge).To(SayError("foo.yml: no such file or directory"))
				Expect(merge).To(ExitWith(1))
			})
		})

		Context("when given a single file", func() {
			basicTemplate, err := ioutil.TempFile(os.TempDir(), "basic.yml")
			if err != nil {
				panic(err)
			}

			basicTemplate.Write([]byte(`
---
foo: bar
`))

			BeforeEach(func() {
				merge, err = cmdtest.Start(exec.Command(spiff, "merge", basicTemplate.Name()))
				Expect(err).NotTo(HaveOccurred())
			})

			It("resolves the template and prints it out", func() {
				Expect(merge).To(Say(`foo: bar`))
				Expect(merge).To(ExitWith(0))
			})
		})
	})
})
