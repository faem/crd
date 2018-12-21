package main

import (
	"crd/pkg/apis/crd.com/v1alpha1"
	cpclientset "crd/pkg/client/clientset/versioned"
	"github.com/tamalsaha/go-oneliners"
	crdapi "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	crdclientset "k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/tools/clientcmd"
	"log"
	"os"
	"os/signal"
	"path/filepath"
	"time"
)

func main() {
	kubeconfigPath := filepath.Join(os.Getenv("HOME"), ".kube/config")
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfigPath)
	if err != nil {
		log.Fatalf("Could not get Kubernetes config: %s", err)
	}

	log.Println("Custom Resource 'CustomPod' Creating. . . . .")
	crdClient, err := crdclientset.NewForConfig(config)
	mycrd := &crdapi.CustomResourceDefinition{
		ObjectMeta: metav1.ObjectMeta{
			Name: "custompods.crd.com",
		},
		Spec: crdapi.CustomResourceDefinitionSpec{
			Group: "crd.com",
			Names: crdapi.CustomResourceDefinitionNames{
				Plural:   "custompods",
				Singular: "custompod",
				ShortNames: []string{
					"cp",
				},
				Kind: "CustomPod",
			},
			Scope: "Namespaced",
			Versions: []crdapi.CustomResourceDefinitionVersion{
				{
					Name:    "v1alpha1",
					Served:  true,
					Storage: true,
				},
			},
		},
	}
	_, err = crdClient.ApiextensionsV1beta1().CustomResourceDefinitions().Create(mycrd)
	if err != nil {
		panic(err)
	}
	time.Sleep(5 * time.Second)
	log.Println("Custom Resource 'CustomPod' Created!")

	defer func() {
		log.Println("Deleting CustomPod. . . . .")
		if err = crdClient.ApiextensionsV1beta1().CustomResourceDefinitions().Delete(
			"custompods.crd.com",
			nil,
		); err != nil {
			panic(err)
		}
		log.Println("CustomPod Deleted")
	}()

	log.Println("Press Ctrl+C to Create an instance of CustomPod. . . .")
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	<-ch

	log.Println("Creating CustomPod 'cpod'. . . . .")
	cpClient, err := cpclientset.NewForConfig(config)
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

	defer func() {
		log.Println("Deleting 'cpod'. . . . .")
		if err := cpClient.CrdV1alpha1().CustomPods("default").Delete(
			"cpod",
			nil,
		); err != nil {
			panic(err)
		}
		log.Println("'cpod' Deleted")
	}()

	log.Println("Press Ctrl+C to Delete the Custom Resource 'CustomPod' and the instance of CustomPod 'cpod'. . . .")
	ch = make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	<-ch
}
