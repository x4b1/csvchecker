package csvchecker

import (
	"errors"
	"reflect"
	"strings"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type CheckerTestSuite struct {
	suite.Suite
	checker   *checker
	separator rune
}

func (suite *CheckerTestSuite) SetupTest() {
	suite.separator = rune(';')
	suite.checker = NewChecker(suite.separator, true)
}

func (suite *CheckerTestSuite) TestCheckerCreatedWithSeparatorAndEmptyValidators() {
	suite.Equal(suite.separator, suite.checker.separator)
	suite.Empty(suite.checker.columns)
	suite.True(suite.checker.withHeader)
}

func (suite *CheckerTestSuite) TestAddNewColumAddsNewItemToColumns() {
	insertColumn := NewColumn(1, new(MockedValidator))
	suite.Empty(suite.checker.columns)
	suite.checker.AddColum(insertColumn)
	suite.Len(suite.checker.columns, 1)
	suite.IsType(insertColumn, suite.checker.columns[0])
}

func (suite *CheckerTestSuite) TestCheckWithInvalidNumberRowsReturnsError() {
	csv := `id;name;text
	123;John;"hello"
	432;Doe;"hello2";122`
	errs := suite.checker.Check(strings.NewReader(csv))

	suite.Len(errs, 1)
	suite.IsType(new(rowError), errs[0])
}

func (suite *CheckerTestSuite) TestCheckWithHeaderNotChecksHeader() {
	validatorMock := new(MockedValidator)
	validatorMock.On("Validate", mock.AnythingOfType("string")).Return(nil)
	suite.checker.AddColum(NewColumn(1, validatorMock))
	csv := `id;name;text
	123;John;"hello"`

	suite.checker.Check(strings.NewReader(csv))

	validatorMock.AssertNumberOfCalls(suite.Suite.T(), "Validate", 1)
}

func (suite *CheckerTestSuite) TestCheckColumnValidationReturnsError() {
	validatorMock := new(MockedValidator)
	validatorMock.On("Validate", mock.AnythingOfType("string")).Return(errors.New("Paco"))
	suite.checker.AddColum(NewColumn(1, validatorMock))
	csv := `id;name;text
	123;John;"hello"`
	errs := suite.checker.Check(strings.NewReader(csv))

	suite.Len(errs, 1)

	testError := errs[0]
	suite.IsType(new(colError), testError)
	r := reflect.ValueOf(testError)
	line := reflect.Indirect(r).FieldByName("line")
	column := reflect.Indirect(r).FieldByName("col")
	suite.Equal(2, int(line.Int()))
	suite.Equal(1, int(column.Int()))
}

func TestCheckerSuite(t *testing.T) {
	suite.Run(t, new(CheckerTestSuite))
}

type MockedValidator struct {
	mock.Mock
}

func (m *MockedValidator) Validate(s string) error {
	args := m.Called(s)
	return args.Error(0)
}
