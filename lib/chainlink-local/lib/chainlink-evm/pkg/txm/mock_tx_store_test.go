// Code generated by mockery v2.53.3. DO NOT EDIT.

package txm

import (
	context "context"

	common "github.com/ethereum/go-ethereum/common"

	mock "github.com/stretchr/testify/mock"

	types "github.com/smartcontractkit/chainlink-evm/pkg/txm/types"
)

// mockTxStore is an autogenerated mock type for the TxStore type
type mockTxStore struct {
	mock.Mock
}

type mockTxStore_Expecter struct {
	mock *mock.Mock
}

func (_m *mockTxStore) EXPECT() *mockTxStore_Expecter {
	return &mockTxStore_Expecter{mock: &_m.Mock}
}

// AbandonPendingTransactions provides a mock function with given fields: _a0, _a1
func (_m *mockTxStore) AbandonPendingTransactions(_a0 context.Context, _a1 common.Address) error {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for AbandonPendingTransactions")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, common.Address) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// mockTxStore_AbandonPendingTransactions_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AbandonPendingTransactions'
type mockTxStore_AbandonPendingTransactions_Call struct {
	*mock.Call
}

// AbandonPendingTransactions is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 common.Address
func (_e *mockTxStore_Expecter) AbandonPendingTransactions(_a0 interface{}, _a1 interface{}) *mockTxStore_AbandonPendingTransactions_Call {
	return &mockTxStore_AbandonPendingTransactions_Call{Call: _e.mock.On("AbandonPendingTransactions", _a0, _a1)}
}

func (_c *mockTxStore_AbandonPendingTransactions_Call) Run(run func(_a0 context.Context, _a1 common.Address)) *mockTxStore_AbandonPendingTransactions_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(common.Address))
	})
	return _c
}

func (_c *mockTxStore_AbandonPendingTransactions_Call) Return(_a0 error) *mockTxStore_AbandonPendingTransactions_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *mockTxStore_AbandonPendingTransactions_Call) RunAndReturn(run func(context.Context, common.Address) error) *mockTxStore_AbandonPendingTransactions_Call {
	_c.Call.Return(run)
	return _c
}

// AppendAttemptToTransaction provides a mock function with given fields: _a0, _a1, _a2, _a3
func (_m *mockTxStore) AppendAttemptToTransaction(_a0 context.Context, _a1 uint64, _a2 common.Address, _a3 *types.Attempt) error {
	ret := _m.Called(_a0, _a1, _a2, _a3)

	if len(ret) == 0 {
		panic("no return value specified for AppendAttemptToTransaction")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uint64, common.Address, *types.Attempt) error); ok {
		r0 = rf(_a0, _a1, _a2, _a3)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// mockTxStore_AppendAttemptToTransaction_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AppendAttemptToTransaction'
type mockTxStore_AppendAttemptToTransaction_Call struct {
	*mock.Call
}

// AppendAttemptToTransaction is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 uint64
//   - _a2 common.Address
//   - _a3 *types.Attempt
func (_e *mockTxStore_Expecter) AppendAttemptToTransaction(_a0 interface{}, _a1 interface{}, _a2 interface{}, _a3 interface{}) *mockTxStore_AppendAttemptToTransaction_Call {
	return &mockTxStore_AppendAttemptToTransaction_Call{Call: _e.mock.On("AppendAttemptToTransaction", _a0, _a1, _a2, _a3)}
}

func (_c *mockTxStore_AppendAttemptToTransaction_Call) Run(run func(_a0 context.Context, _a1 uint64, _a2 common.Address, _a3 *types.Attempt)) *mockTxStore_AppendAttemptToTransaction_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uint64), args[2].(common.Address), args[3].(*types.Attempt))
	})
	return _c
}

func (_c *mockTxStore_AppendAttemptToTransaction_Call) Return(_a0 error) *mockTxStore_AppendAttemptToTransaction_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *mockTxStore_AppendAttemptToTransaction_Call) RunAndReturn(run func(context.Context, uint64, common.Address, *types.Attempt) error) *mockTxStore_AppendAttemptToTransaction_Call {
	_c.Call.Return(run)
	return _c
}

