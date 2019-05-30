import { getRepository } from 'typeorm'
import { User } from '../entity/user.entity'
import { Context } from 'koa'

export async function getUserInfo(ctx: Context) {
  const UserModel = getRepository(User)
  const { unionId } = ctx.state.user
  const user = await UserModel.findOne({ unionId })
  ctx.body = user
}
