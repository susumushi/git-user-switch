package profile

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProfile(t *testing.T) {
	conf := "gitus_test_profile"
	t.Run("entry", func(t *testing.T) {
		ps := Profiles{}
		err := ps.Flush(conf)
		assert.NoError(t, err)
		err = ps.Load(conf)
		assert.NoError(t, err)

		err = ps.Set(Profile{
			Name:     "example.testuser2",
			Email:    "test2@example.com",
			NickName: "etu2",
		})
		assert.NoError(t, err)

		err = ps.Save(conf)
		assert.NoError(t, err)
		err = ps.Flush(conf)
		assert.NoError(t, err)
	})
}
