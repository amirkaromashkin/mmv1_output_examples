// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file in
//     .github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------

package compute

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/cai2hcl/cai2hcl"
	cai2hclConfig "github.com/GoogleCloudPlatform/terraform-google-conversion/v2/cai2hcl/config"
	cai2hclHelper "github.com/GoogleCloudPlatform/terraform-google-conversion/v2/cai2hcl/helper"
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/caiasset"
	tpg "github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	tfschema "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/tpgresource"
	transport_tpg "github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/transport"
	apiComputeV1 "google.golang.org/api/compute/v1"
)

// Whether the port should be set or not
func validatePortSpec(diff *schema.ResourceDiff, blockName string) error {
	block := diff.Get(blockName + ".0").(map[string]interface{})
	portSpec := block["port_specification"]
	portName := block["port_name"]
	port := block["port"]

	hasPort := (port != nil && port != 0)
	noName := (portName == nil || portName == "")

	if portSpec == "USE_NAMED_PORT" && hasPort {
		return fmt.Errorf("Error in %s: port cannot be specified when using port_specification USE_NAMED_PORT.", blockName)
	}
	if portSpec == "USE_NAMED_PORT" && noName {
		return fmt.Errorf("Error in %s: Must specify port_name when using USE_NAMED_PORT as port_specification.", blockName)
	}

	if portSpec == "USE_SERVING_PORT" && hasPort {
		return fmt.Errorf("Error in %s: port cannot be specified when using port_specification USE_SERVING_PORT.", blockName)
	}
	if portSpec == "USE_SERVING_PORT" && !noName {
		return fmt.Errorf("Error in %s: port_name cannot be specified when using port_specification USE_SERVING_PORT.", blockName)
	}

	return nil
}

func healthCheckCustomizeDiff(_ context.Context, diff *schema.ResourceDiff, v interface{}) error {
	if diff.Get("http_health_check") != nil {
		return validatePortSpec(diff, "http_health_check")
	}
	if diff.Get("https_health_check") != nil {
		return validatePortSpec(diff, "https_health_check")
	}
	if diff.Get("http2_health_check") != nil {
		return validatePortSpec(diff, "http2_health_check")
	}
	if diff.Get("tcp_health_check") != nil {
		return validatePortSpec(diff, "tcp_health_check")
	}
	if diff.Get("ssl_health_check") != nil {
		return validatePortSpec(diff, "ssl_health_check")
	}

	return nil
}

func portDiffSuppress(k, old, new string, d *schema.ResourceData) bool {
	b := strings.Split(k, ".")
	if len(b) > 2 {
		attr := b[2]

		if attr == "port" {
			var defaultPort int64

			blockType := b[0]

			switch blockType {
			case "http_health_check":
				defaultPort = 80
			case "https_health_check":
				defaultPort = 443
			case "http2_health_check":
				defaultPort = 443
			case "tcp_health_check":
				defaultPort = 80
			case "ssl_health_check":
				defaultPort = 443
			}

			oldPort, _ := strconv.Atoi(old)
			newPort, _ := strconv.Atoi(new)

			portSpec := d.Get(b[0] + ".0.port_specification")
			if int64(oldPort) == defaultPort && newPort == 0 && (portSpec == "USE_FIXED_PORT" || portSpec == "") {
				return true
			}
		}
	}

	return false
}

const ComputeHealthCheckAssetType string = "compute.googleapis.com/HealthCheck"

type ComputeHealthCheckConverter struct {
	name   string
	schema map[string]*tfschema.Schema
}

func NewComputeHealthCheckConverter(name string) cai2hcl.Converter {
	schema := tpg.Provider().ResourcesMap[name].Schema

	return &ComputeHealthCheckConverter{
		name:   name,
		schema: schema,
	}
}

func (c *ComputeHealthCheckConverter) Convert(assets []*caiasset.Asset) ([]*cai2hcl.HCLResourceBlock, error) {
	var blocks []*cai2hcl.HCLResourceBlock
	config, error := cai2hclConfig.NewConfig()
	if error != nil {
		return nil, error
	}

	for _, asset := range assets {
		if asset == nil {
			continue
		}
		if asset.Resource != nil && asset.Resource.Data != nil {
			block, err := c.convertResourceData(asset, config)
			if err != nil {
				return nil, err
			}
			blocks = append(blocks, block)
		}
	}
	return blocks, nil
}

