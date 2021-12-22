package business

import "errors"

var (
	mockBusiness = &Business{}
	mockError    = errors.New("error")
)

type RepoTest struct {
	data    []Business
	wantErr error
}

func NewRepoTest(wantErr error) *RepoTest {
	return &RepoTest{
		data:    []Business{{}},
		wantErr: wantErr,
	}
}

func (r *RepoTest) Create(data *Business) error    { data = &r.data[0]; return r.wantErr }
func (r *RepoTest) ReadOne(data *Business) error   { data = &r.data[0]; return r.wantErr }
func (r *RepoTest) ReadAll(data *[]Business) error { data = &r.data; return r.wantErr }
func (r *RepoTest) Update(data *Business) error    { data = &r.data[0]; return r.wantErr }
func (r *RepoTest) Delete(data *Business) error    { data = &r.data[0]; return r.wantErr }
