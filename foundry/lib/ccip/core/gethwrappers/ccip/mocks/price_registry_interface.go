// Code generated by mockery v2.28.1. DO NOT EDIT.

package mock_contracts

import (
	big "math/big"

	bind "github.com/ethereum/go-ethereum/accounts/abi/bind"
	common "github.com/ethereum/go-ethereum/common"

	event "github.com/ethereum/go-ethereum/event"

	generated "github.com/smartcontractkit/chainlink/v2/core/gethwrappers/generated"

	mock "github.com/stretchr/testify/mock"

	price_registry "github.com/smartcontractkit/chainlink/v2/core/gethwrappers/ccip/generated/price_registry"

	types "github.com/ethereum/go-ethereum/core/types"
)

// PriceRegistryInterface is an autogenerated mock type for the PriceRegistryInterface type
type PriceRegistryInterface struct {
	mock.Mock
}

// AcceptOwnership provides a mock function with given fields: opts
func (_m *PriceRegistryInterface) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	ret := _m.Called(opts)

	var r0 *types.Transaction
	var r1 error
	if rf, ok := ret.Get(0).(func(*bind.TransactOpts) (*types.Transaction, error)); ok {
		return rf(opts)
	}
	if rf, ok := ret.Get(0).(func(*bind.TransactOpts) *types.Transaction); ok {
		r0 = rf(opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Transaction)
		}
	}

	if rf, ok := ret.Get(1).(func(*bind.TransactOpts) error); ok {
		r1 = rf(opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Address provides a mock function with given fields:
func (_m *PriceRegistryInterface) Address() common.Address {
	ret := _m.Called()

	var r0 common.Address
	if rf, ok := ret.Get(0).(func() common.Address); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(common.Address)
		}
	}

	return r0
}

