package tax

import "github.com/stretchr/testify/mock"

type TaxRepositoryMock struct {
	mock.Mock
}

func (m *TaxRepositoryMock) Save(tax float64) error {
	ret := m.Called(tax)
	return ret.Error(0)
}
