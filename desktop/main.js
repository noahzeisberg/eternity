const { app, BrowserWindow } = require('electron')

const createWindow = () => {
    const window = new BrowserWindow({
        width: 1200,
        height: 750,
        title: "Eternity - Desktop Edition",
        titleBarStyle: 'hidden',
        titleBarOverlay: {
            color: '#000000',
            symbolColor: '#f5f3ff',
            height: 45
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