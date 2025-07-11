minikube_init() {
    minikube start --cpus=4 --memory=4096

    minikube addons enable ingress
    minikube addons enable dashboard
    minikube addons enable metrics-server

    minikube dashboard &
}

minikube_init

cd k8s/postgresql
kubectl apply -f secret.yaml; kubectl create -f pv.yaml; kubectl create -f pvc.yaml; kubectl apply -f configmap.yaml; kubectl apply -f deployment.yaml; kubectl apply -f service.yaml;

cd ../backend
kubectl wait --for=jsonpath='{.status.phase}'=Running $(kubectl get pods -o=name) --timeout=300s
kubectl apply -f secret.yaml; kubectl create -f deployment.yaml; kubectl create -f service.yaml; 

cd ../frontend
kubectl wait --for=jsonpath='{.status.phase}'=Running $(kubectl get pods -o=name) --timeout=300s
kubectl create -f deployment.yaml; kubectl create -f service.yaml;

cd ..
kubectl apply -f ingress.yaml