// CreateEmptyUnconfirmedTransaction provides a mock function with given fields: _a0, _a1, _a2, _a3
func (_m *mockTxStore) CreateEmptyUnconfirmedTransaction(_a0 context.Context, _a1 common.Address, _a2 uint64, _a3 uint64) (*types.Transaction, error) {
	ret := _m.Called(_a0, _a1, _a2, _a3)

	if len(ret) == 0 {
		panic("no return value specified for CreateEmptyUnconfirmedTransaction")
	}

	var r0 *types.Transaction
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, common.Address, uint64, uint64) (*types.Transaction, error)); ok {
		return rf(_a0, _a1, _a2, _a3)
	}
	if rf, ok := ret.Get(0).(func(context.Context, common.Address, uint64, uint64) *types.Transaction); ok {
		r0 = rf(_a0, _a1, _a2, _a3)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Transaction)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, common.Address, uint64, uint64) error); ok {
		r1 = rf(_a0, _a1, _a2, _a3)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// mockTxStore_CreateEmptyUnconfirmedTransaction_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateEmptyUnconfirmedTransaction'
type mockTxStore_CreateEmptyUnconfirmedTransaction_Call struct {
	*mock.Call
}

// CreateEmptyUnconfirmedTransaction is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 common.Address
//   - _a2 uint64
//   - _a3 uint64
func (_e *mockTxStore_Expecter) CreateEmptyUnconfirmedTransaction(_a0 interface{}, _a1 interface{}, _a2 interface{}, _a3 interface{}) *mockTxStore_CreateEmptyUnconfirmedTransaction_Call {
	return &mockTxStore_CreateEmptyUnconfirmedTransaction_Call{Call: _e.mock.On("CreateEmptyUnconfirmedTransaction", _a0, _a1, _a2, _a3)}
}

func (_c *mockTxStore_CreateEmptyUnconfirmedTransaction_Call) Run(run func(_a0 context.Context, _a1 common.Address, _a2 uint64, _a3 uint64)) *mockTxStore_CreateEmptyUnconfirmedTransaction_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(common.Address), args[2].(uint64), args[3].(uint64))
	})
	return _c
}

func (_c *mockTxStore_CreateEmptyUnconfirmedTransaction_Call) Return(_a0 *types.Transaction, _a1 error) *mockTxStore_CreateEmptyUnconfirmedTransaction_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *mockTxStore_CreateEmptyUnconfirmedTransaction_Call) RunAndReturn(run func(context.Context, common.Address, uint64, uint64) (*types.Transaction, error)) *mockTxStore_CreateEmptyUnconfirmedTransaction_Call {
	_c.Call.Return(run)
	return _c
}

// CreateTransaction provides a mock function with given fields: _a0, _a1
func (_m *mockTxStore) CreateTransaction(_a0 context.Context, _a1 *types.TxRequest) (*types.Transaction, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for CreateTransaction")
	}

	var r0 *types.Transaction
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *types.TxRequest) (*types.Transaction, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *types.TxRequest) *types.Transaction); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Transaction)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *types.TxRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// mockTxStore_CreateTransaction_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateTransaction'
type mockTxStore_CreateTransaction_Call struct {
	*mock.Call
}

// CreateTransaction is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 *types.TxRequest
func (_e *mockTxStore_Expecter) CreateTransaction(_a0 interface{}, _a1 interface{}) *mockTxStore_CreateTransaction_Call {
	return &mockTxStore_CreateTransaction_Call{Call: _e.mock.On("CreateTransaction", _a0, _a1)}
}

func (_c *mockTxStore_CreateTransaction_Call) Run(run func(_a0 context.Context, _a1 *types.TxRequest)) *mockTxStore_CreateTransaction_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*types.TxRequest))
	})
	return _c
}

