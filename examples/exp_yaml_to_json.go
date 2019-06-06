package main

import (
	"encoding/json"
	"fmt"
	"github.com/ghodss/yaml"
	"log"
	"strings"
)


func testYamlToJson(a string)  {

	// yaml 转 json
	yamlBytes, err := yaml.YAMLToJSON([]byte(a))
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}
	fmt.Println(string(yamlBytes))
}

func testYamlToMap(a string)  {
	// yaml 转 map
	mapObj := make(map[string]interface{})
	err := yaml.Unmarshal([]byte(a), &mapObj)
	if err != nil {
		log.Fatalln(err)
	}
	//fmt.Println(mapObj)
	// 读取嵌套map
	//fmt.Println(mapObj["image"].(map[interface{}]interface{})["repository"])

	// map 转 yaml
	yamlBytes, err := yaml.Marshal(&mapObj)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}
	fmt.Println(string(yamlBytes))

	// 打印转换内容 - Json 缩进
	jsonIndentBytes, errJI := json.MarshalIndent(a, "", "\t") //转换成JSON返回的是byte[]
	if errJI != nil {
		fmt.Println(errJI.Error())
	}
	fmt.Println(string(jsonIndentBytes))
}

func main() {
	a := "\n---\n# Source: mariadb/templates/secrets.yaml\napiVersion: v1\nkind: Secret\nmetadata:\n  name: mariadb\n  labels:\n    app: \"mariadb\"\n    chart: \"mariadb-5.11.3\"\n    release: \"mariadb\"\n    heritage: \"Tiller\"\ntype: Opaque\ndata:\n  mariadb-root-password: \"OWRwU2RDQTNzYQ==\"\n  \n  mariadb-replication-password: \"Slp0ZG5GQjdTTQ==\"\n---\n# Source: mariadb/templates/master-configmap.yaml\napiVersion: v1\nkind: ConfigMap\nmetadata:\n  name: mariadb-master\n  labels:\n    app: \"mariadb\"\n    component: \"master\"\n    chart: \"mariadb-5.11.3\"\n    release: \"mariadb\"\n    heritage: \"Tiller\"\ndata:\n  my.cnf: |-\n    [mysqld]\n    skip-name-resolve\n    explicit_defaults_for_timestamp\n    basedir=/opt/bitnami/mariadb\n    port=3306\n    socket=/opt/bitnami/mariadb/tmp/mysql.sock\n    tmpdir=/opt/bitnami/mariadb/tmp\n    max_allowed_packet=16M\n    bind-address=0.0.0.0\n    pid-file=/opt/bitnami/mariadb/tmp/mysqld.pid\n    log-error=/opt/bitnami/mariadb/logs/mysqld.log\n    character-set-server=UTF8\n    collation-server=utf8_general_ci\n    \n    [client]\n    port=3306\n    socket=/opt/bitnami/mariadb/tmp/mysql.sock\n    default-character-set=UTF8\n    \n    [manager]\n    port=3306\n    socket=/opt/bitnami/mariadb/tmp/mysql.sock\n    pid-file=/opt/bitnami/mariadb/tmp/mysqld.pid\n---\n# Source: mariadb/templates/slave-configmap.yaml\napiVersion: v1\nkind: ConfigMap\nmetadata:\n  name: mariadb-slave\n  labels:\n    app: \"mariadb\"\n    component: \"slave\"\n    chart: \"mariadb-5.11.3\"\n    release: \"mariadb\"\n    heritage: \"Tiller\"\ndata:\n  my.cnf: |-\n    [mysqld]\n    skip-name-resolve\n    explicit_defaults_for_timestamp\n    basedir=/opt/bitnami/mariadb\n    port=3306\n    socket=/opt/bitnami/mariadb/tmp/mysql.sock\n    tmpdir=/opt/bitnami/mariadb/tmp\n    max_allowed_packet=16M\n    bind-address=0.0.0.0\n    pid-file=/opt/bitnami/mariadb/tmp/mysqld.pid\n    log-error=/opt/bitnami/mariadb/logs/mysqld.log\n    character-set-server=UTF8\n    collation-server=utf8_general_ci\n    \n    [client]\n    port=3306\n    socket=/opt/bitnami/mariadb/tmp/mysql.sock\n    default-character-set=UTF8\n    \n    [manager]\n    port=3306\n    socket=/opt/bitnami/mariadb/tmp/mysql.sock\n    pid-file=/opt/bitnami/mariadb/tmp/mysqld.pid\n---\n# Source: mariadb/templates/tests.yaml\napiVersion: v1\nkind: ConfigMap\nmetadata:\n  name: mariadb-tests\ndata:\n  run.sh: |-\n    @test \"Testing MariaDB is accessible\" {\n      mysql -h mariadb -uroot -p$MARIADB_ROOT_PASSWORD -e 'show databases;'\n    }\n---\n# Source: mariadb/templates/master-svc.yaml\napiVersion: v1\nkind: Service\nmetadata:\n  name: mariadb\n  labels:\n    app: \"mariadb\"\n    component: \"master\"\n    chart: \"mariadb-5.11.3\"\n    release: \"mariadb\"\n    heritage: \"Tiller\"\nspec:\n  type: ClusterIP\n  ports:\n  - name: mysql\n    port: 3306\n    targetPort: mysql\n  selector:\n    app: \"mariadb\"\n    component: \"master\"\n    release: \"mariadb\"\n---\n# Source: mariadb/templates/slave-svc.yaml\napiVersion: v1\nkind: Service\nmetadata:\n  name: mariadb-slave\n  labels:\n    app: \"mariadb\"\n    chart: \"mariadb-5.11.3\"\n    component: \"slave\"\n    release: \"mariadb\"\n    heritage: \"Tiller\"\nspec:\n  type: ClusterIP\n  ports:\n  - name: mysql\n    port: 3306\n    targetPort: mysql\n  selector:\n    app: \"mariadb\"\n    component: \"slave\"\n    release: \"mariadb\"\n---\n# Source: mariadb/templates/master-statefulset.yaml\napiVersion: apps/v1beta1\nkind: StatefulSet\nmetadata:\n  name: mariadb-master\n  labels:\n    app: \"mariadb\"\n    chart: \"mariadb-5.11.3\"\n    component: \"master\"\n    release: \"mariadb\"\n    heritage: \"Tiller\"\nspec:\n  selector:\n    matchLabels:\n      release: \"mariadb\"\n      component: \"master\"\n      app: \"mariadb\"\n  serviceName: \"mariadb-master\"\n  replicas: 1\n  updateStrategy:\n    type: RollingUpdate\n  template:\n    metadata:\n      labels:\n        app: \"mariadb\"\n        component: \"master\"\n        release: \"mariadb\"\n        chart: \"mariadb-5.11.3\"\n    spec:\n      serviceAccountName: \"default\"\n      securityContext:\n        fsGroup: 1001\n        runAsUser: 1001\n      affinity:\n        podAntiAffinity:\n          preferredDuringSchedulingIgnoredDuringExecution:\n          - weight: 1\n            podAffinityTerm:\n              topologyKey: kubernetes.io/hostname\n              labelSelector:\n                matchLabels:\n                  app: \"mariadb\"\n                  release: \"mariadb\"      \n      containers:\n      - name: \"mariadb\"\n        image: docker.io/bitnami/mariadb:10.1.40\n        imagePullPolicy: \"IfNotPresent\"\n        env:\n        - name: MARIADB_ROOT_PASSWORD\n          valueFrom:\n            secretKeyRef:\n              name: mariadb\n              key: mariadb-root-password\n        - name: MARIADB_DATABASE\n          value: \"my_database\"\n        - name: MARIADB_REPLICATION_MODE\n          value: \"master\"\n        - name: MARIADB_REPLICATION_USER\n          value: \"replicator\"\n        - name: MARIADB_REPLICATION_PASSWORD\n          valueFrom:\n            secretKeyRef:\n              name: mariadb\n              key: mariadb-replication-password\n        ports:\n        - name: mysql\n          containerPort: 3306\n        livenessProbe:\n          exec:\n            command: [\"sh\", \"-c\", \"exec mysqladmin status -uroot -p$MARIADB_ROOT_PASSWORD\"]\n          initialDelaySeconds: 120\n          periodSeconds: 10\n          timeoutSeconds: 1\n          successThreshold: 1\n          failureThreshold: 3\n        readinessProbe:\n          exec:\n            command: [\"sh\", \"-c\", \"exec mysqladmin status -uroot -p$MARIADB_ROOT_PASSWORD\"]\n          initialDelaySeconds: 30\n          periodSeconds: 10\n          timeoutSeconds: 1\n          successThreshold: 1\n          failureThreshold: 3\n        resources:\n          {}\n          \n        volumeMounts:\n        - name: data\n          mountPath: /bitnami/mariadb\n        - name: config\n          mountPath: /opt/bitnami/mariadb/conf/my.cnf\n          subPath: my.cnf\n      volumes:\n        - name: config\n          configMap:\n            name: mariadb-master\n  volumeClaimTemplates:\n    - metadata:\n        name: data\n        labels:\n          app: \"mariadb\"\n          component: \"master\"\n          release: \"mariadb\"\n          heritage: \"Tiller\"\n      spec:\n        accessModes:\n          - \"ReadWriteOnce\"\n        resources:\n          requests:\n            storage: \"8Gi\"\n---\n# Source: mariadb/templates/slave-statefulset.yaml\napiVersion: apps/v1beta1\nkind: StatefulSet\nmetadata:\n  name: mariadb-slave\n  labels:\n    app: \"mariadb\"\n    chart: \"mariadb-5.11.3\"\n    component: \"slave\"\n    release: \"mariadb\"\n    heritage: \"Tiller\"\nspec:\n  selector:\n    matchLabels:\n      release: \"mariadb\"\n      component: \"slave\"\n      app: \"mariadb\"\n  serviceName: \"mariadb-slave\"\n  replicas: 1\n  updateStrategy:\n    type: RollingUpdate\n  template:\n    metadata:\n      labels:\n        app: \"mariadb\"\n        component: \"slave\"\n        release: \"mariadb\"\n        chart: \"mariadb-5.11.3\"\n    spec:\n      serviceAccountName: \"default\"\n      securityContext:\n        fsGroup: 1001\n        runAsUser: 1001\n      affinity:\n        podAntiAffinity:\n          preferredDuringSchedulingIgnoredDuringExecution:\n          - weight: 1\n            podAffinityTerm:\n              topologyKey: kubernetes.io/hostname\n              labelSelector:\n                matchLabels:\n                  app: \"mariadb\"\n                  release: \"mariadb\"      \n      containers:\n      - name: \"mariadb\"\n        image: docker.io/bitnami/mariadb:10.1.40\n        imagePullPolicy: \"IfNotPresent\"\n        env:\n        - name: MARIADB_REPLICATION_MODE\n          value: \"slave\"\n        - name: MARIADB_MASTER_HOST\n          value: mariadb\n        - name: MARIADB_MASTER_PORT_NUMBER\n          value: \"3306\"\n        - name: MARIADB_MASTER_ROOT_USER\n          value: \"root\"\n        - name: MARIADB_MASTER_ROOT_PASSWORD\n          valueFrom:\n            secretKeyRef:\n              name: mariadb\n              key: mariadb-root-password\n        - name: MARIADB_REPLICATION_USER\n          value: \"replicator\"\n        - name: MARIADB_REPLICATION_PASSWORD\n          valueFrom:\n            secretKeyRef:\n              name: mariadb\n              key: mariadb-replication-password\n        ports:\n        - name: mysql\n          containerPort: 3306\n        livenessProbe:\n          exec:\n            command: [\"sh\", \"-c\", \"exec mysqladmin status -uroot -p$MARIADB_MASTER_ROOT_PASSWORD\"]\n          initialDelaySeconds: 120\n          periodSeconds: 10\n          timeoutSeconds: 1\n          successThreshold: 1\n          failureThreshold: 3\n        readinessProbe:\n          exec:\n            command: [\"sh\", \"-c\", \"exec mysqladmin status -uroot -p$MARIADB_MASTER_ROOT_PASSWORD\"]\n          initialDelaySeconds: 45\n          periodSeconds: 10\n          timeoutSeconds: 1\n          successThreshold: 1\n          failureThreshold: 3\n        resources:\n          {}\n          \n        volumeMounts:\n        - name: data\n          mountPath: /bitnami/mariadb\n        - name: config\n          mountPath: /opt/bitnami/mariadb/conf/my.cnf\n          subPath: my.cnf\n      volumes:\n        - name: config\n          configMap:\n            name: mariadb-slave\n  volumeClaimTemplates:\n    - metadata:\n        name: data\n        labels:\n          app: \"mariadb\"\n          component: \"slave\"\n          release: \"mariadb\"\n          heritage: \"Tiller\"\n      spec:\n        accessModes:\n          - \"ReadWriteOnce\"\n        resources:\n          requests:\n            storage: \"8Gi\""
	b := strings.Split(a, "\n---\n")
	for _, i := range b{
		testYamlToJson(i)
		testYamlToMap(i)
		fmt.Println("----------------------")
	}
}

