const { ipcRenderer } = require('electron')

const tablelistDiv = document.getElementById('tablelistDiv')

ipcRenderer.on('reply-tablelist', function (event, arg) {
  for (let i = 0; i < arg.text.tables.length; i++) {
    console.log(arg.text.tables[i].name, arg.text.tables[i].seq)
    let table = document.createElement('div')
    table.className = 'table'
    table.style.width = '200px'
    table.style.height = '100px'
    table.style.boxShadow = '0 0 5px grey'
    table.style.borderRadius = '5px'
    table.style.margin = '10px'
    let tableName = document.createElement('div')
    tableName.innerText = arg.text.tables[i].name
    tableName.className = 'tablename'
    table.appendChild(tableName)
    tablelistDiv.appendChild(table)
  }
})

ipcRenderer.send('logic_conn', {
  cmd: 'tablelist',
  text: {}
})

