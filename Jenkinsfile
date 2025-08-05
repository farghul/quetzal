pipeline {
    agent { label "cactuar && deploy" }
    options {
        buildDiscarder logRotator(
            artifactDaysToKeepStr: "28",
            artifactNumToKeepStr: "5",
            daysToKeepStr: "56",
            numToKeepStr: "10"
        )
    }
    stages {
        stage('Clean WS') {
            steps {
                cleanWs()
            }
        }
        stage("Checkout Quetzal") {
            steps {
                checkout scmGit(
                    branches: [[name: 'main']],
                    userRemoteConfigs: [[url: 'https://github.com/farghul/quetzal.git']]
                )
            }
        }
        stage("Build Quetzal") {
            steps {
                script {
                    sh "/data/apps/go/bin/go build -o /data/automation/bin/quetzal"
                }
            }
        }
        stage("Checkout DAC") {
            steps {
                checkout scmGit(
                    branches: [[name: 'main']],
                    userRemoteConfigs: [[credentialsId: 'DES-Project', url: 'https://bitbucket.org/bc-gov/desso-automation-conf.git']]
                )
            }
        }
        stage('Run Quetzal') {
            steps {
                script {
                    sh './scripts/plugin/quetzal.sh'
                }
            }
        }
    }
}