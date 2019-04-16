// Code generated by MockGen. DO NOT EDIT.
// Source: storage/interfaces/db_interface.go

// Package mock_interfaces is a generated GoMock package.
package mock_interfaces

import (
	gomock "github.com/golang/mock/gomock"
	models "github.com/user/2019_1_newTeam2/models"
	reflect "reflect"
)

// MockDBInterface is a mock of DBInterface interface
type MockDBInterface struct {
	ctrl     *gomock.Controller
	recorder *MockDBInterfaceMockRecorder
}

// MockDBInterfaceMockRecorder is the mock recorder for MockDBInterface
type MockDBInterfaceMockRecorder struct {
	mock *MockDBInterface
}

// NewMockDBInterface creates a new mock instance
func NewMockDBInterface(ctrl *gomock.Controller) *MockDBInterface {
	mock := &MockDBInterface{ctrl: ctrl}
	mock.recorder = &MockDBInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockDBInterface) EXPECT() *MockDBInterfaceMockRecorder {
	return m.recorder
}

// Login mocks base method
func (m *MockDBInterface) Login(username, password string, secret []byte) (string, string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Login", username, password, secret)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Login indicates an expected call of Login
func (mr *MockDBInterfaceMockRecorder) Login(username, password, secret interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Login", reflect.TypeOf((*MockDBInterface)(nil).Login), username, password, secret)
}

// GetUserByID mocks base method
func (m *MockDBInterface) GetUserByID(userID int) (models.User, bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByID", userID)
	ret0, _ := ret[0].(models.User)
	ret1, _ := ret[1].(bool)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetUserByID indicates an expected call of GetUserByID
func (mr *MockDBInterfaceMockRecorder) GetUserByID(userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByID", reflect.TypeOf((*MockDBInterface)(nil).GetUserByID), userID)
}

// UserRegistration mocks base method
func (m *MockDBInterface) UserRegistration(username, email, password string, langid, pronounceOn int) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UserRegistration", username, email, password, langid, pronounceOn)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UserRegistration indicates an expected call of UserRegistration
func (mr *MockDBInterfaceMockRecorder) UserRegistration(username, email, password, langid, pronounceOn interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UserRegistration", reflect.TypeOf((*MockDBInterface)(nil).UserRegistration), username, email, password, langid, pronounceOn)
}

// DeleteUserById mocks base method
func (m *MockDBInterface) DeleteUserById(userID int) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteUserById", userID)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteUserById indicates an expected call of DeleteUserById
func (mr *MockDBInterfaceMockRecorder) DeleteUserById(userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUserById", reflect.TypeOf((*MockDBInterface)(nil).DeleteUserById), userID)
}

// GetUsers mocks base method
func (m *MockDBInterface) GetUsers(page, rowsNum int) ([]models.UserTableElem, bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUsers", page, rowsNum)
	ret0, _ := ret[0].([]models.UserTableElem)
	ret1, _ := ret[1].(bool)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetUsers indicates an expected call of GetUsers
func (mr *MockDBInterfaceMockRecorder) GetUsers(page, rowsNum interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUsers", reflect.TypeOf((*MockDBInterface)(nil).GetUsers), page, rowsNum)
}

// AddImage mocks base method
func (m *MockDBInterface) AddImage(path string, userID int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddImage", path, userID)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddImage indicates an expected call of AddImage
func (mr *MockDBInterfaceMockRecorder) AddImage(path, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddImage", reflect.TypeOf((*MockDBInterface)(nil).AddImage), path, userID)
}

// UpdateUserById mocks base method
func (m *MockDBInterface) UpdateUserById(userID int, username, email string, langid, pronounceOn int) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUserById", userID, username, email, langid, pronounceOn)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateUserById indicates an expected call of UpdateUserById
func (mr *MockDBInterfaceMockRecorder) UpdateUserById(userID, username, email, langid, pronounceOn interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUserById", reflect.TypeOf((*MockDBInterface)(nil).UpdateUserById), userID, username, email, langid, pronounceOn)
}

// GetLangs mocks base method
func (m *MockDBInterface) GetLangs() ([]models.Language, bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLangs")
	ret0, _ := ret[0].([]models.Language)
	ret1, _ := ret[1].(bool)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetLangs indicates an expected call of GetLangs
func (mr *MockDBInterfaceMockRecorder) GetLangs() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLangs", reflect.TypeOf((*MockDBInterface)(nil).GetLangs))
}

// GetCards mocks base method
func (m *MockDBInterface) GetCards(dictId, page, rowsNum int) ([]models.Card, bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCards", dictId, page, rowsNum)
	ret0, _ := ret[0].([]models.Card)
	ret1, _ := ret[1].(bool)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetCards indicates an expected call of GetCards
func (mr *MockDBInterfaceMockRecorder) GetCards(dictId, page, rowsNum interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCards", reflect.TypeOf((*MockDBInterface)(nil).GetCards), dictId, page, rowsNum)
}

// GetCard mocks base method
func (m *MockDBInterface) GetCard(cardId int) (models.Card, bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCard", cardId)
	ret0, _ := ret[0].(models.Card)
	ret1, _ := ret[1].(bool)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetCard indicates an expected call of GetCard
func (mr *MockDBInterfaceMockRecorder) GetCard(cardId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCard", reflect.TypeOf((*MockDBInterface)(nil).GetCard), cardId)
}

// SetCardToDictionary mocks base method
func (m *MockDBInterface) SetCardToDictionary(dictID int, card models.Card) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetCardToDictionary", dictID, card)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetCardToDictionary indicates an expected call of SetCardToDictionary
func (mr *MockDBInterfaceMockRecorder) SetCardToDictionary(dictID, card interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetCardToDictionary", reflect.TypeOf((*MockDBInterface)(nil).SetCardToDictionary), dictID, card)
}

// DeleteCardInDictionary mocks base method
func (m *MockDBInterface) DeleteCardInDictionary(cardID, dictionaryID int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteCardInDictionary", cardID, dictionaryID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteCardInDictionary indicates an expected call of DeleteCardInDictionary
func (mr *MockDBInterfaceMockRecorder) DeleteCardInDictionary(cardID, dictionaryID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCardInDictionary", reflect.TypeOf((*MockDBInterface)(nil).DeleteCardInDictionary), cardID, dictionaryID)
}

// DictionaryDelete mocks base method
func (m *MockDBInterface) DictionaryDelete(DictID int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DictionaryDelete", DictID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DictionaryDelete indicates an expected call of DictionaryDelete
func (mr *MockDBInterfaceMockRecorder) DictionaryDelete(DictID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DictionaryDelete", reflect.TypeOf((*MockDBInterface)(nil).DictionaryDelete), DictID)
}

// DictionaryCreate mocks base method
func (m *MockDBInterface) DictionaryCreate(UserID int, Name, Description string, Cards []models.Card) (models.DictionaryInfoPrivilege, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DictionaryCreate", UserID, Name, Description, Cards)
	ret0, _ := ret[0].(models.DictionaryInfoPrivilege)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DictionaryCreate indicates an expected call of DictionaryCreate
func (mr *MockDBInterfaceMockRecorder) DictionaryCreate(UserID, Name, Description, Cards interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DictionaryCreate", reflect.TypeOf((*MockDBInterface)(nil).DictionaryCreate), UserID, Name, Description, Cards)
}

// DictionaryUpdate mocks base method
func (m *MockDBInterface) DictionaryUpdate(DictID int, Name, Description string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DictionaryUpdate", DictID, Name, Description)
	ret0, _ := ret[0].(error)
	return ret0
}

// DictionaryUpdate indicates an expected call of DictionaryUpdate
func (mr *MockDBInterfaceMockRecorder) DictionaryUpdate(DictID, Name, Description interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DictionaryUpdate", reflect.TypeOf((*MockDBInterface)(nil).DictionaryUpdate), DictID, Name, Description)
}

// GetDicts mocks base method
func (m *MockDBInterface) GetDicts(userId, page, rowsNum int) ([]models.DictionaryInfo, bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDicts", userId, page, rowsNum)
	ret0, _ := ret[0].([]models.DictionaryInfo)
	ret1, _ := ret[1].(bool)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetDicts indicates an expected call of GetDicts
func (mr *MockDBInterfaceMockRecorder) GetDicts(userId, page, rowsNum interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDicts", reflect.TypeOf((*MockDBInterface)(nil).GetDicts), userId, page, rowsNum)
}

// GetDict mocks base method
func (m *MockDBInterface) GetDict(dictId int) (models.DictionaryInfoPrivilege, bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDict", dictId)
	ret0, _ := ret[0].(models.DictionaryInfoPrivilege)
	ret1, _ := ret[1].(bool)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetDict indicates an expected call of GetDict
func (mr *MockDBInterfaceMockRecorder) GetDict(dictId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDict", reflect.TypeOf((*MockDBInterface)(nil).GetDict), dictId)
}

// BorrowDictById mocks base method
func (m *MockDBInterface) BorrowDictById(dictId, thiefId int) (int, models.DictionaryInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BorrowDictById", dictId, thiefId)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(models.DictionaryInfo)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// BorrowDictById indicates an expected call of BorrowDictById
func (mr *MockDBInterfaceMockRecorder) BorrowDictById(dictId, thiefId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BorrowDictById", reflect.TypeOf((*MockDBInterface)(nil).BorrowDictById), dictId, thiefId)
}

// MockCardManager is a mock of CardManager interface
type MockCardManager struct {
	ctrl     *gomock.Controller
	recorder *MockCardManagerMockRecorder
}

// MockCardManagerMockRecorder is the mock recorder for MockCardManager
type MockCardManagerMockRecorder struct {
	mock *MockCardManager
}

// NewMockCardManager creates a new mock instance
func NewMockCardManager(ctrl *gomock.Controller) *MockCardManager {
	mock := &MockCardManager{ctrl: ctrl}
	mock.recorder = &MockCardManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCardManager) EXPECT() *MockCardManagerMockRecorder {
	return m.recorder
}

// GetCards mocks base method
func (m *MockCardManager) GetCards(dictId, page, rowsNum int) ([]models.Card, bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCards", dictId, page, rowsNum)
	ret0, _ := ret[0].([]models.Card)
	ret1, _ := ret[1].(bool)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetCards indicates an expected call of GetCards
func (mr *MockCardManagerMockRecorder) GetCards(dictId, page, rowsNum interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCards", reflect.TypeOf((*MockCardManager)(nil).GetCards), dictId, page, rowsNum)
}

// GetCard mocks base method
func (m *MockCardManager) GetCard(cardId int) (models.Card, bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCard", cardId)
	ret0, _ := ret[0].(models.Card)
	ret1, _ := ret[1].(bool)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetCard indicates an expected call of GetCard
func (mr *MockCardManagerMockRecorder) GetCard(cardId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCard", reflect.TypeOf((*MockCardManager)(nil).GetCard), cardId)
}

// SetCardToDictionary mocks base method
func (m *MockCardManager) SetCardToDictionary(dictID int, card models.Card) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetCardToDictionary", dictID, card)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetCardToDictionary indicates an expected call of SetCardToDictionary
func (mr *MockCardManagerMockRecorder) SetCardToDictionary(dictID, card interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetCardToDictionary", reflect.TypeOf((*MockCardManager)(nil).SetCardToDictionary), dictID, card)
}

// DeleteCardInDictionary mocks base method
func (m *MockCardManager) DeleteCardInDictionary(cardID, dictionaryID int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteCardInDictionary", cardID, dictionaryID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteCardInDictionary indicates an expected call of DeleteCardInDictionary
func (mr *MockCardManagerMockRecorder) DeleteCardInDictionary(cardID, dictionaryID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCardInDictionary", reflect.TypeOf((*MockCardManager)(nil).DeleteCardInDictionary), cardID, dictionaryID)
}

// MockDictionaryManager is a mock of DictionaryManager interface
type MockDictionaryManager struct {
	ctrl     *gomock.Controller
	recorder *MockDictionaryManagerMockRecorder
}

// MockDictionaryManagerMockRecorder is the mock recorder for MockDictionaryManager
type MockDictionaryManagerMockRecorder struct {
	mock *MockDictionaryManager
}

// NewMockDictionaryManager creates a new mock instance
func NewMockDictionaryManager(ctrl *gomock.Controller) *MockDictionaryManager {
	mock := &MockDictionaryManager{ctrl: ctrl}
	mock.recorder = &MockDictionaryManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockDictionaryManager) EXPECT() *MockDictionaryManagerMockRecorder {
	return m.recorder
}

// DictionaryDelete mocks base method
func (m *MockDictionaryManager) DictionaryDelete(DictID int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DictionaryDelete", DictID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DictionaryDelete indicates an expected call of DictionaryDelete
func (mr *MockDictionaryManagerMockRecorder) DictionaryDelete(DictID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DictionaryDelete", reflect.TypeOf((*MockDictionaryManager)(nil).DictionaryDelete), DictID)
}

// DictionaryCreate mocks base method
func (m *MockDictionaryManager) DictionaryCreate(UserID int, Name, Description string, Cards []models.Card) (models.DictionaryInfoPrivilege, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DictionaryCreate", UserID, Name, Description, Cards)
	ret0, _ := ret[0].(models.DictionaryInfoPrivilege)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DictionaryCreate indicates an expected call of DictionaryCreate
func (mr *MockDictionaryManagerMockRecorder) DictionaryCreate(UserID, Name, Description, Cards interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DictionaryCreate", reflect.TypeOf((*MockDictionaryManager)(nil).DictionaryCreate), UserID, Name, Description, Cards)
}

// DictionaryUpdate mocks base method
func (m *MockDictionaryManager) DictionaryUpdate(DictID int, Name, Description string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DictionaryUpdate", DictID, Name, Description)
	ret0, _ := ret[0].(error)
	return ret0
}

// DictionaryUpdate indicates an expected call of DictionaryUpdate
func (mr *MockDictionaryManagerMockRecorder) DictionaryUpdate(DictID, Name, Description interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DictionaryUpdate", reflect.TypeOf((*MockDictionaryManager)(nil).DictionaryUpdate), DictID, Name, Description)
}

// GetDicts mocks base method
func (m *MockDictionaryManager) GetDicts(userId, page, rowsNum int) ([]models.DictionaryInfo, bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDicts", userId, page, rowsNum)
	ret0, _ := ret[0].([]models.DictionaryInfo)
	ret1, _ := ret[1].(bool)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetDicts indicates an expected call of GetDicts
func (mr *MockDictionaryManagerMockRecorder) GetDicts(userId, page, rowsNum interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDicts", reflect.TypeOf((*MockDictionaryManager)(nil).GetDicts), userId, page, rowsNum)
}

// GetDict mocks base method
func (m *MockDictionaryManager) GetDict(dictId int) (models.DictionaryInfoPrivilege, bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDict", dictId)
	ret0, _ := ret[0].(models.DictionaryInfoPrivilege)
	ret1, _ := ret[1].(bool)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetDict indicates an expected call of GetDict
func (mr *MockDictionaryManagerMockRecorder) GetDict(dictId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDict", reflect.TypeOf((*MockDictionaryManager)(nil).GetDict), dictId)
}

// BorrowDictById mocks base method
func (m *MockDictionaryManager) BorrowDictById(dictId, thiefId int) (int, models.DictionaryInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BorrowDictById", dictId, thiefId)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(models.DictionaryInfo)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// BorrowDictById indicates an expected call of BorrowDictById
func (mr *MockDictionaryManagerMockRecorder) BorrowDictById(dictId, thiefId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BorrowDictById", reflect.TypeOf((*MockDictionaryManager)(nil).BorrowDictById), dictId, thiefId)
}

// MockLanguageManager is a mock of LanguageManager interface
type MockLanguageManager struct {
	ctrl     *gomock.Controller
	recorder *MockLanguageManagerMockRecorder
}

// MockLanguageManagerMockRecorder is the mock recorder for MockLanguageManager
type MockLanguageManagerMockRecorder struct {
	mock *MockLanguageManager
}

// NewMockLanguageManager creates a new mock instance
func NewMockLanguageManager(ctrl *gomock.Controller) *MockLanguageManager {
	mock := &MockLanguageManager{ctrl: ctrl}
	mock.recorder = &MockLanguageManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockLanguageManager) EXPECT() *MockLanguageManagerMockRecorder {
	return m.recorder
}

// GetLangs mocks base method
func (m *MockLanguageManager) GetLangs() ([]models.Language, bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLangs")
	ret0, _ := ret[0].([]models.Language)
	ret1, _ := ret[1].(bool)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetLangs indicates an expected call of GetLangs
func (mr *MockLanguageManagerMockRecorder) GetLangs() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLangs", reflect.TypeOf((*MockLanguageManager)(nil).GetLangs))
}

// MockUserManager is a mock of UserManager interface
type MockUserManager struct {
	ctrl     *gomock.Controller
	recorder *MockUserManagerMockRecorder
}

// MockUserManagerMockRecorder is the mock recorder for MockUserManager
type MockUserManagerMockRecorder struct {
	mock *MockUserManager
}

// NewMockUserManager creates a new mock instance
func NewMockUserManager(ctrl *gomock.Controller) *MockUserManager {
	mock := &MockUserManager{ctrl: ctrl}
	mock.recorder = &MockUserManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockUserManager) EXPECT() *MockUserManagerMockRecorder {
	return m.recorder
}

// Login mocks base method
func (m *MockUserManager) Login(username, password string, secret []byte) (string, string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Login", username, password, secret)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Login indicates an expected call of Login
func (mr *MockUserManagerMockRecorder) Login(username, password, secret interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Login", reflect.TypeOf((*MockUserManager)(nil).Login), username, password, secret)
}

// GetUserByID mocks base method
func (m *MockUserManager) GetUserByID(userID int) (models.User, bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByID", userID)
	ret0, _ := ret[0].(models.User)
	ret1, _ := ret[1].(bool)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetUserByID indicates an expected call of GetUserByID
func (mr *MockUserManagerMockRecorder) GetUserByID(userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByID", reflect.TypeOf((*MockUserManager)(nil).GetUserByID), userID)
}

// UserRegistration mocks base method
func (m *MockUserManager) UserRegistration(username, email, password string, langid, pronounceOn int) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UserRegistration", username, email, password, langid, pronounceOn)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UserRegistration indicates an expected call of UserRegistration
func (mr *MockUserManagerMockRecorder) UserRegistration(username, email, password, langid, pronounceOn interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UserRegistration", reflect.TypeOf((*MockUserManager)(nil).UserRegistration), username, email, password, langid, pronounceOn)
}

// DeleteUserById mocks base method
func (m *MockUserManager) DeleteUserById(userID int) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteUserById", userID)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteUserById indicates an expected call of DeleteUserById
func (mr *MockUserManagerMockRecorder) DeleteUserById(userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUserById", reflect.TypeOf((*MockUserManager)(nil).DeleteUserById), userID)
}

// GetUsers mocks base method
func (m *MockUserManager) GetUsers(page, rowsNum int) ([]models.UserTableElem, bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUsers", page, rowsNum)
	ret0, _ := ret[0].([]models.UserTableElem)
	ret1, _ := ret[1].(bool)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetUsers indicates an expected call of GetUsers
func (mr *MockUserManagerMockRecorder) GetUsers(page, rowsNum interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUsers", reflect.TypeOf((*MockUserManager)(nil).GetUsers), page, rowsNum)
}

// AddImage mocks base method
func (m *MockUserManager) AddImage(path string, userID int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddImage", path, userID)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddImage indicates an expected call of AddImage
func (mr *MockUserManagerMockRecorder) AddImage(path, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddImage", reflect.TypeOf((*MockUserManager)(nil).AddImage), path, userID)
}

// UpdateUserById mocks base method
func (m *MockUserManager) UpdateUserById(userID int, username, email string, langid, pronounceOn int) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUserById", userID, username, email, langid, pronounceOn)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateUserById indicates an expected call of UpdateUserById
func (mr *MockUserManagerMockRecorder) UpdateUserById(userID, username, email, langid, pronounceOn interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUserById", reflect.TypeOf((*MockUserManager)(nil).UpdateUserById), userID, username, email, langid, pronounceOn)
}
