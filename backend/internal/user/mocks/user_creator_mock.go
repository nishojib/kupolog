// Code generated by http://github.com/gojuno/minimock (v3.3.13). DO NOT EDIT.

package mocks

//go:generate minimock -i github.com/nishojib/ffxivdailies/internal/user.userCreator -o user_creator_mock.go -n UserCreatorMock -p mocks

import (
	"context"
	"sync"
	mm_atomic "sync/atomic"
	mm_time "time"

	"github.com/gojuno/minimock/v3"
	mm_user "github.com/nishojib/ffxivdailies/internal/user"
)

// UserCreatorMock implements user.userCreator
type UserCreatorMock struct {
	t          minimock.Tester
	finishOnce sync.Once

	funcGetUserByProviderID          func(ctx context.Context, providerAccountID string) (u1 mm_user.User, err error)
	inspectFuncGetUserByProviderID   func(ctx context.Context, providerAccountID string)
	afterGetUserByProviderIDCounter  uint64
	beforeGetUserByProviderIDCounter uint64
	GetUserByProviderIDMock          mUserCreatorMockGetUserByProviderID

	funcInsertAndLinkAccount          func(ctx context.Context, user *mm_user.User, account *mm_user.Account) (err error)
	inspectFuncInsertAndLinkAccount   func(ctx context.Context, user *mm_user.User, account *mm_user.Account)
	afterInsertAndLinkAccountCounter  uint64
	beforeInsertAndLinkAccountCounter uint64
	InsertAndLinkAccountMock          mUserCreatorMockInsertAndLinkAccount
}

// NewUserCreatorMock returns a mock for user.userCreator
func NewUserCreatorMock(t minimock.Tester) *UserCreatorMock {
	m := &UserCreatorMock{t: t}

	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.GetUserByProviderIDMock = mUserCreatorMockGetUserByProviderID{mock: m}
	m.GetUserByProviderIDMock.callArgs = []*UserCreatorMockGetUserByProviderIDParams{}

	m.InsertAndLinkAccountMock = mUserCreatorMockInsertAndLinkAccount{mock: m}
	m.InsertAndLinkAccountMock.callArgs = []*UserCreatorMockInsertAndLinkAccountParams{}

	t.Cleanup(m.MinimockFinish)

	return m
}

type mUserCreatorMockGetUserByProviderID struct {
	optional           bool
	mock               *UserCreatorMock
	defaultExpectation *UserCreatorMockGetUserByProviderIDExpectation
	expectations       []*UserCreatorMockGetUserByProviderIDExpectation

	callArgs []*UserCreatorMockGetUserByProviderIDParams
	mutex    sync.RWMutex

	expectedInvocations uint64
}

// UserCreatorMockGetUserByProviderIDExpectation specifies expectation struct of the userCreator.GetUserByProviderID
type UserCreatorMockGetUserByProviderIDExpectation struct {
	mock      *UserCreatorMock
	params    *UserCreatorMockGetUserByProviderIDParams
	paramPtrs *UserCreatorMockGetUserByProviderIDParamPtrs
	results   *UserCreatorMockGetUserByProviderIDResults
	Counter   uint64
}

// UserCreatorMockGetUserByProviderIDParams contains parameters of the userCreator.GetUserByProviderID
type UserCreatorMockGetUserByProviderIDParams struct {
	ctx               context.Context
	providerAccountID string
}

// UserCreatorMockGetUserByProviderIDParamPtrs contains pointers to parameters of the userCreator.GetUserByProviderID
type UserCreatorMockGetUserByProviderIDParamPtrs struct {
	ctx               *context.Context
	providerAccountID *string
}

// UserCreatorMockGetUserByProviderIDResults contains results of the userCreator.GetUserByProviderID
type UserCreatorMockGetUserByProviderIDResults struct {
	u1  mm_user.User
	err error
}

// Marks this method to be optional. The default behavior of any method with Return() is '1 or more', meaning
// the test will fail minimock's automatic final call check if the mocked method was not called at least once.
// Optional() makes method check to work in '0 or more' mode.
// It is NOT RECOMMENDED to use this option unless you really need it, as default behaviour helps to
// catch the problems when the expected method call is totally skipped during test run.
func (mmGetUserByProviderID *mUserCreatorMockGetUserByProviderID) Optional() *mUserCreatorMockGetUserByProviderID {
	mmGetUserByProviderID.optional = true
	return mmGetUserByProviderID
}

