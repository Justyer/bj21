import { ipcMain } from "electron";
import { LogicConn, bizKey } from "@/logic/server/logic_conn";

class Action {
    constructor(win) {
        this.userinfo = { name: "", token: "" }
        this.win = win
        this.logicConn
    }
    dialLogicConn(addr) {
        try {
            this.logicConn = new LogicConn(addr);
            return true
        } catch (e) {
            return false
        }
    }
    onLogicConnRequest() {
        ipcMain.on('logicconn', (event, arg) => {
            if (arg.cmd !== 'login') {
                arg.text['token'] = this.userinfo.token
            }
            this.logicConn.requestFromRenderer(arg)
        })
    }
    onLogicConnReply() {
        this.logicConn.listenReply((arg) => {
            if (arg.cmd === "login") {
                this.userinfo.name = arg.text.name
                this.userinfo.token = arg.text.token
            }
            const key = bizKey(arg.cmd)
            this.win.webContents.send(key, arg)
        });
    }
    onBackend() {
        ipcMain.on('mainlogic', (event, arg) => {
            switch (arg.cmd) {
                case "login":
                    const isSucc = this.dialLogicConn(arg.text.addr)
                    if (isSucc) {
                        this.onLogicConnReply()
                        this.onLogicConnRequest()
                        this.logicConn.requestFromRenderer({
                            cmd: arg.cmd,
                            text: {
                                name: arg.text.name
                            }
                        })
                    }
                    break;
                case "get-user-info":
                    event.returnValue = this.userinfo
                    break;
                case 'exit':
                    win.close()
                    break;
                default:
                    break;
            }
        })
    }
}

export { Action };