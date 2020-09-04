package main

import (
	"context"
	"fmt"
	"github.com/kelseyhightower/envconfig"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"regexp"

	_ "k8s.io/client-go/plugin/pkg/client/auth"
)

type env struct {
	Pod string `required:"true" split_words:"true"`
	Namespace string `required:"true" split_words:"true"`
}

func main() {
	var cfg env
	// Parse env variables to our spec.
	if err := envconfig.Process("Cleaner", &cfg); err != nil {
		panic(err)
	}

	regex := regexp.MustCompile(cfg.Pod)

	config, err := rest.InClusterConfig()
	if err != nil {
		fmt.Println(err)
		panic(err.Error())
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	// create the clientset
	pods, err := clientset.CoreV1().Pods(cfg.Namespace).List(context.Background(), metav1.ListOptions{})
	if err != nil {
		fmt.Println(err)
	}

	for _, pod := range pods.Items {
		match := regex.MatchString(pod.Name)
		if match == true {
			clientset.CoreV1().Pods(cfg.Namespace).Delete(context.Background(), pod.Name, metav1.DeleteOptions{})
			fmt.Println("Deleting pod", pod.Name)
		}
	}
}