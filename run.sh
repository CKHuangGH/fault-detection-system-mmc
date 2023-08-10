kubectl apply -f cluster_role.yaml
kubectl apply -f cluster_role_binding.yaml
kubectl apply -f service_account.yaml

kubectl apply -f deploy.yaml