import { prop } from 'typegoose'

export enum MenstrualStatus {
  disabled,
  enable,
}

export class Menstrual {
  @prop()
  public status: MenstrualStatus = MenstrualStatus.disabled

  @prop()
  public lastDate: Date = new Date()

  @prop()
  public cycle: number = 28

  @prop()
  public duration: number = 6
}
