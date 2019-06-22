import { getManager } from 'typeorm'
import { User } from '../entity/user.entity'
import { Context } from 'koa'
import _ from 'lodash'

export async function getUserData(ctx: Context) {
  const { userId } = ctx.state.user
  const manager = getManager()
  const user = await manager.findOne(User, userId)
  ctx.body = _.pick(user, ['id', 'menstrual'])
}
