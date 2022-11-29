pipeline {
    agent any

    stages {
        stage('CI') {
            parallel {
                stage('Linter') {
                    agent {
                        docker { image 'golangci/golangci-lint' }
                    }
                    when {
                        not {
                            branch 'main'
                        }
                    }
                    steps {
                        sh 'golangci-lint run --exclude-use-default=false ./...'
                    }
                }
                stage('Unit Testing') {
                    agent {
                        docker { image 'golang:1.18-alpine' }
                    }
                    when {
                        not {
                            branch 'main'
                        }
                    }
                    steps {
                        sh 'go test ./...'
                    }
                }
            }
            // stage('Build') {
            //     agent any
            //     environment {

            //     }
            //     steps {

            //     }
            // }
        }
    }
}