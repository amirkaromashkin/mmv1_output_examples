// Package cai2hcl converts CAI assets to hcl bytes.
package cai2hcl

import (
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/cai2hcl/custom"
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/cai2hcl/generated"
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/cai2hcl/generated/converters/common"
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/caiasset"
)

func Convert(assets []*caiasset.Asset, options *generated.ConvertOptions) ([]byte, error) {
	// Map which contains mapping between CAI Asset type and TF resource name.
	converterNames := mergeNameMaps(custom.ConverterNames, generated.ConverterNames)

	// Map which contains converters by their TF resource name.
	converterMap := mergeConverterMaps(custom.ConverterMap, generated.ConverterMap)

	return generated.Convert(assets, converterNames, converterMap, options)
}

func mergeNameMaps(m1 map[string]string, m2 map[string]string) map[string]string {
	merged := make(map[string]string)
	for k, v := range m1 {
		merged[k] = v
	}
	for key, value := range m2 {
		merged[key] = value
	}
	return merged
}

func mergeConverterMaps(m1 map[string]common.Converter, m2 map[string]common.Converter) map[string]common.Converter {
	merged := make(map[string]common.Converter)
	for k, v := range m1 {
		merged[k] = v
	}
	for key, value := range m2 {
		merged[key] = value
	}
	return merged
}
