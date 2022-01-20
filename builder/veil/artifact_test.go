package veil

import (
	"github.com/stretchr/testify/require"
	"testing"

	packersdk "github.com/hashicorp/packer-plugin-sdk/packer"
	"github.com/jsc-masshtab/veil-api-client-go/veil"
)

const templateID = "286dd44a-ec6b-4789-b192-804f08f04b4c"

func TestArtifact_Impl(t *testing.T) {
	var raw interface{} = &Artifact{}

	if _, ok := raw.(packersdk.Artifact); !ok {
		t.Fatalf("Artifact does not implement packersdk.Artifact")
	}
}

func TestArtifactId(t *testing.T) {
	//a := &Artifact{
	//	client: nil,
	//	config: nil,
	//	template: &veil.CreateTemplateResponse{
	//		Id: "286dd44a-ec6b-4789-b192-804f08f04b4c",
	//	},
	//}
	client := veil.NewClient("", "", false)
	config := new(veil.DomainCreateConfig)
	config.DomainId = templateID
	config.VerboseName = veil.NameGenerator("domain")
	config.MemoryCount = 50
	domain, _, err := client.Domain.Create(*config)
	require.Nil(t, err)
	a := &Artifact{
		client:   client,
		config:   nil,
		template: domain,
	}
	if a.Id() != templateID {
		t.Fatalf("artifact ID should match: %s", templateID)
	}
}

func TestArtifactString(t *testing.T) {
	//a := &Artifact{
	//	client: nil,
	//	config: nil,
	//	template: &veil.CreateTemplateResponse{
	//		Name: "packer-foobar",
	//	},
	//}
	client := veil.NewClient("", "", false)
	config := new(veil.DomainCreateConfig)
	config.DomainId = templateID
	config.VerboseName = "packer-foobar"
	config.MemoryCount = 50
	domain, _, err := client.Domain.Create(*config)
	require.Nil(t, err)
	expected := "A template was created: packer-foobar"
	a := &Artifact{
		client:   nil,
		config:   nil,
		template: domain,
	}
	if a.String() != expected {
		t.Fatalf("artifact string should match: %s", expected)
	}
}

func TestArtifactState_StateData(t *testing.T) {
	expectedData := "this is the data"
	artifact := &Artifact{
		StateData: map[string]interface{}{"state_data": expectedData},
	}

	// Valid state
	result := artifact.State("state_data")
	if result != expectedData {
		t.Fatalf("Bad: State data was %s instead of %s", result, expectedData)
	}

	// Invalid state
	result = artifact.State("invalid_key")
	if result != nil {
		t.Fatalf("Bad: State should be nil for invalid state data name")
	}

	// Nil StateData should not fail and should return nil
	artifact = &Artifact{}
	result = artifact.State("key")
	if result != nil {
		t.Fatalf("Bad: State should be nil for nil StateData")
	}
}
