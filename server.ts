import app from './app'
import createConnection from './database.connection'
;(async () => {
  // Create Connection to Database
  try {
    await createConnection()
  } catch (error) {
    console.error(error)
  }

  /** Application working port */
  const PORT: number = Number(process.env.APP_PORT) || 3050

  /** Application listener host (for local area network) */
  const HOSTNAME = process.env.NODE_ENV === 'development' ? '0.0.0.0' : '127.0.0.1'

  // launch App
  app.listen(PORT, HOSTNAME, () => {
    console.log(`koa is started on ${PORT}`)
  })
})()