// ApplyFeeTokensUpdates provides a mock function with given fields: opts, feeTokensToAdd, feeTokensToRemove
func (_m *PriceRegistryInterface) ApplyFeeTokensUpdates(opts *bind.TransactOpts, feeTokensToAdd []common.Address, feeTokensToRemove []common.Address) (*types.Transaction, error) {
	ret := _m.Called(opts, feeTokensToAdd, feeTokensToRemove)

	var r0 *types.Transaction
	var r1 error
	if rf, ok := ret.Get(0).(func(*bind.TransactOpts, []common.Address, []common.Address) (*types.Transaction, error)); ok {
		return rf(opts, feeTokensToAdd, feeTokensToRemove)
	}
	if rf, ok := ret.Get(0).(func(*bind.TransactOpts, []common.Address, []common.Address) *types.Transaction); ok {
		r0 = rf(opts, feeTokensToAdd, feeTokensToRemove)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Transaction)
		}
	}

	if rf, ok := ret.Get(1).(func(*bind.TransactOpts, []common.Address, []common.Address) error); ok {
		r1 = rf(opts, feeTokensToAdd, feeTokensToRemove)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ApplyPriceUpdatersUpdates provides a mock function with given fields: opts, priceUpdatersToAdd, priceUpdatersToRemove
func (_m *PriceRegistryInterface) ApplyPriceUpdatersUpdates(opts *bind.TransactOpts, priceUpdatersToAdd []common.Address, priceUpdatersToRemove []common.Address) (*types.Transaction, error) {
	ret := _m.Called(opts, priceUpdatersToAdd, priceUpdatersToRemove)

	var r0 *types.Transaction
	var r1 error
	if rf, ok := ret.Get(0).(func(*bind.TransactOpts, []common.Address, []common.Address) (*types.Transaction, error)); ok {
		return rf(opts, priceUpdatersToAdd, priceUpdatersToRemove)
	}
	if rf, ok := ret.Get(0).(func(*bind.TransactOpts, []common.Address, []common.Address) *types.Transaction); ok {
		r0 = rf(opts, priceUpdatersToAdd, priceUpdatersToRemove)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Transaction)
		}
	}

	if rf, ok := ret.Get(1).(func(*bind.TransactOpts, []common.Address, []common.Address) error); ok {
		r1 = rf(opts, priceUpdatersToAdd, priceUpdatersToRemove)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ConvertTokenAmount provides a mock function with given fields: opts, fromToken, fromTokenAmount, toToken
func (_m *PriceRegistryInterface) ConvertTokenAmount(opts *bind.CallOpts, fromToken common.Address, fromTokenAmount *big.Int, toToken common.Address) (*big.Int, error) {
	ret := _m.Called(opts, fromToken, fromTokenAmount, toToken)

	var r0 *big.Int
	var r1 error
	if rf, ok := ret.Get(0).(func(*bind.CallOpts, common.Address, *big.Int, common.Address) (*big.Int, error)); ok {
		return rf(opts, fromToken, fromTokenAmount, toToken)
	}
	if rf, ok := ret.Get(0).(func(*bind.CallOpts, common.Address, *big.Int, common.Address) *big.Int); ok {
		r0 = rf(opts, fromToken, fromTokenAmount, toToken)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*big.Int)
		}
	}

	if rf, ok := ret.Get(1).(func(*bind.CallOpts, common.Address, *big.Int, common.Address) error); ok {
		r1 = rf(opts, fromToken, fromTokenAmount, toToken)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FilterFeeTokenAdded provides a mock function with given fields: opts, feeToken
func (_m *PriceRegistryInterface) FilterFeeTokenAdded(opts *bind.FilterOpts, feeToken []common.Address) (*price_registry.PriceRegistryFeeTokenAddedIterator, error) {
	ret := _m.Called(opts, feeToken)

	var r0 *price_registry.PriceRegistryFeeTokenAddedIterator
	var r1 error
	if rf, ok := ret.Get(0).(func(*bind.FilterOpts, []common.Address) (*price_registry.PriceRegistryFeeTokenAddedIterator, error)); ok {
		return rf(opts, feeToken)
	}
	if rf, ok := ret.Get(0).(func(*bind.FilterOpts, []common.Address) *price_registry.PriceRegistryFeeTokenAddedIterator); ok {
		r0 = rf(opts, feeToken)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*price_registry.PriceRegistryFeeTokenAddedIterator)
		}
	}

	if rf, ok := ret.Get(1).(func(*bind.FilterOpts, []common.Address) error); ok {
		r1 = rf(opts, feeToken)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FilterFeeTokenRemoved provides a mock function with given fields: opts, feeToken
func (_m *PriceRegistryInterface) FilterFeeTokenRemoved(opts *bind.FilterOpts, feeToken []common.Address) (*price_registry.PriceRegistryFeeTokenRemovedIterator, error) {
	ret := _m.Called(opts, feeToken)

	var r0 *price_registry.PriceRegistryFeeTokenRemovedIterator
	var r1 error
	if rf, ok := ret.Get(0).(func(*bind.FilterOpts, []common.Address) (*price_registry.PriceRegistryFeeTokenRemovedIterator, error)); ok {
		return rf(opts, feeToken)
	}
	if rf, ok := ret.Get(0).(func(*bind.FilterOpts, []common.Address) *price_registry.PriceRegistryFeeTokenRemovedIterator); ok {
		r0 = rf(opts, feeToken)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*price_registry.PriceRegistryFeeTokenRemovedIterator)
		}
	}

	if rf, ok := ret.Get(1).(func(*bind.FilterOpts, []common.Address) error); ok {
		r1 = rf(opts, feeToken)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FilterOwnershipTransferRequested provides a mock function with given fields: opts, from, to
func (_m *PriceRegistryInterface) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*price_registry.PriceRegistryOwnershipTransferRequestedIterator, error) {
	ret := _m.Called(opts, from, to)

	var r0 *price_registry.PriceRegistryOwnershipTransferRequestedIterator
	var r1 error
	if rf, ok := ret.Get(0).(func(*bind.FilterOpts, []common.Address, []common.Address) (*price_registry.PriceRegistryOwnershipTransferRequestedIterator, error)); ok {
		return rf(opts, from, to)
	}
	if rf, ok := ret.Get(0).(func(*bind.FilterOpts, []common.Address, []common.Address) *price_registry.PriceRegistryOwnershipTransferRequestedIterator); ok {
		r0 = rf(opts, from, to)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*price_registry.PriceRegistryOwnershipTransferRequestedIterator)
		}
	}

	if rf, ok := ret.Get(1).(func(*bind.FilterOpts, []common.Address, []common.Address) error); ok {
		r1 = rf(opts, from, to)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FilterOwnershipTransferred provides a mock function with given fields: opts, from, to
func (_m *PriceRegistryInterface) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*price_registry.PriceRegistryOwnershipTransferredIterator, error) {
	ret := _m.Called(opts, from, to)

	var r0 *price_registry.PriceRegistryOwnershipTransferredIterator
	var r1 error
	if rf, ok := ret.Get(0).(func(*bind.FilterOpts, []common.Address, []common.Address) (*price_registry.PriceRegistryOwnershipTransferredIterator, error)); ok {
		return rf(opts, from, to)
	}
	if rf, ok := ret.Get(0).(func(*bind.FilterOpts, []common.Address, []common.Address) *price_registry.PriceRegistryOwnershipTransferredIterator); ok {
		r0 = rf(opts, from, to)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*price_registry.PriceRegistryOwnershipTransferredIterator)
		}
	}

	if rf, ok := ret.Get(1).(func(*bind.FilterOpts, []common.Address, []common.Address) error); ok {
		r1 = rf(opts, from, to)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FilterPriceUpdaterRemoved provides a mock function with given fields: opts, priceUpdater
func (_m *PriceRegistryInterface) FilterPriceUpdaterRemoved(opts *bind.FilterOpts, priceUpdater []common.Address) (*price_registry.PriceRegistryPriceUpdaterRemovedIterator, error) {
	ret := _m.Called(opts, priceUpdater)

	var r0 *price_registry.PriceRegistryPriceUpdaterRemovedIterator
	var r1 error
	if rf, ok := ret.Get(0).(func(*bind.FilterOpts, []common.Address) (*price_registry.PriceRegistryPriceUpdaterRemovedIterator, error)); ok {
		return rf(opts, priceUpdater)
	}
	if rf, ok := ret.Get(0).(func(*bind.FilterOpts, []common.Address) *price_registry.PriceRegistryPriceUpdaterRemovedIterator); ok {
		r0 = rf(opts, priceUpdater)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*price_registry.PriceRegistryPriceUpdaterRemovedIterator)
		}
	}

	if rf, ok := ret.Get(1).(func(*bind.FilterOpts, []common.Address) error); ok {
		r1 = rf(opts, priceUpdater)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FilterPriceUpdaterSet provides a mock function with given fields: opts, priceUpdater
func (_m *PriceRegistryInterface) FilterPriceUpdaterSet(opts *bind.FilterOpts, priceUpdater []common.Address) (*price_registry.PriceRegistryPriceUpdaterSetIterator, error) {
	ret := _m.Called(opts, priceUpdater)

	var r0 *price_registry.PriceRegistryPriceUpdaterSetIterator
	var r1 error
	if rf, ok := ret.Get(0).(func(*bind.FilterOpts, []common.Address) (*price_registry.PriceRegistryPriceUpdaterSetIterator, error)); ok {
		return rf(opts, priceUpdater)
	}
	if rf, ok := ret.Get(0).(func(*bind.FilterOpts, []common.Address) *price_registry.PriceRegistryPriceUpdaterSetIterator); ok {
		r0 = rf(opts, priceUpdater)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*price_registry.PriceRegistryPriceUpdaterSetIterator)
		}
	}

	if rf, ok := ret.Get(1).(func(*bind.FilterOpts, []common.Address) error); ok {
		r1 = rf(opts, priceUpdater)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FilterUsdPerTokenUpdated provides a mock function with given fields: opts, token
func (_m *PriceRegistryInterface) FilterUsdPerTokenUpdated(opts *bind.FilterOpts, token []common.Address) (*price_registry.PriceRegistryUsdPerTokenUpdatedIterator, error) {
	ret := _m.Called(opts, token)

	var r0 *price_registry.PriceRegistryUsdPerTokenUpdatedIterator
	var r1 error
	if rf, ok := ret.Get(0).(func(*bind.FilterOpts, []common.Address) (*price_registry.PriceRegistryUsdPerTokenUpdatedIterator, error)); ok {
		return rf(opts, token)
	}
	if rf, ok := ret.Get(0).(func(*bind.FilterOpts, []common.Address) *price_registry.PriceRegistryUsdPerTokenUpdatedIterator); ok {
		r0 = rf(opts, token)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*price_registry.PriceRegistryUsdPerTokenUpdatedIterator)
		}
	}

	if rf, ok := ret.Get(1).(func(*bind.FilterOpts, []common.Address) error); ok {
		r1 = rf(opts, token)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FilterUsdPerUnitGasUpdated provides a mock function with given fields: opts, destChain
func (_m *PriceRegistryInterface) FilterUsdPerUnitGasUpdated(opts *bind.FilterOpts, destChain []uint64) (*price_registry.PriceRegistryUsdPerUnitGasUpdatedIterator, error) {
	ret := _m.Called(opts, destChain)

	var r0 *price_registry.PriceRegistryUsdPerUnitGasUpdatedIterator
	var r1 error
	if rf, ok := ret.Get(0).(func(*bind.FilterOpts, []uint64) (*price_registry.PriceRegistryUsdPerUnitGasUpdatedIterator, error)); ok {
		return rf(opts, destChain)
	}
	if rf, ok := ret.Get(0).(func(*bind.FilterOpts, []uint64) *price_registry.PriceRegistryUsdPerUnitGasUpdatedIterator); ok {
		r0 = rf(opts, destChain)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*price_registry.PriceRegistryUsdPerUnitGasUpdatedIterator)
		}
	}

	if rf, ok := ret.Get(1).(func(*bind.FilterOpts, []uint64) error); ok {
		r1 = rf(opts, destChain)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetDestinationChainGasPrice provides a mock function with given fields: opts, destChainSelector
func (_m *PriceRegistryInterface) GetDestinationChainGasPrice(opts *bind.CallOpts, destChainSelector uint64) (price_registry.InternalTimestampedPackedUint224, error) {
	ret := _m.Called(opts, destChainSelector)

	var r0 price_registry.InternalTimestampedPackedUint224
	var r1 error
	if rf, ok := ret.Get(0).(func(*bind.CallOpts, uint64) (price_registry.InternalTimestampedPackedUint224, error)); ok {
		return rf(opts, destChainSelector)
	}
	if rf, ok := ret.Get(0).(func(*bind.CallOpts, uint64) price_registry.InternalTimestampedPackedUint224); ok {
		r0 = rf(opts, destChainSelector)
	} else {
		r0 = ret.Get(0).(price_registry.InternalTimestampedPackedUint224)
	}

	if rf, ok := ret.Get(1).(func(*bind.CallOpts, uint64) error); ok {
		r1 = rf(opts, destChainSelector)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetFeeTokens provides a mock function with given fields: opts
func (_m *PriceRegistryInterface) GetFeeTokens(opts *bind.CallOpts) ([]common.Address, error) {
	ret := _m.Called(opts)

	var r0 []common.Address
	var r1 error
	if rf, ok := ret.Get(0).(func(*bind.CallOpts) ([]common.Address, error)); ok {
		return rf(opts)
	}
	if rf, ok := ret.Get(0).(func(*bind.CallOpts) []common.Address); ok {
		r0 = rf(opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]common.Address)
		}
	}

	if rf, ok := ret.Get(1).(func(*bind.CallOpts) error); ok {
		r1 = rf(opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetPriceUpdaters provides a mock function with given fields: opts
func (_m *PriceRegistryInterface) GetPriceUpdaters(opts *bind.CallOpts) ([]common.Address, error) {
	ret := _m.Called(opts)

	var r0 []common.Address
	var r1 error
	if rf, ok := ret.Get(0).(func(*bind.CallOpts) ([]common.Address, error)); ok {
		return rf(opts)
	}
	if rf, ok := ret.Get(0).(func(*bind.CallOpts) []common.Address); ok {
		r0 = rf(opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]common.Address)
		}
	}

	if rf, ok := ret.Get(1).(func(*bind.CallOpts) error); ok {
		r1 = rf(opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetStalenessThreshold provides a mock function with given fields: opts
func (_m *PriceRegistryInterface) GetStalenessThreshold(opts *bind.CallOpts) (*big.Int, error) {
	ret := _m.Called(opts)

	var r0 *big.Int
	var r1 error
	if rf, ok := ret.Get(0).(func(*bind.CallOpts) (*big.Int, error)); ok {
		return rf(opts)
	}
	if rf, ok := ret.Get(0).(func(*bind.CallOpts) *big.Int); ok {
		r0 = rf(opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*big.Int)
		}
	}

	if rf, ok := ret.Get(1).(func(*bind.CallOpts) error); ok {
		r1 = rf(opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTokenAndGasPrices provides a mock function with given fields: opts, token, destChainSelector
func (_m *PriceRegistryInterface) GetTokenAndGasPrices(opts *bind.CallOpts, token common.Address, destChainSelector uint64) (price_registry.GetTokenAndGasPrices, error) {
	ret := _m.Called(opts, token, destChainSelector)

	var r0 price_registry.GetTokenAndGasPrices
	var r1 error
	if rf, ok := ret.Get(0).(func(*bind.CallOpts, common.Address, uint64) (price_registry.GetTokenAndGasPrices, error)); ok {
		return rf(opts, token, destChainSelector)
	}
	if rf, ok := ret.Get(0).(func(*bind.CallOpts, common.Address, uint64) price_registry.GetTokenAndGasPrices); ok {
		r0 = rf(opts, token, destChainSelector)
	} else {
		r0 = ret.Get(0).(price_registry.GetTokenAndGasPrices)
	}

	if rf, ok := ret.Get(1).(func(*bind.CallOpts, common.Address, uint64) error); ok {
		r1 = rf(opts, token, destChainSelector)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTokenPrice provides a mock function with given fields: opts, token
func (_m *PriceRegistryInterface) GetTokenPrice(opts *bind.CallOpts, token common.Address) (price_registry.InternalTimestampedPackedUint224, error) {
	ret := _m.Called(opts, token)

	var r0 price_registry.InternalTimestampedPackedUint224
	var r1 error
	if rf, ok := ret.Get(0).(func(*bind.CallOpts, common.Address) (price_registry.InternalTimestampedPackedUint224, error)); ok {
		return rf(opts, token)
	}
	if rf, ok := ret.Get(0).(func(*bind.CallOpts, common.Address) price_registry.InternalTimestampedPackedUint224); ok {
		r0 = rf(opts, token)
	} else {
		r0 = ret.Get(0).(price_registry.InternalTimestampedPackedUint224)
	}

	if rf, ok := ret.Get(1).(func(*bind.CallOpts, common.Address) error); ok {
		r1 = rf(opts, token)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTokenPrices provides a mock function with given fields: opts, tokens
func (_m *PriceRegistryInterface) GetTokenPrices(opts *bind.CallOpts, tokens []common.Address) ([]price_registry.InternalTimestampedPackedUint224, error) {
	ret := _m.Called(opts, tokens)

	var r0 []price_registry.InternalTimestampedPackedUint224
	var r1 error
	if rf, ok := ret.Get(0).(func(*bind.CallOpts, []common.Address) ([]price_registry.InternalTimestampedPackedUint224, error)); ok {
		return rf(opts, tokens)
	}
	if rf, ok := ret.Get(0).(func(*bind.CallOpts, []common.Address) []price_registry.InternalTimestampedPackedUint224); ok {
		r0 = rf(opts, tokens)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]price_registry.InternalTimestampedPackedUint224)
		}
	}

	if rf, ok := ret.Get(1).(func(*bind.CallOpts, []common.Address) error); ok {
		r1 = rf(opts, tokens)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetValidatedTokenPrice provides a mock function with given fields: opts, token
func (_m *PriceRegistryInterface) GetValidatedTokenPrice(opts *bind.CallOpts, token common.Address) (*big.Int, error) {
	ret := _m.Called(opts, token)

	var r0 *big.Int
	var r1 error
	if rf, ok := ret.Get(0).(func(*bind.CallOpts, common.Address) (*big.Int, error)); ok {
		return rf(opts, token)
	}
	if rf, ok := ret.Get(0).(func(*bind.CallOpts, common.Address) *big.Int); ok {
		r0 = rf(opts, token)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*big.Int)
		}
	}

	if rf, ok := ret.Get(1).(func(*bind.CallOpts, common.Address) error); ok {
		r1 = rf(opts, token)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Owner provides a mock function with given fields: opts
func (_m *PriceRegistryInterface) Owner(opts *bind.CallOpts) (common.Address, error) {
	ret := _m.Called(opts)

	var r0 common.Address
	var r1 error
	if rf, ok := ret.Get(0).(func(*bind.CallOpts) (common.Address, error)); ok {
		return rf(opts)
	}
	if rf, ok := ret.Get(0).(func(*bind.CallOpts) common.Address); ok {
		r0 = rf(opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(common.Address)
		}
	}

	if rf, ok := ret.Get(1).(func(*bind.CallOpts) error); ok {
		r1 = rf(opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ParseFeeTokenAdded provides a mock function with given fields: log
func (_m *PriceRegistryInterface) ParseFeeTokenAdded(log types.Log) (*price_registry.PriceRegistryFeeTokenAdded, error) {
	ret := _m.Called(log)

	var r0 *price_registry.PriceRegistryFeeTokenAdded
	var r1 error
	if rf, ok := ret.Get(0).(func(types.Log) (*price_registry.PriceRegistryFeeTokenAdded, error)); ok {
		return rf(log)
	}
	if rf, ok := ret.Get(0).(func(types.Log) *price_registry.PriceRegistryFeeTokenAdded); ok {
		r0 = rf(log)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*price_registry.PriceRegistryFeeTokenAdded)
		}
	}

	if rf, ok := ret.Get(1).(func(types.Log) error); ok {
		r1 = rf(log)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ParseFeeTokenRemoved provides a mock function with given fields: log
func (_m *PriceRegistryInterface) ParseFeeTokenRemoved(log types.Log) (*price_registry.PriceRegistryFeeTokenRemoved, error) {
	ret := _m.Called(log)

	var r0 *price_registry.PriceRegistryFeeTokenRemoved
	var r1 error
	if rf, ok := ret.Get(0).(func(types.Log) (*price_registry.PriceRegistryFeeTokenRemoved, error)); ok {
		return rf(log)
	}
	if rf, ok := ret.Get(0).(func(types.Log) *price_registry.PriceRegistryFeeTokenRemoved); ok {
		r0 = rf(log)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*price_registry.PriceRegistryFeeTokenRemoved)
		}
	}

	if rf, ok := ret.Get(1).(func(types.Log) error); ok {
		r1 = rf(log)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ParseLog provides a mock function with given fields: log
func (_m *PriceRegistryInterface) ParseLog(log types.Log) (generated.AbigenLog, error) {
	ret := _m.Called(log)

	var r0 generated.AbigenLog
	var r1 error
	if rf, ok := ret.Get(0).(func(types.Log) (generated.AbigenLog, error)); ok {
		return rf(log)
	}
	if rf, ok := ret.Get(0).(func(types.Log) generated.AbigenLog); ok {
		r0 = rf(log)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(generated.AbigenLog)
		}
	}

	if rf, ok := ret.Get(1).(func(types.Log) error); ok {
		r1 = rf(log)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ParseOwnershipTransferRequested provides a mock function with given fields: log
func (_m *PriceRegistryInterface) ParseOwnershipTransferRequested(log types.Log) (*price_registry.PriceRegistryOwnershipTransferRequested, error) {
	ret := _m.Called(log)

	var r0 *price_registry.PriceRegistryOwnershipTransferRequested
	var r1 error
	if rf, ok := ret.Get(0).(func(types.Log) (*price_registry.PriceRegistryOwnershipTransferRequested, error)); ok {
		return rf(log)
	}
	if rf, ok := ret.Get(0).(func(types.Log) *price_registry.PriceRegistryOwnershipTransferRequested); ok {
		r0 = rf(log)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*price_registry.PriceRegistryOwnershipTransferRequested)
		}
	}

	if rf, ok := ret.Get(1).(func(types.Log) error); ok {
		r1 = rf(log)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ParseOwnershipTransferred provides a mock function with given fields: log
func (_m *PriceRegistryInterface) ParseOwnershipTransferred(log types.Log) (*price_registry.PriceRegistryOwnershipTransferred, error) {
	ret := _m.Called(log)

	var r0 *price_registry.PriceRegistryOwnershipTransferred
	var r1 error
	if rf, ok := ret.Get(0).(func(types.Log) (*price_registry.PriceRegistryOwnershipTransferred, error)); ok {
		return rf(log)
	}
	if rf, ok := ret.Get(0).(func(types.Log) *price_registry.PriceRegistryOwnershipTransferred); ok {
		r0 = rf(log)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*price_registry.PriceRegistryOwnershipTransferred)
		}
	}

	if rf, ok := ret.Get(1).(func(types.Log) error); ok {
		r1 = rf(log)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ParsePriceUpdaterRemoved provides a mock function with given fields: log
func (_m *PriceRegistryInterface) ParsePriceUpdaterRemoved(log types.Log) (*price_registry.PriceRegistryPriceUpdaterRemoved, error) {
	ret := _m.Called(log)

	var r0 *price_registry.PriceRegistryPriceUpdaterRemoved
	var r1 error
	if rf, ok := ret.Get(0).(func(types.Log) (*price_registry.PriceRegistryPriceUpdaterRemoved, error)); ok {
		return rf(log)
	}
	if rf, ok := ret.Get(0).(func(types.Log) *price_registry.PriceRegistryPriceUpdaterRemoved); ok {
		r0 = rf(log)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*price_registry.PriceRegistryPriceUpdaterRemoved)
		}
	}

	if rf, ok := ret.Get(1).(func(types.Log) error); ok {
		r1 = rf(log)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ParsePriceUpdaterSet provides a mock function with given fields: log
func (_m *PriceRegistryInterface) ParsePriceUpdaterSet(log types.Log) (*price_registry.PriceRegistryPriceUpdaterSet, error) {
	ret := _m.Called(log)

	var r0 *price_registry.PriceRegistryPriceUpdaterSet
	var r1 error
	if rf, ok := ret.Get(0).(func(types.Log) (*price_registry.PriceRegistryPriceUpdaterSet, error)); ok {
		return rf(log)
	}
	if rf, ok := ret.Get(0).(func(types.Log) *price_registry.PriceRegistryPriceUpdaterSet); ok {
		r0 = rf(log)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*price_registry.PriceRegistryPriceUpdaterSet)
		}
	}

	if rf, ok := ret.Get(1).(func(types.Log) error); ok {
		r1 = rf(log)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ParseUsdPerTokenUpdated provides a mock function with given fields: log
func (_m *PriceRegistryInterface) ParseUsdPerTokenUpdated(log types.Log) (*price_registry.PriceRegistryUsdPerTokenUpdated, error) {
	ret := _m.Called(log)

	var r0 *price_registry.PriceRegistryUsdPerTokenUpdated
	var r1 error
	if rf, ok := ret.Get(0).(func(types.Log) (*price_registry.PriceRegistryUsdPerTokenUpdated, error)); ok {
		return rf(log)
	}
	if rf, ok := ret.Get(0).(func(types.Log) *price_registry.PriceRegistryUsdPerTokenUpdated); ok {
		r0 = rf(log)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*price_registry.PriceRegistryUsdPerTokenUpdated)
		}
	}

	if rf, ok := ret.Get(1).(func(types.Log) error); ok {
		r1 = rf(log)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ParseUsdPerUnitGasUpdated provides a mock function with given fields: log
func (_m *PriceRegistryInterface) ParseUsdPerUnitGasUpdated(log types.Log) (*price_registry.PriceRegistryUsdPerUnitGasUpdated, error) {
	ret := _m.Called(log)

	var r0 *price_registry.PriceRegistryUsdPerUnitGasUpdated
	var r1 error
	if rf, ok := ret.Get(0).(func(types.Log) (*price_registry.PriceRegistryUsdPerUnitGasUpdated, error)); ok {
		return rf(log)
	}
	if rf, ok := ret.Get(0).(func(types.Log) *price_registry.PriceRegistryUsdPerUnitGasUpdated); ok {
		r0 = rf(log)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*price_registry.PriceRegistryUsdPerUnitGasUpdated)
		}
	}

	if rf, ok := ret.Get(1).(func(types.Log) error); ok {
		r1 = rf(log)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// TransferOwnership provides a mock function with given fields: opts, to
func (_m *PriceRegistryInterface) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	ret := _m.Called(opts, to)

	var r0 *types.Transaction
	var r1 error
	if rf, ok := ret.Get(0).(func(*bind.TransactOpts, common.Address) (*types.Transaction, error)); ok {
		return rf(opts, to)
	}
	if rf, ok := ret.Get(0).(func(*bind.TransactOpts, common.Address) *types.Transaction); ok {
		r0 = rf(opts, to)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Transaction)
		}
	}

	if rf, ok := ret.Get(1).(func(*bind.TransactOpts, common.Address) error); ok {
		r1 = rf(opts, to)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// TypeAndVersion provides a mock function with given fields: opts
func (_m *PriceRegistryInterface) TypeAndVersion(opts *bind.CallOpts) (string, error) {
	ret := _m.Called(opts)

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(*bind.CallOpts) (string, error)); ok {
		return rf(opts)
	}
	if rf, ok := ret.Get(0).(func(*bind.CallOpts) string); ok {
		r0 = rf(opts)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(*bind.CallOpts) error); ok {
		r1 = rf(opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdatePrices provides a mock function with given fields: opts, priceUpdates
func (_m *PriceRegistryInterface) UpdatePrices(opts *bind.TransactOpts, priceUpdates price_registry.InternalPriceUpdates) (*types.Transaction, error) {
	ret := _m.Called(opts, priceUpdates)

	var r0 *types.Transaction
	var r1 error
	if rf, ok := ret.Get(0).(func(*bind.TransactOpts, price_registry.InternalPriceUpdates) (*types.Transaction, error)); ok {
		return rf(opts, priceUpdates)
	}
	if rf, ok := ret.Get(0).(func(*bind.TransactOpts, price_registry.InternalPriceUpdates) *types.Transaction); ok {
		r0 = rf(opts, priceUpdates)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Transaction)
		}
	}

	if rf, ok := ret.Get(1).(func(*bind.TransactOpts, price_registry.InternalPriceUpdates) error); ok {
		r1 = rf(opts, priceUpdates)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// WatchFeeTokenAdded provides a mock function with given fields: opts, sink, feeToken
func (_m *PriceRegistryInterface) WatchFeeTokenAdded(opts *bind.WatchOpts, sink chan<- *price_registry.PriceRegistryFeeTokenAdded, feeToken []common.Address) (event.Subscription, error) {
	ret := _m.Called(opts, sink, feeToken)

	var r0 event.Subscription
	var r1 error
	if rf, ok := ret.Get(0).(func(*bind.WatchOpts, chan<- *price_registry.PriceRegistryFeeTokenAdded, []common.Address) (event.Subscription, error)); ok {
		return rf(opts, sink, feeToken)
	}
	if rf, ok := ret.Get(0).(func(*bind.WatchOpts, chan<- *price_registry.PriceRegistryFeeTokenAdded, []common.Address) event.Subscription); ok {
		r0 = rf(opts, sink, feeToken)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(event.Subscription)
		}
	}

	if rf, ok := ret.Get(1).(func(*bind.WatchOpts, chan<- *price_registry.PriceRegistryFeeTokenAdded, []common.Address) error); ok {
		r1 = rf(opts, sink, feeToken)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// WatchFeeTokenRemoved provides a mock function with given fields: opts, sink, feeToken
func (_m *PriceRegistryInterface) WatchFeeTokenRemoved(opts *bind.WatchOpts, sink chan<- *price_registry.PriceRegistryFeeTokenRemoved, feeToken []common.Address) (event.Subscription, error) {
	ret := _m.Called(opts, sink, feeToken)

	var r0 event.Subscription
	var r1 error
	if rf, ok := ret.Get(0).(func(*bind.WatchOpts, chan<- *price_registry.PriceRegistryFeeTokenRemoved, []common.Address) (event.Subscription, error)); ok {
		return rf(opts, sink, feeToken)
	}
	if rf, ok := ret.Get(0).(func(*bind.WatchOpts, chan<- *price_registry.PriceRegistryFeeTokenRemoved, []common.Address) event.Subscription); ok {
		r0 = rf(opts, sink, feeToken)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(event.Subscription)
		}
	}

	if rf, ok := ret.Get(1).(func(*bind.WatchOpts, chan<- *price_registry.PriceRegistryFeeTokenRemoved, []common.Address) error); ok {
		r1 = rf(opts, sink, feeToken)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// WatchOwnershipTransferRequested provides a mock function with given fields: opts, sink, from, to
func (_m *PriceRegistryInterface) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *price_registry.PriceRegistryOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {
	ret := _m.Called(opts, sink, from, to)

	var r0 event.Subscription
	var r1 error
	if rf, ok := ret.Get(0).(func(*bind.WatchOpts, chan<- *price_registry.PriceRegistryOwnershipTransferRequested, []common.Address, []common.Address) (event.Subscription, error)); ok {
		return rf(opts, sink, from, to)
	}
	if rf, ok := ret.Get(0).(func(*bind.WatchOpts, chan<- *price_registry.PriceRegistryOwnershipTransferRequested, []common.Address, []common.Address) event.Subscription); ok {
		r0 = rf(opts, sink, from, to)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(event.Subscription)
		}
	}

	if rf, ok := ret.Get(1).(func(*bind.WatchOpts, chan<- *price_registry.PriceRegistryOwnershipTransferRequested, []common.Address, []common.Address) error); ok {
		r1 = rf(opts, sink, from, to)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// WatchOwnershipTransferred provides a mock function with given fields: opts, sink, from, to
func (_m *PriceRegistryInterface) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *price_registry.PriceRegistryOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {
	ret := _m.Called(opts, sink, from, to)

	var r0 event.Subscription
	var r1 error
	if rf, ok := ret.Get(0).(func(*bind.WatchOpts, chan<- *price_registry.PriceRegistryOwnershipTransferred, []common.Address, []common.Address) (event.Subscription, error)); ok {
		return rf(opts, sink, from, to)
	}
	if rf, ok := ret.Get(0).(func(*bind.WatchOpts, chan<- *price_registry.PriceRegistryOwnershipTransferred, []common.Address, []common.Address) event.Subscription); ok {
		r0 = rf(opts, sink, from, to)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(event.Subscription)
		}
	}

	if rf, ok := ret.Get(1).(func(*bind.WatchOpts, chan<- *price_registry.PriceRegistryOwnershipTransferred, []common.Address, []common.Address) error); ok {
		r1 = rf(opts, sink, from, to)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// WatchPriceUpdaterRemoved provides a mock function with given fields: opts, sink, priceUpdater
func (_m *PriceRegistryInterface) WatchPriceUpdaterRemoved(opts *bind.WatchOpts, sink chan<- *price_registry.PriceRegistryPriceUpdaterRemoved, priceUpdater []common.Address) (event.Subscription, error) {
	ret := _m.Called(opts, sink, priceUpdater)

	var r0 event.Subscription
	var r1 error
	if rf, ok := ret.Get(0).(func(*bind.WatchOpts, chan<- *price_registry.PriceRegistryPriceUpdaterRemoved, []common.Address) (event.Subscription, error)); ok {
		return rf(opts, sink, priceUpdater)
	}
	if rf, ok := ret.Get(0).(func(*bind.WatchOpts, chan<- *price_registry.PriceRegistryPriceUpdaterRemoved, []common.Address) event.Subscription); ok {
		r0 = rf(opts, sink, priceUpdater)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(event.Subscription)
		}
	}

	if rf, ok := ret.Get(1).(func(*bind.WatchOpts, chan<- *price_registry.PriceRegistryPriceUpdaterRemoved, []common.Address) error); ok {
		r1 = rf(opts, sink, priceUpdater)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// WatchPriceUpdaterSet provides a mock function with given fields: opts, sink, priceUpdater
func (_m *PriceRegistryInterface) WatchPriceUpdaterSet(opts *bind.WatchOpts, sink chan<- *price_registry.PriceRegistryPriceUpdaterSet, priceUpdater []common.Address) (event.Subscription, error) {
	ret := _m.Called(opts, sink, priceUpdater)

	var r0 event.Subscription
	var r1 error
	if rf, ok := ret.Get(0).(func(*bind.WatchOpts, chan<- *price_registry.PriceRegistryPriceUpdaterSet, []common.Address) (event.Subscription, error)); ok {
		return rf(opts, sink, priceUpdater)
	}
	if rf, ok := ret.Get(0).(func(*bind.WatchOpts, chan<- *price_registry.PriceRegistryPriceUpdaterSet, []common.Address) event.Subscription); ok {
		r0 = rf(opts, sink, priceUpdater)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(event.Subscription)
		}
	}

	if rf, ok := ret.Get(1).(func(*bind.WatchOpts, chan<- *price_registry.PriceRegistryPriceUpdaterSet, []common.Address) error); ok {
		r1 = rf(opts, sink, priceUpdater)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// WatchUsdPerTokenUpdated provides a mock function with given fields: opts, sink, token
func (_m *PriceRegistryInterface) WatchUsdPerTokenUpdated(opts *bind.WatchOpts, sink chan<- *price_registry.PriceRegistryUsdPerTokenUpdated, token []common.Address) (event.Subscription, error) {
	ret := _m.Called(opts, sink, token)

	var r0 event.Subscription
	var r1 error
	if rf, ok := ret.Get(0).(func(*bind.WatchOpts, chan<- *price_registry.PriceRegistryUsdPerTokenUpdated, []common.Address) (event.Subscription, error)); ok {
		return rf(opts, sink, token)
	}
	if rf, ok := ret.Get(0).(func(*bind.WatchOpts, chan<- *price_registry.PriceRegistryUsdPerTokenUpdated, []common.Address) event.Subscription); ok {
		r0 = rf(opts, sink, token)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(event.Subscription)
		}
	}

	if rf, ok := ret.Get(1).(func(*bind.WatchOpts, chan<- *price_registry.PriceRegistryUsdPerTokenUpdated, []common.Address) error); ok {
		r1 = rf(opts, sink, token)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// WatchUsdPerUnitGasUpdated provides a mock function with given fields: opts, sink, destChain
func (_m *PriceRegistryInterface) WatchUsdPerUnitGasUpdated(opts *bind.WatchOpts, sink chan<- *price_registry.PriceRegistryUsdPerUnitGasUpdated, destChain []uint64) (event.Subscription, error) {
	ret := _m.Called(opts, sink, destChain)

	var r0 event.Subscription
	var r1 error
	if rf, ok := ret.Get(0).(func(*bind.WatchOpts, chan<- *price_registry.PriceRegistryUsdPerUnitGasUpdated, []uint64) (event.Subscription, error)); ok {
		return rf(opts, sink, destChain)
	}
	if rf, ok := ret.Get(0).(func(*bind.WatchOpts, chan<- *price_registry.PriceRegistryUsdPerUnitGasUpdated, []uint64) event.Subscription); ok {
		r0 = rf(opts, sink, destChain)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(event.Subscription)
		}
	}

	if rf, ok := ret.Get(1).(func(*bind.WatchOpts, chan<- *price_registry.PriceRegistryUsdPerUnitGasUpdated, []uint64) error); ok {
		r1 = rf(opts, sink, destChain)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewPriceRegistryInterface interface {
	mock.TestingT
	Cleanup(func())
}

// NewPriceRegistryInterface creates a new instance of PriceRegistryInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewPriceRegistryInterface(t mockConstructorTestingTNewPriceRegistryInterface) *PriceRegistryInterface {
	mock := &PriceRegistryInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
