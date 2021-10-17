const { app, BrowserWindow, ipcMain } = require('electron')
const path = require('path')
const grpc = require('@grpc/grpc-js')
const messages = require('./api/bj21/v1/bj21_pb')
const services = require('./api/bj21/v1/bj21_grpc_pb')

// *** 变量定义区 ***

const client = new services.BlackJackClient('0.0.0.0:9009', grpc.credentials.createInsecure())
const call = client.logicConn()
var win // main windows.
const temporaryEventPool = {} // 临时事件池，处理icpMain接收到消息后往服务器发送异步请求，在grpc的call.on中使用
var token = '' // 用户登录后的临时token，暂时每次都需要登录


// *** 常规函数区 ***

createWindow = () => {
    win = new BrowserWindow({
        width: 1280,
        height: 720,
        webPreferences: {
            preload: path.join(__dirname, 'preload.js'),
            nodeIntegration: true,
            enableRemoteModule: true,
            contextIsolation: false,
        },
    })
    win.loadFile('page/login.html')
    win.webContents.openDevTools()
}

marshalGrpcBytes = (dic) => {
    return new Buffer.from(JSON.stringify(dic), 'utf8').toString("base64")
}

bizKey = (biz) => {
    return 'biz-' + biz
}

// *** 监听区 ***

app.whenReady().then(() => {
    createWindow()
})

app.on('window-all-closed', () => {
    call.end()
    client.close()
    if (process.platform !== 'darwin') app.quit()
})

// 接收服务器返回的消息
call.on('data', (note) => {
    console.log(note.getCmd(), new Buffer.from(note.getText()).toString())

    let cmd = note.getCmd()
    let textdic = JSON.parse(new TextDecoder().decode(note.getText()))
    console.log(textdic)
    if (note.getCmd() == "login") { 
        token = textdic.token
    }
    const key = bizKey(note.getCmd())
    temporaryEventPool[key].sender.send(key, {
        cmd: cmd,
        text: textdic,
    })
    delete temporaryEventPool[key]
})

// 请求服务器接口
ipcMain.on('logic_conn', (event, arg) => {
    const msg = new messages.Message()
    msg.setCmd(arg.cmd)
    if (arg.cmd !== 'login') {
        arg.text['token'] = token
    }
    msg.setText(marshalGrpcBytes(arg.text))
    call.write(msg)
    temporaryEventPool[bizKey(arg.cmd)] = event
})

ipcMain.on('change_scene', (event, arg)=> {
    switch (arg.scene) {
        case "tablelist":
            win.loadFile('page/tablelist.html')
            break;
    
        default:
            break;
    }
})