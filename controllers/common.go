package controllers

import (
	"context"
	"fmt"
	"sync"

	"github.com/kuadrant/policy-machinery/machinery"
	"github.com/samber/lo"
	"k8s.io/apimachinery/pkg/api/meta"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
	gatewayapiv1alpha2 "sigs.k8s.io/gateway-api/apis/v1alpha2"

	kuadrantv1alpha1 "github.com/kuadrant/kuadrant-operator/api/v1alpha1"
	kuadrantv1beta3 "github.com/kuadrant/kuadrant-operator/api/v1beta3"
)

const (
	KuadrantAppName                = "kuadrant"
	PolicyAffectedConditionPattern = "kuadrant.io/%sAffected" // Policy kinds are expected to be named XPolicy
)

var (
	AppLabelKey   = "app"
	AppLabelValue = KuadrantAppName
)

func CommonLabels() map[string]string {
	return map[string]string{
		AppLabelKey:                    AppLabelValue,
		"app.kubernetes.io/component":  KuadrantAppName,
		"app.kubernetes.io/managed-by": "kuadrant-operator",
		"app.kubernetes.io/instance":   KuadrantAppName,
		"app.kubernetes.io/name":       KuadrantAppName,
		"app.kubernetes.io/part-of":    KuadrantAppName,
	}
}

func PolicyAffectedCondition(policyKind string, policies []machinery.Policy) metav1.Condition {
	condition := metav1.Condition{
		Type:   PolicyAffectedConditionType(policyKind),
		Status: metav1.ConditionTrue,
		Reason: string(gatewayapiv1alpha2.PolicyReasonAccepted),
		Message: fmt.Sprintf("Object affected by %s %s", policyKind, lo.Map(policies, func(item machinery.Policy, index int) client.ObjectKey {
			return client.ObjectKey{Name: item.GetName(), Namespace: item.GetNamespace()}
		})),
	}

	return condition
}

func PolicyAffectedConditionType(policyKind string) string {
	return fmt.Sprintf(PolicyAffectedConditionPattern, policyKind)
}

func IsPolicyAccepted(ctx context.Context, p machinery.Policy, s *sync.Map) bool {
	switch t := p.(type) {
	case *kuadrantv1beta3.AuthPolicy:
		// TODO: update when information is available in state
		return meta.IsStatusConditionTrue(t.GetStatus().GetConditions(), string(gatewayapiv1alpha2.PolicyConditionAccepted))
	case *kuadrantv1beta3.RateLimitPolicy:
		return isRateLimitPolicyAcceptedFunc(s)(p)
	case *kuadrantv1alpha1.TLSPolicy:
		isValid, _ := IsTLSPolicyValid(ctx, s, p.(*kuadrantv1alpha1.TLSPolicy))
		return isValid
	case *kuadrantv1alpha1.DNSPolicy:
		isValid, _ := dnsPolicyAcceptedStatusFunc(s)(p)
		return isValid
	default:
		return false
	}
}
