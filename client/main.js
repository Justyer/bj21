const { app, BrowserWindow } = require('electron')
const path = require('path')

const grpc = require('@grpc/grpc-js')
const bjproto = require('./proto')



function createWindow() {
    const win = new BrowserWindow({
        width: 1280,
        height: 720,
        webPreferences: {
            preload: path.join(__dirname, 'preload.js'),
            nodeIntegration: true
        }
    })

    win.loadFile('index.html')
}

app.whenReady().then(() => {
    createWindow()
    grpc_conn()
})

app.on('window-all-closed', function () {
    if (process.platform !== 'darwin') app.quit()
})

function grpc_conn() {
    var client = new bjproto.BlackJack('0.0.0.0:9009', grpc.credentials.createInsecure())
    client.enterRoom({ room_id: '123' }, function (err, response) {
        if (err) {
            console.error('Error: ', err)
        } else {
            console.log('Message: ', response.room_token)
        }
    })
}