"use strict";

import { app, protocol, BrowserWindow, ipcMain, Menu } from "electron";
import { createProtocol } from "vue-cli-plugin-electron-builder/lib";
import installExtension, { VUEJS3_DEVTOOLS } from "electron-devtools-installer";
const isDevelopment = process.env.NODE_ENV !== "production";

// Scheme must be registered before the app is ready
protocol.registerSchemesAsPrivileged([
  { scheme: "app", privileges: { secure: true, standard: true } },
]);

const grpc = require('@grpc/grpc-js')
const messages = require('./api/bj21/v1/bj21_pb')
const services = require('./api/bj21/v1/bj21_grpc_pb')
const { marshalGrpcBytes, bizKey } = require('./utils/igrpc')
const client = new services.BlackJackClient('0.0.0.0:9009', grpc.credentials.createInsecure())
const call = client.logicConn()
let win // main windows.
const userinfo = {name: "", token: ''} // 用户登录后的临时token，暂时每次都需要登录

async function createWindow() {
  // Create the browser window.
  win = new BrowserWindow({
    width: 1280,
    height: 720,
    autoHideMenuBar: true,
    frame: false,
    webPreferences: {
      // Use pluginOptions.nodeIntegration, leave this alone
      // See nklayman.github.io/vue-cli-plugin-electron-builder/guide/security.html#node-integration for more info
      nodeIntegration: process.env.ELECTRON_NODE_INTEGRATION,
      contextIsolation: !process.env.ELECTRON_NODE_INTEGRATION,
    },
  });
  // Menu.setApplicationMenu(null)

  if (process.env.WEBPACK_DEV_SERVER_URL) {
    // Load the url of the dev server if in development mode
    await win.loadURL(process.env.WEBPACK_DEV_SERVER_URL);
    // if (!process.env.IS_TEST) win.webContents.openDevTools();
  } else {
    createProtocol("app");
    // Load the index.html when not in development
    win.loadURL("app://./index.html");
  }
}

// Quit when all windows are closed.
app.on("window-all-closed", () => {
  console.log("quit soon.")
  // On macOS it is common for applications and their menu bar
  // to stay active until the user quits explicitly with Cmd + Q
  if (process.platform !== "darwin") {
    app.quit();
  }
});

app.on("activate", () => {
  // On macOS it's common to re-create a window in the app when the
  // dock icon is clicked and there are no other windows open.
  if (BrowserWindow.getAllWindows().length === 0) createWindow();
});

// This method will be called when Electron has finished
// initialization and is ready to create browser windows.
// Some APIs can only be used after this event occurs.
app.on("ready", async () => {
  if (isDevelopment && !process.env.IS_TEST) {
    // Install Vue Devtools
    try {
      await installExtension(VUEJS3_DEVTOOLS);
    } catch (e) {
      console.error("Vue Devtools failed to install:", e.toString());
    }
  }
  createWindow();
});

// 接收服务器返回的消息
call.on('data', (note) => {
  console.log(note.getCmd(), new Buffer.from(note.getText()).toString())

  let cmd = note.getCmd()
  let textdic = JSON.parse(new TextDecoder().decode(note.getText()))
  console.log(textdic)
  if (note.getCmd() == "login") {
    userinfo['name'] = textdic.name
    userinfo['token'] = textdic.token
  }
  const key = bizKey(note.getCmd())
  win.webContents.send(key, {
    cmd: cmd,
    text: textdic,
  })
})

// 请求服务器接口
ipcMain.on('logic-conn', (event, arg) => {
  const msg = new messages.Message()
  msg.setCmd(arg.cmd)
  if (arg.cmd !== 'login') {
    arg.text['token'] = userinfo['token']
  }
  msg.setText(marshalGrpcBytes(arg.text))
  call.write(msg)
})

ipcMain.on('backend-sys', (event, arg) => {
  switch (arg.cmd) {
    case "get-user-info":
      event.returnValue = userinfo
      break;
    case 'exit':
      win.close()
      break;
    default:
      break;
  }
})

// Exit cleanly on request from parent process in development mode.
if (isDevelopment) {
  if (process.platform === "win32") {
    process.on("message", (data) => {
      if (data === "graceful-exit") {
        app.quit();
      }
    });
  } else {
    process.on("SIGTERM", () => {
      app.quit();
    });
  }
}
