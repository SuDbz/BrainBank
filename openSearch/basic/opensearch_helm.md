# Enabling and Using OpenSearch and Dashboard with Helm

## Table of Contents
1. [Introduction](#introduction)
2. [Prerequisites](#prerequisites)
3. [Installing Helm](#installing-helm)
4. [Adding OpenSearch Helm Repository](#adding-opensearch-helm-repository)
5. [Deploying OpenSearch](#deploying-opensearch)
6. [Configuring OpenSearch](#configuring-opensearch)
7. [Accessing OpenSearch Dashboard](#accessing-opensearch-dashboard)
8. [Enabling OpenSearch as a Service](#enabling-opensearch-as-a-service)
9. [Conclusion](#conclusion)

## Introduction
This guide provides instructions on how to enable and use OpenSearch and its dashboard using Helm.

## Prerequisites
- Kubernetes cluster
- kubectl configured
- Helm installed

## Installing Helm
To install Helm, follow the official [Helm installation guide](https://helm.sh/docs/intro/install/).

## Adding OpenSearch Helm Repository
Add the OpenSearch Helm repository to your Helm configuration:
```sh
helm repo add opensearch https://opensearch-project.github.io/helm-charts/
helm repo update
```

## Deploying OpenSearch
Deploy OpenSearch using the Helm chart:
```sh
helm install my-opensearch opensearch/opensearch
```

## Configuring OpenSearch
You can customize the OpenSearch deployment by creating a `values.yaml` file and specifying your configuration. For example:
```yaml
clusterName: "my-opensearch-cluster"
nodeGroup: "master"
```
Then deploy using:
```sh
helm install my-opensearch -f values.yaml opensearch/opensearch
```

## Accessing OpenSearch Dashboard
To access the OpenSearch dashboard, you need to port-forward the service:
```sh
kubectl port-forward svc/my-opensearch-dashboards 5601:5601
```
Then, open your browser and navigate to `http://localhost:5601`.

## Enabling OpenSearch as a Service
To ensure OpenSearch starts on boot and runs continuously, you can enable it as a service:
```sh
kubectl create -f opensearch-service.yaml
```
Make sure your `opensearch-service.yaml` is properly configured to manage the OpenSearch service.

## Conclusion
You have successfully enabled and accessed OpenSearch and its dashboard using Helm.