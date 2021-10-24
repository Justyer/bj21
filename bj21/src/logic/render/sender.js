import { ipcRenderer } from "electron";
import { bizKey } from "@/logic/server/logic_conn";

const logicConnSender = (msg) => {
    ipcRenderer.send("logicconn", msg);
}

const logicConnReply = (biz, callback) => {
    ipcRenderer.on(bizKey(biz), callback);
}

export { logicConnSender, logicConnReply };