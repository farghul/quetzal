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
        stage("Pull Changes") {
            steps {
                lock("satis-rebuild-resource") {
                    dir("/data/automation/github/quetzal") {
                        sh '''#!/bin/bash
                        source ~/.bashrc
                        git fetch --all
                        git pull
                        '''
                    }
                }
            }
        }
        stage("Build Quetzal") {
            steps {
                lock("satis-rebuild-resource") {
                    dir("/data/automation/github/quetzal") {
                        sh "/data/apps/go/bin/go build -o /data/automation/bin/quetzal"
                    }
                }
            }
        }
        stage("Run Quetzal") {
            steps {
                lock("satis-rebuild-resource") {
                    timeout(time: 5, unit: "MINUTES") {
                        retry(2) {
                            dir("/data/automation/bitbucket/desso-automation-conf/scripts") {
                                sh "quetzal.sh"
                            }
                        }
                    }
                }
            }
        }
    }
}