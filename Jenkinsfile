pipeline {
    agent any

    stages {
        stage("Build") {
            steps {
                go build -o "proj" cmd/main.go
            }
        }
    }
}