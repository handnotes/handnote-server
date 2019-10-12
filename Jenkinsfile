pipeline {
  agent any
  stages {
    stage('检出') {
      steps {
        checkout([$class: 'GitSCM', branches: [[name: env.GIT_BUILD_REF]], 
            userRemoteConfigs: [[url: env.GIT_REPO_URL, credentialsId: env.CREDENTIALS_ID]]])
      }
    }

    stage('测试') {
      steps {
        echo '单元测试中...'
        sh 'GIN_MODE=test go test ./test'
        echo '单元测试完成.'
      }
    }

    stage('构建文档') {
      steps {
        echo '构建中...'
        sh 'swagger generate spec -o ./swagger.yml'
        echo '构建完成.'
      }
    }

    stage('部署') {
      steps {
        echo '部署中...'
        echo '部署完成'
      }
    }
  }
}
