package main

// +kubebuilder:rbac:groups="",resources=deployments;statefulsets;replicasets;pods,verbs=get;list;watch;create;update;delete;patch
// +kubebuilder:rbac:groups="",resources=deployments/status;statefulsets/status;replicasets/status;pods/status,verbs=get;list;watch;create;update;delete;patch
// +kubebuilder:rbac:groups="",resources=services;endpoints,verbs=get;list;watch;create;update;delete;patch
// +kubebuilder:rbac:groups="",resources=configmaps;secrets,verbs=get;list;watch;create;update;delete;patch
// +kubebuilder:rbac:groups=networking.k8s.io;extensions,resources=ingresses,ingresses/status,verbs=get;list;watch;create;update;delete;patch
// +kubebuilder:rbac:groups=rbac.authorization.k8s.io,resources=serviceaccounts;roles;clusterroles;rolebinding;clusterrolebindings,verbs=get;list;watch;create;update;delete;patch

func main() {}
