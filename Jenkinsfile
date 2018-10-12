pipeline {
  agent any
  stages {
    stage('Build Step') {
      steps {
        git(url: 'https://github.com/iPaoKung/Helloworld.git', branch: 'master', poll: true)
      }
    }
  }
}