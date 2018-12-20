package main

import (
	"crd/pkg/apis/crd.com/v1alpha1"
	cpclientset "crd/pkg/client/clientset/versioned"
	"github.com/tamalsaha/go-oneliners"
	"k8s.io/apiextensions-apiserver/pkg/apis/apiextensions"
	"k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	"k8s.io/apiextensions-apiserver/pkg/apiserver"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/tools/clientcmd"
	"log"
	"os"
	"os/signal"
	"path/filepath"
)

func main() {
	kubeconfigPath := filepath.Join(os.Getenv("HOME"), ".kube/config")

	config, err := clientcmd.BuildConfigFromFlags("", kubeconfigPath)
	if err != nil {
		log.Fatalf("Could not get Kubernetes config: %s", err)
	}
	cpClient, err := cpclientset.NewForConfig(config)

	//crdClient, err :=
	crd := v1beta1.CustomResourceDefinition{
		Spec:   v1beta1.CustomResourceDefinitionSpec{},
		Status: v1beta1.CustomResourceDefinitionStatus{},
	}

	//crdClient.
	log.Println("Creating CustomPod 'cpod'. . . . .")

	customPod := &v1alpha1.CustomPod{
		TypeMeta: metav1.TypeMeta{
			Kind:       "CustomPod",
			APIVersion: "crd.com/v1alpha1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name: "cpod",
			Labels: map[string]string{
				"app": "cpod"},
		},
		Spec: v1alpha1.CustomPodSpec{
			Containers: []v1alpha1.Container{
				{
					Name:  "api-latest",
					Image: "fahimabrar/api:latest",
				},
				{
					Name:  "api-alpine",
					Image: "fahimabrar/api:alpine",
				},
			},
		},
	}


	cp, err := cpClient.CrdV1alpha1().CustomPods("default").Create(customPod)
	if err != nil {
		panic(err)
	}

	log.Println("'cpod' Created!\n")
	oneliners.PrettyJson(cp)

	ch := make(chan os.Signal,1)
	signal.Notify(ch, os.Interrupt)
	<-ch

	log.Println("Deleting CustomPod. . . . .")
	if err := cpClient.CrdV1alpha1().CustomPods("default").Delete(
		"cpod",
		nil,
	); err != nil {
		panic(err)
	}
	log.Println("CustomPod Deleted")
}
