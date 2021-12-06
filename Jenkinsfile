pipeline {
  agent any
  options {
    timestamps()
    timeout(time: 3600, unit: 'SECONDS')
  }
  environment{
    BUILD_PATH="/home/jenkins/gopath/src/github.com/cloudtrust/keycloak-client/v2"
  }
  stages {
    stage('Build') {
      agent {
        label 'jenkins-slave-go-ct'
      }
      steps {
        script {
          sh 'printenv'
          def isBranch = ""
          if (!env.CHANGE_ID) {
            isBranch = " || true"
          }
          withCredentials([usernamePassword(credentialsId: 'cloudtrust-cicd-sonarqube', usernameVariable: 'USER', passwordVariable: 'PASS')]) {
            sh """
              set -eo pipefail

              mkdir -p "${BUILD_PATH}"
              cp -r "${WORKSPACE}/." "${BUILD_PATH}/"
              cd "${BUILD_PATH}"

              golint ./... | tee golint.out || true

              go generate ./...
              go mod vendor

              go test -coverprofile=coverage.out -json ./... | tee report.json
              go tool cover -func=coverage.out
              bash -c \"go vet ./... > >(cat) 2> >(tee govet.out)\" || true
              gometalinter --vendor --disable=gotype --disable=golint --disable=vet --disable=gocyclo --exclude=/usr/local/go/src --deadline=300s ./... | tee gometalinter.out || true

              go list -json -deps | nancy -no-color || true

              JAVA_TOOL_OPTIONS="" sonar-scanner \
                -Dsonar.host.url=https://sonarqube-cloudtrust-cicd.openshift.west.ch.elca-cloud.com \
                -Dsonar.login="${USER}" \
                -Dsonar.password="${PASS}" \
                -Dsonar.sourceEncoding=UTF-8 \
                -Dsonar.projectKey=keycloak-client \
                -Dsonar.projectName=keycloak-client \
                -Dsonar.projectVersion="${env.GIT_COMMIT}" \
                -Dsonar.sources=. \
                -Dsonar.exclusions=**/*_test.go,**/vendor/**,**/mock/** \
                -Dsonar.tests=. \
                -Dsonar.test.inclusions=**/*_test.go \
                -Dsonar.test.exclusions=**/vendor/** \
                -Dsonar.go.coverage.reportPaths=./coverage.out \
                -Dsonar.go.tests.reportPaths=./report.json \
                -Dsonar.go.govet.reportPaths=./govet.out \
                -Dsonar.go.golint.reportPaths=./golint.out \
                -Dsonar.go.gometalinter.reportPaths=./gometalinter.out ${isBranch}
            """
          }
        }
      }
    }
  }
}