// Expect sets up expected params for userCreator.GetUserByProviderID
func (mmGetUserByProviderID *mUserCreatorMockGetUserByProviderID) Expect(ctx context.Context, providerAccountID string) *mUserCreatorMockGetUserByProviderID {
	if mmGetUserByProviderID.mock.funcGetUserByProviderID != nil {
		mmGetUserByProviderID.mock.t.Fatalf("UserCreatorMock.GetUserByProviderID mock is already set by Set")
	}

	if mmGetUserByProviderID.defaultExpectation == nil {
		mmGetUserByProviderID.defaultExpectation = &UserCreatorMockGetUserByProviderIDExpectation{}
	}

	if mmGetUserByProviderID.defaultExpectation.paramPtrs != nil {
		mmGetUserByProviderID.mock.t.Fatalf("UserCreatorMock.GetUserByProviderID mock is already set by ExpectParams functions")
	}

	mmGetUserByProviderID.defaultExpectation.params = &UserCreatorMockGetUserByProviderIDParams{ctx, providerAccountID}
	for _, e := range mmGetUserByProviderID.expectations {
		if minimock.Equal(e.params, mmGetUserByProviderID.defaultExpectation.params) {
			mmGetUserByProviderID.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmGetUserByProviderID.defaultExpectation.params)
		}
	}

	return mmGetUserByProviderID
}

// ExpectCtxParam1 sets up expected param ctx for userCreator.GetUserByProviderID
func (mmGetUserByProviderID *mUserCreatorMockGetUserByProviderID) ExpectCtxParam1(ctx context.Context) *mUserCreatorMockGetUserByProviderID {
	if mmGetUserByProviderID.mock.funcGetUserByProviderID != nil {
		mmGetUserByProviderID.mock.t.Fatalf("UserCreatorMock.GetUserByProviderID mock is already set by Set")
	}

	if mmGetUserByProviderID.defaultExpectation == nil {
		mmGetUserByProviderID.defaultExpectation = &UserCreatorMockGetUserByProviderIDExpectation{}
	}

	if mmGetUserByProviderID.defaultExpectation.params != nil {
		mmGetUserByProviderID.mock.t.Fatalf("UserCreatorMock.GetUserByProviderID mock is already set by Expect")
	}

	if mmGetUserByProviderID.defaultExpectation.paramPtrs == nil {
		mmGetUserByProviderID.defaultExpectation.paramPtrs = &UserCreatorMockGetUserByProviderIDParamPtrs{}
	}
	mmGetUserByProviderID.defaultExpectation.paramPtrs.ctx = &ctx

	return mmGetUserByProviderID
}

// ExpectProviderAccountIDParam2 sets up expected param providerAccountID for userCreator.GetUserByProviderID
func (mmGetUserByProviderID *mUserCreatorMockGetUserByProviderID) ExpectProviderAccountIDParam2(providerAccountID string) *mUserCreatorMockGetUserByProviderID {
	if mmGetUserByProviderID.mock.funcGetUserByProviderID != nil {
		mmGetUserByProviderID.mock.t.Fatalf("UserCreatorMock.GetUserByProviderID mock is already set by Set")
	}

	if mmGetUserByProviderID.defaultExpectation == nil {
		mmGetUserByProviderID.defaultExpectation = &UserCreatorMockGetUserByProviderIDExpectation{}
	}

	if mmGetUserByProviderID.defaultExpectation.params != nil {
		mmGetUserByProviderID.mock.t.Fatalf("UserCreatorMock.GetUserByProviderID mock is already set by Expect")
	}

	if mmGetUserByProviderID.defaultExpectation.paramPtrs == nil {
		mmGetUserByProviderID.defaultExpectation.paramPtrs = &UserCreatorMockGetUserByProviderIDParamPtrs{}
	}
	mmGetUserByProviderID.defaultExpectation.paramPtrs.providerAccountID = &providerAccountID

	return mmGetUserByProviderID
}

// Inspect accepts an inspector function that has same arguments as the userCreator.GetUserByProviderID
func (mmGetUserByProviderID *mUserCreatorMockGetUserByProviderID) Inspect(f func(ctx context.Context, providerAccountID string)) *mUserCreatorMockGetUserByProviderID {
	if mmGetUserByProviderID.mock.inspectFuncGetUserByProviderID != nil {
		mmGetUserByProviderID.mock.t.Fatalf("Inspect function is already set for UserCreatorMock.GetUserByProviderID")
	}

	mmGetUserByProviderID.mock.inspectFuncGetUserByProviderID = f

	return mmGetUserByProviderID
}

// Return sets up results that will be returned by userCreator.GetUserByProviderID
func (mmGetUserByProviderID *mUserCreatorMockGetUserByProviderID) Return(u1 mm_user.User, err error) *UserCreatorMock {
	if mmGetUserByProviderID.mock.funcGetUserByProviderID != nil {
		mmGetUserByProviderID.mock.t.Fatalf("UserCreatorMock.GetUserByProviderID mock is already set by Set")
	}

	if mmGetUserByProviderID.defaultExpectation == nil {
		mmGetUserByProviderID.defaultExpectation = &UserCreatorMockGetUserByProviderIDExpectation{mock: mmGetUserByProviderID.mock}
	}
	mmGetUserByProviderID.defaultExpectation.results = &UserCreatorMockGetUserByProviderIDResults{u1, err}
	return mmGetUserByProviderID.mock
}

