package gospel

import (
	"github.com/lynxbat/gospel/truths"
)

type gospel struct {
	truths *truths.Truths
}

func New() *gospel {
	return &gospel{truths: truths.New()}
}

func (g *gospel) ListTruths() map[string]truths.Truth {
	return g.truths.List()
}

func (g *gospel) GatherTruths() map[string]interface{} {
	return g.truths.Gather()
}

// Gather

// GatherByTag

// GatherByName
