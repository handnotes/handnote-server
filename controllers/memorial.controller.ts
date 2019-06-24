import { Context } from 'koa'
import { getMongoRepository } from 'typeorm'
import { User } from '../entity/user.entity'
import { Memorial } from '../entity/memorial.entity'

export async function createMemorial(ctx: Context) {
  const { userId } = ctx.state.user
  const { type, date, person } = ctx.request.body
  const userRepo = getMongoRepository(User)
  const memorial = new Memorial(type, date, person)

  const { modifiedCount } = await userRepo.updateOne(
    { id: userId },
    {
      $push: {
        memorials: memorial,
      },
    },
  )
  ctx.assert(modifiedCount, 500, '未更新')
  ctx.status = 201
}
