package functions_test

import (
	"bytes"
	"fmt"
	"os"
	"strings"
	"testing"

	shell "github.com/ipfs/go-ipfs-api"
	"github.com/sonr-io/sonr/pkg/functions"
	"github.com/stretchr/testify/assert"
)

func Test_Functions(t *testing.T) {
	t.Skip("cannot run in CI")
	shell := shell.NewLocalShell()
	filepath := "/Users/joshlong/Documents/fun-projects/test/main"
	t.Run("Should store file and be in cache", func(t *testing.T) {
		file, err := os.ReadFile(filepath)
		assert.NoError(t, err)
		f := functions.NewFunction(&file, "")

		executor := functions.New(shell)
		bldr := strings.Builder{}
		err = executor.Execute(f, &bldr)
		assert.NoError(t, err)
	})

	t.Run("Should store file and be in cache and execute", func(t *testing.T) {
		file, err := os.ReadFile(filepath)
		assert.NoError(t, err)
		f := functions.NewFunction(&file, "")
		b, err := f.Marshal()
		assert.NoError(t, err)
		cid, err := shell.Add(bytes.NewBuffer(b))

		assert.NoError(t, err)
		executor := functions.New(shell)
		bldr := strings.Builder{}
		err = executor.GetAndExecute(cid, &bldr)
		assert.NoError(t, err)
		str := bldr.String()
		fmt.Print(str)
	})
}
