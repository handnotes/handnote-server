import { getRepository } from 'typeorm'
import { User } from '../entity/user.entity'
import { Context } from 'koa'

class AuthController {
  async login(ctx: Context) {
    const UserModel = getRepository(User)
    const input = ctx.request.body
    let user = await UserModel.findOne({ openId: input.openId })
    if (!user) {
      user = UserModel.create(input) as any
      user = await UserModel.save(user as any)
    }
    ctx.body = user
  }
}

export default new AuthController()
