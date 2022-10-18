# kubetool
Tool to checks few operation on the K8s resource

This project creates a CLI tool using Golang which will be able to perform the following operations:

Scale the number of pods
Upgrade the version of the nginx server
Deletes the pods
Continuously watches for the desired number of replicas on the cluster and takes action accordingly

This project also illustrates RBAC in K8s in particaular how one can create a user, give permissions to perform all actions on the ‘pod’ and ‘deployment’ resource in the ‘test’ namespace. Assign the required role.
Include authentication using a kubeconfig file
Following is the usage of the CLI
Example usage:
 To Scale and upgrade nginx deployment:
kubedeploy --kubeconfig <Absolute Path> –scale 3 –upgrade <version> --name nginx

To delete the nginx deployment:
kubedeploy delete –name nginx

To watch the nginx pods:
kubedeploy watch –name nginx
  -kubeconfig string
        (optional) absolute path to the kubeconfig file (default "/home/edgec/.kube/config")
  -name string
        Nginx deployment name (default "nginx")
  -scale int
        Number replicas for the Nginx Pod Deployment (default 1)
  -upgrade float
        Default Nginx version deployed (default 1.14)

The application is also containerized and can be run as the pod on a K8s cluster:
Command to deploy the pod:
1. Swicth to minikube contex:
kubectl config use-context minikube
2. Deploy the nginx service
kubectl apply -f nginx-deploy.yml
3. Switch back to the user(pyn) context
kubectl config use-context pyn-context
4.Execute above commands to different operations


Optional:  You should also be able to execute the CLI tool by deploying it as a pod on the cluster.

 

NOTE: All the commands are tested on minikube
Reference: https://github.com/kubernetes/client-go
