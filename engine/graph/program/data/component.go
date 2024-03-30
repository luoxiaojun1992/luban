package data

import "github.com/luoxiaojun1992/luban/engine/graph/program/node/component"

type Component struct {
	ComponentType component.ComponentType `json:"component_type"`
	Attrs         map[string]interface{}  `json:"attrs"`
}
