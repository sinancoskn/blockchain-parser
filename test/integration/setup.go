package integration

import (
	"blockchain-parser/internal/app"

	"go.uber.org/dig"
)

const (
	TEST_ADDRESS_1 = "0x3f4f49de2a6108ce51b6a9489278dd5c57baa6b6"
	TEST_ADDRESS_2 = "0x411ee650a394b22a1d684834f2728b6b71e0fe50"
)

var (
	Container *dig.Container
)

func SetupTest() {
	if Container == nil {
		Container = app.InitializeContainer()
		if Container == nil {
			panic("Failed to initialize DI container")
		}
	}
}