// Set uses given function f to mock the userCreator.GetUserByProviderID method
func (mmGetUserByProviderID *mUserCreatorMockGetUserByProviderID) Set(f func(ctx context.Context, providerAccountID string) (u1 mm_user.User, err error)) *UserCreatorMock {
	if mmGetUserByProviderID.defaultExpectation != nil {
		mmGetUserByProviderID.mock.t.Fatalf("Default expectation is already set for the userCreator.GetUserByProviderID method")
	}

	if len(mmGetUserByProviderID.expectations) > 0 {
		mmGetUserByProviderID.mock.t.Fatalf("Some expectations are already set for the userCreator.GetUserByProviderID method")
	}

	mmGetUserByProviderID.mock.funcGetUserByProviderID = f
	return mmGetUserByProviderID.mock
}

// When sets expectation for the userCreator.GetUserByProviderID which will trigger the result defined by the following
// Then helper
func (mmGetUserByProviderID *mUserCreatorMockGetUserByProviderID) When(ctx context.Context, providerAccountID string) *UserCreatorMockGetUserByProviderIDExpectation {
	if mmGetUserByProviderID.mock.funcGetUserByProviderID != nil {
		mmGetUserByProviderID.mock.t.Fatalf("UserCreatorMock.GetUserByProviderID mock is already set by Set")
	}

	expectation := &UserCreatorMockGetUserByProviderIDExpectation{
		mock:   mmGetUserByProviderID.mock,
		params: &UserCreatorMockGetUserByProviderIDParams{ctx, providerAccountID},
	}
	mmGetUserByProviderID.expectations = append(mmGetUserByProviderID.expectations, expectation)
	return expectation
}

// Then sets up userCreator.GetUserByProviderID return parameters for the expectation previously defined by the When method
func (e *UserCreatorMockGetUserByProviderIDExpectation) Then(u1 mm_user.User, err error) *UserCreatorMock {
	e.results = &UserCreatorMockGetUserByProviderIDResults{u1, err}
	return e.mock
}

// Times sets number of times userCreator.GetUserByProviderID should be invoked
func (mmGetUserByProviderID *mUserCreatorMockGetUserByProviderID) Times(n uint64) *mUserCreatorMockGetUserByProviderID {
	if n == 0 {
		mmGetUserByProviderID.mock.t.Fatalf("Times of UserCreatorMock.GetUserByProviderID mock can not be zero")
	}
	mm_atomic.StoreUint64(&mmGetUserByProviderID.expectedInvocations, n)
	return mmGetUserByProviderID
}

func (mmGetUserByProviderID *mUserCreatorMockGetUserByProviderID) invocationsDone() bool {
	if len(mmGetUserByProviderID.expectations) == 0 && mmGetUserByProviderID.defaultExpectation == nil && mmGetUserByProviderID.mock.funcGetUserByProviderID == nil {
		return true
	}

	totalInvocations := mm_atomic.LoadUint64(&mmGetUserByProviderID.mock.afterGetUserByProviderIDCounter)
	expectedInvocations := mm_atomic.LoadUint64(&mmGetUserByProviderID.expectedInvocations)

	return totalInvocations > 0 && (expectedInvocations == 0 || expectedInvocations == totalInvocations)
}

// GetUserByProviderID implements user.userCreator
func (mmGetUserByProviderID *UserCreatorMock) GetUserByProviderID(ctx context.Context, providerAccountID string) (u1 mm_user.User, err error) {
	mm_atomic.AddUint64(&mmGetUserByProviderID.beforeGetUserByProviderIDCounter, 1)
	defer mm_atomic.AddUint64(&mmGetUserByProviderID.afterGetUserByProviderIDCounter, 1)

	if mmGetUserByProviderID.inspectFuncGetUserByProviderID != nil {
		mmGetUserByProviderID.inspectFuncGetUserByProviderID(ctx, providerAccountID)
	}

	mm_params := UserCreatorMockGetUserByProviderIDParams{ctx, providerAccountID}

	// Record call args
	mmGetUserByProviderID.GetUserByProviderIDMock.mutex.Lock()
	mmGetUserByProviderID.GetUserByProviderIDMock.callArgs = append(mmGetUserByProviderID.GetUserByProviderIDMock.callArgs, &mm_params)
	mmGetUserByProviderID.GetUserByProviderIDMock.mutex.Unlock()

	for _, e := range mmGetUserByProviderID.GetUserByProviderIDMock.expectations {
		if minimock.Equal(*e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.u1, e.results.err
		}
	}

	if mmGetUserByProviderID.GetUserByProviderIDMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmGetUserByProviderID.GetUserByProviderIDMock.defaultExpectation.Counter, 1)
		mm_want := mmGetUserByProviderID.GetUserByProviderIDMock.defaultExpectation.params
		mm_want_ptrs := mmGetUserByProviderID.GetUserByProviderIDMock.defaultExpectation.paramPtrs

		mm_got := UserCreatorMockGetUserByProviderIDParams{ctx, providerAccountID}

		if mm_want_ptrs != nil {

			if mm_want_ptrs.ctx != nil && !minimock.Equal(*mm_want_ptrs.ctx, mm_got.ctx) {
				mmGetUserByProviderID.t.Errorf("UserCreatorMock.GetUserByProviderID got unexpected parameter ctx, want: %#v, got: %#v%s\n", *mm_want_ptrs.ctx, mm_got.ctx, minimock.Diff(*mm_want_ptrs.ctx, mm_got.ctx))
			}

			if mm_want_ptrs.providerAccountID != nil && !minimock.Equal(*mm_want_ptrs.providerAccountID, mm_got.providerAccountID) {
				mmGetUserByProviderID.t.Errorf("UserCreatorMock.GetUserByProviderID got unexpected parameter providerAccountID, want: %#v, got: %#v%s\n", *mm_want_ptrs.providerAccountID, mm_got.providerAccountID, minimock.Diff(*mm_want_ptrs.providerAccountID, mm_got.providerAccountID))
			}

		} else if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmGetUserByProviderID.t.Errorf("UserCreatorMock.GetUserByProviderID got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmGetUserByProviderID.GetUserByProviderIDMock.defaultExpectation.results
		if mm_results == nil {
			mmGetUserByProviderID.t.Fatal("No results are set for the UserCreatorMock.GetUserByProviderID")
		}
		return (*mm_results).u1, (*mm_results).err
	}
	if mmGetUserByProviderID.funcGetUserByProviderID != nil {
		return mmGetUserByProviderID.funcGetUserByProviderID(ctx, providerAccountID)
	}
	mmGetUserByProviderID.t.Fatalf("Unexpected call to UserCreatorMock.GetUserByProviderID. %v %v", ctx, providerAccountID)
	return
}

