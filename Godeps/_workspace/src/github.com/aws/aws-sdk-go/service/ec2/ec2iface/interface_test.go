// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

package ec2iface_test

import (
	"testing"

	"github.com/99designs/iamy/Godeps/_workspace/src/github.com/aws/aws-sdk-go/service/ec2"
	"github.com/99designs/iamy/Godeps/_workspace/src/github.com/aws/aws-sdk-go/service/ec2/ec2iface"
	"github.com/stretchr/testify/assert"
)

func TestInterface(t *testing.T) {
	assert.Implements(t, (*ec2iface.EC2API)(nil), ec2.New(nil))
}