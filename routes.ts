import Router from 'koa-router'
import Auth, { jwtSecret } from './controllers/auth.controller'
import jwt from 'koa-jwt'
import User from './controllers/user.controller'

export const router = new Router({ prefix: '/api' })
export const protectedRouter = new Router({ prefix: '/api' })
protectedRouter.use(jwt({ secret: jwtSecret }))

protectedRouter.get('/user', User.getUserInfo)

router.post('/auth/login', Auth.login)
