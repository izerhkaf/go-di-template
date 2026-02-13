package config

import "github.com/google/wire"

var ProviderSet = wire.NewSet(
	NewConfig,
)

var ProviderSetTest = wire.NewSet(
	NewConfigTest,
)
