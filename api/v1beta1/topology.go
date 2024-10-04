package v1beta1

import (
	authorinov1beta1 "github.com/kuadrant/authorino-operator/api/v1beta1"
	limitadorv1alpha1 "github.com/kuadrant/limitador-operator/api/v1alpha1"
	"github.com/kuadrant/policy-machinery/controller"
	"github.com/kuadrant/policy-machinery/machinery"
	"github.com/samber/lo"
	"k8s.io/apimachinery/pkg/runtime/schema"
	gwapiv1 "sigs.k8s.io/gateway-api/apis/v1"
)

var (
	AuthorinoGroupKind = schema.GroupKind{Group: authorinov1beta1.GroupVersion.Group, Kind: "Authorino"}
	KuadrantGroupKind  = schema.GroupKind{Group: GroupVersion.Group, Kind: "Kuadrant"}
	LimitadorGroupKind = schema.GroupKind{Group: limitadorv1alpha1.GroupVersion.Group, Kind: "Limitador"}

	AuthorinosResource = authorinov1beta1.GroupVersion.WithResource("authorinos")
	KuadrantsResource  = GroupVersion.WithResource("kuadrants")
	LimitadorsResource = limitadorv1alpha1.GroupVersion.WithResource("limitadors")
)

var _ machinery.Object = &Kuadrant{}

func (p *Kuadrant) GetLocator() string {
	return machinery.LocatorFromObject(p)
}

func LinkKuadrantToGatewayClasses(objs controller.Store) machinery.LinkFunc {
	kuadrants := lo.Map(objs.FilterByGroupKind(KuadrantGroupKind), controller.ObjectAs[*Kuadrant])

	return machinery.LinkFunc{
		From: KuadrantGroupKind,
		To:   schema.GroupKind{Group: gwapiv1.GroupVersion.Group, Kind: "GatewayClass"},
		Func: func(_ machinery.Object) []machinery.Object {
			parents := make([]machinery.Object, len(kuadrants))
			for _, parent := range kuadrants {
				parents = append(parents, parent)
			}
			return parents
		},
	}
}

func LinkKuadrantToLimitador(objs controller.Store) machinery.LinkFunc {
	kuadrants := lo.Map(objs.FilterByGroupKind(KuadrantGroupKind), controller.ObjectAs[machinery.Object])

	return machinery.LinkFunc{
		From: KuadrantGroupKind,
		To:   LimitadorGroupKind,
		Func: func(child machinery.Object) []machinery.Object {
			return lo.Filter(kuadrants, func(kuadrant machinery.Object, _ int) bool {
				return kuadrant.GetNamespace() == child.GetNamespace() && child.GetName() == "limitador"
			})
		},
	}
}

func LinkKuadrantToAuthorino(objs controller.Store) machinery.LinkFunc {
	kuadrants := lo.Map(objs.FilterByGroupKind(KuadrantGroupKind), controller.ObjectAs[machinery.Object])

	return machinery.LinkFunc{
		From: KuadrantGroupKind,
		To:   AuthorinoGroupKind,
		Func: func(child machinery.Object) []machinery.Object {
			return lo.Filter(kuadrants, func(kuadrant machinery.Object, _ int) bool {
				return kuadrant.GetNamespace() == child.GetNamespace() && child.GetName() == "authorino"
			})
		},
	}
}
