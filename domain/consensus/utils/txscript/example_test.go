// Copyright (c) 2014-2016 The btcsuite developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package txscript_test

import (
	"encoding/hex"
	"fmt"
	"github.com/kaspikr/kaspid/domain/consensus/model/externalapi"

	"github.com/kaspikr/kaspid/domain/consensus/utils/txscript"
	"github.com/kaspikr/kaspid/domain/dagconfig"
	"github.com/kaspikr/kaspid/util"
)

// This example demonstrates creating a script which pays to a kaspi address.
// It also prints the created script hex and uses the DisasmString function to
// display the disassembled script.
func ExamplePayToAddrScript() {
	// Parse the address to send the coins to into a util.Address
	// which is useful to ensure the accuracy of the address and determine
	// the address type. It is also required for the upcoming call to
	// PayToAddrScript.
	addressStr := "kaspi:qzy6cf82zzah2xh5jwtz8nx9u4gdj6zzke8gljs0v055ksmnl424u50xfc8mh"
	address, err := util.DecodeAddress(addressStr, util.Bech32PrefixKaspi)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Create a public key script that pays to the address.
	script, err := txscript.PayToAddrScript(address)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Script Hex: %x\n", script.Script)

	disasm, err := txscript.DisasmString(script.Version, script.Script)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Script Disassembly:", disasm)

	// Output:
	// Script Hex: 2089ac24ea10bb751af4939623ccc5e550d96842b64e8fca0f63e94b4373fd555eac
	// Script Disassembly: 89ac24ea10bb751af4939623ccc5e550d96842b64e8fca0f63e94b4373fd555e OP_CHECKSIG
}

// This example demonstrates extracting information from a standard public key
// script.
func ExampleExtractScriptPubKeyAddress() {
	// Start with a standard pay-to-pubkey script.
	scriptHex := "200b331220d9d652a4470ce774d7af86516247fa9e198ae2a7b97163047404ee00ac"
	script, err := hex.DecodeString(scriptHex)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Extract and print details from the script.
	scriptClass, address, err := txscript.ExtractScriptPubKeyAddress(
		&externalapi.ScriptPublicKey{
			Script:  script,
			Version: 0,
		}, &dagconfig.MainnetParams)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Script Class:", scriptClass)
	fmt.Println("Address:", address)

	// Output:
	// Script Class: pubkey
	// Address: kaspi:qq9nxy3qm8t99fz8pnnhf4a0segky3l6ncvc4c48h9ckxpr5qnhqq439jethw
}
