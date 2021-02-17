// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

// Code generated by beats/dev-tools/cmd/asset/asset.go - DO NOT EDIT.

package system

import (
	"github.com/elastic/beats/v7/libbeat/asset"
)

func init() {
	if err := asset.SetFields("auditbeat", "system", asset.ModuleFieldsPri, AssetSystem); err != nil {
		panic(err)
	}
}

// AssetSystem returns asset data.
// This is the base64 encoded gzipped contents of module/system.
func AssetSystem() string {
	return "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"
}
