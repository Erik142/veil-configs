package config

import (
	"fmt"
	yaml "gopkg.in/yaml.v2"
)

// PKI represents the pki section of the Nebula configuration.
type PKI struct {
	CA   string `yaml:"ca"`
	Cert string `yaml:"cert"`
	Key  string `yaml:"key"`
}

// Inbound represents an inbound firewall rule.
type Inbound struct {
	Port string `yaml:"port"`
	Proto string `yaml:"proto"`
	Host string `yaml:"host"`
}

// Firewall represents the firewall section of the Nebula configuration.
type Firewall struct {
	Inbound []Inbound `yaml:"inbound"`
}

// Tun represents the tun section of the Nebula configuration.
type Tun struct {
	Dev string `yaml:"dev"`
	DropLocalBroadcast bool `yaml:"drop_local_broadcast"`
}

// Logging represents the logging section of the Nebula configuration.
type Logging struct {
	Level string `yaml:"level"`
	LogFile string `yaml:"log_file"`
}

// NebulaConfig represents the overall structure of a Nebula configuration file.
type NebulaConfig struct {
	PKI      PKI      `yaml:"pki"`
	Firewall Firewall `yaml:"firewall"`
	Tun      Tun      `yaml:"tun"`
	Logging  Logging  `yaml:"logging"`
}

// ConfigStore defines the interface for storing and retrieving Nebula configurations.
type ConfigStore interface {
	GetConfig(clientID string) (string, error)
}

// InMemoryConfigStore implements ConfigStore with an in-memory map.
type InMemoryConfigStore struct {
	configs map[string]NebulaConfig
}

// NewInMemoryConfigStore creates a new InMemoryConfigStore.
func NewInMemoryConfigStore() *InMemoryConfigStore {
	return &InMemoryConfigStore{
		configs: map[string]NebulaConfig{
			"client1": NebulaConfig{
				PKI: PKI{
					CA:   "-----BEGIN NEBULA CA CERT-----\nclient1_ca_cert_content\n-----END NEBULA CA CERT-----",
					Cert: "-----BEGIN NEBULA CERT-----\nclient1_cert_content\n-----END NEBULA CERT-----",
					Key:  "-----BEGIN NEBULA KEY-----\nclient1_key_content\n-----END NEBULA KEY-----",
				},
				Firewall: Firewall{
					Inbound: []Inbound{
						{Port: "any", Proto: "any", Host: "any"},
					},
				},
				Tun: Tun{
					Dev:                "nebula1",
					DropLocalBroadcast: true,
				},
				Logging: Logging{
					Level:   "info",
					LogFile: "/var/log/nebula.log",
				},
			},
			"client2": NebulaConfig{
				PKI: PKI{
					CA:   "-----BEGIN NEBULA CA CERT-----\nclient2_ca_cert_content\n-----END NEBULA CA CERT-----",
					Cert: "-----BEGIN NEBULA CERT-----\nclient2_cert_content\n-----END NEBULA CERT-----",
					Key:  "-----BEGIN NEBULA KEY-----\nclient2_key_content\n-----END NEBULA KEY-----",
				},
				Firewall: Firewall{
					Inbound: []Inbound{
						{Port: "any", Proto: "any", Host: "any"},
					},
				},
				Tun: Tun{
					Dev:                "nebula2",
					DropLocalBroadcast: true,
				},
				Logging: Logging{
					Level:   "info",
					LogFile: "/var/log/nebula.log",
				},
			},
			"test-client": NebulaConfig{
				PKI: PKI{
					CA:   "-----BEGIN NEBULA CA CERT-----\ntest_client_ca_cert_content\n-----END NEBULA CA CERT-----",
					Cert: "-----BEGIN NEBULA CERT-----\ntest_client_cert_content\n-----END NEBULA CERT-----",
					Key:  "-----BEGIN NEBULA KEY-----\ntest_client_key_content\n-----END NEBULA KEY-----",
				},
				Firewall: Firewall{
					Inbound: []Inbound{
						{Port: "any", Proto: "any", Host: "any"},
					},
				},
				Tun: Tun{
					Dev:                "nebula_test",
					DropLocalBroadcast: true,
				},
				Logging: Logging{
					Level:   "debug",
					LogFile: "/var/log/nebula_test.log",
				},
			},
		},
	}
}

// GetConfig retrieves the configuration for a given client ID.
func (s *InMemoryConfigStore) GetConfig(clientID string) (string, error) {
	config, ok := s.configs[clientID]
	if !ok {
		return "", fmt.Errorf("configuration not found for client ID: %s", clientID)
	}

	// Marshal the NebulaConfig struct to YAML
	yamlConfig, err := yaml.Marshal(config)
	if err != nil {
		return "", fmt.Errorf("failed to marshal config to YAML: %v", err)
	}

	return string(yamlConfig), nil
}