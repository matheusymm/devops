cd k8s/
kubectl delete -f ingress.yaml

cd frontend/
kubectl delete -f deployment.yaml; kubectl delete -f service.yaml;  

cd ../backend/
kubectl delete -f deployment.yaml; kubectl delete -f service.yaml; kubectl delete -f secret.yaml;

cd ../postgresql/
kubectl delete -f deployment.yaml; kubectl delete -f service.yaml; kubectl delete -f configmap.yaml; kubectl delete -f secret.yaml; kubectl delete -f pvc.yaml; kubectl delete pv postgresql-pv-volume