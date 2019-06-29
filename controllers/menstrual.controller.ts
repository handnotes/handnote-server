import { Context } from 'koa'
import { UserModel } from '../model/user.model'

/** 设置生理期 */
export async function setMenstrual(ctx: Context) {
  const { userId } = ctx.state.user
  const { enable, lastDate, cycle, duration } = ctx.request.body
  await UserModel.update(
    { _id: userId },
    {
      menstrual: {
        status: enable ? 1 : 0,
        cycle,
        duration,
        lastDate: new Date(lastDate),
      },
    },
  )
  ctx.status = 204
}
