package common

import (
	"encoding/json"
	"fmt"
	"strings"

	hashicorpcty "github.com/hashicorp/go-cty/cty"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
	"github.com/zclconf/go-cty/cty"
	ctyjson "github.com/zclconf/go-cty/cty/json"
)

func parseFieldValue(str string, field string) string {
	strList := strings.Split(str, "/")
	for ix, item := range strList {
		if item == field && ix+1 < len(strList) {
			return strList[ix+1]
		}
	}
	return ""
}

func ConvertSchemaSetToArray(schemaSet *schema.Set) []interface{} {
	return schemaSet.List()
}

func ParseFieldValue(str string, field string) string {
	return parseFieldValue(str, field)
}

func decodeJSON(data map[string]interface{}, v interface{}) error {
	b, err := json.Marshal(data)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	return nil
}

// Decodes the map object into the target struct.
func DecodeJSON(data map[string]interface{}, v interface{}) error {
	return decodeJSON(data, v)
}

func normalize(obj interface{}) interface{} {
	switch obj.(type) {
	case []interface{}:
		arr := obj.([]interface{})
		newArr := make([]interface{}, len(arr))
		for i := range arr {
			newArr[i] = normalize(arr[i])
		}

		return newArr
	case map[string]interface{}:
		mp := obj.(map[string]interface{})
		newMap := map[string]interface{}{}
		for key, value := range mp {
			newMap[key] = normalize(value)
		}
		return newMap
	case *schema.Set:
		return obj.(*schema.Set).List()
	default:
		return obj
	}
}

func mapToCtyValWithSchema(m map[string]interface{}, s map[string]*schema.Schema) (cty.Value, error) {
	m = normalize(m).(map[string]interface{})

	b, err := json.Marshal(&m)
	if err != nil {
		return cty.NilVal, fmt.Errorf("error marshaling map as JSON: %v", err)
	}
	ty, err := hashicorpCtyTypeToZclconfCtyType(schema.InternalMap(s).CoreConfigSchema().ImpliedType())
	if err != nil {
		return cty.NilVal, fmt.Errorf("error casting type: %v", err)
	}
	ret, err := ctyjson.Unmarshal(b, ty)
	if err != nil {
		return cty.NilVal, fmt.Errorf("error unmarshaling JSON as cty.Value: %v", err)
	}

	return ret, nil
}

func MapToCtyValWithSchema(m map[string]interface{}, s map[string]*schema.Schema) (cty.Value, error) {
	return mapToCtyValWithSchema(m, s)
}

func hashicorpCtyTypeToZclconfCtyType(t hashicorpcty.Type) (cty.Type, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return cty.NilType, err
	}
	var ret cty.Type
	if err := json.Unmarshal(b, &ret); err != nil {
		return cty.NilType, err
	}
	return ret, nil
}

func hashicorpCtyTypeToZclconfCtyValue(t hashicorpcty.Value) (cty.Value, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return cty.NilVal, err
	}
	var ret cty.Value
	if err := json.Unmarshal(b, &ret); err != nil {
		return cty.NilVal, err
	}
	return ret, nil
}

func NewConfig() *transport_tpg.Config {
	return &transport_tpg.Config{
		Project:   "",
		Zone:      "",
		Region:    "",
		UserAgent: "",
	}
}