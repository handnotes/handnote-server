import mongoose from 'mongoose'
import { arrayProp, ModelType, pre, prop, staticMethod, Typegoose } from 'typegoose'
import { Memorial } from './memorial.model'
import { Menstrual } from './menstrual.model'

const unique = true
const required = true

enum Gender {
  SECRET = 0,
  MALE,
  FEMALE,
}

@pre<User>('save', function(next) {
  this.updatedAt = new Date()
  next()
})
export class User extends Typegoose {
  @staticMethod
  public static findByOpenId(this: ModelType<User> & typeof User, openId: string) {
    return this.findOne({ openId })
  }
  @prop({ unique, required })
  public openId!: string

  @prop()
  public sessionKey?: string

  @prop()
  public name?: string

  @prop()
  public email?: string

  @prop()
  public avatar?: string

  @prop({ enum: Gender, default: Gender.SECRET })
  public gender: Gender = Gender.SECRET

  @prop()
  public address?: string

  @prop({ _id: false })
  public menstrual?: Menstrual

  @arrayProp({ items: Memorial, _id: false })
  public memorials?: Memorial[]

  @prop({ required, default: new Date() })
  public createdAt!: Date

  @prop({ required, default: new Date() })
  public updatedAt!: Date
}

export const UserModel = new User().getModelForClass(User, {
  existingConnection: mongoose.connection,
})
