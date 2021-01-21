package geographer

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestServices(t *testing.T) {

	assert := assert.New(t)

	s, err := GetServices(context.TODO())

	assert.NoError(err)

	r := s.Services()

	assert.EqualValues([]string{
		"eu-west-1",
		"us-east-1",
		"us-west-2",
	}, r["workmail"])

	assert.EqualValues([]string{
		"eu-west-1",
		"us-east-1",
		"us-west-2",
	}, r["workmail"].Intersection())

	assert.EqualValues([]string{
		"eu-west-1",
		"us-east-1",
	}, r["workmail"].Intersection("eu-west-1", "us-east-1", "no-such-region"))

}

func TestStaticServices(t *testing.T) {

	assert := assert.New(t)

	r := Services

	assert.EqualValues([]string{
		"eu-west-1",
		"us-east-1",
		"us-west-2",
	}, r["workmail"])

	assert.EqualValues([]string{
		"eu-west-1",
		"us-east-1",
		"us-west-2",
	}, r["workmail"].Intersection())

	assert.EqualValues([]string{
		"eu-west-1",
		"us-east-1",
	}, r["workmail"].Intersection("eu-west-1", "us-east-1", "no-such-region"))

}
