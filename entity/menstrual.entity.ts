import { Entity, Column, PrimaryColumn } from 'typeorm'

export enum MenstrualStatus {
  disabled,
  enable,
}

@Entity()
export class Menstrual {
  @PrimaryColumn()
  userId!: number

  @Column({ default: MenstrualStatus.disabled })
  status: MenstrualStatus = MenstrualStatus.disabled

  @Column({ default: () => 'CURRENT_TIMESTAMP' })
  lastDate: Date = new Date()

  @Column({ default: 28 })
  cycle: number = 28

  @Column({ default: 6 })
  duration: number = 6
}
