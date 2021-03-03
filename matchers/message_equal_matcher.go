// SPDX-License-Identifier: Apache-2.0

package matchers

import (
	"fmt"

	"github.com/google/go-cmp/cmp"
	"github.com/onsi/gomega/format"
	"google.golang.org/protobuf/encoding/prototext"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/testing/protocmp"
)

type MessageEqualMatcher struct {
	Expected proto.Message
}

func (matcher *MessageEqualMatcher) Match(actual interface{}) (bool, error) {
	if actual == nil && matcher.Expected == nil {
		return false, fmt.Errorf("refusing to compare <nil> to <nil>.\nBe explicit and use BeNil() instead")
	} else if matcher.Expected == nil {
		return false, fmt.Errorf("refusing to compare message to <nil>.\nBe explicit and use BeNil() instead")
	} else if actual == nil {
		return false, nil
	}

	message, ok := actual.(proto.Message)
	if !ok {
		return false, fmt.Errorf("matcher expects a proto message.  Got:\n%s", format.Object(actual, 1))
	}

	return proto.Equal(message, matcher.Expected), nil
}

func (matcher *MessageEqualMatcher) FailureMessage(actual interface{}) string {
	actualMessage, _ := actual.(proto.Message)

	return format.Message(prototext.Format(actualMessage), "to equal", prototext.Format(matcher.Expected)) +
		matcher.mismatchMessage(actualMessage)
}

func (matcher *MessageEqualMatcher) NegatedFailureMessage(actual interface{}) string {
	actualMessage, _ := actual.(proto.Message)

	return format.Message(prototext.Format(actualMessage), "not to equal", prototext.Format(matcher.Expected))
}

func (matcher *MessageEqualMatcher) mismatchMessage(actual proto.Message) string {
	diff := cmp.Diff(actual, matcher.Expected, protocmp.Transform())

	return "\nMismatch (-want +got)\n" + format.IndentString(diff, 1)
}