// GetUserByProviderIDAfterCounter returns a count of finished UserCreatorMock.GetUserByProviderID invocations
func (mmGetUserByProviderID *UserCreatorMock) GetUserByProviderIDAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmGetUserByProviderID.afterGetUserByProviderIDCounter)
}

// GetUserByProviderIDBeforeCounter returns a count of UserCreatorMock.GetUserByProviderID invocations
func (mmGetUserByProviderID *UserCreatorMock) GetUserByProviderIDBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmGetUserByProviderID.beforeGetUserByProviderIDCounter)
}

// Calls returns a list of arguments used in each call to UserCreatorMock.GetUserByProviderID.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmGetUserByProviderID *mUserCreatorMockGetUserByProviderID) Calls() []*UserCreatorMockGetUserByProviderIDParams {
	mmGetUserByProviderID.mutex.RLock()

	argCopy := make([]*UserCreatorMockGetUserByProviderIDParams, len(mmGetUserByProviderID.callArgs))
	copy(argCopy, mmGetUserByProviderID.callArgs)

	mmGetUserByProviderID.mutex.RUnlock()

	return argCopy
}

// MinimockGetUserByProviderIDDone returns true if the count of the GetUserByProviderID invocations corresponds
// the number of defined expectations
func (m *UserCreatorMock) MinimockGetUserByProviderIDDone() bool {
	if m.GetUserByProviderIDMock.optional {
		// Optional methods provide '0 or more' call count restriction.
		return true
	}

	for _, e := range m.GetUserByProviderIDMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	return m.GetUserByProviderIDMock.invocationsDone()
}

