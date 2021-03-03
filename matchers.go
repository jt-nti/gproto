// SPDX-License-Identifier: Apache-2.0

package gproto

import (
	"github.com/jt-nti/gproto/matchers"
	"github.com/onsi/gomega/types"
	"google.golang.org/protobuf/proto"
)

//Equal succeeds if actual is a Protobuf message that matches
//the expected message.  The messages are compared via proto.Equal
func Equal(expected proto.Message) types.GomegaMatcher {
	return &matchers.MessageEqualMatcher{
		Expected: expected,
	}
}

// TODO add matches, e.g.
// func EqualCmp(expected interface{}, options ...cmp.Option) types.GomegaMatcher {
// 	return &matchers.MessageEqualCmpMatcher{
// 		Expected: expected,
// 		Options:  options,
// 	}
// }
