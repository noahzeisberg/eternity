const { app, BrowserWindow } = require('electron')

const createWindow = () => {
    const window = new BrowserWindow({
        width: 900,
        height: 650,
        title: "Eternity - Desktop Edition",
        titleBarStyle: 'hidden',
        titleBarOverlay: {
            color: '#000000',
            symbolColor: '#f5f3ff',
            height: 40
        },
        webPreferences: {
            webviewTag: true
        }
    })

    window.loadFile("index.html")
}

app.whenReady().then(() => {
    createWindow()
})

app.on('window-all-closed', () => {
    if (process.platform !== 'darwin') app.quit()
})