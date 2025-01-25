pipeline {
    agent { label 'cactuar && deploy' }
    options {
        buildDiscarder logRotator(
            artifactDaysToKeepStr: "28",
            artifactNumToKeepStr: "5",
            daysToKeepStr: "56",
            numToKeepStr: "10"
        )
    }
    stages {
        stage('Sync') {
            steps {
                lock('satis-rebuild-resource') {
                    dir("/data/scripts/automation/github/quetzal") {
                        sh 'git pull'
                    }
                }
            }
        }
        stage('Build') {
            steps {
                lock('satis-rebuild-resource') {
                    dir("/data/scripts/automation/github/quetzal") {
                        sh '/data/apps/go/bin/go build -o /data/scripts/automation/programs/quetzal .'
                    }
                }
            }
        }
        stage('Premium') {
            steps {
                lock('satis-rebuild-resource') {
                    timeout(time: 5, unit: 'MINUTES') {
                        retry(2) {
                            sh '/data/scripts/automation/scripts/run_quetzal.sh'
                        }
                    }
                }
            }
        }
    }
}