package cmd

import (
	"bytes"
	"os/exec"
	"testing"

	"github.com/shynome/err0"
	"github.com/shynome/err0/try"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	compile()
	m.Run()
}

func TestCall(t *testing.T) {
	var err error
	defer err0.Then(&err, nil, func() {
		t.Error(err)
	})
	input := []byte(`{"name":"host"}`)
	stdin := bytes.NewBuffer(input)
	stdout := new(bytes.Buffer)
	stderr := new(bytes.Buffer)

	rootCmd.SetIn(stdin)
	rootCmd.SetOut(stdout)
	rootCmd.SetErr(stderr)
	rootCmd.SetArgs([]string{"../guest/guest.wasm", "hello"})

	try.To(rootCmd.Execute())

	assert.Equal(t, "", stderr.String())
	assert.Equal(t, "hello host", stdout.String())
}

func compile() {
	cmd := exec.Command("go", "generate", "../guest")
	try.To(cmd.Run())
}
