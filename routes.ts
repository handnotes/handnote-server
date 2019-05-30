import Router from 'koa-router'
import { jwtSecret } from './controllers/auth.controller'
import jwt from 'koa-jwt'
import { login } from './controllers/wechat.controller'

/** 无需登录即可访问的接口 */
export const router = new Router({ prefix: '/api' })

// 根据 code 获取 accessToken 并刷新 sessionKey
router.get('/wechat/login', login)

/** 受授权保护的接口 */
export const protectedRouter = new Router({ prefix: '/api' })
protectedRouter.use(jwt({ secret: jwtSecret }))

// protectedRouter.get('/user', User.getUserInfo)
