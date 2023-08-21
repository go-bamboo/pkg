package kuberegistry

import (
	"github.com/go-bamboo/pkg/registry"
	"github.com/go-bamboo/pkg/registry/core"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"path/filepath"
)

func init() {
	registry.Register("kube", Create)
}

func getClientSet() (*kubernetes.Clientset, error) {
	restConfig, err := rest.InClusterConfig()
	home := homedir.HomeDir()

	if err != nil {
		kubeconfig := filepath.Join(home, ".kube", "config")
		restConfig, err = clientcmd.BuildConfigFromFlags("", kubeconfig)
		if err != nil {
			return nil, err
		}
	}
	clientSet, err := kubernetes.NewForConfig(restConfig)
	if err != nil {
		return nil, err
	}
	return clientSet, nil
}

func Create(c *registry.Conf) (core.Registrar, core.Discovery, error) {
	cli, err := getClientSet()
	if err != nil {
		panic(err)
	}
	r := NewRegistry(cli)
	return r, r, nil
}
