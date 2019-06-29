import { UserModel } from '../model/user.model'
import { Context } from 'koa'
import _ from 'lodash'

export async function getUserData(ctx: Context) {
  const { userId } = ctx.state.user
  const user = await UserModel.findOne({ _id: userId })
  ctx.assert(user, 401, 'invalid access token')
  ctx.body = _.pick(user, ['id', 'gender', 'menstrual', 'memorials'])
}
