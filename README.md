# Gomega Matchers for Protobufs

`gproto` provides [Gomega](https://github.com/onsi/gomega) matchers to help write effective assertions against Protobufs.

Based on the [EqualProto Matcher gomega issue](https://github.com/onsi/gomega/issues/292), a [Hyperledger Fabric review comment](https://github.com/hyperledger/fabric/pull/2395#discussion_r580397470), and a [equal_cmp_matcher.go Gist](https://gist.github.com/jaslong/8852fb0ae5367484957dc9b6c33924d3).

## Installation

`go get github.com/jt-nti/gproto`

## Usage

```go
//...
import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/jt-nti/gproto"
	"google.golang.org/protobuf/types/known/timestamppb"
)
//...
Context("when something something protobuf", func() {
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
})
//...
```

## TODO

- Add more tests for more complex messages
- Add more matchers based on the following packages:
    - https://pkg.go.dev/github.com/google/go-cmp/cmp
    - https://pkg.go.dev/google.golang.org/protobuf/testing/protocmp
