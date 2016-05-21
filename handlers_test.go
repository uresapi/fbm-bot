package main_test

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/uresapi/fbm-bot"
)

var _ = Describe("Handlers", func() {
	var (
		resp   *httptest.ResponseRecorder
		Url    *url.URL
		err    error
		req    *http.Request
		params url.Values
	)

	BeforeEach(func() {
		resp = httptest.NewRecorder()
		Url, err = url.Parse("")
		Expect(err).NotTo(HaveOccurred())

		Url.Path += "webhook/"
		params = url.Values{}
	})

	It("responds with 200", func() {
		params.Add("hub.verify_token", "ures_official_verification_token")
		params.Add("hub.challenge", "very-challenging")
		Url.RawQuery = params.Encode()

		req, err = http.NewRequest(
			"GET", Url.String(), nil,
		)
		Expect(err).NotTo(HaveOccurred())

		Verify(resp, req)
		Expect(resp.Code).To(Equal(200))

		Expect(ioutil.ReadAll(resp.Body)).To(ContainSubstring(("very-challenging")))
	})

	Context("when the token is invalid", func() {
		It("should respond with 500", func() {
			params.Add("hub.verify_token", "ures_unofficial_verification_token")

			Url.RawQuery = params.Encode()

			req, err = http.NewRequest(
				"GET", Url.String(), nil,
			)
			Expect(err).NotTo(HaveOccurred())

			Verify(resp, req)
			Expect(resp.Code).To(Equal(500))
		})
	})
})
