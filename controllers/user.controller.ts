import { getRepository } from 'typeorm'
import { User } from '../entity/user.entity'
import { Context } from 'koa'
import { Menstrual } from '../entity/menstrual.entity'

export async function getUserData(ctx: Context) {
  const { userId } = ctx.state.user
  const [user, menstrual] = await Promise.all([
    getRepository(User).findOne({ id: userId }),
    getRepository(Menstrual).findOne({ userId }),
  ])
  ctx.body = { user, menstrual }
}
