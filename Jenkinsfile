pipeline {
    agent any

    tools { go '1.19' }

    stages {
        stage("Build") {
            steps {
                sh "go version"
                sh "go build -o 'proj' cmd/main.go"
                sh "ls -la"
            }
        }
    }
}