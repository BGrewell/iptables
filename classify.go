package iptables

import (
	"fmt"
)

const (
	TargetClassifyStr string = "--set-class"
)

type TargetClassify struct {
	Major int `json:"major" yaml:"major" xml:"major"`
	Minor int `json:"minor" yaml:"minor" xml:"minor"`
}

func (t TargetClassify) String() string {
	return TargetJump{
		Value: fmt.Sprintf("CLASSIFY %s %d:%d", TargetClassifyStr, t.Major, t.Minor),
	}.String()
}

func (t TargetClassify) Validate(rule Rule) error {
	return nil
}
