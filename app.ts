import 'dotenv/config'
import bodyParser from 'koa-bodyparser'
import jwt from 'koa-jwt'
import Koa from 'koa'
import { router, protectedRouter } from './routes'

// Initial Application
const app = new Koa()

// Register error handler
app.use(async (ctx, next) => {
  try {
    await next()
    if (ctx.status >= 200 && ctx.status < 300 && !ctx.body) {
      ctx.body = {}
    }
  } catch (error) {
    ctx.status = error.statusCode || error.status || 500
    error = {
      status: ctx.status,
      message: error.message,
    }
    ctx.body = { error }
    ctx.app.emit('error', error, ctx)
  }
})

// Register body parser
app.use(bodyParser())

// Register json web token
const secret = process.env.JWT_SRCRET || 'HandNote!'
app.use(jwt({ secret, passthrough: true }).unless({ path: [/^\/api\/auth/] }))

// Register normal routes
app.use(router.routes())
app.use(router.allowedMethods())

// Register protected routes
app.use(protectedRouter.routes())
app.use(protectedRouter.allowedMethods())

// Application error logging
app.on('error', console.error)

export default app
