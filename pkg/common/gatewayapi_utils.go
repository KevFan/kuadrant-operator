package common

import (
	"context"
	"fmt"
	"reflect"
	"strings"

	"github.com/kuadrant/kuadrant-operator/pkg/library/common"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/api/meta"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/utils/ptr"
	"sigs.k8s.io/controller-runtime/pkg/client"
	gatewayapiv1alpha2 "sigs.k8s.io/gateway-api/apis/v1alpha2"
	gatewayapiv1beta1 "sigs.k8s.io/gateway-api/apis/v1beta1"
)

type HTTPRouteRule struct {
	Paths   []string
	Methods []string
	Hosts   []string
}

func IsTargetRefHTTPRoute(targetRef gatewayapiv1alpha2.PolicyTargetReference) bool {
	return targetRef.Kind == ("HTTPRoute")
}

func IsTargetRefGateway(targetRef gatewayapiv1alpha2.PolicyTargetReference) bool {
	return targetRef.Kind == ("Gateway")
}

func RouteHTTPMethodToRuleMethod(httpMethod *gatewayapiv1beta1.HTTPMethod) []string {
	if httpMethod == nil {
		return nil
	}

	return []string{string(*httpMethod)}
}

func RouteHostnames(route *gatewayapiv1beta1.HTTPRoute) []string {
	if route == nil {
		return nil
	}

	if len(route.Spec.Hostnames) == 0 {
		return []string{"*"}
	}

	hosts := make([]string, 0, len(route.Spec.Hostnames))

	for _, hostname := range route.Spec.Hostnames {
		hosts = append(hosts, string(hostname))
	}

	return hosts
}

// RulesFromHTTPRoute computes a list of rules from the HTTPRoute object
func RulesFromHTTPRoute(route *gatewayapiv1beta1.HTTPRoute) []HTTPRouteRule {
	if route == nil {
		return nil
	}

	var rules []HTTPRouteRule

	for routeRuleIdx := range route.Spec.Rules {
		for matchIdx := range route.Spec.Rules[routeRuleIdx].Matches {
			match := &route.Spec.Rules[routeRuleIdx].Matches[matchIdx]

			rule := HTTPRouteRule{
				Hosts:   RouteHostnames(route),
				Methods: RouteHTTPMethodToRuleMethod(match.Method),
				Paths:   routePathMatchToRulePath(match.Path),
			}

			if len(rule.Methods) != 0 || len(rule.Paths) != 0 {
				// Only append rule when there are methods or path rules
				// a valid rule must include HTTPRoute hostnames as well
				rules = append(rules, rule)
			}
		}
	}

	// If no rules compiled from the route, at least one rule for the hosts
	if len(rules) == 0 {
		rules = []HTTPRouteRule{{Hosts: RouteHostnames(route)}}
	}

	return rules
}

type HTTPRouteRuleSelector struct {
	*gatewayapiv1beta1.HTTPRouteMatch
}

func (s *HTTPRouteRuleSelector) Selects(rule gatewayapiv1beta1.HTTPRouteRule) bool {
	if s.HTTPRouteMatch == nil {
		return true
	}

	_, found := common.Find(rule.Matches, func(ruleMatch gatewayapiv1beta1.HTTPRouteMatch) bool {
		// path
		if s.Path != nil && !reflect.DeepEqual(s.Path, ruleMatch.Path) {
			return false
		}

		// method
		if s.Method != nil && !reflect.DeepEqual(s.Method, ruleMatch.Method) {
			return false
		}

		// headers
		for _, header := range s.Headers {
			if _, found := common.Find(ruleMatch.Headers, func(otherHeader gatewayapiv1beta1.HTTPHeaderMatch) bool {
				return reflect.DeepEqual(header, otherHeader)
			}); !found {
				return false
			}
		}

		// query params
		for _, param := range s.QueryParams {
			if _, found := common.Find(ruleMatch.QueryParams, func(otherParam gatewayapiv1beta1.HTTPQueryParamMatch) bool {
				return reflect.DeepEqual(param, otherParam)
			}); !found {
				return false
			}
		}

		return true
	})

	return found
}

