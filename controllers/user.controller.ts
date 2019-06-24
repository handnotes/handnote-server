import { getMongoRepository } from 'typeorm'
import { User } from '../entity/user.entity'
import { Context } from 'koa'
import _ from 'lodash'

export async function getUserData(ctx: Context) {
  const { userId } = ctx.state.user
  const manager = getMongoRepository(User)
  const user = await manager.findOne(userId)
  ctx.assert(user, 401, 'invalid access token')
  ctx.body = _.pick(user, ['id', 'menstrual'])
}
