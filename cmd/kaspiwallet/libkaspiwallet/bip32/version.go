package bip32

import "github.com/pkg/errors"

// BitcoinMainnetPrivate is the version that is used for
// bitcoin mainnet bip32 private extended keys.
// Ecnodes to xprv in base58.
var BitcoinMainnetPrivate = [4]byte{
	0x04,
	0x88,
	0xad,
	0xe4,
}

// BitcoinMainnetPublic is the version that is used for
// bitcoin mainnet bip32 public extended keys.
// Ecnodes to xpub in base58.
var BitcoinMainnetPublic = [4]byte{
	0x04,
	0x88,
	0xb2,
	0x1e,
}

// KaspiMainnetPrivate is the version that is used for
// kaspi mainnet bip32 private extended keys.
// Ecnodes to xprv in base58.
var KaspiMainnetPrivate = [4]byte{
	0x03,
	0x8f,
	0x2e,
	0xf4,
}

// KaspiMainnetPublic is the version that is used for
// kaspi mainnet bip32 public extended keys.
// Ecnodes to kpub in base58.
var KaspiMainnetPublic = [4]byte{
	0x03,
	0x8f,
	0x33,
	0x2e,
}

// KaspiTestnetPrivate is the version that is used for
// kaspi testnet bip32 public extended keys.
// Ecnodes to ktrv in base58.
var KaspiTestnetPrivate = [4]byte{
	0x03,
	0x90,
	0x9e,
	0x07,
}

// KaspiTestnetPublic is the version that is used for
// kaspi testnet bip32 public extended keys.
// Ecnodes to ktub in base58.
var KaspiTestnetPublic = [4]byte{
	0x03,
	0x90,
	0xa2,
	0x41,
}

// KaspiDevnetPrivate is the version that is used for
// kaspi devnet bip32 public extended keys.
// Ecnodes to kdrv in base58.
var KaspiDevnetPrivate = [4]byte{
	0x03,
	0x8b,
	0x3d,
	0x80,
}

// KaspiDevnetPublic is the version that is used for
// kaspi devnet bip32 public extended keys.
// Ecnodes to xdub in base58.
var KaspiDevnetPublic = [4]byte{
	0x03,
	0x8b,
	0x41,
	0xba,
}

// KaspiSimnetPrivate is the version that is used for
// kaspi simnet bip32 public extended keys.
// Ecnodes to ksrv in base58.
var KaspiSimnetPrivate = [4]byte{
	0x03,
	0x90,
	0x42,
	0x42,
}

// KaspiSimnetPublic is the version that is used for
// kaspi simnet bip32 public extended keys.
// Ecnodes to xsub in base58.
var KaspiSimnetPublic = [4]byte{
	0x03,
	0x90,
	0x46,
	0x7d,
}

func toPublicVersion(version [4]byte) ([4]byte, error) {
	switch version {
	case BitcoinMainnetPrivate:
		return BitcoinMainnetPublic, nil
	case KaspiMainnetPrivate:
		return KaspiMainnetPublic, nil
	case KaspiTestnetPrivate:
		return KaspiTestnetPublic, nil
	case KaspiDevnetPrivate:
		return KaspiDevnetPublic, nil
	case KaspiSimnetPrivate:
		return KaspiSimnetPublic, nil
	}

	return [4]byte{}, errors.Errorf("unknown version %x", version)
}

func isPrivateVersion(version [4]byte) bool {
	switch version {
	case BitcoinMainnetPrivate:
		return true
	case KaspiMainnetPrivate:
		return true
	case KaspiTestnetPrivate:
		return true
	case KaspiDevnetPrivate:
		return true
	case KaspiSimnetPrivate:
		return true
	}

	return false
}
