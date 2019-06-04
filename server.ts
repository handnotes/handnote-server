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

  // launch App
  app.listen(PORT, () => {
    console.log(`koa is started on ${PORT}`)
  })
})()
