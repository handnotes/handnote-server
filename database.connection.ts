import { createConnection, getConnectionOptions, ConnectionOptions } from 'typeorm'

const entitiesPath = 'entity/*.entity.ts'

export default async () => {
  // Get connection options from '.env.local' or '.env'
  const connectOpts = await getConnectionOptions()

  // Overriding environment options
  Object.assign(connectOpts, {
    logging: true,
    synchronize: true,
    entities: [entitiesPath],
  } as ConnectionOptions)

  // Create Connection
  return createConnection(connectOpts)
}
