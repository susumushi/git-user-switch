package gituser

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetLocalRepo(t *testing.T) {
	t.Run("if git repository dose not exist, then return error", func(t *testing.T) {
		err := os.Chdir(`/tmp`)
		assert.NoError(t, err)

		r, err := getLocalRepo()
		assert.Nil(t, r)
		assert.Error(t, err)
	})
}
