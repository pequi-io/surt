package storage

import (
	"testing"
	"time"

	"github.com/pequi-io/surt/pkg/entity"
	"github.com/pequi-io/surt/pkg/storage/services/mockstorage"
	"github.com/pequi-io/surt/pkg/util"
	"github.com/stretchr/testify/assert"
)

var (
	objTags = map[string]string{
		"SURT_LAST_SCAN":   time.Now().String(),
		"SURT_SCAN_STATUS": "CLEAN",
	}
	obj = entity.Object{
		RawFilePath: "/tmp/myobject.zip",
		Tags:        objTags,
	}
	objWithoutPath = entity.Object{
		RawFilePath: "",
		Tags:        map[string]string{},
	}
	objToFail = entity.Object{
		RawFilePath: "fail",
		Tags:        objTags,
	}
)

func TestGetObject(t *testing.T) {

	service := mockstorage.New()
	s := New(service)

	body, err := s.GetObject(&obj)
	assert.Nil(t, err)
	assert.Equal(t, "fake", util.ByteToString(body), "object content/body should be equal")

	_, err = s.GetObject(&objWithoutPath)
	assert.NotNil(t, err)

	_, err = s.GetObject(&objToFail)
	assert.NotNil(t, err)

}

func TestGetObjectTags(t *testing.T) {

	service := mockstorage.New()
	s := New(service)

	tags, err := s.GetObjectTags(&obj)
	assert.Nil(t, err)
	assert.Equal(t, 2, len(tags), "number of tags should be equal")

	_, err = s.GetObjectTags(&objWithoutPath)
	assert.NotNil(t, err)

	_, err = s.GetObjectTags(&objToFail)
	assert.NotNil(t, err)
}

func TestSetObjectTags(t *testing.T) {

	service := mockstorage.New()
	s := New(service)

	err := s.SetObjectTags(&obj)
	assert.Nil(t, err)

	err = s.SetObjectTags(&objWithoutPath)
	assert.NotNil(t, err)

	err = s.SetObjectTags(&objToFail)
	assert.NotNil(t, err)
}
