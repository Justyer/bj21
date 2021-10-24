const grpc = require('@grpc/grpc-js')

class LogicConn {
    constructor(addr) {
        this.messages = require('./api/bj21/v1/bj21_pb')
        this.services = require('./api/bj21/v1/bj21_grpc_pb')
        this.client = new this.services.BlackJackClient(addr, grpc.credentials.createInsecure())
        this.call = this.client.logicConn()
    }
    listenReply(callback) {
        this.call.on('data', (note) => {
            console.log(note.getCmd(), new Buffer.from(note.getText()).toString())
            let cmd = note.getCmd()
            let textdic = JSON.parse(new TextDecoder().decode(note.getText()))
            console.log(textdic)
            callback({
                cmd: cmd,
                text: textdic,
            })
        })
    }
    requestFromRenderer(arg) {
        const msg = new this.messages.Message()
        msg.setCmd(arg.cmd)
        msg.setText(this.marshalGrpcBytes(arg.text))
        this.call.write(msg)
    }
    marshalGrpcBytes(dic) {
        return new Buffer.from(JSON.stringify(dic), 'utf8').toString("base64")
    }

}

const bizKey = (biz) => {
    return 'reply-' + biz
}

export { LogicConn, bizKey };