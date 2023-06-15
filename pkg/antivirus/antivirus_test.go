package antivirus

import (
	"testing"

	"github.com/pequi-io/surt/pkg/antivirus/engine/mockengine"
	"github.com/pequi-io/surt/pkg/entity"
	"github.com/stretchr/testify/assert"
)

var (
	obj = entity.Object{
		Content: []byte("fake"),
	}
	emptyContent = entity.Object{
		Content: make([]byte, 0, 1),
	}
)

func TestNewAV(t *testing.T) {

	avengine := mockengine.New()
	av := New(avengine)

	res, err := av.ScanObject(&obj)
	assert.Nil(t, err)
	assert.Equal(t, "CLEAN", res[0].Status, "av scan result should be equal")

	h, err := av.GetHealthStatus()
	assert.Nil(t, err)
	assert.Equal(t, "healthy", h, "av engien healthcheck status should be equal")

}

func TestScanWithError(t *testing.T) {

	avengine := mockengine.New()
	av := New(avengine)

	res, err := av.ScanObject(&emptyContent)
	assert.NotNil(t, err)
	assert.Equal(t, []entity.Result(nil), res, "av scan result should be equal")

}
