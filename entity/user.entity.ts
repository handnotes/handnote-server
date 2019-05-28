import { PrimaryGeneratedColumn, Column, Entity, CreateDateColumn, UpdateDateColumn } from 'typeorm'

const nullable = true
const unique = true

@Entity()
export class User {
  @PrimaryGeneratedColumn()
  id!: number

  @Column({ unique })
  openId!: string

  @Column({ nullable, default: '' })
  name?: string

  @Column({ nullable })
  avatar?: string

  @Column({ type: 'tinyint', nullable })
  gender?: number

  @Column({ nullable })
  address?: string

  @CreateDateColumn()
  createdAt!: Date

  @UpdateDateColumn()
  updatedAt!: Date
}
