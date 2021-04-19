package iptables

import (
	"fmt"
	"strings"
)

const (
	TargetSNatStr string = "--to-source"
)

type TargetSNat struct {
	SourceIp        string `json:"source_ip" yaml:"source_ip" xml:"source_ip"`
	SourceIpRange   string `json:"source_ip_range" yaml:"source_ip_range" xml:"source_ip_range"`
	SourcePort      string `json:"source_port" yaml:"source_port" xml:"source_port"`
	SourcePortRange string `json:"source_port_range" yaml:"source_port_range" xml:"source_port_range"`
}

func (t TargetSNat) String() string {
	parts := make([]string, 0)
	parts = append(parts, "SNAT")
	parts = append(parts, TargetSNatStr)
	if t.SourceIpRange != "" {
		parts = append(parts, t.SourceIpRange)
	} else {
		parts = append(parts, t.SourceIp)
	}
	if t.SourcePortRange != "" {
		parts = append(parts, fmt.Sprintf(":%s", t.SourcePortRange))
	} else if t.SourcePort != "" {
		parts = append(parts, fmt.Sprintf(":%s", t.SourcePort))
	}

	return TargetJump{
		Value: strings.Join(parts, " "),
	}.String()
}

// Returns if the target is valid when applied with the specified rule
func (t TargetSNat) Validate(rule Rule) error {
	// Only valid on the mangle table
	if rule.Table != TableNat {
		return fmt.Errorf("target SNAT is only valid on the 'nat' table")
	}
	if rule.Chain != ChainPostRouting {
		return fmt.Errorf("target SNAT is only valid on the 'POSTROUTING' chain") //TODO: need more smarts as this is valid if it's on a chain that is jumped to from POSTROUTING I believe
	}
	return nil
}
