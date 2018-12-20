package main

import (
	"crd/pkg/apis/crd/v1alpha1"
	cpclientset "crd/pkg/client/clientset/versioned"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/tools/clientcmd"
	"log"
	"os"
	"path/filepath"
)

func main() {
	kubeconfigPath := filepath.Join(os.Getenv("HOME"), ".kube/config")

	config, err := clientcmd.BuildConfigFromFlags("", kubeconfigPath)
	if err != nil {
		log.Fatalf("Could not get Kubernetes config: %s", err)
	}
	cpClient, err := cpclientset.NewForConfig(config)

	log.Println("Creating CustomPod. . . . .")

	customPod := v1alpha1.CustomPod{
		ObjectMeta: v1.ObjectMeta{
			Name:            "cpod",
			Labels: map[string]string{
				"app":"cpod",
				},
		},
		Spec: v1alpha1.CustomPodSpec{
			Containers: []v1alpha1.Container{
				{
					Name: "api-latest",
					Image: "fahimabrar/api:latest",
				},
				{
					Name: "api-alpine",
					Image:"fahimabrar/api:alpine",

				},

			},
		},
	}

	_, err = cpClient.CrdV1alpha1().CustomPods("default").Create(&customPod)
	log.Println("CustomPod Created!\n %v",customPod)
}
