package parse_spec_bench

import (
	"encoding/json"
	"github.com/containerd/containerd/oci"
	"os"
	"testing"
)

// ReadSpec deserializes JSON into an OCI runtime Spec from a given path.
func ReadSpec(path string) (*oci.Spec, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	var s oci.Spec
	if err := json.NewDecoder(f).Decode(&s); err != nil {
		return nil, err
	}
	return &s, nil
}

func BenchmarkParseFullSpec(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := ReadSpec("config.json")
		if err != nil {
			b.Error(err)
		}
	}
}

type annoSpec struct {
	Annotations map[string]string `json:"annotations,omitempty"`
}

func readAnnotations(path string) (*annoSpec, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	var s annoSpec
	if err := json.NewDecoder(f).Decode(&s); err != nil {
		return nil, err
	}
	return &s, nil
}

func BenchmarkParseAnnotations(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := readAnnotations("config.json")
		if err != nil {
			b.Error(err)
		}
	}
}

type hookSpec struct {
	Root struct {
		Path string `json:"path"`
	} `json:"root"`
}

func readHookSpec(path string) (*hookSpec, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	var s hookSpec
	if err := json.NewDecoder(f).Decode(&s); err != nil {
		return nil, err
	}
	return &s, nil
}

func BenchmarkParseHooks(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := readHookSpec("config.json")
		if err != nil {
			b.Error(err)
		}
	}
}
