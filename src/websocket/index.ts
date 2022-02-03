import { Server } from 'http'
import { Server as HTTPSServer } from 'https'
import * as WebSocket from 'ws'
export class WebSocketServer{
    wsServer
    constructor(server: Server| HTTPSServer){
        this.wsServer = new WebSocket.Server({
            server,
            path:'/ws/doc'
        })
    }
}