func (_c *mockTxStore_CreateTransaction_Call) Return(_a0 *types.Transaction, _a1 error) *mockTxStore_CreateTransaction_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *mockTxStore_CreateTransaction_Call) RunAndReturn(run func(context.Context, *types.TxRequest) (*types.Transaction, error)) *mockTxStore_CreateTransaction_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteAttemptForUnconfirmedTx provides a mock function with given fields: _a0, _a1, _a2, _a3
func (_m *mockTxStore) DeleteAttemptForUnconfirmedTx(_a0 context.Context, _a1 uint64, _a2 *types.Attempt, _a3 common.Address) error {
	ret := _m.Called(_a0, _a1, _a2, _a3)

	if len(ret) == 0 {
		panic("no return value specified for DeleteAttemptForUnconfirmedTx")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uint64, *types.Attempt, common.Address) error); ok {
		r0 = rf(_a0, _a1, _a2, _a3)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// mockTxStore_DeleteAttemptForUnconfirmedTx_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteAttemptForUnconfirmedTx'
type mockTxStore_DeleteAttemptForUnconfirmedTx_Call struct {
	*mock.Call
}

// DeleteAttemptForUnconfirmedTx is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 uint64
//   - _a2 *types.Attempt
//   - _a3 common.Address
func (_e *mockTxStore_Expecter) DeleteAttemptForUnconfirmedTx(_a0 interface{}, _a1 interface{}, _a2 interface{}, _a3 interface{}) *mockTxStore_DeleteAttemptForUnconfirmedTx_Call {
	return &mockTxStore_DeleteAttemptForUnconfirmedTx_Call{Call: _e.mock.On("DeleteAttemptForUnconfirmedTx", _a0, _a1, _a2, _a3)}
}

func (_c *mockTxStore_DeleteAttemptForUnconfirmedTx_Call) Run(run func(_a0 context.Context, _a1 uint64, _a2 *types.Attempt, _a3 common.Address)) *mockTxStore_DeleteAttemptForUnconfirmedTx_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uint64), args[2].(*types.Attempt), args[3].(common.Address))
	})
	return _c
}

func (_c *mockTxStore_DeleteAttemptForUnconfirmedTx_Call) Return(_a0 error) *mockTxStore_DeleteAttemptForUnconfirmedTx_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *mockTxStore_DeleteAttemptForUnconfirmedTx_Call) RunAndReturn(run func(context.Context, uint64, *types.Attempt, common.Address) error) *mockTxStore_DeleteAttemptForUnconfirmedTx_Call {
	_c.Call.Return(run)
	return _c
}

// FetchUnconfirmedTransactionAtNonceWithCount provides a mock function with given fields: _a0, _a1, _a2
func (_m *mockTxStore) FetchUnconfirmedTransactionAtNonceWithCount(_a0 context.Context, _a1 uint64, _a2 common.Address) (*types.Transaction, int, error) {
	ret := _m.Called(_a0, _a1, _a2)

	if len(ret) == 0 {
		panic("no return value specified for FetchUnconfirmedTransactionAtNonceWithCount")
	}

	var r0 *types.Transaction
	var r1 int
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, uint64, common.Address) (*types.Transaction, int, error)); ok {
		return rf(_a0, _a1, _a2)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uint64, common.Address) *types.Transaction); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Transaction)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uint64, common.Address) int); ok {
		r1 = rf(_a0, _a1, _a2)
	} else {
		r1 = ret.Get(1).(int)
	}

	if rf, ok := ret.Get(2).(func(context.Context, uint64, common.Address) error); ok {
		r2 = rf(_a0, _a1, _a2)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// mockTxStore_FetchUnconfirmedTransactionAtNonceWithCount_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FetchUnconfirmedTransactionAtNonceWithCount'
type mockTxStore_FetchUnconfirmedTransactionAtNonceWithCount_Call struct {
	*mock.Call
}

// FetchUnconfirmedTransactionAtNonceWithCount is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 uint64
//   - _a2 common.Address
func (_e *mockTxStore_Expecter) FetchUnconfirmedTransactionAtNonceWithCount(_a0 interface{}, _a1 interface{}, _a2 interface{}) *mockTxStore_FetchUnconfirmedTransactionAtNonceWithCount_Call {
	return &mockTxStore_FetchUnconfirmedTransactionAtNonceWithCount_Call{Call: _e.mock.On("FetchUnconfirmedTransactionAtNonceWithCount", _a0, _a1, _a2)}
}

func (_c *mockTxStore_FetchUnconfirmedTransactionAtNonceWithCount_Call) Run(run func(_a0 context.Context, _a1 uint64, _a2 common.Address)) *mockTxStore_FetchUnconfirmedTransactionAtNonceWithCount_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uint64), args[2].(common.Address))
	})
	return _c
}

