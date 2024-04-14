package constants

import "math"

const (
	// BlockVersion represents the current block version
	BlockVersion uint16 = 1

	// MaxTransactionVersion is the current latest supported transaction version.
	MaxTransactionVersion uint16 = 0

	// MaxScriptPublicKeyVersion is the current latest supported public key script version.
	MaxScriptPublicKeyVersion uint16 = 0

	// SompiPerKaspi is the number of sompi in one kaspi (1 KAS).
	SompiPerKaspi = 100_000_000

	// MaxSompi is the maximum transaction amount allowed in sompi.
	// 10초당 1블럭. 500개 보상일경우 117년. 반감기 없을때.
	MaxSompi = uint64(180_000_000_000 * SompiPerKaspi)

	// MaxTxInSequenceNum is the maximum sequence number the sequence field
	// of a transaction input can be.
	MaxTxInSequenceNum uint64 = math.MaxUint64

	// SequenceLockTimeDisabled is a flag that if set on a transaction
	// input's sequence number, the sequence number will not be interpreted
	// as a relative locktime.
	SequenceLockTimeDisabled uint64 = 1 << 63

	// SequenceLockTimeMask is a mask that extracts the relative locktime
	// when masked against the transaction input sequence number.
	SequenceLockTimeMask uint64 = 0x00000000ffffffff

	// LockTimeThreshold is the number below which a lock time is
	// interpreted to be a DAA score.
	LockTimeThreshold = 5e11 // Tue Nov 5 00:53:20 1985 UTC

	// UnacceptedDAAScore is used to for UTXOEntries that were created by transactions in the mempool, or otherwise
	// not-yet-accepted transactions.
	UnacceptedDAAScore = math.MaxUint64
)
