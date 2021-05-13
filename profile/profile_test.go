package profile

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProfile(t *testing.T) {
	t.Run("entry", func(t *testing.T) {
		ps := Profiles{}
		err := ps.Load()
		assert.NoError(t, err)

		err = ps.Set(Profile{
			Name:     "example.testuser2",
			Email:    "test2@example.com",
			NickName: "etu2",
		})
		assert.NoError(t, err)

		err = ps.Save()
		assert.NoError(t, err)
		err = ps.Flush()
		assert.NoError(t, err)
	})
}
