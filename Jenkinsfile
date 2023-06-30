pipeline {
    agent any

    tools { go '1.19' }

    stages {
        stage("Pre-build") {
            steps {
                sh "go mod tidy"
            }
        }
        stage("Build") {
            steps {
                sh "go build -o 'proj' cmd/main.go"
                sh "ls -la | grep proj"
            }
        }
        stage("Test") {
            steps {
                sh "go test ./..."
            }
        }
    }
    post {
        success {
            echo "Successfully done"
        }
        failure {
            echo "Some error occurred"
        }
    }
}