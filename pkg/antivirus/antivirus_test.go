package antivirus

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/surt-io/surt/pkg/antivirus/engine/fake"
	"github.com/surt-io/surt/pkg/types"
)

var (
	obj = types.Object{
		Content: []byte("fake"),
	}
	emptyContent = types.Object{
		Content: make([]byte, 0, 1),
	}
)

func TestNewAV(t *testing.T) {

	avengine := fake.New()
	av := New(avengine)

	res, err := av.Scan(&obj)
	assert.Nil(t, err)
	assert.Equal(t, "CLEAN", res[0].Status, "av scan result should be equal")

	h, err := av.GetHealthStatus()
	assert.Nil(t, err)
	assert.Equal(t, "healthy", h, "av engien healthcheck status should be equal")

}

func TestScanWithError(t *testing.T) {

	avengine := fake.New()
	av := New(avengine)

	res, err := av.Scan(&emptyContent)
	assert.NotNil(t, err)
	assert.Equal(t, []types.Result(nil), res, "av scan result should be equal")

}
