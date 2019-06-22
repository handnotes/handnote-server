// import cache from 'memory-cache'
import { Context } from 'koa'
import axios from 'axios'
import { User } from '../entity/user.entity'
import { getManager } from 'typeorm'
import { getToken } from './auth.controller'
import { Menstrual } from '../entity/menstrual.entity'

const appid = process.env.WECHAT_MP_APPID || ''
const secret = process.env.WECHAT_MP_SECRET || ''

export async function login(ctx: Context) {
  ctx.assert(appid, 500, '服务器未配置 appId')
  ctx.assert(secret, 500, '服务器未配置 secret')

  const code = ctx.query.code
  ctx.assert(code, 400, 'code 不存在')

  let openId: string
  let sessionKey: string

  // 向微信服务器获取 session key
  try {
    const url = 'https://api.weixin.qq.com/sns/jscode2session'
    const params = { grant_type: 'authorization_code', appid, secret, js_code: code }
    interface Returns {
      openid: string
      session_key: string
    }
    const res = (await axios.get<Returns>(url, { params })).data
    openId = res.openid
    sessionKey = res.session_key
  } catch (error) {
    console.error(error.response)
    return ctx.throw(502, '获取微信授权失败')
  }

  // 存储 session key 用于后续请求
  const manager = getManager()
  let user = await manager.findOne(User, { openId })
  console.error(user)
  // 没有就创建一个
  if (!user) {
    user = manager.create(User, { openId })
    user.menstrual = new Menstrual(user.id)
    ctx.status = 201
  } else {
    ctx.status = 200
  }
  user.sessionKey = sessionKey
  user = await manager.save(user)

  const accessToken = getToken(user.id, openId, sessionKey)
  return (ctx.body = { accessToken })
}

/*
export async function loginWithEncryptPayload(ctx: Context) {
  type RequestBody = wx.GetUserInfoSuccessCallbackResult & { openId: string }
  const { openId, encryptedData, iv } = ctx.request.body as RequestBody
  ctx.assert(openId, 400, '用户状态错误: openId 不存在')

  const UserSessionModel = getRepository(UserSession)
  const userSession = await UserSessionModel.findOne({ openId })
  ctx.assert(userSession, 400, '用户状态错误: userSession 不存在')
  const { sessionKey } = userSession!
  ctx.assert(userSession, 400, '用户状态错误: sessionKey 不存在')
  const pc = new WXBizDataCrypt(appid, sessionKey!)
  const openData = pc.decryptData(encryptedData, iv)
  ctx.body = { openData }
}
*/

/*
// 获取微信服务端 access_token
export async function getWxAccessToken(ctx: Context) {
  ctx.assert(appid, 500, '服务器未配置 appId')
  ctx.assert(secret, 500, '服务器未配置 secret')

  // 从缓存读取
  let accessToken = cache.get('wxAccessToken') as string
  if (accessToken) return accessToken

  try {
    const params = { grant_type: 'client_credential', appid, secret }
    const url = 'https://api.weixin.qq.com/cgi-bin/token'
    interface Returns {
      access_token: string
      expires_in: number
    }
    const { access_token, expires_in } = (await axios.get<Returns>(url, { params })).data
    accessToken = access_token
    // 保存到缓存
    cache.put('wxAccessToken', accessToken, expires_in * 1000)
  } catch (error) {
    console.error(error)
    ctx.throw(502, '获取微信授权失败')
  }
  return accessToken
}
*/
