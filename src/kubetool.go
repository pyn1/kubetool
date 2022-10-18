package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"time"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	kubeinformers "k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

var (
	kubeconfig      *string
	nginxDeployName *string
	scaleNginx      *int
	upgradeNginx    *float64
	watchNginx      *bool
)

func CLIUsage() {
	fmt.Printf("Example usage: \n To Scale and upgrade nginx deployment:\nkubedeploy --kubeconfig <Absolute Path> –scale 3 –upgrade <version> --name nginx \n\n To delete the nginx deployment:\nkubedeploy delete –name nginx\n\nTo watch the nginx pods:\nkubedeploy watch –name nginx\n")
	flag.PrintDefaults()
}

func main() {

	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	scaleNginx = flag.Int("scale", 1, "Number replicas for the Nginx Pod Deployment")
	upgradeNginx = flag.Float64("upgrade", 1.14, "Default Nginx version deployed")
	nginxDeployName = flag.String("name", "nginx", "Nginx deployment name")
	//Override the usage of CLI for the user
	flag.Usage = CLIUsage
	flag.Parse()

	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		fmt.Println("Error: Building config from flags: ", err)
		config, err = rest.InClusterConfig()
		if err != nil {
			fmt.Println("Error: Getting Incluster config", err)
			os.Exit(1)
		}
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		fmt.Println("Error: Getting the clisent set", err)
	}

	if len(os.Args) < 2 {
		fmt.Println("Error: Wrong number of arguments, please execute with help option(-h/--help)")
		os.Exit(1)
	}

	//Scale the pods
	if *scaleNginx != 1 {
		fmt.Println("\n\n Scaling...")
		s, err := clientset.AppsV1().
			Deployments("test").
			GetScale(context.TODO(), *nginxDeployName, metav1.GetOptions{})
		if err != nil {
			fmt.Println(err)
		}

		sc := *s
		sc.Spec.Replicas = int32(*scaleNginx)

		us, err := clientset.AppsV1().
			Deployments("test").
			UpdateScale(context.TODO(),
				*nginxDeployName, &sc, metav1.UpdateOptions{})
		if err != nil {
			fmt.Println(err)
		}
		//Updated object
		fmt.Println(*us)

		fmt.Println("\n\n Scaled.")
	}

	// Delete Deployment
	if os.Args[1] == "delete" && *nginxDeployName == "nginx" {

		fmt.Println("Deleting deployment...")
		deletePolicy := metav1.DeletePropagationForeground
		if err := clientset.AppsV1().
			Deployments("test").Delete(context.TODO(), *nginxDeployName, metav1.DeleteOptions{
			PropagationPolicy: &deletePolicy,
		}); err != nil {
			panic(err)
		}
		fmt.Println("Deleted deployment.")
	}
	//pods,err := clientset.CoreV1().Pods("test").List(context.Background(), metav1.ListOptions{})

	if os.Args[1] == "watch" && *nginxDeployName == "nginx" {
		kubeInformerFactory := kubeinformers.NewSharedInformerFactory(clientset, time.Second*30)
		podsInformer := kubeInformerFactory.Core().V1().Pods().Informer()

		podsInformer.AddEventHandler(cache.ResourceEventHandlerFuncs{
			AddFunc: func(obj interface{}) {
				fmt.Println("Service added")
				pod := obj.(*v1.Pod)
				fmt.Printf("pod added: %s/%s", pod.Namespace, pod.Name)

			},
			DeleteFunc: func(obj interface{}) {
				fmt.Println("Service deleted")
				pod := obj.(*v1.Pod)
				fmt.Printf("pod deleted: %s/%s", pod.Namespace, pod.Name)
			},
			UpdateFunc: func(oldObj, newObj interface{}) {
				fmt.Println("Service Updated")
				podO := oldObj.(*v1.Pod)
				podN := newObj.(*v1.Pod)
				fmt.Printf("Old pod: %s/%s", podO.Namespace, podO.Name)
				fmt.Printf("New pod updated: %s/%s", podN.Namespace, podN.Name)
			},
		})

		stop := make(chan struct{})
		defer close(stop)
		kubeInformerFactory.Start(stop)
		for {
			time.Sleep(time.Second)
		}

	}
	//fmt.Println(pods)
}