func (c *ComputeHealthCheckConverter) convertResourceData(asset *caiasset.Asset, config *transport_tpg.Config) (*cai2hcl.HCLResourceBlock, error) {
	if asset == nil || asset.Resource == nil || asset.Resource.Data == nil {
		return nil, fmt.Errorf("asset resource data is nil")
	}

	var resource *apiComputeV1.HealthCheck
	if err := cai2hclHelper.DecodeJSON(asset.Resource.Data, &resource); err != nil {
		return nil, err
	}

	hcl, _ := resourceComputeHealthCheckRead(resource, config)

	ctyVal, err := cai2hclHelper.MapToCtyValWithSchema(hcl, c.schema)
	if err != nil {
		return nil, err
	}
	return &cai2hcl.HCLResourceBlock{
		Labels: []string{c.name, resource.Name},
		Value:  ctyVal,
	}, nil
}

func resourceComputeHealthCheckRead(res *apiComputeV1.HealthCheck, config map[string]*tfschema.Schema) (map[string]interface{}, error) {
	result := make(map[string]interface{})

	result["check_interval_sec"] = flattenComputeHealthCheckCheckIntervalSec(res.CheckIntervalSec, config)
	result["creation_timestamp"] = flattenComputeHealthCheckCreationTimestamp(res.CreationTimestamp, config)
	result["description"] = flattenComputeHealthCheckDescription(res.Description, config)
	result["healthy_threshold"] = flattenComputeHealthCheckHealthyThreshold(res.HealthyThreshold, config)
	result["name"] = flattenComputeHealthCheckName(res.Name, config)
	result["timeout_sec"] = flattenComputeHealthCheckTimeoutSec(res.TimeoutSec, config)
	result["unhealthy_threshold"] = flattenComputeHealthCheckUnhealthyThreshold(res.UnhealthyThreshold, config)
	result["type"] = flattenComputeHealthCheckType(res.Type, config)
	result["http_health_check"] = flattenComputeHealthCheckHttpHealthCheck(res.HttpHealthCheck, config)
	result["https_health_check"] = flattenComputeHealthCheckHttpsHealthCheck(res.HttpsHealthCheck, config)
	result["tcp_health_check"] = flattenComputeHealthCheckTcpHealthCheck(res.TcpHealthCheck, config)
	result["ssl_health_check"] = flattenComputeHealthCheckSslHealthCheck(res.SslHealthCheck, config)
	result["http2_health_check"] = flattenComputeHealthCheckHttp2HealthCheck(res.Http2HealthCheck, config)
	result["grpc_health_check"] = flattenComputeHealthCheckGrpcHealthCheck(res.GrpcHealthCheck, config)
	result["log_config"] = flattenComputeHealthCheckLogConfig(res.LogConfig, config)

	result["self_link"] = tpgresource.ConvertSelfLinkToV1(res.SelfLink)

	return result, nil
}

func flattenComputeHealthCheckCheckIntervalSec(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	// Handles the string fixed64 format
	if strVal, ok := v.(string); ok {
		if intVal, err := tpgresource.StringToFixed64(strVal); err == nil {
			return intVal
		}
	}

	// number values are represented as float64
	if floatVal, ok := v.(float64); ok {
		intVal := int(floatVal)
		return intVal
	}

	return v // let terraform core handle it otherwise
}