// HTTPRouteRuleToString prints the matches of a  HTTPRouteRule as string
func HTTPRouteRuleToString(rule gatewayapiv1beta1.HTTPRouteRule) string {
	matches := common.Map(rule.Matches, HTTPRouteMatchToString)
	return fmt.Sprintf("{matches:[%s]}", strings.Join(matches, ","))
}

func HTTPRouteMatchToString(match gatewayapiv1beta1.HTTPRouteMatch) string {
	var patterns []string
	if method := match.Method; method != nil {
		patterns = append(patterns, fmt.Sprintf("method:%v", HTTPMethodToString(method)))
	}
	if path := match.Path; path != nil {
		patterns = append(patterns, fmt.Sprintf("path:%s", HTTPPathMatchToString(path)))
	}
	if len(match.QueryParams) > 0 {
		queryParams := common.Map(match.QueryParams, HTTPQueryParamMatchToString)
		patterns = append(patterns, fmt.Sprintf("queryParams:[%s]", strings.Join(queryParams, ",")))
	}
	if len(match.Headers) > 0 {
		headers := common.Map(match.Headers, HTTPHeaderMatchToString)
		patterns = append(patterns, fmt.Sprintf("headers:[%s]", strings.Join(headers, ",")))
	}
	return fmt.Sprintf("{%s}", strings.Join(patterns, ","))
}

func HTTPPathMatchToString(path *gatewayapiv1beta1.HTTPPathMatch) string {
	if path == nil {
		return "*"
	}
	if path.Type != nil {
		switch *path.Type {
		case gatewayapiv1beta1.PathMatchExact:
			return *path.Value
		case gatewayapiv1beta1.PathMatchRegularExpression:
			return fmt.Sprintf("~/%s/", *path.Value)
		}
	}
	return fmt.Sprintf("%s*", *path.Value)
}

func HTTPHeaderMatchToString(header gatewayapiv1beta1.HTTPHeaderMatch) string {
	if header.Type != nil {
		switch *header.Type {
		case gatewayapiv1beta1.HeaderMatchRegularExpression:
			return fmt.Sprintf("{%s:~/%s/}", header.Name, header.Value)
		}
	}
	return fmt.Sprintf("{%s:%s}", header.Name, header.Value)
}

func HTTPQueryParamMatchToString(queryParam gatewayapiv1beta1.HTTPQueryParamMatch) string {
	if queryParam.Type != nil {
		switch *queryParam.Type {
		case gatewayapiv1beta1.QueryParamMatchRegularExpression:
			return fmt.Sprintf("{%s:~/%s/}", queryParam.Name, queryParam.Value)
		}
	}
	return fmt.Sprintf("{%s:%s}", queryParam.Name, queryParam.Value)
}

func HTTPMethodToString(method *gatewayapiv1beta1.HTTPMethod) string {
	if method == nil {
		return "*"
	}
	return string(*method)
}

func GetKuadrantNamespaceFromPolicyTargetRef(ctx context.Context, cli client.Client, policy common.KuadrantPolicy) (string, error) {
	targetRef := policy.GetTargetRef()
	gwNamespacedName := types.NamespacedName{Namespace: string(ptr.Deref(targetRef.Namespace, policy.GetWrappedNamespace())), Name: string(targetRef.Name)}
	if IsTargetRefHTTPRoute(targetRef) {
		route := &gatewayapiv1beta1.HTTPRoute{}
		if err := cli.Get(
			ctx,
			types.NamespacedName{Namespace: string(ptr.Deref(targetRef.Namespace, policy.GetWrappedNamespace())), Name: string(targetRef.Name)},
			route,
		); err != nil {
			return "", err
		}
		// First should be OK considering there's 1 Kuadrant instance per cluster and all are tagged
		parentRef := route.Spec.ParentRefs[0]
		gwNamespacedName = types.NamespacedName{Namespace: string(*parentRef.Namespace), Name: string(parentRef.Name)}
	}
	gw := &gatewayapiv1beta1.Gateway{}
	if err := cli.Get(ctx, gwNamespacedName, gw); err != nil {
		return "", err
	}
	return GetKuadrantNamespace(gw)
}

func GetKuadrantNamespaceFromPolicy(policy common.KuadrantPolicy) (string, bool) {
	if kuadrantNamespace, isSet := policy.GetAnnotations()[common.KuadrantNamespaceLabel]; isSet {
		return kuadrantNamespace, true
	}
	return "", false
}

