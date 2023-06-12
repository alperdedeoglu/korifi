package payloads_test

import (
	"testing"

	apierrors "code.cloudfoundry.org/korifi/api/errors"
	"code.cloudfoundry.org/korifi/api/handlers"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var validator *handlers.DecoderValidator

var _ = BeforeEach(func() {
	var err error
	validator, err = handlers.NewDefaultDecoderValidator()
	Expect(err).NotTo(HaveOccurred())
})

func TestPayloads(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Payloads Suite")
}

func expectUnprocessableEntityError(err error, detail string) {
	ExpectWithOffset(1, err).To(HaveOccurred())
	ExpectWithOffset(1, err).To(BeAssignableToTypeOf(apierrors.UnprocessableEntityError{}))
	ExpectWithOffset(1, err.(apierrors.UnprocessableEntityError).Detail()).To(ContainSubstring(detail))
}
