package custom

import (
	converters "github.com/GoogleCloudPlatform/terraform-google-conversion/v2/cai2hcl/custom/converters"
	computeConverters "github.com/GoogleCloudPlatform/terraform-google-conversion/v2/cai2hcl/custom/converters/compute"
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/cai2hcl/generated/converters/common"
	tfschema "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta"
)

var ConverterNames = map[string]string{
	computeConverters.ComputeInstanceAssetType: "google_compute_instance",
	converters.ProjectAssetType:                "google_project",
	converters.ProjectBillingAssetType:         "google_project",
}

var ConverterMap map[string]common.Converter

func init() {
	var schemaProvider = tpg.Provider()

	var factoryMap = map[string]func(schema map[string]*tfschema.Schema, name string) common.Converter{
		"google_compute_instance": computeConverters.NewComputeInstanceConverter,
		"google_project":          converters.NewProjectConverter,
	}

	ConverterMap = make(map[string]common.Converter, len(factoryMap))
	for name, factory := range factoryMap {
		schema := schemaProvider.ResourcesMap[name].Schema

		ConverterMap[name] = factory(schema, name)
	}
}
