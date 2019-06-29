import { prop } from 'typegoose'

export enum MenstrualStatus {
  disabled,
  enable,
}

export class Menstrual {
  @prop()
  status: MenstrualStatus = MenstrualStatus.disabled

  @prop()
  lastDate: Date = new Date()

  @prop()
  cycle: number = 28

  @prop()
  duration: number = 6
}
