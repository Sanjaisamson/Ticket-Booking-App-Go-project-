// mock_ctx.go
package mock

import (
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/mock"
)

type MockFiberCtx struct {
	mock.Mock
}

func (m *MockFiberCtx) Status(status int) *fiber.Ctx {
	args := m.Called(status)
	return args.Get(0).(*fiber.Ctx)
}

// Implement other fiber.Ctx interface methods here

type MockCtx struct {
	FiberCtx *MockFiberCtx
	mock.Mock
}

func (m *MockCtx) BodyParser(obj interface{}) error {
	args := m.Called(obj)
	return args.Error(0)
}

func (m *MockCtx) JSON(obj interface{}) error {
	args := m.Called(obj)
	return args.Error(0)
}
