pipeline {
    agent any

    stages {
        stage("Build") {
            steps {
                sh "go build -o 'proj' cmd/main.go"
            }
        }
    }
}