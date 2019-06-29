import { Context } from 'koa'
import { UserModel } from '../model/user.model'
import { Memorial } from '../model/memorial.model'

export async function createMemorial(ctx: Context) {
  const { userId } = ctx.state.user
  const { type, date, person } = ctx.request.body
  const memorial = new Memorial(type, date, person)

  await UserModel.updateOne(
    { _id: userId },
    {
      $push: {
        memorials: memorial,
      },
    },
  )
  ctx.status = 201
  ctx.body = { memorial }
}

export async function getAllMemorials(ctx: Context) {
  const { userId } = ctx.state.user
  const user = await UserModel.findById(userId, 'memorials')
  if (!user) return ctx.throw(400, '用户不存在')
  ctx.body = user.memorials
}