// MinimockGetUserByProviderIDInspect logs each unmet expectation
func (m *UserCreatorMock) MinimockGetUserByProviderIDInspect() {
	for _, e := range m.GetUserByProviderIDMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to UserCreatorMock.GetUserByProviderID with params: %#v", *e.params)
		}
	}

	afterGetUserByProviderIDCounter := mm_atomic.LoadUint64(&m.afterGetUserByProviderIDCounter)
	// if default expectation was set then invocations count should be greater than zero
	if m.GetUserByProviderIDMock.defaultExpectation != nil && afterGetUserByProviderIDCounter < 1 {
		if m.GetUserByProviderIDMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to UserCreatorMock.GetUserByProviderID")
		} else {
			m.t.Errorf("Expected call to UserCreatorMock.GetUserByProviderID with params: %#v", *m.GetUserByProviderIDMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcGetUserByProviderID != nil && afterGetUserByProviderIDCounter < 1 {
		m.t.Error("Expected call to UserCreatorMock.GetUserByProviderID")
	}

	if !m.GetUserByProviderIDMock.invocationsDone() && afterGetUserByProviderIDCounter > 0 {
		m.t.Errorf("Expected %d calls to UserCreatorMock.GetUserByProviderID but found %d calls",
			mm_atomic.LoadUint64(&m.GetUserByProviderIDMock.expectedInvocations), afterGetUserByProviderIDCounter)
	}
}

type mUserCreatorMockInsertAndLinkAccount struct {
	optional           bool
	mock               *UserCreatorMock
	defaultExpectation *UserCreatorMockInsertAndLinkAccountExpectation
	expectations       []*UserCreatorMockInsertAndLinkAccountExpectation

	callArgs []*UserCreatorMockInsertAndLinkAccountParams
	mutex    sync.RWMutex

	expectedInvocations uint64
}

// UserCreatorMockInsertAndLinkAccountExpectation specifies expectation struct of the userCreator.InsertAndLinkAccount
type UserCreatorMockInsertAndLinkAccountExpectation struct {
	mock      *UserCreatorMock
	params    *UserCreatorMockInsertAndLinkAccountParams
	paramPtrs *UserCreatorMockInsertAndLinkAccountParamPtrs
	results   *UserCreatorMockInsertAndLinkAccountResults
	Counter   uint64
}

// UserCreatorMockInsertAndLinkAccountParams contains parameters of the userCreator.InsertAndLinkAccount
type UserCreatorMockInsertAndLinkAccountParams struct {
	ctx     context.Context
	user    *mm_user.User
	account *mm_user.Account
}

// UserCreatorMockInsertAndLinkAccountParamPtrs contains pointers to parameters of the userCreator.InsertAndLinkAccount
type UserCreatorMockInsertAndLinkAccountParamPtrs struct {
	ctx     *context.Context
	user    **mm_user.User
	account **mm_user.Account
}

// UserCreatorMockInsertAndLinkAccountResults contains results of the userCreator.InsertAndLinkAccount
type UserCreatorMockInsertAndLinkAccountResults struct {
	err error
}

// Marks this method to be optional. The default behavior of any method with Return() is '1 or more', meaning
// the test will fail minimock's automatic final call check if the mocked method was not called at least once.
// Optional() makes method check to work in '0 or more' mode.
// It is NOT RECOMMENDED to use this option unless you really need it, as default behaviour helps to
// catch the problems when the expected method call is totally skipped during test run.
func (mmInsertAndLinkAccount *mUserCreatorMockInsertAndLinkAccount) Optional() *mUserCreatorMockInsertAndLinkAccount {
	mmInsertAndLinkAccount.optional = true
	return mmInsertAndLinkAccount
}

// Expect sets up expected params for userCreator.InsertAndLinkAccount
func (mmInsertAndLinkAccount *mUserCreatorMockInsertAndLinkAccount) Expect(ctx context.Context, user *mm_user.User, account *mm_user.Account) *mUserCreatorMockInsertAndLinkAccount {
	if mmInsertAndLinkAccount.mock.funcInsertAndLinkAccount != nil {
		mmInsertAndLinkAccount.mock.t.Fatalf("UserCreatorMock.InsertAndLinkAccount mock is already set by Set")
	}

	if mmInsertAndLinkAccount.defaultExpectation == nil {
		mmInsertAndLinkAccount.defaultExpectation = &UserCreatorMockInsertAndLinkAccountExpectation{}
	}

	if mmInsertAndLinkAccount.defaultExpectation.paramPtrs != nil {
		mmInsertAndLinkAccount.mock.t.Fatalf("UserCreatorMock.InsertAndLinkAccount mock is already set by ExpectParams functions")
	}

	mmInsertAndLinkAccount.defaultExpectation.params = &UserCreatorMockInsertAndLinkAccountParams{ctx, user, account}
	for _, e := range mmInsertAndLinkAccount.expectations {
		if minimock.Equal(e.params, mmInsertAndLinkAccount.defaultExpectation.params) {
			mmInsertAndLinkAccount.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmInsertAndLinkAccount.defaultExpectation.params)
		}
	}

	return mmInsertAndLinkAccount
}

// ExpectCtxParam1 sets up expected param ctx for userCreator.InsertAndLinkAccount
func (mmInsertAndLinkAccount *mUserCreatorMockInsertAndLinkAccount) ExpectCtxParam1(ctx context.Context) *mUserCreatorMockInsertAndLinkAccount {
	if mmInsertAndLinkAccount.mock.funcInsertAndLinkAccount != nil {
		mmInsertAndLinkAccount.mock.t.Fatalf("UserCreatorMock.InsertAndLinkAccount mock is already set by Set")
	}

	if mmInsertAndLinkAccount.defaultExpectation == nil {
		mmInsertAndLinkAccount.defaultExpectation = &UserCreatorMockInsertAndLinkAccountExpectation{}
	}

	if mmInsertAndLinkAccount.defaultExpectation.params != nil {
		mmInsertAndLinkAccount.mock.t.Fatalf("UserCreatorMock.InsertAndLinkAccount mock is already set by Expect")
	}

	if mmInsertAndLinkAccount.defaultExpectation.paramPtrs == nil {
		mmInsertAndLinkAccount.defaultExpectation.paramPtrs = &UserCreatorMockInsertAndLinkAccountParamPtrs{}
	}
	mmInsertAndLinkAccount.defaultExpectation.paramPtrs.ctx = &ctx

	return mmInsertAndLinkAccount
}

