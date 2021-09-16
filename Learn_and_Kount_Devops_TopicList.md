# Updated Topic List for "Learn&Kount/DevOps", starting Sept. 29

Topics:

* gRPC at Kount
    * Brief technical overview and recommended learning resources
    * Kount's Service Catalog: patterns and selected under-the-hood details
    * Techniques for compiling proto files with or without the Service Catalog pipeline
    * Techniques for using proto files in development without the Service Catalog
    * Unary RPC vs streaming RPC
    * gRPC metadata: client-side/server-side
    * Interceptors and middleware for gRPC server/client
    * Kount's gRPC middleware: Overview
    * gRPC gateway for REST compatibility: How it works and how it's used at Kount
    * Customizing gRPC-gateway JSON marshaling if necessary
    * gRPC healthcheck; gRPC-service Dockerfiles
    * Deploying a new gRPC server with or without a template project

* Cloud infra at Kount
    * Brief intro to how Kount uses AWS
    * Brief intro to important AWS services that Kount uses
    * Intro to IAM and access policies
    * Intro to AWS CLI and how Kount configures its cloud access for users
    * Review of networking basics
    * Intro to subnets and security groups
    * Intro to AWS accounts, VPC/networking, and cross-account access config (including transit-gateway) at Kount
    * Intro to DNS and Route53 at Kount
    * Deeper look at EC2
    * Deeper look at load balancing and autoscaling
    * Deeper look at containers/Docker
    * Dockerfiles at Kount: selected special topics and debugging case-studies
    * Intro to ECR

* CI/CD at Kount
    * Gitlab technical intro
    * Kount's Gitlab server
    * Kount's Gitlab runners & config
    * Job images, ECR, and Platform Dockerfile projects
    * Environments at Kount
    * Project setup and config at Kount
    * CI/CD variables (both Gitlab-defined and custom) and secrets
    * Your personal space in Kount GitLab; forking projects
    * Brief intro to the GitLab API and GitLab webhooks/plugins
    * Some useful git tactics for Kount projects
    * Garden overview and topics / case studies

* Infrastructure-as-Code (IaC) at Kount
    * Intro to IaC concepts and tooling options
    * Intro to common IaC options for AWS
    * Intro to Cloudformation; tips/gotchas
    * Cloudformation at Kount
    * Intro to Terraform; tips/gotchas
    * The Terraform platform at Kount
    * Hands-on infra with Terraform
        * Take-home projects
        * Case studies

* Technical overview of K8s
    * What a container orchestrator *is* and why it's needed
    * The container orchestrator ecosystem and how K8s fits in
    * K8s components and how they work together
    * kubectl
    * K8s networking overview
    * Pods vs containers
    * ReplicaSets and Deployments
    * Healthchecks and resource requests/limits
    * Services and Service types
    * Secrets and ConfigMaps
    * Ingress
    * DaemonSets
    * Jobs & CronJobs
    * Metrics and autoscaling
    * Helm
    * Istio

* Practical guide to Kount's K8s
    * Kount's K8s clusters
    * Kount's Ingress controllers
    * Kount's microservices helm chart
    * Configuring container environment & data
    * Working with helm
    * Troubleshooting K8s resources
    * Troubleshooting helm releases
    * Configuring your pod autoscaling
    * Right-sizing your resource requests/limits

* Misc infra & CI/CD topics
    * Overview of selected AWS services and how Kount uses them
        * KMS
        * Lambda
        * DynamoDB
        * RDS
        * Elasticsearch
    * Datadog at Kount: Selected concepts and tools
    * Snowflake at Kount: "Journaling" overview and selected topics
    * Intro to TLS and PKI
        * TLS and PKI at Kount
    * Sonarqube at Kount

* (Probably skip): Overview of the infra & deployment architecture/toolkit for typical Kount projects
    * Brief anatomy of a Kount Go backend service and its deployment
        * Brief intro to Docker & K8s
        * Brief intro to k8s app config & how Kount middleware is used for app config and server/logger/etc setup
        * Brief intro to CI/CD options and how Kount's choices fit into the larger CI/CD ecosystem
    * Brief anatomy of a Kount backend orchestrator
    * Brief anatomy of Kount's Request Router
        * Brief intro to API gateway concept
        * Overview of Request Router
        * Overview of AWS API Gateway and similar cloud api-gateway alternatives
    * Brief anatomy of a Kount JS frontend service and its deployment
        * Brief intro to options for deploying a frontend/Angular app
        * How Kount uses S3 for frontend apps
        * Brief intro to Cloudfront and how Kount uses it