package main

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/kaspikr/kaspid/infrastructure/network/netadapter/server/grpcserver/protowire"
)

var commandTypes = []reflect.Type{
	reflect.TypeOf(protowire.KaspidMessage_AddPeerRequest{}),
	reflect.TypeOf(protowire.KaspidMessage_GetConnectedPeerInfoRequest{}),
	reflect.TypeOf(protowire.KaspidMessage_GetPeerAddressesRequest{}),
	reflect.TypeOf(protowire.KaspidMessage_GetCurrentNetworkRequest{}),
	reflect.TypeOf(protowire.KaspidMessage_GetInfoRequest{}),

	reflect.TypeOf(protowire.KaspidMessage_GetBlockRequest{}),
	reflect.TypeOf(protowire.KaspidMessage_GetBlocksRequest{}),
	reflect.TypeOf(protowire.KaspidMessage_GetHeadersRequest{}),
	reflect.TypeOf(protowire.KaspidMessage_GetBlockCountRequest{}),
	reflect.TypeOf(protowire.KaspidMessage_GetBlockDagInfoRequest{}),
	reflect.TypeOf(protowire.KaspidMessage_GetSelectedTipHashRequest{}),
	reflect.TypeOf(protowire.KaspidMessage_GetVirtualSelectedParentBlueScoreRequest{}),
	reflect.TypeOf(protowire.KaspidMessage_GetVirtualSelectedParentChainFromBlockRequest{}),
	reflect.TypeOf(protowire.KaspidMessage_ResolveFinalityConflictRequest{}),
	reflect.TypeOf(protowire.KaspidMessage_EstimateNetworkHashesPerSecondRequest{}),

	reflect.TypeOf(protowire.KaspidMessage_GetBlockTemplateRequest{}),
	reflect.TypeOf(protowire.KaspidMessage_SubmitBlockRequest{}),

	reflect.TypeOf(protowire.KaspidMessage_GetMempoolEntryRequest{}),
	reflect.TypeOf(protowire.KaspidMessage_GetMempoolEntriesRequest{}),
	reflect.TypeOf(protowire.KaspidMessage_GetMempoolEntriesByAddressesRequest{}),

	reflect.TypeOf(protowire.KaspidMessage_SubmitTransactionRequest{}),

	reflect.TypeOf(protowire.KaspidMessage_GetUtxosByAddressesRequest{}),
	reflect.TypeOf(protowire.KaspidMessage_GetBalanceByAddressRequest{}),
	reflect.TypeOf(protowire.KaspidMessage_GetCoinSupplyRequest{}),

	reflect.TypeOf(protowire.KaspidMessage_BanRequest{}),
	reflect.TypeOf(protowire.KaspidMessage_UnbanRequest{}),
}

type commandDescription struct {
	name       string
	parameters []*parameterDescription
	typeof     reflect.Type
}

type parameterDescription struct {
	name   string
	typeof reflect.Type
}

func commandDescriptions() []*commandDescription {
	commandDescriptions := make([]*commandDescription, len(commandTypes))

	for i, commandTypeWrapped := range commandTypes {
		commandType := unwrapCommandType(commandTypeWrapped)

		name := strings.TrimSuffix(commandType.Name(), "RequestMessage")
		numFields := commandType.NumField()

		var parameters []*parameterDescription
		for i := 0; i < numFields; i++ {
			field := commandType.Field(i)

			if !isFieldExported(field) {
				continue
			}

			parameters = append(parameters, &parameterDescription{
				name:   field.Name,
				typeof: field.Type,
			})
		}
		commandDescriptions[i] = &commandDescription{
			name:       name,
			parameters: parameters,
			typeof:     commandTypeWrapped,
		}
	}

	return commandDescriptions
}

func (cd *commandDescription) help() string {
	sb := &strings.Builder{}
	sb.WriteString(cd.name)
	for _, parameter := range cd.parameters {
		_, _ = fmt.Fprintf(sb, " [%s]", parameter.name)
	}
	return sb.String()
}