func (_c *mockTxStore_FetchUnconfirmedTransactionAtNonceWithCount_Call) Return(_a0 *types.Transaction, _a1 int, _a2 error) *mockTxStore_FetchUnconfirmedTransactionAtNonceWithCount_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *mockTxStore_FetchUnconfirmedTransactionAtNonceWithCount_Call) RunAndReturn(run func(context.Context, uint64, common.Address) (*types.Transaction, int, error)) *mockTxStore_FetchUnconfirmedTransactionAtNonceWithCount_Call {
	_c.Call.Return(run)
	return _c
}

// MarkConfirmedAndReorgedTransactions provides a mock function with given fields: _a0, _a1, _a2
func (_m *mockTxStore) MarkConfirmedAndReorgedTransactions(_a0 context.Context, _a1 uint64, _a2 common.Address) ([]*types.Transaction, []uint64, error) {
	ret := _m.Called(_a0, _a1, _a2)

	if len(ret) == 0 {
		panic("no return value specified for MarkConfirmedAndReorgedTransactions")
	}

	var r0 []*types.Transaction
	var r1 []uint64
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, uint64, common.Address) ([]*types.Transaction, []uint64, error)); ok {
		return rf(_a0, _a1, _a2)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uint64, common.Address) []*types.Transaction); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*types.Transaction)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uint64, common.Address) []uint64); ok {
		r1 = rf(_a0, _a1, _a2)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).([]uint64)
		}
	}

	if rf, ok := ret.Get(2).(func(context.Context, uint64, common.Address) error); ok {
		r2 = rf(_a0, _a1, _a2)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// mockTxStore_MarkConfirmedAndReorgedTransactions_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'MarkConfirmedAndReorgedTransactions'
type mockTxStore_MarkConfirmedAndReorgedTransactions_Call struct {
	*mock.Call
}

// MarkConfirmedAndReorgedTransactions is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 uint64
//   - _a2 common.Address
func (_e *mockTxStore_Expecter) MarkConfirmedAndReorgedTransactions(_a0 interface{}, _a1 interface{}, _a2 interface{}) *mockTxStore_MarkConfirmedAndReorgedTransactions_Call {
	return &mockTxStore_MarkConfirmedAndReorgedTransactions_Call{Call: _e.mock.On("MarkConfirmedAndReorgedTransactions", _a0, _a1, _a2)}
}

func (_c *mockTxStore_MarkConfirmedAndReorgedTransactions_Call) Run(run func(_a0 context.Context, _a1 uint64, _a2 common.Address)) *mockTxStore_MarkConfirmedAndReorgedTransactions_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uint64), args[2].(common.Address))
	})
	return _c
}

func (_c *mockTxStore_MarkConfirmedAndReorgedTransactions_Call) Return(_a0 []*types.Transaction, _a1 []uint64, _a2 error) *mockTxStore_MarkConfirmedAndReorgedTransactions_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *mockTxStore_MarkConfirmedAndReorgedTransactions_Call) RunAndReturn(run func(context.Context, uint64, common.Address) ([]*types.Transaction, []uint64, error)) *mockTxStore_MarkConfirmedAndReorgedTransactions_Call {
	_c.Call.Return(run)
	return _c
}

