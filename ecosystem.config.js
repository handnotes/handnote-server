module.exports = {
  apps: [
    {
      name: 'dev',
      script: './src/index.ts',

      instances: 1,
      exec_mode: 'fork',
      autorestart: true,
      watch: true,
      ignore_watch: ['node_modules', 'public', 'logs'],
      listen_timeout: 8000,
      kill_timeout: 1600,
      max_memory_restart: '1G',
      env: {
        NODE_ENV: 'development',
      },
    },
    {
      name: 'prod',
      script: './src/index.ts',

      instances: 1,
      exec_mode: 'fork',
      autorestart: true,
      watch: false,
      out_file: './logs/app.log',
      error_file: './logs/error.log',
      log_date_format: 'YYYY-MM-DD HH:mm Z',
      combine_logs: true,
      listen_timeout: 8000,
      kill_timeout: 1600,
      max_memory_restart: '1G',
      env: {
        NODE_ENV: 'production',
      },
    },
  ],

  deploy: {
    production: {
      user: 'node',
      host: '212.83.163.1',
      ref: 'origin/master',
      repo: 'git@github.com:repo.git',
      path: '/var/www/production',
      'post-deploy': 'npm install && pm2 reload ecosystem.config.js --env production',
    },
  },
}
