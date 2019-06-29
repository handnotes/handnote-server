import Router from 'koa-router'
import jwt from 'koa-jwt'
import { jwtSecret } from './controllers/auth.controller'
import { login } from './controllers/wechat.controller'
import { setMenstrual } from './controllers/menstrual.controller'
import { getUserData } from './controllers/user.controller'
import { createMemorial, getAllMemorials } from './controllers/memorial.controller'

/** 无需登录即可访问的接口 */
export const router = new Router({ prefix: '/api' })

// 根据 code 获取 accessToken 并刷新 sessionKey
router.get('/wechat/login', login)

/** 受授权保护的接口 */
export const protectedRouter = new Router({ prefix: '/api' })
protectedRouter.use(jwt({ secret: jwtSecret }))

// 获取聚合数据
protectedRouter.get('/user', getUserData)

// 设置生理期
protectedRouter.put('/menstrual', setMenstrual)

// 添加纪念日
protectedRouter.post('/memorial', createMemorial)
protectedRouter.get('/memorial', getAllMemorials)
