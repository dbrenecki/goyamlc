package cmd

import (
	"go/ast"
	"go/token"
	"os"
	"testing"

	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	typesPath = "../test/data/types.go"
)

func Test_generate(t *testing.T) {
	err := configureLogger("info")
	require.NoError(t, err)

	testHook := func(fs *token.FileSet, f *ast.File) (err error) {
		out, err := os.Create("../test/data/types.txt")
		if err != nil {
			return err
		}

		defer func() {
			if tmpErr := out.Close(); tmpErr != nil || err != nil {
				err = tmpErr
			}
		}()

		log.Info().Msg("create ast obj tree")

		err = ast.Fprint(out, fs, f, nil)
		if err != nil {
			return err
		}
		return nil
	}

	err = Generate(typesPath, testHook)
	assert.NoError(t, err)
}
