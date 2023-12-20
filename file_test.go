package emf

import (
	"bytes"
	"image/png"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestEMFDraw(t *testing.T) {
	testcases := []struct {
		filename string
		size     int
	}{
		{
			filename: "testdata/x.emf",
			size:     961,
		},
		{
			filename: "testdata/test-000.emf",
			size:     3888,
		},
		{
			filename: "testdata/test-001.emf",
			size:     1942,
		},
	}

	for _, tc := range testcases {
		fBytes, err := ioutil.ReadFile(tc.filename)
		require.NoError(t, err)

		emfFile, err := ReadFile(fBytes)
		require.NotNil(t, emfFile)
		require.NoError(t, err)

		img := emfFile.Draw()

		fw := new(bytes.Buffer)

		err = png.Encode(fw, img)
		require.NoError(t, err)

		require.Equal(t, tc.size, len(fw.Bytes()))
	}
}
