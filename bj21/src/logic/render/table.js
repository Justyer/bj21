import { ipcRenderer } from "electron";
import { isUndefined } from "../../utils/judge";
import { bizKey } from "../server/logic_conn";

// 主要返回值处理成业务需要的结构
const onTableInfo = (callback) => {
    ipcRenderer.on(bizKey("tableinfo"), (event, arg) => {
        let pyou = { name: "<nil>", hands: [{ num: 0, id: "youid" }], sum_text: '' },
            pme = { name: "<nil>", hands: [{ num: 0, id: "meid" }], sum_text: '' };

        // 确定自己是哪个player
        let replypyou = {}, replypme = {};
        if (arg.text.table.me == 1) {
            replypyou = arg.text.table.p2
            replypme = arg.text.table.p1
        } else {
            replypyou = arg.text.table.p1
            replypme = arg.text.table.p2
        }

        if (!isUndefined(replypyou)) {
            pyou.name = replypyou.name;
            if (!isUndefined(replypyou.cards)) {
                pyou.hands = replypyou.cards
                let sum = 0;
                for (let i = 0; i < replypyou.cards.length; i++) {
                    sum = sum + replypyou.cards[i].num
                }
                pyou.sum_text = sum + "/21";
            }
        }
        if (!isUndefined(replypme)) {
            pme.name = replypme.name;
            if (!isUndefined(replypme.cards)) {
                pme.hands = replypme.cards
                let sum = 0;
                for (let i = 0; i < replypme.cards.length; i++) {
                    sum = sum + replypme.cards[i].num
                }
                pyou.sum_text = sum + "/21";
            }
        }

        callback({
            name: arg.text.table.name,
            seq: arg.text.table.seq,
            player_you: pyou,
            player_me: pme,
        });
    });
}

export { onTableInfo };