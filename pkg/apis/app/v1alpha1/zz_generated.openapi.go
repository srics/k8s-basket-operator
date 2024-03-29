// +build !ignore_autogenerated

// This file was autogenerated by openapi-gen. Do not edit it manually!

package v1alpha1

import (
	spec "github.com/go-openapi/spec"
	common "k8s.io/kube-openapi/pkg/common"
)

func GetOpenAPIDefinitions(ref common.ReferenceCallback) map[string]common.OpenAPIDefinition {
	return map[string]common.OpenAPIDefinition{
		"github.com/srics/basket-operator/pkg/apis/app/v1alpha1.Basket":       schema_pkg_apis_app_v1alpha1_Basket(ref),
		"github.com/srics/basket-operator/pkg/apis/app/v1alpha1.BasketSpec":   schema_pkg_apis_app_v1alpha1_BasketSpec(ref),
		"github.com/srics/basket-operator/pkg/apis/app/v1alpha1.BasketStatus": schema_pkg_apis_app_v1alpha1_BasketStatus(ref),
	}
}

func schema_pkg_apis_app_v1alpha1_Basket(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "Basket is the Schema for the baskets API",
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/srics/basket-operator/pkg/apis/app/v1alpha1.BasketSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/srics/basket-operator/pkg/apis/app/v1alpha1.BasketStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/srics/basket-operator/pkg/apis/app/v1alpha1.BasketSpec", "github.com/srics/basket-operator/pkg/apis/app/v1alpha1.BasketStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_app_v1alpha1_BasketSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "BasketSpec defines the desired state of Basket",
				Properties:  map[string]spec.Schema{},
			},
		},
		Dependencies: []string{},
	}
}

func schema_pkg_apis_app_v1alpha1_BasketStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "BasketStatus defines the observed state of Basket",
				Properties:  map[string]spec.Schema{},
			},
		},
		Dependencies: []string{},
	}
}