// ExpectUserParam2 sets up expected param user for userCreator.InsertAndLinkAccount
func (mmInsertAndLinkAccount *mUserCreatorMockInsertAndLinkAccount) ExpectUserParam2(user *mm_user.User) *mUserCreatorMockInsertAndLinkAccount {
	if mmInsertAndLinkAccount.mock.funcInsertAndLinkAccount != nil {
		mmInsertAndLinkAccount.mock.t.Fatalf("UserCreatorMock.InsertAndLinkAccount mock is already set by Set")
	}

	if mmInsertAndLinkAccount.defaultExpectation == nil {
		mmInsertAndLinkAccount.defaultExpectation = &UserCreatorMockInsertAndLinkAccountExpectation{}
	}

	if mmInsertAndLinkAccount.defaultExpectation.params != nil {
		mmInsertAndLinkAccount.mock.t.Fatalf("UserCreatorMock.InsertAndLinkAccount mock is already set by Expect")
	}

	if mmInsertAndLinkAccount.defaultExpectation.paramPtrs == nil {
		mmInsertAndLinkAccount.defaultExpectation.paramPtrs = &UserCreatorMockInsertAndLinkAccountParamPtrs{}
	}
	mmInsertAndLinkAccount.defaultExpectation.paramPtrs.user = &user

	return mmInsertAndLinkAccount
}

// ExpectAccountParam3 sets up expected param account for userCreator.InsertAndLinkAccount
func (mmInsertAndLinkAccount *mUserCreatorMockInsertAndLinkAccount) ExpectAccountParam3(account *mm_user.Account) *mUserCreatorMockInsertAndLinkAccount {
	if mmInsertAndLinkAccount.mock.funcInsertAndLinkAccount != nil {
		mmInsertAndLinkAccount.mock.t.Fatalf("UserCreatorMock.InsertAndLinkAccount mock is already set by Set")
	}

	if mmInsertAndLinkAccount.defaultExpectation == nil {
		mmInsertAndLinkAccount.defaultExpectation = &UserCreatorMockInsertAndLinkAccountExpectation{}
	}

	if mmInsertAndLinkAccount.defaultExpectation.params != nil {
		mmInsertAndLinkAccount.mock.t.Fatalf("UserCreatorMock.InsertAndLinkAccount mock is already set by Expect")
	}

	if mmInsertAndLinkAccount.defaultExpectation.paramPtrs == nil {
		mmInsertAndLinkAccount.defaultExpectation.paramPtrs = &UserCreatorMockInsertAndLinkAccountParamPtrs{}
	}
	mmInsertAndLinkAccount.defaultExpectation.paramPtrs.account = &account

	return mmInsertAndLinkAccount
}

// Inspect accepts an inspector function that has same arguments as the userCreator.InsertAndLinkAccount
func (mmInsertAndLinkAccount *mUserCreatorMockInsertAndLinkAccount) Inspect(f func(ctx context.Context, user *mm_user.User, account *mm_user.Account)) *mUserCreatorMockInsertAndLinkAccount {
	if mmInsertAndLinkAccount.mock.inspectFuncInsertAndLinkAccount != nil {
		mmInsertAndLinkAccount.mock.t.Fatalf("Inspect function is already set for UserCreatorMock.InsertAndLinkAccount")
	}

	mmInsertAndLinkAccount.mock.inspectFuncInsertAndLinkAccount = f

	return mmInsertAndLinkAccount
}

// Return sets up results that will be returned by userCreator.InsertAndLinkAccount
func (mmInsertAndLinkAccount *mUserCreatorMockInsertAndLinkAccount) Return(err error) *UserCreatorMock {
	if mmInsertAndLinkAccount.mock.funcInsertAndLinkAccount != nil {
		mmInsertAndLinkAccount.mock.t.Fatalf("UserCreatorMock.InsertAndLinkAccount mock is already set by Set")
	}

	if mmInsertAndLinkAccount.defaultExpectation == nil {
		mmInsertAndLinkAccount.defaultExpectation = &UserCreatorMockInsertAndLinkAccountExpectation{mock: mmInsertAndLinkAccount.mock}
	}
	mmInsertAndLinkAccount.defaultExpectation.results = &UserCreatorMockInsertAndLinkAccountResults{err}
	return mmInsertAndLinkAccount.mock
}

// Set uses given function f to mock the userCreator.InsertAndLinkAccount method
func (mmInsertAndLinkAccount *mUserCreatorMockInsertAndLinkAccount) Set(f func(ctx context.Context, user *mm_user.User, account *mm_user.Account) (err error)) *UserCreatorMock {
	if mmInsertAndLinkAccount.defaultExpectation != nil {
		mmInsertAndLinkAccount.mock.t.Fatalf("Default expectation is already set for the userCreator.InsertAndLinkAccount method")
	}

	if len(mmInsertAndLinkAccount.expectations) > 0 {
		mmInsertAndLinkAccount.mock.t.Fatalf("Some expectations are already set for the userCreator.InsertAndLinkAccount method")
	}

	mmInsertAndLinkAccount.mock.funcInsertAndLinkAccount = f
	return mmInsertAndLinkAccount.mock
}

