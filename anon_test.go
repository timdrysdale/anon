package anon

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {

	_, err := New("./test/identity.csv")
	assert.NoError(t, err)
}

func TestGetAnonymous(t *testing.T) {

	a, err := New("./test/identity.csv")
	assert.NoError(t, err)

	anon, err := a.GetAnonymous("s00000000")
	assert.Equal(t, anon, "B999999")

	anon, err = a.GetAnonymous("s00000003")
	assert.Equal(t, anon, "B999996")

	anon, err = a.GetAnonymous("s00000006")
	assert.Equal(t, anon, "B999994")

	anon, err = a.GetAnonymous("NotInDictionary")
	assert.Error(t, err)

}

func TestGetIdentity(t *testing.T) {

	a, err := New("./test/identity.csv")
	assert.NoError(t, err)

	id, err := a.GetIdentity("B999999")
	assert.Equal(t, id, "s00000000")

	id, err = a.GetIdentity("B999997")
	assert.Equal(t, id, "s00000002")

	id, err = a.GetIdentity("B999995")
	assert.Equal(t, id, "s00000005")

	id, err = a.GetIdentity("NotInDictionary")
	assert.Error(t, err)

}

func TestGetLength(t *testing.T) {

	a, err := New("./test/identity.csv")
	assert.NoError(t, err)

	assert.Equal(t, a.GetLength(), 6)

}
