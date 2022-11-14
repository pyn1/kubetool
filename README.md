# kubetool
Tool to check few operation on the K8s resource

This project creates a CLI tool using Golang which will be able to perform the following operations:

1.Scale the number of pods  
2. Upgrade the version of the nginx server  
3. Deletes the pods  
4. Continuously watches for the desired number of replicas on the cluster and takes action accordingly 


This project also illustrates RBAC in K8s in particaular how one can create a user, give permissions to perform all actions on the ‘pod’ and ‘deployment’ resource in the ‘test’ namespace and assign the required role.
Include authentication using a kubeconfig file

Prequisite:
Create the user and certs for same, the repo provides a sample certs under cert directory. Please run below commands to create an example user and certs for same.  
1. cd cert/  
2. openssl genrsa -out <user>.key 2048  
3. openssl req -new -key pyn.key -out <user>.csr -subj "/CN=<user>/O=<group>"  
4. openssl x509 -req -in <user>.csr -CA ~/.minikube/ca.crt -CAkey ~/.minikube/ca.key -CAcreateserial -out <user>.crt -days 500

Create user context and update the minikube config
1. kubectl config set-credentials <user> --client-certificate=<user>.crt --client-key=<user>.key  
2. kubectl config set-context <user>-context --cluster=minikube --user=<user>  

Switch back to minikube and deploy sample nginx server  
1. kubectl config use-context minikube

Apply role and rolebinding
1. kubectl apply -f role.yaml
2. kubectl apply -f role-binding.yml

NOTE: Change the user name in the role-binding.yml based on user name


USAGE:
Following is the usage of the CLI
 
To Scale nginx deployment:
kubedeploy --kubeconfig <Absolute Path> --scale 3 --name nginx

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
1. Switch to minikube contex:
kubectl config use-context minikube
2. Deploy the nginx service
kubectl apply -f nginx-deploy.yml
3. Switch back to the user(pyn) context
kubectl config use-context pyn-context
4.Execute above commands to different operations

NOTE: All the commands are tested on minikube  
Reference: https://github.com/kubernetes/client-go
