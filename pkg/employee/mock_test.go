package employee

import "errors"

var (
	mockEmployee = &Employee{}
	mockError    = errors.New("error")
)

type RepoTest struct {
	data    []Employee
	wantErr error
}

func NewRepoTest(wantErr error) *RepoTest {
	return &RepoTest{
		data:    []Employee{{}},
		wantErr: wantErr,
	}
}

func (r *RepoTest) Create(data *Employee) error    { data = &r.data[0]; return r.wantErr }
func (r *RepoTest) ReadOne(data *Employee) error   { data = &r.data[0]; return r.wantErr }
func (r *RepoTest) ReadAll(data *[]Employee) error { data = &r.data; return r.wantErr }
func (r *RepoTest) Update(data *Employee) error    { data = &r.data[0]; return r.wantErr }
func (r *RepoTest) Delete(data *Employee) error    { data = &r.data[0]; return r.wantErr }
