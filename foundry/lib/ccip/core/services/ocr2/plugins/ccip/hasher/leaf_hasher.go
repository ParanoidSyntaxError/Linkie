package hasher

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/abihelpers"
	"github.com/smartcontractkit/chainlink/v2/core/utils"
)

type LeafHasherInterface[H Hash] interface {
	HashLeaf(log types.Log) (H, error)
}

var LeafDomainSeparator = [1]byte{0x00}

func GetMetaDataHash[H Hash](ctx Ctx[H], prefix [32]byte, sourceChainId uint64, onRampId common.Address, destChainId uint64) H {
	paddedOnRamp := onRampId.Hash()
	return ctx.Hash(utils.ConcatBytes(prefix[:], math.U256Bytes(big.NewInt(0).SetUint64(sourceChainId)), math.U256Bytes(big.NewInt(0).SetUint64(destChainId)), paddedOnRamp[:]))
}

type LeafHasher struct {
	metaDataHash [32]byte
	ctx          Ctx[[32]byte]
}

func NewLeafHasher(sourceChainId uint64, destChainId uint64, onRampId common.Address, ctx Ctx[[32]byte]) *LeafHasher {
	return &LeafHasher{
		metaDataHash: GetMetaDataHash(ctx, ctx.Hash([]byte("EVM2EVMMessageEvent")), sourceChainId, onRampId, destChainId),
		ctx:          ctx,
	}
}

var _ LeafHasherInterface[[32]byte] = &LeafHasher{}

func (t *LeafHasher) HashLeaf(log types.Log) ([32]byte, error) {
	message, err := abihelpers.DecodeOffRampMessage(log.Data)
	if err != nil {
		return [32]byte{}, err
	}

	encodedTokens, err := abihelpers.TokenAmountsArgs.PackValues([]interface{}{message.TokenAmounts})
	if err != nil {
		return [32]byte{}, err
	}

	packedValues, err := utils.ABIEncode(
		`[
{"name": "leafDomainSeparator","type":"bytes1"},
{"name": "metadataHash", "type":"bytes32"},
{"name": "sequenceNumber", "type":"uint64"},
{"name": "nonce", "type":"uint64"},
{"name": "sender", "type":"address"},
{"name": "receiver", "type":"address"},
{"name": "dataHash", "type":"bytes32"},
{"name": "tokenAmountsHash", "type":"bytes32"},
{"name": "gasLimit", "type":"uint256"},
{"name": "strict", "type":"bool"},
{"name": "feeToken","type": "address"},
{"name": "feeTokenAmount","type": "uint256"}
]`,
		LeafDomainSeparator,
		t.metaDataHash,
		message.SequenceNumber,
		message.Nonce,
		message.Sender,
		message.Receiver,
		t.ctx.Hash(message.Data),
		t.ctx.Hash(encodedTokens),
		message.GasLimit,
		message.Strict,
		message.FeeToken,
		message.FeeTokenAmount,
	)
	if err != nil {
		return [32]byte{}, err
	}
	return t.ctx.Hash(packedValues), nil
}
