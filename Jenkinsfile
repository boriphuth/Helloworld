pipeline {
  agent any
  stages {
    stage('Build Steps') {
      parallel {
        stage('Pull Code Step') {
          steps {
            git(url: 'https://github.com/iPaoKung/Helloworld.git', branch: 'master', poll: true)
          }
        }
        stage('Build Code Step') {
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