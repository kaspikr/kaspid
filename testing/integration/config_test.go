package integration

import (
	"io/ioutil"
	"testing"
	"time"

	"github.com/kaspikr/kaspid/domain/dagconfig"
	"github.com/kaspikr/kaspid/infrastructure/config"
)

const (
	p2pAddress1 = "127.0.0.1:54321"
	p2pAddress2 = "127.0.0.1:54322"
	p2pAddress3 = "127.0.0.1:54323"
	p2pAddress4 = "127.0.0.1:54324"
	p2pAddress5 = "127.0.0.1:54325"

	rpcAddress1 = "127.0.0.1:12345"
	rpcAddress2 = "127.0.0.1:12346"
	rpcAddress3 = "127.0.0.1:12347"
	rpcAddress4 = "127.0.0.1:12348"
	rpcAddress5 = "127.0.0.1:12349"

	miningAddress1           = "kaspisim:qzj4l226yct3x6s8ztc0eah39qlzsyjyd5xxmkm7v95pl9k0whcr6zxl24066"
	miningAddress1PrivateKey = "64c2d5ea29364bc1a5215b10aeabe167ce514486caa283ebdaa1f7c38dba4526"

	miningAddress2           = "kaspisim:qzj4l226yct3x6s8ztc0eah39qlzsyjyd5xxmkm7v95pl9k0whcr6zxl24066"
	miningAddress2PrivateKey = "64c2d5ea29364bc1a5215b10aeabe167ce514486caa283ebdaa1f7c38dba4526"

	miningAddress3           = "kaspisim:qrggqrqnajm8qkrflrzcjchj4n32nddsfe7jmlatgl28vmdkxf54wykjyus4r"
	miningAddress3PrivateKey = "51bd4a0a65fe05788c9fb6ee4905fad879c60d82715d7dfae6d0e5f719409363"

	defaultTimeout = 30 * time.Second
)

func setConfig(t *testing.T, harness *appHarness, protocolVersion uint32) {
	harness.config = commonConfig()
	harness.config.AppDir = randomDirectory(t)
	harness.config.Listeners = []string{harness.p2pAddress}
	harness.config.RPCListeners = []string{harness.rpcAddress}
	harness.config.UTXOIndex = harness.utxoIndex
	harness.config.AllowSubmitBlockWhenNotSynced = true
	if protocolVersion != 0 {
		harness.config.ProtocolVersion = protocolVersion
	}

	if harness.overrideDAGParams != nil {
		harness.config.ActiveNetParams = harness.overrideDAGParams
	}
}

func commonConfig() *config.Config {
	commonConfig := config.DefaultConfig()

	*commonConfig.ActiveNetParams = dagconfig.SimnetParams // Copy so that we can make changes safely
	commonConfig.ActiveNetParams.BlockCoinbaseMaturity = 10
	commonConfig.TargetOutboundPeers = 0
	commonConfig.DisableDNSSeed = true
	commonConfig.Simnet = true

	return commonConfig
}

func randomDirectory(t *testing.T) string {
	dir, err := ioutil.TempDir("", "integration-test")
	if err != nil {
		t.Fatalf("Error creating temporary directory for test: %+v", err)
	}

	return dir
}
