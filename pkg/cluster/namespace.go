package cluster

import (
	"container/list"
	"context"
	"fmt"
	"path/filepath"

	"github.com/sirupsen/logrus"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

func ListNamespaces() *list.List {
	// var kubeConfig *string
	ctx := context.Background()
	// if home := homedir.HomeDir(); home != "" {
	// 	kubeConfig = flag.String("kubeConfig", filepath.Join(home, ".kube", "config"), "absolute path to the kubeConfig")
	// } else {
	// 	kubeConfig = flag.String("kubeConfig", "", "absolute path to the kubeConfig")
	// }
	home := homedir.HomeDir()
	config, err := clientcmd.BuildConfigFromFlags("", filepath.Join(home, ".kube", "config"))
	if err != nil {
		logrus.Fatal(err)
		panic(err)
	}
	clientSet, err := kubernetes.NewForConfig(config)
	if err != nil {
		logrus.Fatal(err)
		panic(err)
	}
	namespaceList, err := clientSet.CoreV1().Namespaces().List(ctx, metav1.ListOptions{})
	if err != nil {
		logrus.Fatal(err)
		panic(err)
	}
	namespaces := namespaceList.Items
	result := list.New()
	for _, namespace := range namespaces {
		fmt.Println(namespace.Name)
		result.PushBack(namespace.Name)
	}
	return result
}
