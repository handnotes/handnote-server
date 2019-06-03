import { Context } from 'koa'
import { getRepository } from 'typeorm'
import { Menstrual, MenstrualStatus } from '../entity/menstrual.entity'

export async function getMenstrual(ctx: Context) {
  const { userId } = ctx.state.user
  ctx.assert(userId, 401, '用户状态错误')

  const MenstrualModel = getRepository(Menstrual)
  let menstrual = await MenstrualModel.findOne({ userId })

  if (!menstrual) {
    return (ctx.body = { status: MenstrualStatus.disabled })
  }
  return (ctx.body = menstrual)
}

export async function setMenstrual(ctx: Context) {
  const { userId } = ctx.state.user
  const { enable, lastDate, cycle, duration } = ctx.request.body
  console.log(enable)
  let menstrual = await getRepository(Menstrual).findOne({ userId })
  if (menstrual) {
    ctx.status = 200
  } else {
    ctx.status = 201
    menstrual = await getRepository(Menstrual).create({ userId })
  }
  menstrual.status = enable ? 1 : 0
  menstrual.lastDate = new Date(lastDate)
  menstrual.cycle = cycle
  menstrual.duration = duration
  menstrual = await getRepository(Menstrual).save(menstrual)
  ctx.body = menstrual
}
