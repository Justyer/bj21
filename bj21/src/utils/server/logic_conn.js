const grpc = require('@grpc/grpc-js')

const { marshalGrpcBytes, bizKey } = require('./utils/igrpc')
const addr = '0.0.0.0:9009'


class LogicConn {
    constructor(addr) {
        this.messages = require('./api/bj21/v1/bj21_pb')
        this.services = require('./api/bj21/v1/bj21_grpc_pb')
        this.client = new services.BlackJackClient(addr, grpc.credentials.createInsecure())
        this.call = client.logicConn()
        this.userinfo = {name: "", token: ''}
    }
    listenReply(callback) {
        this.call.on('data', (note) => {
            console.log(note.getCmd(), new Buffer.from(note.getText()).toString())

            let cmd = note.getCmd()
            let textdic = JSON.parse(new TextDecoder().decode(note.getText()))
            console.log(textdic)
            if (note.getCmd() == "login") {
                this.userinfo['name'] = textdic.name
                this.userinfo['token'] = textdic.token
            }
            // const key = bizKey(note.getCmd())
            callback({
                cmd: cmd,
                text: textdic,
            })
            // win.webContents.send(key, {
            //     cmd: cmd,
            //     text: textdic,
            // })
        })
    }
}

module.exports.LogicConn = LogicConn