import { useExpressServer } from 'routing-controllers'
import 'reflect-metadata'
import express from 'express'
import { DocController } from './controller/doc/DocController'
import { WebSocketServer } from './websocket'

let app = express()

useExpressServer(app, {
    routePrefix: '/v1',
    controllers:[DocController]
})

const PORT = process.env.PORT || 30009
const server = app.listen(PORT, () => {
    console.log(`Start server on ${PORT}`)
})
new WebSocketServer(server)
