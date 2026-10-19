package cli

import (
	"strings"
	"testing"

	"github.com/growthbook/cli/v3/internal/sdk/models/components"
)

func TestTSType(t *testing.T) {
	cases := map[components.FeatureV2ValueType]string{
		components.FeatureV2ValueTypeBoolean: "boolean",
		components.FeatureV2ValueTypeNumber:  "number",
		components.FeatureV2ValueTypeString:  "string",
		components.FeatureV2ValueTypeJSON:    "Record<string, unknown> | unknown[]",
		"unexpected":                         "unknown",
	}
	for vt, want := range cases {
		if got := tsType(vt); got != want {
			t.Errorf("tsType(%q) = %q, want %q", vt, got, want)
		}
	}
}

func TestRenderAppFeaturesTSSortsAndTypesFeatures(t *testing.T) {
	out := renderAppFeaturesTS([]components.FeatureV2{
		{ID: "b-config", ValueType: components.FeatureV2ValueTypeJSON},
		{ID: "a-enabled", ValueType: components.FeatureV2ValueTypeBoolean},
	})
	want := "export type AppFeatures = {\n" +
		"  'a-enabled': boolean;\n" +
		"  'b-config': Record<string, unknown> | unknown[];\n" +
		"}\n"
	if !strings.HasSuffix(out, want) {
		t.Errorf("renderAppFeaturesTS output does not end with:\n%s\ngot:\n%s", want, out)
	}
}
