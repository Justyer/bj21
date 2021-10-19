const { ipcRenderer } = require('electron')

ipcRenderer.on('reply-login', function (event, arg) {
  // 切换场景：牌桌列表
  ipcRenderer.send('change_scene', {scene: 'tablelist'})
})

const loginBtn = document.getElementById('loginBtn')

loginBtn.addEventListener('click', () => {
  const loginInput = document.getElementById('loginInput')
  const reply = ipcRenderer.send('logic_conn', {
    cmd: 'login',
    text: {
      name: loginInput.value
    }
  })
  console.log(reply)
})