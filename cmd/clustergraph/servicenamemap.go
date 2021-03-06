package main

import (
	serviceparser "github.com/fabric8-analytics/poc-ocp-upgrade-prediction/pkg/serviceparser"
)

type ServiceCompileTimeCalls struct {
	servicename string
	edges       []serviceparser.CompileEdge
}

type ServiceMapping struct {
	servicename string
	cmdname     []string
}

var ServicePackageMap = map[string]ServiceMapping{
	"cluster-api-provider-aws":                      ServiceMapping{servicename: "aws-machine-controllers", cmdname: []string{"././cmd/aws-actuator", "./cmd/manager"}},
	"cluster-api-provider-azure":                    ServiceMapping{servicename: "cluster-api-provider-azure", cmdname: []string{"./cmd/clusterctl", "manager"}},
	"cluster-api-provider-baremetal":                ServiceMapping{servicename: "cluster-api-provider-baremetal", cmdname: []string{"./cmd/manager"}},
	"origin":                                        ServiceMapping{servicename: "hypershift", cmdname: []string{"./cmd/hypershift", "./cmd/oc"}},
	"etcd":                                          ServiceMapping{servicename: "etcd", cmdname: []string{"./"}},
	"kubernetes":                                    ServiceMapping{servicename: "hyperkube", cmdname: []string{"./cmd/hyperkube"}},
	"cloud-credential-operator":                     ServiceMapping{servicename: "cloud-credential-operator", cmdname: []string{"./cmd/manager"}},
	"cluster-authentication-operator":               ServiceMapping{servicename: "cluster-authentication-operator", cmdname: []string{"./cmd/authentication-operator"}},
	"cluster-autoscaler-operator":                   ServiceMapping{servicename: "cluster-autoscaler-operator", cmdname: []string{"./cmd/main"}},
	"cluster-bootstrap":                             ServiceMapping{servicename: "cluster-bootstrap", cmdname: []string{"./cmd/cluster-bootstrap"}},
	"cluster-config-operator":                       ServiceMapping{servicename: "cluster-config-operator", cmdname: []string{"./cmd/cluster-config-operator"}},
	"cluster-dns-operator":                          ServiceMapping{servicename: "cluster-dns-operator", cmdname: []string{"./cmd/dns-operator"}},
	"cluster-image-registry-operator":               ServiceMapping{servicename: "cluster-image-registry-operator", cmdname: []string{"./cmd/cluster-image-registry-operator"}},
	"cluster-ingress-operator":                      ServiceMapping{servicename: "cluster-ingress-operator", cmdname: []string{"./cmd/ingress-operator"}},
	"cluster-kube-apiserver-operator":               ServiceMapping{servicename: "cluster-kube-apiserver-operator", cmdname: []string{"./cmd/cluster-kube-apiserver-operator"}},
	"cluster-kube-controller-manager-operator":      ServiceMapping{servicename: "aws-machine-controllers", cmdname: []string{"./cmd/aws-machine-controllers"}},
	"cluster-kube-scheduler-operator":               ServiceMapping{servicename: "cluster-kube-scheduler-operator", cmdname: []string{"./cmd/cluster-kube-scheduler-operator"}},
	"cluster-monitoring-operator":                   ServiceMapping{servicename: "cluster-monitoring-operator", cmdname: []string{"./cmd/operator"}},
	"cluster-network-operator":                      ServiceMapping{servicename: "cluster-network-operator", cmdname: []string{"./cmd/cluster-network-operator"}},
	"cluster-node-tuning-operator":                  ServiceMapping{servicename: "cluster-node-tuning-operator", cmdname: []string{"./cmd/manager"}},
	"cluster-openshift-apiserver-operator":          ServiceMapping{servicename: "cluster-openshift-apiserver-operator", cmdname: []string{"./cmd/cluster-openshift-apiserver-operator"}},
	"cluster-openshift-controller-manager-operator": ServiceMapping{servicename: "cluster-openshift-controller-manager-operator", cmdname: []string{"./cmd/cluster-openshift-controller-manager-operator"}},
	"cluster-samples-operator":                      ServiceMapping{servicename: "cluster-samples-operator", cmdname: []string{"./cmd/cluster-samples-operator"}},
	"cluster-storage-operator":                      ServiceMapping{servicename: "cluster-storage-operator", cmdname: []string{"./cmd/manager"}},
	"cluster-svcat-apiserver-operator":              ServiceMapping{servicename: "cluster-svcat-apiserver-operator", cmdname: []string{"./cmd/cluster-svcat-apiserver-operator"}},
	"cluster-svcat-controller-manager-operator":     ServiceMapping{servicename: "cluster-svcat-controller-manager-operator", cmdname: []string{"./cmd/cluster-svcat-controller-manager-operator"}},
	"cluster-version-operator":                      ServiceMapping{servicename: "cluster-version-operator", cmdname: []string{"cmd"}},
	"console":                                       ServiceMapping{servicename: "console", cmdname: []string{"./cmd/bridge"}},
	"console-operator":                              ServiceMapping{servicename: "console-operator", cmdname: []string{"./cmd/console"}},
	"builder":                                       ServiceMapping{servicename: "builder", cmdname: []string{"cmd"}},
	"image-registry":                                ServiceMapping{servicename: "image-registry", cmdname: []string{"./cmd/dockerregistry"}},
	"router":                                        ServiceMapping{servicename: "router", cmdname: []string{"./cmd/openshift-router"}},
	"installer":                                     ServiceMapping{servicename: "installer", cmdname: []string{"./cmd/openshift-install"}},
	"k8s-prometheus-adapter":                        ServiceMapping{servicename: "k8s-prometheus-adapter", cmdname: []string{"./cmd/adapter"}},
	"kubecsr":                                       ServiceMapping{servicename: "kubecsr", cmdname: []string{"./cmd/kube-aws-approver", "./cmd/kube-client-agent", "./cmd/kube-etcd-signer-server"}},
	"cluster-api-provider-libvirt":                  ServiceMapping{servicename: "cluster-api-provider-libvirt", cmdname: []string{"./cmd/libvirt-actuator", "manager"}},
	"machine-api-operator":                          ServiceMapping{servicename: "machine-api-operator", cmdname: []string{"./cmd/machine-api-operator", "./cmd/machine-healthcheck", "./cmd/nodelink-controller"}},
	"machine-config-operator":                       ServiceMapping{servicename: "machine-config-operator", cmdname: []string{"./cmd/gcp-routes-controller", "./cmd/machine-config-controller", "./cmd/machine-config-daemon", "./cmd/machine-config-operator", "./cmd/machine-config-server", "./cmd/setup-etcd-environment"}},
	"must-gather":                                   ServiceMapping{servicename: "must-gather", cmdname: []string{"./cmd/openshift-dev-helpers", "./cmd/openshift-must-gather"}},
	"cluster-api-provider-openstack":                ServiceMapping{servicename: "cluster-api-provider-openstack", cmdname: []string{"./cmd/clusterctl", "./cmd/manager"}},
	"operator-lifecycle-manager":                    ServiceMapping{servicename: "operator-lifecycle-manager", cmdname: []string{"./cmd/catalog", "./cmd/package-server", "./cmd/olm"}},
	"operator-marketplace":                          ServiceMapping{servicename: "operator-marketplace", cmdname: []string{"./cmd/manager"}},
	"operator-registry":                             ServiceMapping{servicename: "operator-registry", cmdname: []string{"./cmd/appregistry-server", "./cmd/registry-server", "./cmd/initializer", "./cmd/configmap-server"}},
	"prometheus":                                    ServiceMapping{servicename: "prometheus", cmdname: []string{"./cmd/prometheus", "./cmd/promtool"}},
	"prometheus-alertmanager":                       ServiceMapping{servicename: "prometheus-alertmanager", cmdname: []string{"./cmd/alertmanager", "./cmd/amtool"}},
	"prometheus-operator": ServiceMapping{servicename: "prometheus-operator	", cmdname: []string{"./cmd/lint", "./cmd/operator", "./cmd/po-crdgen", "./cmd/po-docgen", "./cmd/po-rule-migration", "./cmd/prometheus-config-reloader"}},
	"service-ca-operator":         ServiceMapping{servicename: "service-ca-operator", cmdname: []string{"./cmd/service-ca-operator"}},
	"service-catalog":             ServiceMapping{servicename: "service-catalog", cmdname: []string{"./cmd/apiserver/app/server", "./cmd/controller-manager/app", "./cmd/healthcheck", "./cmd/service-catalog", "./cmd/svcat"}},
	"sriov-cni":                   ServiceMapping{servicename: "sriov-cni", cmdname: []string{"./cmd/sriov"}},
	"sriov-network-device-plugin": ServiceMapping{servicename: "sriov-network-device-plugin", cmdname: []string{"./cmd/sriovdp"}},
	"telemeter":                   ServiceMapping{servicename: "telemeter", cmdname: []string{"./cmd/authorization-server", "./cmd/telemeter-benchmark", "./cmd/telemeter-client", "./cmd/telemeter-server"}},
}
