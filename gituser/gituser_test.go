package gituser

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetLocalRepo(t *testing.T) {
	t.Run("if git repository dose not exist, then return error", func(t *testing.T) {
		os.Chdir(`/tmp`)
		_, err := getLocalRepo()
		assert.Error(t, err)
	})
}
