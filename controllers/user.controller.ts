import { getRepository } from 'typeorm'
import { User } from '../entity/user.entity'
import { Context } from 'koa'

export class UserController {
  async getUserInfo(ctx: Context) {
    const UserModel = getRepository(User)
    const { openId } = ctx.state.user
    const user = await UserModel.findOne({ openId })
    ctx.body = user
  }
}

export default new UserController()
