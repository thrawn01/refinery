package config

import (
	"encoding/json"
	"time"
)

// MockConfig will respond with whatever config it's set to do during
// initialization
type MockConfig struct {
	Callbacks                            []func()
	GetAPIKeysErr                        error
	GetAPIKeysVal                        []string
	GetCollectorTypeErr                  error
	GetCollectorTypeVal                  string
	GetInMemoryCollectorCacheCapacityErr error
	GetInMemoryCollectorCacheCapacityVal InMemoryCollectorCacheCapacity
	GetHoneycombAPIErr                   error
	GetHoneycombAPIVal                   string
	GetListenAddrErr                     error
	GetListenAddrVal                     string
	GetPeerListenAddrErr                 error
	GetPeerListenAddrVal                 string
	GetLoggerTypeErr                     error
	GetLoggerTypeVal                     string
	GetHoneycombLoggerConfigErr          error
	GetHoneycombLoggerConfigVal          HoneycombLoggerConfig
	GetLoggingLevelErr                   error
	GetLoggingLevelVal                   string
	GetOtherConfigErr                    error
	// GetOtherConfigVal must be a JSON representation of the config struct to be populated.
	GetOtherConfigVal             string
	GetPeersErr                   error
	GetPeersVal                   []string
	GetRedisHostErr               error
	GetRedisHostVal               string
	GetSamplerTypeErr             error
	GetSamplerTypeVal             interface{}
	GetMetricsTypeErr             error
	GetMetricsTypeVal             string
	GetHoneycombMetricsConfigErr  error
	GetHoneycombMetricsConfigVal  HoneycombMetricsConfig
	GetPrometheusMetricsConfigErr error
	GetPrometheusMetricsConfigVal PrometheusMetricsConfig
	GetSendDelayErr               error
	GetSendDelayVal               time.Duration
	GetTraceTimeoutErr            error
	GetTraceTimeoutVal            time.Duration
	GetUpstreamBufferSizeVal      int
	GetPeerBufferSizeVal          int
	SendTickerVal                 time.Duration
	IdentifierInterfaceName       string
	UseIPV6Identifier             bool
	RedisIdentifier               string
	PeerManagementType            string
	DebugServiceAddr              string
	DryRun                        bool
}

func (m *MockConfig) ReloadConfig() {
	for _, callback := range m.Callbacks {
		callback()
	}
}
func (m *MockConfig) RegisterReloadCallback(callback func()) {
	m.Callbacks = append(m.Callbacks, callback)
}
func (m *MockConfig) GetAPIKeys() ([]string, error) { return m.GetAPIKeysVal, m.GetAPIKeysErr }
func (m *MockConfig) GetCollectorType() (string, error) {
	return m.GetCollectorTypeVal, m.GetCollectorTypeErr
}
func (m *MockConfig) GetInMemCollectorCacheCapacity() (InMemoryCollectorCacheCapacity, error) {
	return m.GetInMemoryCollectorCacheCapacityVal, m.GetInMemoryCollectorCacheCapacityErr
}
func (m *MockConfig) GetHoneycombAPI() (string, error) {
	return m.GetHoneycombAPIVal, m.GetHoneycombAPIErr
}
func (m *MockConfig) GetListenAddr() (string, error) { return m.GetListenAddrVal, m.GetListenAddrErr }
func (m *MockConfig) GetPeerListenAddr() (string, error) {
	return m.GetPeerListenAddrVal, m.GetPeerListenAddrErr
}
func (m *MockConfig) GetLoggerType() (string, error) { return m.GetLoggerTypeVal, m.GetLoggerTypeErr }
func (m *MockConfig) GetHoneycombLoggerConfig() (HoneycombLoggerConfig, error) {
	return m.GetHoneycombLoggerConfigVal, m.GetHoneycombLoggerConfigErr
}
func (m *MockConfig) GetLoggingLevel() (string, error) {
	return m.GetLoggingLevelVal, m.GetLoggingLevelErr
}
func (m *MockConfig) GetOtherConfig(name string, iface interface{}) error {
	err := json.Unmarshal([]byte(m.GetOtherConfigVal), iface)
	if err != nil {
		return err
	}
	return m.GetOtherConfigErr
}
func (m *MockConfig) GetPeers() ([]string, error)   { return m.GetPeersVal, m.GetPeersErr }
func (m *MockConfig) GetRedisHost() (string, error) { return m.GetRedisHostVal, m.GetRedisHostErr }
func (m *MockConfig) GetMetricsType() (string, error) {
	return m.GetMetricsTypeVal, m.GetMetricsTypeErr
}
func (m *MockConfig) GetHoneycombMetricsConfig() (HoneycombMetricsConfig, error) {
	return m.GetHoneycombMetricsConfigVal, m.GetHoneycombMetricsConfigErr
}
func (m *MockConfig) GetPrometheusMetricsConfig() (PrometheusMetricsConfig, error) {
	return m.GetPrometheusMetricsConfigVal, m.GetPrometheusMetricsConfigErr
}
func (m *MockConfig) GetSendDelay() (time.Duration, error) {
	return m.GetSendDelayVal, m.GetSendDelayErr
}
func (m *MockConfig) GetTraceTimeout() (time.Duration, error) {
	return m.GetTraceTimeoutVal, m.GetTraceTimeoutErr
}

// TODO: allow per-dataset mock values
func (m *MockConfig) GetSamplerConfigForDataset(dataset string) (interface{}, error) {
	return m.GetSamplerTypeVal, m.GetSamplerTypeErr
}

func (m *MockConfig) GetUpstreamBufferSize() int {
	return m.GetUpstreamBufferSizeVal
}
func (m *MockConfig) GetPeerBufferSize() int {
	return m.GetPeerBufferSizeVal
}

func (m *MockConfig) GetIdentifierInterfaceName() (string, error) {
	return m.IdentifierInterfaceName, nil
}

func (m *MockConfig) GetUseIPV6Identifier() (bool, error) {
	return m.UseIPV6Identifier, nil
}

func (m *MockConfig) GetRedisIdentifier() (string, error) {
	return m.RedisIdentifier, nil
}

func (m *MockConfig) GetSendTickerValue() time.Duration {
	return m.SendTickerVal
}

func (m *MockConfig) GetPeerManagementType() (string, error) {
	return m.PeerManagementType, nil
}

func (m *MockConfig) GetDebugServiceAddr() (string, error) {
	return m.DebugServiceAddr, nil
}

func (m *MockConfig) GetIsDryRun() bool {
	return m.DryRun
}
