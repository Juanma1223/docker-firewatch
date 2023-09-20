pipeline {
    agent any
    environment {
        DOCKER_TAG = 'docker-firewatch'
        DOCKER_API_VERSION = '1.41'
        PORT='4200'
        CONFIG_FILES_DIR = '/go/bin/'
    }
    stages {
        stage('Build') {
            steps {
                echo 'Building..'
                sh 'docker build -t $DOCKER_TAG .'
            }
        }
        stage('Deploy') {
            steps {
                echo 'Stopping previous version...'
                sh 'docker stop $DOCKER_TAG || echo Nothing to stop'
                sh 'docker rm $DOCKER_TAG || echo Nothing to delete'
                echo 'Deploying....'
                sh 'docker run -d -e CONFIG_FILES_DIR -e PORT -v /var/run/docker.sock:/var/run/docker.sock -e DOCKER_API_VERSION --net=host --name $DOCKER_TAG $DOCKER_TAG'
            }
        }
    }
}