func GetKuadrantNamespace(obj client.Object) (string, error) {
	if !IsKuadrantManaged(obj) {
		return "", errors.NewInternalError(fmt.Errorf("object %T is not Kuadrant managed", obj))
	}
	return obj.GetAnnotations()[common.KuadrantNamespaceLabel], nil
}

func IsKuadrantManaged(obj client.Object) bool {
	_, isSet := obj.GetAnnotations()[common.KuadrantNamespaceLabel]
	return isSet
}

func AnnotateObject(obj client.Object, namespace string) {
	annotations := obj.GetAnnotations()
	if len(annotations) == 0 {
		obj.SetAnnotations(
			map[string]string{
				common.KuadrantNamespaceLabel: namespace,
			},
		)
	} else {
		if !IsKuadrantManaged(obj) {
			annotations[common.KuadrantNamespaceLabel] = namespace
			obj.SetAnnotations(annotations)
		}
	}
}

func DeleteKuadrantAnnotationFromGateway(gw *gatewayapiv1beta1.Gateway, namespace string) {
	annotations := gw.GetAnnotations()
	if IsKuadrantManaged(gw) && annotations[common.KuadrantNamespaceLabel] == namespace {
		delete(gw.Annotations, common.KuadrantNamespaceLabel)
	}
}

// routePathMatchToRulePath converts HTTPRoute pathmatch rule to kuadrant's rule path
func routePathMatchToRulePath(pathMatch *gatewayapiv1beta1.HTTPPathMatch) []string {
	if pathMatch == nil {
		return nil
	}

	// Only support for Exact and Prefix match
	if pathMatch.Type != nil && *pathMatch.Type != gatewayapiv1beta1.PathMatchPathPrefix &&
		*pathMatch.Type != gatewayapiv1beta1.PathMatchExact {
		return nil
	}

	// Exact path match
	suffix := ""
	if pathMatch.Type == nil || *pathMatch.Type == gatewayapiv1beta1.PathMatchPathPrefix {
		// defaults to path prefix match type
		suffix = "*"
	}

	val := "/"
	if pathMatch.Value != nil {
		val = *pathMatch.Value
	}

	return []string{val + suffix}
}

func GetGatewayWorkloadSelector(ctx context.Context, cli client.Client, gateway *gatewayapiv1beta1.Gateway) (map[string]string, error) {
	address, found := common.Find(
		gateway.Status.Addresses,
		func(address gatewayapiv1beta1.GatewayAddress) bool {
			return address.Type != nil && *address.Type == gatewayapiv1beta1.HostnameAddressType
		},
	)
	if !found {
		return nil, fmt.Errorf("cannot find service Hostname in the Gateway status")
	}
	serviceNameParts := strings.Split(address.Value, ".")
	serviceKey := client.ObjectKey{
		Name:      serviceNameParts[0],
		Namespace: serviceNameParts[1],
	}
	return GetServiceWorkloadSelector(ctx, cli, serviceKey)
}

func IsHTTPRouteAccepted(httpRoute *gatewayapiv1beta1.HTTPRoute) bool {
	if httpRoute == nil {
		return false
	}

	if len(httpRoute.Spec.CommonRouteSpec.ParentRefs) == 0 {
		return false
	}

	// Check HTTProute parents (gateways) in the status object
	// if any of the current parent gateways reports not "Admitted", return false
	for _, parentRef := range httpRoute.Spec.CommonRouteSpec.ParentRefs {
		routeParentStatus := func(pRef gatewayapiv1beta1.ParentReference) *gatewayapiv1beta1.RouteParentStatus {
			for idx := range httpRoute.Status.RouteStatus.Parents {
				if reflect.DeepEqual(pRef, httpRoute.Status.RouteStatus.Parents[idx].ParentRef) {
					return &httpRoute.Status.RouteStatus.Parents[idx]
				}
			}

			return nil
		}(parentRef)

		if routeParentStatus == nil {
			return false
		}

		if meta.IsStatusConditionFalse(routeParentStatus.Conditions, "Accepted") {
			return false
		}
	}

	return true
}
