package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewInMemoryConfigStore(t *testing.T) {
	store := NewInMemoryConfigStore()
	assert.NotNil(t, store)
	assert.NotNil(t, store.configs)
	assert.Equal(t, 3, len(store.configs))
}

func TestInMemoryConfigStore_GetConfig(t *testing.T) {
	store := NewInMemoryConfigStore()

	// Test existing client
	config, err := store.GetConfig("client1")
	assert.NoError(t, err)
	assert.Contains(t, config, "client1_ca_cert_content")
	assert.Contains(t, config, "client1_cert_content")
	assert.Contains(t, config, "client1_key_content")
	assert.Contains(t, config, "dev: nebula1")

	// Test non-existing client
	config, err = store.GetConfig("nonexistent_client")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "configuration not found")
	assert.Empty(t, config)

	// Test YAML marshalling
	expectedYAML := `pki:
  ca: |-
    -----BEGIN NEBULA CA CERT-----
    client1_ca_cert_content
    -----END NEBULA CA CERT-----
  cert: |-
    -----BEGIN NEBULA CERT-----
    client1_cert_content
    -----END NEBULA CERT-----
  key: |-
    -----BEGIN NEBULA KEY-----
    client1_key_content
    -----END NEBULA KEY-----
firewall:
  inbound:
  - port: any
    proto: any
    host: any
tun:
  dev: nebula1
  drop_local_broadcast: true
logging:
  level: info
  log_file: /var/log/nebula.log
`
	config, err = store.GetConfig("client1")
	assert.NoError(t, err)
	assert.Equal(t, expectedYAML, config)

	// Test existing test-client ID
	config, err = store.GetConfig("test-client")
	assert.NoError(t, err)
	assert.Contains(t, config, "test_client_ca_cert_content")
	assert.Contains(t, config, "test_client_cert_content")
	assert.Contains(t, config, "test_client_key_content")
	assert.Contains(t, config, "dev: nebula_test")

	// Test YAML marshalling for test-client
	expectedTestClientYAML := `pki:
  ca: |-
    -----BEGIN NEBULA CA CERT-----
    test_client_ca_cert_content
    -----END NEBULA CA CERT-----
  cert: |-
    -----BEGIN NEBULA CERT-----
    test_client_cert_content
    -----END NEBULA CERT-----
  key: |-
    -----BEGIN NEBULA KEY-----
    test_client_key_content
    -----END NEBULA KEY-----
firewall:
  inbound:
  - port: any
    proto: any
    host: any
tun:
  dev: nebula_test
  drop_local_broadcast: true
logging:
  level: debug
  log_file: /var/log/nebula_test.log
`
	config, err = store.GetConfig("test-client")
	assert.NoError(t, err)
	assert.Equal(t, expectedTestClientYAML, config)
}