// MarkTxFatal provides a mock function with given fields: _a0, _a1, _a2
func (_m *mockTxStore) MarkTxFatal(_a0 context.Context, _a1 *types.Transaction, _a2 common.Address) error {
	ret := _m.Called(_a0, _a1, _a2)

	if len(ret) == 0 {
		panic("no return value specified for MarkTxFatal")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *types.Transaction, common.Address) error); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// mockTxStore_MarkTxFatal_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'MarkTxFatal'
type mockTxStore_MarkTxFatal_Call struct {
	*mock.Call
}

// MarkTxFatal is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 *types.Transaction
//   - _a2 common.Address
func (_e *mockTxStore_Expecter) MarkTxFatal(_a0 interface{}, _a1 interface{}, _a2 interface{}) *mockTxStore_MarkTxFatal_Call {
	return &mockTxStore_MarkTxFatal_Call{Call: _e.mock.On("MarkTxFatal", _a0, _a1, _a2)}
}

func (_c *mockTxStore_MarkTxFatal_Call) Run(run func(_a0 context.Context, _a1 *types.Transaction, _a2 common.Address)) *mockTxStore_MarkTxFatal_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*types.Transaction), args[2].(common.Address))
	})
	return _c
}

func (_c *mockTxStore_MarkTxFatal_Call) Return(_a0 error) *mockTxStore_MarkTxFatal_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *mockTxStore_MarkTxFatal_Call) RunAndReturn(run func(context.Context, *types.Transaction, common.Address) error) *mockTxStore_MarkTxFatal_Call {
	_c.Call.Return(run)
	return _c
}

// MarkUnconfirmedTransactionPurgeable provides a mock function with given fields: _a0, _a1, _a2
func (_m *mockTxStore) MarkUnconfirmedTransactionPurgeable(_a0 context.Context, _a1 uint64, _a2 common.Address) error {
	ret := _m.Called(_a0, _a1, _a2)

	if len(ret) == 0 {
		panic("no return value specified for MarkUnconfirmedTransactionPurgeable")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uint64, common.Address) error); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// mockTxStore_MarkUnconfirmedTransactionPurgeable_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'MarkUnconfirmedTransactionPurgeable'
type mockTxStore_MarkUnconfirmedTransactionPurgeable_Call struct {
	*mock.Call
}

// MarkUnconfirmedTransactionPurgeable is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 uint64
//   - _a2 common.Address
func (_e *mockTxStore_Expecter) MarkUnconfirmedTransactionPurgeable(_a0 interface{}, _a1 interface{}, _a2 interface{}) *mockTxStore_MarkUnconfirmedTransactionPurgeable_Call {
	return &mockTxStore_MarkUnconfirmedTransactionPurgeable_Call{Call: _e.mock.On("MarkUnconfirmedTransactionPurgeable", _a0, _a1, _a2)}
}

func (_c *mockTxStore_MarkUnconfirmedTransactionPurgeable_Call) Run(run func(_a0 context.Context, _a1 uint64, _a2 common.Address)) *mockTxStore_MarkUnconfirmedTransactionPurgeable_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uint64), args[2].(common.Address))
	})
	return _c
}

