pipeline {
    agent any
    environment {
        PROJECT = "go-common"
        SERVICE = "go-common"
        DOCKER_COMPOSE_RESTFUL = "192.168.1.8:8089"
        IMAGE = "clibing/go-common"
        TAG = "latest"
    }
    stages {
        stage("Env-Print") {
            steps {
                sh 'printenv'
                sh 'docker info'
            }
        }

        stage('Docker') {
            steps {
                sh 'docker build -t ${IMAGE}:${TAG} -f build/Dockerfile .'
            }
        }
    }

    post {
        always {
            echo 'This will always run, nothind to do. skip'
            /* deleteDir()  clean up our workspace */
        }
        success {
            echo 'This will run only if successful, will auto request remote deploy'
            sh "curl --request POST \'${DOCKER_COMPOSE_RESTFUL}/deploy\' --header \'Content-Type: application/json\' --data-raw '{\"command\": \"docker-compose\", \"parameter\": [ \"up\", \"-d\", \"go-common\" ], \"type\": 2, \"project\": \"${PROJECT}\", \"service\": \"${SERVICE}\", \"image\": \"${IMAGE}\", \"tag\": \"${TAG}\"}'"
            echo "send message to mattermost: image: $IMAGE:$TAG"
            sh "curl --location --request POST 'https://mattermost.linuxcrypt.cn/api/v4/posts' --header 'Authorization: Bearer $BEARER_ACCESS_TOKEN' --header 'Content-Type: application/json' --data-raw '{\"channel_id\": \"tt1hxzf7xbdpdxz8qg6xhrn8do\", \"message\": \"deploy success! image: $IMAGE:$TAG\"}'"
        }
        failure {
            echo 'This will run only if failed'
            mail to: '458914534@qq.com',
                 subject: "Failed Pipeline: ${currentBuild.fullDisplayName}",
                 body: "Something is wrong with ${env.BUILD_URL}"
        }
        unstable {
            echo 'This will run only if the run was marked as unstable'
        }
        changed {
            echo 'This will run only if the state of the Pipeline has changed'
            echo 'For example, if the Pipeline was previously failing but is now successful'
        }
    }
}
