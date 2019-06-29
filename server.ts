import app from './app'

/** Application working port */
const PORT: number = Number(process.env.APP_PORT) || 3050

/** Application listener host (for local area network) */
const HOSTNAME = process.env.NODE_ENV === 'development' ? '0.0.0.0' : '127.0.0.1'

// launch App
app.listen(PORT, HOSTNAME, () => {
  console.log(`koa is started on ${HOSTNAME}:${PORT}`)
})
