package employee

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDefaultService_Create(t *testing.T) {
	s := NewDefaultService(NewRepoTest(nil))
	assert.NoError(t, s.Create(mockEmployee))

	s = NewDefaultService(NewRepoTest(mockError))
	assert.Error(t, s.Create(mockEmployee))
}

func TestDefaultService_ReadOne(t *testing.T) {
	s := NewDefaultService(NewRepoTest(nil))
	assert.NoError(t, s.ReadOne(mockEmployee))

	s = NewDefaultService(NewRepoTest(mockError))
	assert.Error(t, s.ReadOne(mockEmployee))
}

func TestDefaultService_ReadAll(t *testing.T) {
	s := NewDefaultService(NewRepoTest(nil))
	assert.NoError(t, s.ReadAll(&[]Employee{}))

	s = NewDefaultService(NewRepoTest(mockError))
	assert.Error(t, s.ReadAll(&[]Employee{}))
}

func TestDefaultService_Update(t *testing.T) {
	s := NewDefaultService(NewRepoTest(nil))
	assert.NoError(t, s.Update(mockEmployee))

	s = NewDefaultService(NewRepoTest(mockError))
	assert.Error(t, s.Update(mockEmployee))
}

func TestDefaultService_Delete(t *testing.T) {
	s := NewDefaultService(NewRepoTest(nil))
	assert.NoError(t, s.Delete(mockEmployee))

	s = NewDefaultService(NewRepoTest(mockError))
	assert.Error(t, s.Delete(mockEmployee))
}