func (_c *mockTxStore_MarkUnconfirmedTransactionPurgeable_Call) Return(_a0 error) *mockTxStore_MarkUnconfirmedTransactionPurgeable_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *mockTxStore_MarkUnconfirmedTransactionPurgeable_Call) RunAndReturn(run func(context.Context, uint64, common.Address) error) *mockTxStore_MarkUnconfirmedTransactionPurgeable_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateTransactionBroadcast provides a mock function with given fields: _a0, _a1, _a2, _a3, _a4
func (_m *mockTxStore) UpdateTransactionBroadcast(_a0 context.Context, _a1 uint64, _a2 uint64, _a3 common.Hash, _a4 common.Address) error {
	ret := _m.Called(_a0, _a1, _a2, _a3, _a4)

	if len(ret) == 0 {
		panic("no return value specified for UpdateTransactionBroadcast")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uint64, uint64, common.Hash, common.Address) error); ok {
		r0 = rf(_a0, _a1, _a2, _a3, _a4)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// mockTxStore_UpdateTransactionBroadcast_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateTransactionBroadcast'
type mockTxStore_UpdateTransactionBroadcast_Call struct {
	*mock.Call
}

// UpdateTransactionBroadcast is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 uint64
//   - _a2 uint64
//   - _a3 common.Hash
//   - _a4 common.Address
func (_e *mockTxStore_Expecter) UpdateTransactionBroadcast(_a0 interface{}, _a1 interface{}, _a2 interface{}, _a3 interface{}, _a4 interface{}) *mockTxStore_UpdateTransactionBroadcast_Call {
	return &mockTxStore_UpdateTransactionBroadcast_Call{Call: _e.mock.On("UpdateTransactionBroadcast", _a0, _a1, _a2, _a3, _a4)}
}

func (_c *mockTxStore_UpdateTransactionBroadcast_Call) Run(run func(_a0 context.Context, _a1 uint64, _a2 uint64, _a3 common.Hash, _a4 common.Address)) *mockTxStore_UpdateTransactionBroadcast_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uint64), args[2].(uint64), args[3].(common.Hash), args[4].(common.Address))
	})
	return _c
}

func (_c *mockTxStore_UpdateTransactionBroadcast_Call) Return(_a0 error) *mockTxStore_UpdateTransactionBroadcast_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *mockTxStore_UpdateTransactionBroadcast_Call) RunAndReturn(run func(context.Context, uint64, uint64, common.Hash, common.Address) error) *mockTxStore_UpdateTransactionBroadcast_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateUnstartedTransactionWithNonce provides a mock function with given fields: _a0, _a1, _a2
func (_m *mockTxStore) UpdateUnstartedTransactionWithNonce(_a0 context.Context, _a1 common.Address, _a2 uint64) (*types.Transaction, error) {
	ret := _m.Called(_a0, _a1, _a2)

	if len(ret) == 0 {
		panic("no return value specified for UpdateUnstartedTransactionWithNonce")
	}

	var r0 *types.Transaction
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, common.Address, uint64) (*types.Transaction, error)); ok {
		return rf(_a0, _a1, _a2)
	}
	if rf, ok := ret.Get(0).(func(context.Context, common.Address, uint64) *types.Transaction); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Transaction)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, common.Address, uint64) error); ok {
		r1 = rf(_a0, _a1, _a2)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// mockTxStore_UpdateUnstartedTransactionWithNonce_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateUnstartedTransactionWithNonce'
type mockTxStore_UpdateUnstartedTransactionWithNonce_Call struct {
	*mock.Call
}

// UpdateUnstartedTransactionWithNonce is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 common.Address
//   - _a2 uint64
func (_e *mockTxStore_Expecter) UpdateUnstartedTransactionWithNonce(_a0 interface{}, _a1 interface{}, _a2 interface{}) *mockTxStore_UpdateUnstartedTransactionWithNonce_Call {
	return &mockTxStore_UpdateUnstartedTransactionWithNonce_Call{Call: _e.mock.On("UpdateUnstartedTransactionWithNonce", _a0, _a1, _a2)}
}

func (_c *mockTxStore_UpdateUnstartedTransactionWithNonce_Call) Run(run func(_a0 context.Context, _a1 common.Address, _a2 uint64)) *mockTxStore_UpdateUnstartedTransactionWithNonce_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(common.Address), args[2].(uint64))
	})
	return _c
}

func (_c *mockTxStore_UpdateUnstartedTransactionWithNonce_Call) Return(_a0 *types.Transaction, _a1 error) *mockTxStore_UpdateUnstartedTransactionWithNonce_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *mockTxStore_UpdateUnstartedTransactionWithNonce_Call) RunAndReturn(run func(context.Context, common.Address, uint64) (*types.Transaction, error)) *mockTxStore_UpdateUnstartedTransactionWithNonce_Call {
	_c.Call.Return(run)
	return _c
}

// newMockTxStore creates a new instance of mockTxStore. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func newMockTxStore(t interface {
	mock.TestingT
	Cleanup(func())
}) *mockTxStore {
	mock := &mockTxStore{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