// When sets expectation for the userCreator.InsertAndLinkAccount which will trigger the result defined by the following
// Then helper
func (mmInsertAndLinkAccount *mUserCreatorMockInsertAndLinkAccount) When(ctx context.Context, user *mm_user.User, account *mm_user.Account) *UserCreatorMockInsertAndLinkAccountExpectation {
	if mmInsertAndLinkAccount.mock.funcInsertAndLinkAccount != nil {
		mmInsertAndLinkAccount.mock.t.Fatalf("UserCreatorMock.InsertAndLinkAccount mock is already set by Set")
	}

	expectation := &UserCreatorMockInsertAndLinkAccountExpectation{
		mock:   mmInsertAndLinkAccount.mock,
		params: &UserCreatorMockInsertAndLinkAccountParams{ctx, user, account},
	}
	mmInsertAndLinkAccount.expectations = append(mmInsertAndLinkAccount.expectations, expectation)
	return expectation
}

// Then sets up userCreator.InsertAndLinkAccount return parameters for the expectation previously defined by the When method
func (e *UserCreatorMockInsertAndLinkAccountExpectation) Then(err error) *UserCreatorMock {
	e.results = &UserCreatorMockInsertAndLinkAccountResults{err}
	return e.mock
}

// Times sets number of times userCreator.InsertAndLinkAccount should be invoked
func (mmInsertAndLinkAccount *mUserCreatorMockInsertAndLinkAccount) Times(n uint64) *mUserCreatorMockInsertAndLinkAccount {
	if n == 0 {
		mmInsertAndLinkAccount.mock.t.Fatalf("Times of UserCreatorMock.InsertAndLinkAccount mock can not be zero")
	}
	mm_atomic.StoreUint64(&mmInsertAndLinkAccount.expectedInvocations, n)
	return mmInsertAndLinkAccount
}

func (mmInsertAndLinkAccount *mUserCreatorMockInsertAndLinkAccount) invocationsDone() bool {
	if len(mmInsertAndLinkAccount.expectations) == 0 && mmInsertAndLinkAccount.defaultExpectation == nil && mmInsertAndLinkAccount.mock.funcInsertAndLinkAccount == nil {
		return true
	}

	totalInvocations := mm_atomic.LoadUint64(&mmInsertAndLinkAccount.mock.afterInsertAndLinkAccountCounter)
	expectedInvocations := mm_atomic.LoadUint64(&mmInsertAndLinkAccount.expectedInvocations)

	return totalInvocations > 0 && (expectedInvocations == 0 || expectedInvocations == totalInvocations)
}

// InsertAndLinkAccount implements user.userCreator
func (mmInsertAndLinkAccount *UserCreatorMock) InsertAndLinkAccount(ctx context.Context, user *mm_user.User, account *mm_user.Account) (err error) {
	mm_atomic.AddUint64(&mmInsertAndLinkAccount.beforeInsertAndLinkAccountCounter, 1)
	defer mm_atomic.AddUint64(&mmInsertAndLinkAccount.afterInsertAndLinkAccountCounter, 1)

	if mmInsertAndLinkAccount.inspectFuncInsertAndLinkAccount != nil {
		mmInsertAndLinkAccount.inspectFuncInsertAndLinkAccount(ctx, user, account)
	}

	mm_params := UserCreatorMockInsertAndLinkAccountParams{ctx, user, account}

	// Record call args
	mmInsertAndLinkAccount.InsertAndLinkAccountMock.mutex.Lock()
	mmInsertAndLinkAccount.InsertAndLinkAccountMock.callArgs = append(mmInsertAndLinkAccount.InsertAndLinkAccountMock.callArgs, &mm_params)
	mmInsertAndLinkAccount.InsertAndLinkAccountMock.mutex.Unlock()

	for _, e := range mmInsertAndLinkAccount.InsertAndLinkAccountMock.expectations {
		if minimock.Equal(*e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.err
		}
	}

	if mmInsertAndLinkAccount.InsertAndLinkAccountMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmInsertAndLinkAccount.InsertAndLinkAccountMock.defaultExpectation.Counter, 1)
		mm_want := mmInsertAndLinkAccount.InsertAndLinkAccountMock.defaultExpectation.params
		mm_want_ptrs := mmInsertAndLinkAccount.InsertAndLinkAccountMock.defaultExpectation.paramPtrs

		mm_got := UserCreatorMockInsertAndLinkAccountParams{ctx, user, account}

		if mm_want_ptrs != nil {

			if mm_want_ptrs.ctx != nil && !minimock.Equal(*mm_want_ptrs.ctx, mm_got.ctx) {
				mmInsertAndLinkAccount.t.Errorf("UserCreatorMock.InsertAndLinkAccount got unexpected parameter ctx, want: %#v, got: %#v%s\n", *mm_want_ptrs.ctx, mm_got.ctx, minimock.Diff(*mm_want_ptrs.ctx, mm_got.ctx))
			}

			if mm_want_ptrs.user != nil && !minimock.Equal(*mm_want_ptrs.user, mm_got.user) {
				mmInsertAndLinkAccount.t.Errorf("UserCreatorMock.InsertAndLinkAccount got unexpected parameter user, want: %#v, got: %#v%s\n", *mm_want_ptrs.user, mm_got.user, minimock.Diff(*mm_want_ptrs.user, mm_got.user))
			}

			if mm_want_ptrs.account != nil && !minimock.Equal(*mm_want_ptrs.account, mm_got.account) {
				mmInsertAndLinkAccount.t.Errorf("UserCreatorMock.InsertAndLinkAccount got unexpected parameter account, want: %#v, got: %#v%s\n", *mm_want_ptrs.account, mm_got.account, minimock.Diff(*mm_want_ptrs.account, mm_got.account))
			}

		} else if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmInsertAndLinkAccount.t.Errorf("UserCreatorMock.InsertAndLinkAccount got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmInsertAndLinkAccount.InsertAndLinkAccountMock.defaultExpectation.results
		if mm_results == nil {
			mmInsertAndLinkAccount.t.Fatal("No results are set for the UserCreatorMock.InsertAndLinkAccount")
		}
		return (*mm_results).err
	}
	if mmInsertAndLinkAccount.funcInsertAndLinkAccount != nil {
		return mmInsertAndLinkAccount.funcInsertAndLinkAccount(ctx, user, account)
	}
	mmInsertAndLinkAccount.t.Fatalf("Unexpected call to UserCreatorMock.InsertAndLinkAccount. %v %v %v", ctx, user, account)
	return
}

