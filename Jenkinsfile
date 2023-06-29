pipeline {
    agent any

    stages {
        stage("Build") {
            steps {
//                 sh "ls -la"
                sh "go build -o 'proj' cmd/main.go"
            }
        }
    }
}