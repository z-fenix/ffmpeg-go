package types

//go:generate set-gen -i k8s.io/kubernetes/pkg/util/sets/types

type ReferenceSetTypes struct {
	// These types all cause files to be generated.
	// These types should be reflected in the output of
	// the "//pkg/util/sets:set-gen" genrule.
	a int64
	b int
	c byte
	d string
	e int32
}
