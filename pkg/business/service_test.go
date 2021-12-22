package business

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDefaultService_Create(t *testing.T) {
	s := NewDefaultService(NewRepoTest(nil))
	assert.NoError(t, s.Create(mockBusiness))

	s = NewDefaultService(NewRepoTest(mockError))
	assert.Error(t, s.Create(mockBusiness))
}

func TestDefaultService_ReadOne(t *testing.T) {
	s := NewDefaultService(NewRepoTest(nil))
	assert.NoError(t, s.ReadOne(mockBusiness))

	s = NewDefaultService(NewRepoTest(mockError))
	assert.Error(t, s.ReadOne(mockBusiness))
}

func TestDefaultService_ReadAll(t *testing.T) {
	s := NewDefaultService(NewRepoTest(nil))
	assert.NoError(t, s.ReadAll(&[]Business{}))

	s = NewDefaultService(NewRepoTest(mockError))
	assert.Error(t, s.ReadAll(&[]Business{}))
}

func TestDefaultService_Update(t *testing.T) {
	s := NewDefaultService(NewRepoTest(nil))
	assert.NoError(t, s.Update(mockBusiness))

	s = NewDefaultService(NewRepoTest(mockError))
	assert.Error(t, s.Update(mockBusiness))
}

func TestDefaultService_Delete(t *testing.T) {
	s := NewDefaultService(NewRepoTest(nil))
	assert.NoError(t, s.Delete(mockBusiness))

	s = NewDefaultService(NewRepoTest(mockError))
	assert.Error(t, s.Delete(mockBusiness))
}