// InsertAndLinkAccountAfterCounter returns a count of finished UserCreatorMock.InsertAndLinkAccount invocations
func (mmInsertAndLinkAccount *UserCreatorMock) InsertAndLinkAccountAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmInsertAndLinkAccount.afterInsertAndLinkAccountCounter)
}

// InsertAndLinkAccountBeforeCounter returns a count of UserCreatorMock.InsertAndLinkAccount invocations
func (mmInsertAndLinkAccount *UserCreatorMock) InsertAndLinkAccountBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmInsertAndLinkAccount.beforeInsertAndLinkAccountCounter)
}

// Calls returns a list of arguments used in each call to UserCreatorMock.InsertAndLinkAccount.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmInsertAndLinkAccount *mUserCreatorMockInsertAndLinkAccount) Calls() []*UserCreatorMockInsertAndLinkAccountParams {
	mmInsertAndLinkAccount.mutex.RLock()

	argCopy := make([]*UserCreatorMockInsertAndLinkAccountParams, len(mmInsertAndLinkAccount.callArgs))
	copy(argCopy, mmInsertAndLinkAccount.callArgs)

	mmInsertAndLinkAccount.mutex.RUnlock()

	return argCopy
}

// MinimockInsertAndLinkAccountDone returns true if the count of the InsertAndLinkAccount invocations corresponds
// the number of defined expectations
func (m *UserCreatorMock) MinimockInsertAndLinkAccountDone() bool {
	if m.InsertAndLinkAccountMock.optional {
		// Optional methods provide '0 or more' call count restriction.
		return true
	}

	for _, e := range m.InsertAndLinkAccountMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	return m.InsertAndLinkAccountMock.invocationsDone()
}

// MinimockInsertAndLinkAccountInspect logs each unmet expectation
func (m *UserCreatorMock) MinimockInsertAndLinkAccountInspect() {
	for _, e := range m.InsertAndLinkAccountMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to UserCreatorMock.InsertAndLinkAccount with params: %#v", *e.params)
		}
	}

	afterInsertAndLinkAccountCounter := mm_atomic.LoadUint64(&m.afterInsertAndLinkAccountCounter)
	// if default expectation was set then invocations count should be greater than zero
	if m.InsertAndLinkAccountMock.defaultExpectation != nil && afterInsertAndLinkAccountCounter < 1 {
		if m.InsertAndLinkAccountMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to UserCreatorMock.InsertAndLinkAccount")
		} else {
			m.t.Errorf("Expected call to UserCreatorMock.InsertAndLinkAccount with params: %#v", *m.InsertAndLinkAccountMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcInsertAndLinkAccount != nil && afterInsertAndLinkAccountCounter < 1 {
		m.t.Error("Expected call to UserCreatorMock.InsertAndLinkAccount")
	}

	if !m.InsertAndLinkAccountMock.invocationsDone() && afterInsertAndLinkAccountCounter > 0 {
		m.t.Errorf("Expected %d calls to UserCreatorMock.InsertAndLinkAccount but found %d calls",
			mm_atomic.LoadUint64(&m.InsertAndLinkAccountMock.expectedInvocations), afterInsertAndLinkAccountCounter)
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *UserCreatorMock) MinimockFinish() {
	m.finishOnce.Do(func() {
		if !m.minimockDone() {
			m.MinimockGetUserByProviderIDInspect()

			m.MinimockInsertAndLinkAccountInspect()
		}
	})
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *UserCreatorMock) MinimockWait(timeout mm_time.Duration) {
	timeoutCh := mm_time.After(timeout)
	for {
		if m.minimockDone() {
			return
		}
		select {
		case <-timeoutCh:
			m.MinimockFinish()
			return
		case <-mm_time.After(10 * mm_time.Millisecond):
		}
	}
}

func (m *UserCreatorMock) minimockDone() bool {
	done := true
	return done &&
		m.MinimockGetUserByProviderIDDone() &&
		m.MinimockInsertAndLinkAccountDone()
}
