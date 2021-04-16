/*


Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package templates

// PrometheusTemplate is the template that serves as the base for the prometheus deployed by the operator
const PrometheusTemplate = `
apiVersion: monitoring.coreos.com/v1
kind: Prometheus
spec:
  serviceAccountName: prometheus-k8s
  serviceMonitorSelector:
    matchLabels:
      app: managed-ocs
  ruleSelector:
    matchLabels:
      app: managed-ocs
  podMonitorSelector:
    matchLabels:
      app: managed-ocs
  resources:
    requests:
      cpu: 1
      memory: 200Mi
  enableAdminAPI: false
  alerting:
    alertmanagers:
    - namespace: openshift-storage
      name: alertmanager-operated
      port: web
`