func flattenComputeHealthCheckCreationTimestamp(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeHealthCheckDescription(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeHealthCheckHealthyThreshold(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	// Handles the string fixed64 format
	if strVal, ok := v.(string); ok {
		if intVal, err := tpgresource.StringToFixed64(strVal); err == nil {
			return intVal
		}
	}

	// number values are represented as float64
	if floatVal, ok := v.(float64); ok {
		intVal := int(floatVal)
		return intVal
	}

	return v // let terraform core handle it otherwise
}

func flattenComputeHealthCheckName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeHealthCheckTimeoutSec(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	// Handles the string fixed64 format
	if strVal, ok := v.(string); ok {
		if intVal, err := tpgresource.StringToFixed64(strVal); err == nil {
			return intVal
		}
	}

	// number values are represented as float64
	if floatVal, ok := v.(float64); ok {
		intVal := int(floatVal)
		return intVal
	}

	return v // let terraform core handle it otherwise
}

func flattenComputeHealthCheckUnhealthyThreshold(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	// Handles the string fixed64 format
	if strVal, ok := v.(string); ok {
		if intVal, err := tpgresource.StringToFixed64(strVal); err == nil {
			return intVal
		}
	}

	// number values are represented as float64
	if floatVal, ok := v.(float64); ok {
		intVal := int(floatVal)
		return intVal
	}

	return v // let terraform core handle it otherwise
}

func flattenComputeHealthCheckType(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeHealthCheckHttpHealthCheck(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["host"] =
		flattenComputeHealthCheckHttpHealthCheckHost(original["host"], d, config)
	transformed["request_path"] =
		flattenComputeHealthCheckHttpHealthCheckRequestPath(original["requestPath"], d, config)
	transformed["response"] =
		flattenComputeHealthCheckHttpHealthCheckResponse(original["response"], d, config)
	transformed["port"] =
		flattenComputeHealthCheckHttpHealthCheckPort(original["port"], d, config)
	transformed["port_name"] =
		flattenComputeHealthCheckHttpHealthCheckPortName(original["portName"], d, config)
	transformed["proxy_header"] =
		flattenComputeHealthCheckHttpHealthCheckProxyHeader(original["proxyHeader"], d, config)
	transformed["port_specification"] =
		flattenComputeHealthCheckHttpHealthCheckPortSpecification(original["portSpecification"], d, config)
	return []interface{}{transformed}
}
func flattenComputeHealthCheckHttpHealthCheckHost(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeHealthCheckHttpHealthCheckRequestPath(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeHealthCheckHttpHealthCheckResponse(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeHealthCheckHttpHealthCheckPort(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	// Handles the string fixed64 format
	if strVal, ok := v.(string); ok {
		if intVal, err := tpgresource.StringToFixed64(strVal); err == nil {
			return intVal
		}
	}

	// number values are represented as float64
	if floatVal, ok := v.(float64); ok {
		intVal := int(floatVal)
		return intVal
	}

	return v // let terraform core handle it otherwise
}

func flattenComputeHealthCheckHttpHealthCheckPortName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeHealthCheckHttpHealthCheckProxyHeader(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeHealthCheckHttpHealthCheckPortSpecification(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeHealthCheckHttpsHealthCheck(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["host"] =
		flattenComputeHealthCheckHttpsHealthCheckHost(original["host"], d, config)
	transformed["request_path"] =
		flattenComputeHealthCheckHttpsHealthCheckRequestPath(original["requestPath"], d, config)
	transformed["response"] =
		flattenComputeHealthCheckHttpsHealthCheckResponse(original["response"], d, config)
	transformed["port"] =
		flattenComputeHealthCheckHttpsHealthCheckPort(original["port"], d, config)
	transformed["port_name"] =
		flattenComputeHealthCheckHttpsHealthCheckPortName(original["portName"], d, config)
	transformed["proxy_header"] =
		flattenComputeHealthCheckHttpsHealthCheckProxyHeader(original["proxyHeader"], d, config)
	transformed["port_specification"] =
		flattenComputeHealthCheckHttpsHealthCheckPortSpecification(original["portSpecification"], d, config)
	return []interface{}{transformed}
}
func flattenComputeHealthCheckHttpsHealthCheckHost(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeHealthCheckHttpsHealthCheckRequestPath(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeHealthCheckHttpsHealthCheckResponse(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeHealthCheckHttpsHealthCheckPort(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	// Handles the string fixed64 format
	if strVal, ok := v.(string); ok {
		if intVal, err := tpgresource.StringToFixed64(strVal); err == nil {
			return intVal
		}
	}

	// number values are represented as float64
	if floatVal, ok := v.(float64); ok {
		intVal := int(floatVal)
		return intVal
	}

	return v // let terraform core handle it otherwise
}

func flattenComputeHealthCheckHttpsHealthCheckPortName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeHealthCheckHttpsHealthCheckProxyHeader(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeHealthCheckHttpsHealthCheckPortSpecification(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeHealthCheckTcpHealthCheck(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["request"] =
		flattenComputeHealthCheckTcpHealthCheckRequest(original["request"], d, config)
	transformed["response"] =
		flattenComputeHealthCheckTcpHealthCheckResponse(original["response"], d, config)
	transformed["port"] =
		flattenComputeHealthCheckTcpHealthCheckPort(original["port"], d, config)
	transformed["port_name"] =
		flattenComputeHealthCheckTcpHealthCheckPortName(original["portName"], d, config)
	transformed["proxy_header"] =
		flattenComputeHealthCheckTcpHealthCheckProxyHeader(original["proxyHeader"], d, config)
	transformed["port_specification"] =
		flattenComputeHealthCheckTcpHealthCheckPortSpecification(original["portSpecification"], d, config)
	return []interface{}{transformed}
}
func flattenComputeHealthCheckTcpHealthCheckRequest(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeHealthCheckTcpHealthCheckResponse(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeHealthCheckTcpHealthCheckPort(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	// Handles the string fixed64 format
	if strVal, ok := v.(string); ok {
		if intVal, err := tpgresource.StringToFixed64(strVal); err == nil {
			return intVal
		}
	}

	// number values are represented as float64
	if floatVal, ok := v.(float64); ok {
		intVal := int(floatVal)
		return intVal
	}

	return v // let terraform core handle it otherwise
}

func flattenComputeHealthCheckTcpHealthCheckPortName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeHealthCheckTcpHealthCheckProxyHeader(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeHealthCheckTcpHealthCheckPortSpecification(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeHealthCheckSslHealthCheck(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["request"] =
		flattenComputeHealthCheckSslHealthCheckRequest(original["request"], d, config)
	transformed["response"] =
		flattenComputeHealthCheckSslHealthCheckResponse(original["response"], d, config)
	transformed["port"] =
		flattenComputeHealthCheckSslHealthCheckPort(original["port"], d, config)
	transformed["port_name"] =
		flattenComputeHealthCheckSslHealthCheckPortName(original["portName"], d, config)
	transformed["proxy_header"] =
		flattenComputeHealthCheckSslHealthCheckProxyHeader(original["proxyHeader"], d, config)
	transformed["port_specification"] =
		flattenComputeHealthCheckSslHealthCheckPortSpecification(original["portSpecification"], d, config)
	return []interface{}{transformed}
}
func flattenComputeHealthCheckSslHealthCheckRequest(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeHealthCheckSslHealthCheckResponse(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeHealthCheckSslHealthCheckPort(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	// Handles the string fixed64 format
	if strVal, ok := v.(string); ok {
		if intVal, err := tpgresource.StringToFixed64(strVal); err == nil {
			return intVal
		}
	}

	// number values are represented as float64
	if floatVal, ok := v.(float64); ok {
		intVal := int(floatVal)
		return intVal
	}

	return v // let terraform core handle it otherwise
}

func flattenComputeHealthCheckSslHealthCheckPortName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeHealthCheckSslHealthCheckProxyHeader(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeHealthCheckSslHealthCheckPortSpecification(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeHealthCheckHttp2HealthCheck(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["host"] =
		flattenComputeHealthCheckHttp2HealthCheckHost(original["host"], d, config)
	transformed["request_path"] =
		flattenComputeHealthCheckHttp2HealthCheckRequestPath(original["requestPath"], d, config)
	transformed["response"] =
		flattenComputeHealthCheckHttp2HealthCheckResponse(original["response"], d, config)
	transformed["port"] =
		flattenComputeHealthCheckHttp2HealthCheckPort(original["port"], d, config)
	transformed["port_name"] =
		flattenComputeHealthCheckHttp2HealthCheckPortName(original["portName"], d, config)
	transformed["proxy_header"] =
		flattenComputeHealthCheckHttp2HealthCheckProxyHeader(original["proxyHeader"], d, config)
	transformed["port_specification"] =
		flattenComputeHealthCheckHttp2HealthCheckPortSpecification(original["portSpecification"], d, config)
	return []interface{}{transformed}
}
func flattenComputeHealthCheckHttp2HealthCheckHost(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeHealthCheckHttp2HealthCheckRequestPath(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeHealthCheckHttp2HealthCheckResponse(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeHealthCheckHttp2HealthCheckPort(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	// Handles the string fixed64 format
	if strVal, ok := v.(string); ok {
		if intVal, err := tpgresource.StringToFixed64(strVal); err == nil {
			return intVal
		}
	}

	// number values are represented as float64
	if floatVal, ok := v.(float64); ok {
		intVal := int(floatVal)
		return intVal
	}

	return v // let terraform core handle it otherwise
}

func flattenComputeHealthCheckHttp2HealthCheckPortName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeHealthCheckHttp2HealthCheckProxyHeader(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeHealthCheckHttp2HealthCheckPortSpecification(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeHealthCheckGrpcHealthCheck(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["port"] =
		flattenComputeHealthCheckGrpcHealthCheckPort(original["port"], d, config)
	transformed["port_name"] =
		flattenComputeHealthCheckGrpcHealthCheckPortName(original["portName"], d, config)
	transformed["port_specification"] =
		flattenComputeHealthCheckGrpcHealthCheckPortSpecification(original["portSpecification"], d, config)
	transformed["grpc_service_name"] =
		flattenComputeHealthCheckGrpcHealthCheckGrpcServiceName(original["grpcServiceName"], d, config)
	return []interface{}{transformed}
}
func flattenComputeHealthCheckGrpcHealthCheckPort(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	// Handles the string fixed64 format
	if strVal, ok := v.(string); ok {
		if intVal, err := tpgresource.StringToFixed64(strVal); err == nil {
			return intVal
		}
	}

	// number values are represented as float64
	if floatVal, ok := v.(float64); ok {
		intVal := int(floatVal)
		return intVal
	}

	return v // let terraform core handle it otherwise
}

func flattenComputeHealthCheckGrpcHealthCheckPortName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeHealthCheckGrpcHealthCheckPortSpecification(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeHealthCheckGrpcHealthCheckGrpcServiceName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeHealthCheckLogConfig(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	transformed := make(map[string]interface{})
	if v == nil {
		// Disabled by default, but API will not return object if value is false
		transformed["enable"] = false
		return []interface{}{transformed}
	}

	original := v.(map[string]interface{})
	transformed["enable"] = original["enable"]
	return []interface{}{transformed}
}