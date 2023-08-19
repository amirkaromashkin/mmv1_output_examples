package cai2hcl

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/cai2hcl/generated"
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/caiasset"

	"github.com/google/go-cmp/cmp"
	"go.uber.org/zap"
)

type TestCase struct {
	name         string
	sourceFolder string
}

func TestCai2HclConvert(t *testing.T) {
	cases := []TestCase{}

	for _, name := range getCustomConvertersTestFileNames() {
		cases = append(cases, TestCase{name: name, sourceFolder: "../testdata"})
	}

	for _, name := range getGeneratedConvertersTestFileNames() {
		cases = append(cases, TestCase{name: name, sourceFolder: "./generated/converters/testdata"})
	}

	for i := range cases {
		c := cases[i]

		t.Run(c.name, func(t *testing.T) {
			t.Parallel()

			err := assertTestData(c)
			if err != nil {
				t.Fatal(err)
			}
		})
	}
}

func getCustomConvertersTestFileNames() []string {
	return []string{
		"full_compute_instance",
		"compute_instance_iam",
		"project_create",
		"project_iam",
	}
}

func getGeneratedConvertersTestFileNames() []string {
	generatedTestFiles, err := os.ReadDir("./generated/converters/testdata")
	if err != nil {
		log.Fatal(err)
	}

	namesArr := []string{}
	namesSet := make(map[string]bool)
	for _, file := range generatedTestFiles {
		name := file.Name()
		if _, ok := namesSet[name]; !ok {
			namesArr = append(namesArr, fileNameWithoutExt(file.Name()))
		}

		namesSet[name] = true
	}

	return namesArr
}

func fileNameWithoutExt(fileName string) string {
	return strings.TrimSuffix(fileName, filepath.Ext(fileName))
}

func assertTestData(testCase TestCase) (err error) {
	fileName := testCase.name
	folder := testCase.sourceFolder

	assetFilePath := fmt.Sprintf("%s/%s.json", folder, fileName)
	expectedTfFilePath := fmt.Sprintf("%s/%s.tf", folder, fileName)
	assetPayload, err := os.ReadFile(assetFilePath)
	if err != nil {
		return fmt.Errorf("cannot open %s, got: %s", assetFilePath, err)
	}
	want, err := os.ReadFile(expectedTfFilePath)
	if err != nil {
		return fmt.Errorf("cannot open %s, got: %s", expectedTfFilePath, err)
	}

	var assets []*caiasset.Asset
	if err := json.Unmarshal(assetPayload, &assets); err != nil {
		return fmt.Errorf("cannot unmarshal: %s", err)
	}

	logger, err := zap.NewDevelopment()

	got, err := Convert(assets, &generated.ConvertOptions{
		ErrorLogger: logger,
	})
	if err != nil {
		return err
	}
	if diff := cmp.Diff(string(want), string(got)); diff != "" {
		return fmt.Errorf("cmp.Diff() got diff (-want +got): %s", diff)
	}

	return nil
}
