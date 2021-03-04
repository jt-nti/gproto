// SPDX-License-Identifier: Apache-2.0

package matchers_test

import (
	"github.com/jt-nti/gproto"
	"github.com/jt-nti/gproto/matchers"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var _ = Describe("Equal", func() {
	When("passed an unsupported type", func() {
		It("should error", func() {
			expected := &timestamppb.Timestamp{}
			success, err := (&matchers.MessageEqualMatcher{Expected: expected}).Match("string")
			Expect(success).Should(BeFalse())
			Expect(err).Should(HaveOccurred())
			Expect(err).Should(MatchError("matcher expects a proto message.  Got:\n    <string>: string"))
		})
	})

	When("asserting equality between nil values", func() {
		It("should error if actual and expected values are both nil", func() {
			success, err := (&matchers.MessageEqualMatcher{Expected: nil}).Match(nil)

			Expect(success).Should(BeFalse())
			Expect(err).Should(HaveOccurred())
			Expect(err).Should(MatchError("refusing to compare <nil> to <nil>.\nBe explicit and use BeNil() instead"))
		})

		It("should error if only the expected value is nil", func() {
			actual := &timestamppb.Timestamp{}
			success, err := (&matchers.MessageEqualMatcher{Expected: nil}).Match(actual)

			Expect(success).Should(BeFalse())
			Expect(err).Should(HaveOccurred())
			Expect(err).Should(MatchError("refusing to compare message to <nil>.\nBe explicit and use BeNil() instead"))
		})

		It("should succeed if only the actual value is nil", func() {
			expected := &timestamppb.Timestamp{}
			success, err := (&matchers.MessageEqualMatcher{Expected: expected}).Match(nil)

			Expect(success).Should(BeFalse())
			Expect(err).ShouldNot(HaveOccurred())
		})
	})

	When("asserting equality between messages", func() {
		It("should match when the messages are equal", func() {
			actual := &timestamppb.Timestamp{
				Seconds: 1234567890,
			}
			expected := &timestamppb.Timestamp{
				Seconds: 1234567890,
			}
			Expect(actual).To(gproto.Equal(expected))
		})

		It("should not match when the messages are not equal", func() {
			actual := &timestamppb.Timestamp{
				Seconds: 1234567890,
			}
			expected := &timestamppb.Timestamp{
				Seconds: 9876543210,
			}
			Expect(actual).ToNot(gproto.Equal(expected))
		})

		It("should show a failure when the message are not equal", func() {
			actual := &timestamppb.Timestamp{
				Seconds: 1234567890,
			}
			expected := &timestamppb.Timestamp{
				Seconds: 9876543210,
			}

			failures := InterceptGomegaFailures(func() {
				Expect(actual).To(gproto.Equal(expected))
			})
			Expect(failures).To(ConsistOf(And(
				MatchRegexp("Expected\n\\s+<string>: seconds:\\s+1234567890\n\\s+\nto equal\n\\s+<string>: seconds:\\s+9876543210"),
				MatchRegexp("(\\s|\u00a0)+-(\\s|\u00a0)+\"seconds\": int64\\(1234567890\\),"),
				MatchRegexp("(\\s|\u00a0)+\\+(\\s|\u00a0)+\"seconds\": int64\\(9876543210\\),"),
			)))
		})

		It("should show a negated failure when the messages are equal", func() {
			actual := &timestamppb.Timestamp{
				Seconds: 1234567890,
			}
			expected := &timestamppb.Timestamp{
				Seconds: 1234567890,
			}

			failures := InterceptGomegaFailures(func() {
				Expect(actual).ToNot(gproto.Equal(expected))
			})
			Expect(failures).To(ConsistOf(MatchRegexp("Expected\n\\s+<string>: seconds:\\s+1234567890\n\\s+\nnot to equal\n\\s+<string>: seconds:\\s+1234567890")))
		})
	})
})
