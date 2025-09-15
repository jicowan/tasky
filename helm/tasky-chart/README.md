# Tasky Helm Chart

A Helm chart for deploying the Tasky todo application on Kubernetes.

## Prerequisites

- Kubernetes 1.18+
- Helm 3.0+
- Sealed Secrets Controller (if using SealedSecrets)

## Installing the Chart

To install the chart with the release name `my-tasky`:

```bash
helm install my-tasky ./helm/tasky-chart
```

## Uninstalling the Chart

To uninstall/delete the `my-tasky` deployment:

```bash
helm delete my-tasky
```

## Configuration

The following table lists the configurable parameters of the Tasky chart and their default values.

### Application Configuration

| Parameter | Description | Default |
|-----------|-------------|---------|
| `replicaCount` | Number of replicas | `1` |
| `image.repository` | Image repository | `jicowan/tasky` |
| `image.tag` | Image tag | `latest` |
| `image.pullPolicy` | Image pull policy | `Always` |
| `containerPort` | Container port | `8080` |

### Service Configuration

| Parameter | Description | Default |
|-----------|-------------|---------|
| `service.type` | Service type | `ClusterIP` |
| `service.port` | Service port | `80` |
| `service.targetPort` | Target port | `8080` |
| `service.protocol` | Service protocol | `TCP` |

### Ingress Configuration

| Parameter | Description | Default |
|-----------|-------------|---------|
| `ingress.enabled` | Enable ingress | `true` |
| `ingress.className` | Ingress class name | `""` |
| `ingress.annotations` | Ingress annotations | See values.yaml |
| `ingress.hosts` | Ingress hosts | `[{host: "tasky.jicomusic.com", paths: [{path: "/", pathType: "Prefix"}]}]` |
| `ingress.tls` | Ingress TLS configuration | `[]` |

### RBAC Configuration

| Parameter | Description | Default |
|-----------|-------------|---------|
| `rbac.create` | Create RBAC resources | `true` |
| `rbac.clusterRole` | ClusterRole to bind | `cluster-admin` |
| `serviceAccount.create` | Create service account | `true` |
| `serviceAccount.name` | Service account name | `""` |

### SealedSecret Configuration

| Parameter | Description | Default |
|-----------|-------------|---------|
| `sealedSecret.enabled` | Enable SealedSecret | `true` |
| `sealedSecret.name` | SealedSecret name | `tasky-secrets` |
| `sealedSecret.encryptedData` | Encrypted data | See values.yaml |

### Resource Configuration

| Parameter | Description | Default |
|-----------|-------------|---------|
| `resources` | Resource limits and requests | `{}` |
| `nodeSelector` | Node selector | `{}` |
| `tolerations` | Tolerations | `[]` |
| `affinity` | Affinity | `{}` |

### Probes Configuration

| Parameter | Description | Default |
|-----------|-------------|---------|
| `probes.liveness.enabled` | Enable liveness probe | `false` |
| `probes.readiness.enabled` | Enable readiness probe | `false` |

## Examples

### Basic Installation

```bash
helm install tasky ./helm/tasky-chart
```

### Custom Values

```bash
helm install tasky ./helm/tasky-chart \
  --set replicaCount=3 \
  --set image.tag=v1.2.0 \
  --set ingress.hosts[0].host=my-tasky.example.com
```

### Using Custom Values File

Create a `custom-values.yaml` file:

```yaml
replicaCount: 2
image:
  tag: "v1.2.0"
  pullPolicy: IfNotPresent

ingress:
  hosts:
    - host: my-tasky.example.com
      paths:
        - path: /
          pathType: Prefix

resources:
  limits:
    cpu: 500m
    memory: 512Mi
  requests:
    cpu: 250m
    memory: 256Mi
```

Then install:

```bash
helm install tasky ./helm/tasky-chart -f custom-values.yaml
```

## Security Considerations

- The default configuration uses `cluster-admin` permissions, which should be restricted in production
- Consider enabling security contexts and resource limits
- Update the SealedSecret with your own encrypted data
- Use specific image tags instead of `latest` for production deployments

## Upgrading

To upgrade an existing release:

```bash
helm upgrade tasky ./helm/tasky-chart
```

## Troubleshooting

1. Check pod status:
   ```bash
   kubectl get pods -l app.kubernetes.io/name=tasky
   ```

2. View logs:
   ```bash
   kubectl logs -l app.kubernetes.io/name=tasky
   ```

3. Check service:
   ```bash
   kubectl get svc -l app.kubernetes.io/name=tasky
   ```

4. Verify ingress:
   ```bash
   kubectl get ingress -l app.kubernetes.io/name=tasky
