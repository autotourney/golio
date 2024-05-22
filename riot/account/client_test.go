package account

import (
	"testing"

	"github.com/sirupsen/logrus"

	"github.com/autotourney/golio/api"
	"github.com/autotourney/golio/internal"
	"github.com/autotourney/golio/internal/mock"
)

func TestNewClient(t *testing.T) {
	c := NewClient(internal.NewClient(api.RegionBrasil, "key", mock.NewStatusMockDoer(200), logrus.StandardLogger()))
	if c == nil {
		t.Error("returned nil")
	}
}
