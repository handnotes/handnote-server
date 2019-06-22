import { Context } from 'koa'
import { getRepository } from 'typeorm'
import { User } from '../entity/user.entity'

/** 设置生理期 */
export async function setMenstrual(ctx: Context) {
  const { userId } = ctx.state.user
  const { enable, lastDate, cycle, duration } = ctx.request.body
  const userRepo = getRepository(User)

  await userRepo.update(userId, {
    menstrual: {
      status: enable ? 1 : 0,
      cycle,
      duration,
      lastDate: new Date(lastDate),
    },
  })
  ctx.status = 204
}
