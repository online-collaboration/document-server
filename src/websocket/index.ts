import { Server } from 'http'
import { Server as HTTPSServer } from 'https'
import * as WebSocket from 'ws'
import { setupWSConnection } from './utils'

export class WebSocketServer {
    wsServer
    constructor(server: Server | HTTPSServer) {
    this.wsServer = new WebSocket.Server({
        server,
        path: '/ws/doc',
    })
    this.wsServer.on('connection', setupWSConnection)
    }
}
