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
        stage("Empty_Folder") {
            steps {
                dir('/data/automation/checkouts'){
                    script {
                        deleteDir()
                    }
                }
            }
        }
        stage('Checkout_Quetzal'){
            steps{
                dir('/data/automation/checkouts/quetzal'){
                    git url: 'https://github.com/farghul/quetzal.git' , branch: 'main'
                }
            }
        }
        stage('Build_Quetzal') {
            steps {
                dir('/data/automation/checkouts/quetzal'){
                    script {
                        sh "/data/apps/go/bin/go build -o /data/automation/bin/quetzal"
                    }
                }
            }
        }
        stage("Checkout_DAC") {
            steps{
                dir('/data/automation/checkouts/dac'){
                    git credentialsId: 'DES-Project', url: 'https://bitbucket.org/bc-gov/desso-automation-conf.git', branch: 'main'
                }
            }
        }
        stage('Run_Quetzal') {
            steps {
                dir('/data/automation/checkouts/dac'){
                    script {
                        sh './scripts/plugin/quetzal.sh'
                    }
                }
            }
        }
    }
    post {
        always {
            cleanWs(cleanWhenNotBuilt: false,
                deleteDirs: true,
                disableDeferredWipeout: true,
                notFailBuild: true,
                patterns: [[pattern: '.gitignore', type: 'INCLUDE'], [pattern: '.propsfile', type: 'EXCLUDE']]
            )
        }
    }
}