package libkaspiwallet

import (
	"encoding/hex"

	"github.com/kaspikr/kaspid/app/appmessage"
	"github.com/kaspikr/kaspid/cmd/kaspiwallet/daemon/pb"
	"github.com/kaspikr/kaspid/domain/consensus/model/externalapi"
	"github.com/kaspikr/kaspid/domain/consensus/utils/transactionid"
	"github.com/kaspikr/kaspid/domain/consensus/utils/utxo"
)

// KaspiwalletdUTXOsTolibkaspiwalletUTXOs converts a  []*pb.UtxosByAddressesEntry to a []*libkaspiwallet.UTXO
func KaspiwalletdUTXOsTolibkaspiwalletUTXOs(kaspiwalletdUtxoEntires []*pb.UtxosByAddressesEntry) ([]*UTXO, error) {
	UTXOs := make([]*UTXO, len(kaspiwalletdUtxoEntires))
	for i, entry := range kaspiwalletdUtxoEntires {
		script, err := hex.DecodeString(entry.UtxoEntry.ScriptPublicKey.ScriptPublicKey)
		if err != nil {
			return nil, err
		}
		transactionID, err := transactionid.FromString(entry.Outpoint.TransactionId)
		if err != nil {
			return nil, err
		}
		UTXOs[i] = &UTXO{
			UTXOEntry: utxo.NewUTXOEntry(
				entry.UtxoEntry.Amount,
				&externalapi.ScriptPublicKey{
					Script:  script,
					Version: uint16(entry.UtxoEntry.ScriptPublicKey.Version),
				},
				entry.UtxoEntry.IsCoinbase,
				entry.UtxoEntry.BlockDaaScore,
			),
			Outpoint: &externalapi.DomainOutpoint{
				TransactionID: *transactionID,
				Index:         entry.Outpoint.Index,
			},
		}
	}
	return UTXOs, nil
}

// AppMessageUTXOToKaspiwalletdUTXO converts an appmessage.UTXOsByAddressesEntry to a  pb.UtxosByAddressesEntry
func AppMessageUTXOToKaspiwalletdUTXO(appUTXOsByAddressesEntry *appmessage.UTXOsByAddressesEntry) *pb.UtxosByAddressesEntry {
	return &pb.UtxosByAddressesEntry{
		Outpoint: &pb.Outpoint{
			TransactionId: appUTXOsByAddressesEntry.Outpoint.TransactionID,
			Index:         appUTXOsByAddressesEntry.Outpoint.Index,
		},
		UtxoEntry: &pb.UtxoEntry{
			Amount: appUTXOsByAddressesEntry.UTXOEntry.Amount,
			ScriptPublicKey: &pb.ScriptPublicKey{
				Version:         uint32(appUTXOsByAddressesEntry.UTXOEntry.ScriptPublicKey.Version),
				ScriptPublicKey: appUTXOsByAddressesEntry.UTXOEntry.ScriptPublicKey.Script,
			},
			BlockDaaScore: appUTXOsByAddressesEntry.UTXOEntry.BlockDAAScore,
			IsCoinbase:    appUTXOsByAddressesEntry.UTXOEntry.IsCoinbase,
		},
	}
}
