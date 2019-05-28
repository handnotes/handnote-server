import { getRepository } from 'typeorm'
import { User } from '../entity/user.entity'
import jwt from 'jsonwebtoken'
import { Context } from 'koa'

const expiresIn = 24 * 3600 // 24h
export const jwtSecret = process.env.JWT_SECRET || 'HandNote!'

export class AuthController {
  static getToken(openId: string) {
    return jwt.sign({ openId }, jwtSecret, { expiresIn })
  }
  static parseToken(token: string) {
    return jwt.verify(token, jwtSecret)
  }

  async login(ctx: Context) {
    const UserModel = getRepository(User)
    const input = ctx.request.body
    let user = await UserModel.findOne({ openId: input.openId })
    if (!user) {
      user = UserModel.create(input) as any
      user = await UserModel.save(user as any)
    }
    const token = AuthController.getToken(input.openId)

    ctx.body = {
      token,
      expiresAt: Date.now() + expiresIn * 1000,
    }
  }
}

export default new AuthController()
