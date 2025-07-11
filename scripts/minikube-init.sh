minikube start --cpus=4 --memory=4096

minikube addons enable ingress

minikube addons enable dashboard

minikube addons enable metrics-server

kubectl wait --namespace ingress-nginx \
      --for=condition=ready pod \
      --selector=app.kubernetes.io/component=controller \
      --timeout=120s

minikube dashboard &