import { MigrationInterface, QueryRunner, Table } from 'typeorm'

export class InitUser1559106682613 implements MigrationInterface {
  public async up(queryRunner: QueryRunner): Promise<any> {
    await queryRunner.createTable(
      new Table({
        name: 'user',
        columns: [
          { name: 'id', type: 'int', isPrimary: true },
          { name: 'openId', type: 'varchar', isUnique: true },
          { name: 'email', type: 'varchar', isUnique: true },
          { name: 'name', type: 'varchar', isNullable: true },
          { name: 'avatar', type: 'varchar', isNullable: true },
          { name: 'gender', type: 'tinyint', isNullable: true },
          { name: 'location', type: 'varchar', isNullable: true },
          { name: 'createdAt', type: 'datetime' },
          { name: 'updatedAt', type: 'datetime' },
        ],
      }),
      true,
    )
  }

  public async down(queryRunner: QueryRunner): Promise<any> {
    await queryRunner.dropTable('user')
  }
}
