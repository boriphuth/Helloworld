pipeline {
  agent any
  stages {
    stage('PullCode Step') {
      parallel {
        stage('Build Step') {
          steps {
            git(url: 'https://github.com/iPaoKung/Helloworld.git', branch: 'master', poll: true)
          }
        }
        stage('Build Step') {
          steps {
            sh 'echo "Hello World"'
          }
        }
      }
    }
    stage('Build Image') {
      steps {
        sh 'echo "docker build -t bababa"'
      }
    }
  }
}