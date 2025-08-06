# RKE2 (Rancher Kubernetes Engine 2) - Comprehensive Guide

## Table of Contents
- [What is RKE2?](#what-is-rke2)
- [Architecture Overview](#architecture-overview)
- [Key Features and Advantages](#key-features-and-advantages)
- [When to Use RKE2](#when-to-use-rke2)
- [Installation Options](#installation-options)
- [Quick Start Examples](#quick-start-examples)
- [Security and Compliance](#security-and-compliance)
- [Configuration Options](#configuration-options)
- [High Availability Setup](#high-availability-setup)
- [Comparison with Other Kubernetes Distributions](#comparison-with-other-kubernetes-distributions)
- [Best Practices](#best-practices)
- [Troubleshooting](#troubleshooting)

---

## What is RKE2?

**RKE2** (Rancher Kubernetes Engine 2), also known as **RKE Government**, is Rancher's next-generation, enterprise-ready Kubernetes distribution designed for production workloads with enhanced security and compliance features.

### Core Characteristics:
- **Fully conformant Kubernetes distribution** certified by CNCF
- **Security-first design** with CIS Kubernetes Benchmark compliance
- **FIPS 140-2 compliance** for government and regulated industries
- **Single binary** deployment model for simplicity
- **Container runtime**: Uses `containerd` instead of Docker
- **Static pods** for control plane components instead of Docker containers

---

## Architecture Overview

RKE2 combines the best aspects of both RKE1 and K3s:

### From K3s:
- Simple installation and operation
- Single binary approach
- Ease of deployment

### From RKE1:
- Close alignment with upstream Kubernetes
- Enterprise-grade features
- Production-ready stability

### Core Components:

```
RKE2 Binary
├── containerd (Container Runtime)
├── kubelet (Node Agent)
├── Control Plane Components
│   ├── etcd (Key-Value Store)
│   ├── kube-apiserver
│   ├── kube-controller-manager
│   └── kube-scheduler
└── Add-ons
    ├── CNI (Calico/Flannel/Cilium)
    ├── CoreDNS
    ├── NGINX Ingress
    └── Metrics Server
```

### Technology Stack:
- **Container Runtime**: containerd + runc
- **CNI Options**: Canal (Calico + Flannel), Cilium, or Calico
- **DNS**: CoreDNS
- **Ingress**: NGINX Ingress Controller
- **Monitoring**: Metrics Server
- **Package Management**: Helm Controller

---

## Key Features and Advantages

### 🛡️ Security Features
- **CIS Kubernetes Benchmark compliance** (v1.7, v1.8, v1.9)
- **FIPS 140-2 compliance** for cryptographic modules
- **Pod Security Standards** enforcement
- **Network Policies** for microsegmentation
- **Audit logging** with configurable policies
- **Regular CVE scanning** with Trivy

### ⚡ Operational Benefits
- **Single binary deployment** - No complex installation procedures
- **Automatic service management** with systemd
- **Built-in high availability** support
- **Easy upgrades** and rollbacks
- **Minimal resource overhead**
- **Air-gap installation** support

### 🏢 Enterprise Features
- **Multi-architecture support** (AMD64, ARM64)
- **Windows worker node** support
- **Backup and restore** capabilities
- **Integration with Rancher** management platform
- **Professional support** from SUSE

---

## When to Use RKE2

### ✅ Ideal Use Cases:
1. **Government and regulated industries** requiring FIPS compliance
2. **Enterprise production workloads** needing security compliance
3. **Air-gapped environments** with restricted internet access
4. **Organizations requiring CIS benchmark compliance**
5. **Hybrid cloud deployments** with consistent Kubernetes experience
6. **Teams wanting simplicity** without sacrificing enterprise features

### ❌ When NOT to Use RKE2:
1. **Edge computing** scenarios (use K3s instead)
2. **Resource-constrained environments** (< 2GB RAM)
3. **Development/testing** environments needing rapid iteration
4. **IoT deployments** requiring minimal footprint

---

## Installation Options

### 1. Quick Installation Script

#### Server Node (Control Plane):
```bash
# Download and install RKE2
curl -sfL https://get.rke2.io | sh -

# Enable and start the service
sudo systemctl enable rke2-server.service
sudo systemctl start rke2-server.service

# Follow logs
journalctl -u rke2-server -f
```

#### Agent Node (Worker):
```bash
# Install RKE2 agent
curl -sfL https://get.rke2.io | INSTALL_RKE2_TYPE="agent" sh -

# Create configuration
sudo mkdir -p /etc/rancher/rke2/
sudo vim /etc/rancher/rke2/config.yaml
```

Configuration file content:
```yaml
server: https://<server-ip>:9345
token: <token-from-server>
```

```bash
# Enable and start agent service
sudo systemctl enable rke2-agent.service
sudo systemctl start rke2-agent.service
```

### 2. Air-Gap Installation

```bash
# Download RKE2 binary and images
curl -OLs https://github.com/rancher/rke2/releases/download/v1.28.3+rke2r1/rke2.linux-amd64.tar.gz
curl -OLs https://github.com/rancher/rke2/releases/download/v1.28.3+rke2r1/rke2-images.linux-amd64.tar.gz

# Install
sudo tar xzf rke2.linux-amd64.tar.gz -C /usr/local
sudo mkdir -p /var/lib/rancher/rke2/agent/images/
sudo cp rke2-images.linux-amd64.tar.gz /var/lib/rancher/rke2/agent/images/

# Create systemd service
sudo /usr/local/bin/rke2-server --help  # Generate service file
```

### 3. Windows Agent Installation

```powershell
# Download install script
Invoke-WebRequest -Uri https://raw.githubusercontent.com/rancher/rke2/master/install.ps1 -Outfile install.ps1

# Configure agent
New-Item -Type Directory c:/etc/rancher/rke2 -Force
Set-Content -Path c:/etc/rancher/rke2/config.yaml -Value @"
server: https://<server>:9345
token: <token from server node>
"@

# Configure PATH
$env:PATH+=";c:\var\lib\rancher\rke2\bin;c:\usr\local\bin"

# Run installer
./install.ps1

# Start service
rke2.exe agent service --add
Start-Service rke2
```

---

## Quick Start Examples

### Example 1: Single Node Cluster (Development)

```bash
#!/bin/bash
# install-single-node.sh

# Install RKE2
curl -sfL https://get.rke2.io | sh -

# Configure for single node (optional)
sudo mkdir -p /etc/rancher/rke2
cat <<EOF | sudo tee /etc/rancher/rke2/config.yaml
write-kubeconfig-mode: "0644"
tls-san:
  - "127.0.0.1"
  - "localhost"
node-taint:
  - "node-role.kubernetes.io/control-plane:NoSchedule"
EOF

# Start service
sudo systemctl enable --now rke2-server.service

# Wait for cluster to be ready
until sudo kubectl --kubeconfig /etc/rancher/rke2/rke2.yaml get nodes; do
  echo "Waiting for cluster..."
  sleep 5
done

echo "Cluster ready!"
echo "KUBECONFIG=/etc/rancher/rke2/rke2.yaml kubectl get nodes"
```

### Example 2: High Availability Cluster (3 Masters)

#### First Master Node:
```bash
# config.yaml for first master
cat <<EOF | sudo tee /etc/rancher/rke2/config.yaml
cluster-init: true
write-kubeconfig-mode: "0644"
tls-san:
  - "master1.example.com"
  - "master2.example.com"
  - "master3.example.com"
  - "cluster.example.com"
EOF

# Start first master
sudo systemctl enable --now rke2-server.service
```

#### Additional Master Nodes:
```bash
# config.yaml for additional masters
cat <<EOF | sudo tee /etc/rancher/rke2/config.yaml
server: https://master1.example.com:9345
token: <token-from-first-master>
write-kubeconfig-mode: "0644"
tls-san:
  - "master1.example.com"
  - "master2.example.com"
  - "master3.example.com"
  - "cluster.example.com"
EOF

sudo systemctl enable --now rke2-server.service
```

### Example 3: CIS Hardened Cluster

```yaml
# /etc/rancher/rke2/config.yaml
profile: "cis"
selinux: true
secrets-encryption: true
write-kubeconfig-mode: "0640"
use-service-account-credentials: true
kube-controller-manager-arg:
  - "bind-address=127.0.0.1"
  - "use-service-account-credentials=true"
  - "tls-min-version=VersionTLS12"
  - "tls-cipher-suites=TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305,TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305,TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384"
kube-scheduler-arg:
  - "bind-address=127.0.0.1"
  - "tls-min-version=VersionTLS12"
  - "tls-cipher-suites=TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305,TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305,TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384"
kube-apiserver-arg:
  - "tls-min-version=VersionTLS12"
  - "tls-cipher-suites=TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305,TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305,TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384"
  - "authorization-mode=RBAC,Node"
  - "anonymous-auth=false"
kubelet-arg:
  - "protect-kernel-defaults=true"
  - "read-only-port=0"
  - "authorization-mode=Webhook"
```

Pre-installation requirements for CIS profile:
```bash
# Create etcd user
sudo useradd -r -c "etcd user" -s /sbin/nologin -M etcd -U

# Configure kernel parameters
sudo cp -f /usr/share/rke2/rke2-cis-sysctl.conf /etc/sysctl.d/60-rke2-cis.conf
sudo systemctl restart systemd-sysctl
```

---

## Security and Compliance

### CIS Kubernetes Benchmark Compliance

RKE2 can pass CIS Kubernetes Benchmark with minimal configuration:

```yaml
# Automatic CIS compliance configuration
profile: "cis"
```

This automatically:
- Configures Pod Security Standards
- Applies Network Policies
- Sets appropriate file permissions
- Configures audit logging
- Enforces security contexts

### FIPS 140-2 Compliance

For FIPS compliance:
```bash
# Download FIPS-compliant binary
curl -sfL https://get.rke2.io | INSTALL_RKE2_CHANNEL=stable sh -

# Verify FIPS mode
rke2 --version
# Should show "+fips" in the version string
```

### Security Hardening Checklist

- [ ] Enable CIS profile (`profile: "cis"`)
- [ ] Configure Pod Security Standards
- [ ] Implement Network Policies
- [ ] Enable audit logging
- [ ] Use FIPS-compliant binaries (if required)
- [ ] Configure TLS cipher suites
- [ ] Disable anonymous authentication
- [ ] Use service account credentials
- [ ] Regular security updates
- [ ] CVE scanning of images

---

## Configuration Options

### Common Configuration Parameters

```yaml
# /etc/rancher/rke2/config.yaml

# Basic cluster configuration
token: "my-shared-secret"
server: "https://master.example.com:9345"
datacenter: "dc1"
node-name: "node-1"

# Network configuration
cluster-cidr: "10.42.0.0/16"
service-cidr: "10.43.0.0/16"
cluster-dns: "10.43.0.10"
cni: "calico"  # Options: calico, canal, cilium

# Security configuration
profile: "cis"
selinux: true
secrets-encryption: true
protect-kernel-defaults: true

# Audit configuration
audit-log-path: "/var/log/audit.log"
audit-log-maxage: 30
audit-log-maxbackup: 10
audit-log-maxsize: 100

# Networking
disable:
  - rke2-ingress-nginx  # Disable default ingress
node-ip: "192.168.1.100"
node-external-ip: "203.0.113.100"

# Resource limits
kubelet-arg:
  - "max-pods=110"
  - "cluster-dns=10.43.0.10"
  - "cluster-domain=cluster.local"

# Custom registries
system-default-registry: "my-registry.com"
```

### Advanced Networking Configuration

```yaml
# CNI-specific configurations

# Calico configuration
cni: "calico"
kubelet-arg:
  - "cluster-dns=10.43.0.10"

# Cilium configuration
cni: "cilium"
disable-kube-proxy: true  # Cilium can replace kube-proxy

# Canal (Calico + Flannel) configuration
cni: "canal"
flannel-backend: "vxlan"  # Options: vxlan, host-gw, wireguard
```

---

## High Availability Setup

### External Database (PostgreSQL/MySQL)

```yaml
# config.yaml for HA with external DB
datastore-endpoint: "postgres://username:password@hostname:5432/database"
datastore-cafile: "/path/to/ca.crt"
datastore-certfile: "/path/to/cert.crt"  
datastore-keyfile: "/path/to/key.key"
```

### Embedded etcd Cluster

For production HA with embedded etcd:

```yaml
# First server (bootstrap)
cluster-init: true
tls-san:
  - "server1.example.com"
  - "server2.example.com"  
  - "server3.example.com"
  - "cluster.example.com"

# Additional servers
server: https://server1.example.com:9345
token: <bootstrap-token>
tls-san:
  - "server1.example.com"
  - "server2.example.com"
  - "server3.example.com"
  - "cluster.example.com"
```

### Load Balancer Configuration

Example NGINX configuration for RKE2 cluster:
```nginx
upstream rke2_servers {
    server server1.example.com:6443;
    server server2.example.com:6443;
    server server3.example.com:6443;
}

upstream rke2_registration {
    server server1.example.com:9345;
    server server2.example.com:9345;
    server server3.example.com:9345;
}

server {
    listen 6443;
    proxy_pass rke2_servers;
}

server {
    listen 9345;
    proxy_pass rke2_registration;
}
```

---

## Comparison with Other Kubernetes Distributions

| Feature | RKE2 | K3s | RKE1 | kubeadm |
|---------|------|-----|------|---------|
| **Target Use Case** | Enterprise/Government | Edge/IoT | Enterprise Legacy | DIY Clusters |
| **Installation** | Single Binary | Single Binary | Docker-based | Multi-step |
| **Resource Usage** | Medium | Low | High | Medium |
| **Security Focus** | High (CIS/FIPS) | Medium | Medium | Basic |
| **HA Support** | Built-in | Built-in | External LB Required | Manual Setup |
| **Container Runtime** | containerd | containerd | Docker | configurable |
| **Upgrade Process** | Automated | Automated | Manual/Rancher | Manual |
| **Air-gap Support** | Excellent | Good | Good | Limited |
| **Windows Nodes** | Yes | No | Yes | Yes |

---

## Best Practices

### 🏗️ Deployment Best Practices

1. **Plan your architecture**:
   - Odd number of control plane nodes (3 or 5)
   - Separate etcd nodes for large clusters (>100 nodes)
   - Use external load balancers

2. **Resource allocation**:
   - Minimum 2GB RAM, 2 CPU for server nodes
   - Minimum 1GB RAM, 1 CPU for agent nodes
   - Dedicated disks for etcd (SSD recommended)

3. **Network considerations**:
   - Ensure proper firewall rules
   - Plan IP address ranges (avoid conflicts)
   - Use proper DNS resolution

### 🔒 Security Best Practices

1. **Use CIS profile** for compliance
2. **Enable audit logging** with proper rotation
3. **Regular updates** of RKE2 and system packages
4. **Network segmentation** with Network Policies
5. **RBAC configuration** with least privilege principle
6. **Secret management** with external secret stores

### 📊 Monitoring and Maintenance

1. **Monitor cluster health**:
   ```bash
   # Check cluster status
   kubectl --kubeconfig /etc/rancher/rke2/rke2.yaml get nodes
   kubectl --kubeconfig /etc/rancher/rke2/rke2.yaml get pods -A
   
   # Check RKE2 service status
   systemctl status rke2-server
   journalctl -u rke2-server -f
   ```

2. **Regular backups**:
   ```bash
   # Backup etcd
   rke2 etcd-snapshot save --name backup-$(date +%Y%m%d-%H%M%S)
   
   # List snapshots
   rke2 etcd-snapshot list
   ```

3. **Upgrade process**:
   ```bash
   # Upgrade RKE2
   curl -sfL https://get.rke2.io | INSTALL_RKE2_VERSION=v1.28.3+rke2r1 sh -
   systemctl restart rke2-server
   ```

---

## Troubleshooting

### Common Issues and Solutions

#### 1. Service Won't Start
```bash
# Check service status
systemctl status rke2-server
journalctl -u rke2-server --no-pager

# Common solutions:
# - Check configuration syntax
# - Verify token and server URL
# - Check firewall rules
# - Ensure sufficient resources
```

#### 2. Nodes Not Joining
```bash
# On agent node, check:
journalctl -u rke2-agent --no-pager

# Verify:
# - Server URL is accessible
# - Token is correct
# - Port 9345 is open
# - Time synchronization between nodes
```

#### 3. Pod Networking Issues
```bash
# Check CNI status
kubectl get pods -n kube-system | grep -E "(canal|calico|cilium)"

# Check node network configuration
ip route show
iptables -L
```

#### 4. Certificate Issues
```bash
# Check certificate validity
openssl x509 -in /var/lib/rancher/rke2/server/tls/server-ca.crt -text -noout

# Regenerate certificates (if needed)
rm -rf /var/lib/rancher/rke2/server/tls/
systemctl restart rke2-server
```

### Diagnostic Commands

```bash
# Cluster information
kubectl cluster-info
kubectl get nodes -o wide
kubectl get pods -A

# RKE2 specific
rke2 --version
rke2 certificate list

# System resources
df -h
free -m
systemctl list-units --failed
```

### Log Locations

- **RKE2 Service Logs**: `journalctl -u rke2-server` or `journalctl -u rke2-agent`
- **Containerd Logs**: `journalctl -u containerd`
- **Pod Logs**: `/var/log/pods/`
- **Audit Logs**: `/var/lib/rancher/rke2/server/logs/audit.log`
- **etcd Logs**: Check etcd pod logs in kube-system namespace

---

## Conclusion

RKE2 is an excellent choice for organizations requiring:
- **Security and compliance** (CIS, FIPS)
- **Enterprise-grade** Kubernetes distribution
- **Simple deployment** and operations
- **Government or regulated** environment compatibility

With its single-binary approach, built-in security features, and enterprise support, RKE2 bridges the gap between the simplicity of K3s and the feature richness required for production enterprise workloads.

### Next Steps
1. **Evaluate** your requirements against RKE2 capabilities
2. **Plan** your cluster architecture (HA, networking, security)
3. **Test** deployment in a non-production environment
4. **Implement** monitoring and backup strategies
5. **Train** your team on RKE2-specific operations

---

**Resources:**
- [Official RKE2 Documentation](https://docs.rke2.io/)
- [GitHub Repository](https://github.com/rancher/rke2)
- [SUSE Support](https://www.suse.com/support/)
- [Community Forums](https://forums.rancher.com/)
