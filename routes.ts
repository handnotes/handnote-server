import Router from 'koa-router'
import Auth from './controllers/auth.controller'

export const router = new Router({ prefix: '/api' })
export const protectedRouter = new Router({ prefix: '/api' })

router.get('/user', (ctx: any) => {
  ctx.body = 'hello world'
})

router.post('/auth/login', Auth.login)
