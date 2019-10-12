pipeline {
  agent any

  environment {
    GOPROXY = "https://goproxy.io"
    GIN_MODE = "test"
  }

  stages {
    stage('检出') {
      steps {
        checkout([
            $class           : 'GitSCM',
            branches         : [
                [
                    name: env.GIT_BUILD_REF
                ]
            ],
            userRemoteConfigs: [
                [
                    url          : env.GIT_REPO_URL,
                    credentialsId: env.CREDENTIALS_ID
                ]
            ]
        ])
      }
    }

    stage('安装依赖') {
      steps {
        echo '安装依赖中...'
        sh 'go version'
        sh 'go get'
        echo '依赖安装完成'
      }
    }

    stage('测试') {
      steps {
        echo '单元测试中...'
        sh 'go test ./test'
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
