// get brower

export function GetCurrentBrowser() {
  const ua = navigator.userAgent.toLocaleLowerCase()
  let browserType = ''
  if (ua.match(/msie/) != null || ua.match(/trident/) != null) {
    browserType = 'IE'
  }
  else if (ua.match(/firefox/) != null) {
    browserType = 'firefox'
  }
  else if (ua.match(/ucbrowser/) != null) {
    browserType = 'UC'
  }
  else if (ua.match(/opera/) != null || ua.match(/opr/) != null) {
    browserType = 'opera'
  }
  else if (ua.match(/bidubrowser/) != null) {
    browserType = 'baidu'
  }
  else if (ua.match(/metasr/) != null) {
    browserType = 'sougou'
  }
  else if (ua.match(/tencenttraveler/) != null || ua.match(/qqbrowse/) != null) {
    browserType = 'QQ'
  }
  else if (ua.match(/maxthon/) != null) {
    browserType = 'maxthon'
  }
  else if (ua.match(/chrome/) != null) {
    const is360 = _mime('type', 'application/vnd.chromium.remoting-viewer')
    if (is360)
      browserType = '360'
    else
      browserType = 'chrome'
  }
  else if (ua.match(/safari/) != null) {
    browserType = 'Safari'
  }
  else {
    browserType = 'others'
  }
  return browserType
}

function _mime(option, value) {
  const mimeTypes = navigator.mimeTypes
  for (const mt in mimeTypes) {
    if (mimeTypes[mt][option] === value)
      return true
  }
  return false
}

// get os
export function GetOs() {
  console.log(navigator.platform, 'platform')
  const sUserAgent = navigator.userAgent.toLocaleLowerCase()
  const isWin = (navigator.platform === 'Win32') || (navigator.platform === 'Windows')
  const isMac = (navigator.platform === 'Mac68K') || (navigator.platform === 'MacPPC') || (navigator.platform === 'Macintosh') || (navigator.platform === 'MacIntel')
  if (isMac)
    return 'Mac'
  const isUnix = (navigator.platform === 'x11') && !isWin && !isMac
  if (isUnix)
    return 'Unix'
  const isLinux = (String(navigator.platform).includes('linux'))
  if (isLinux)
    return 'Linux'
  if (isWin) {
    const isWin2K = sUserAgent.includes('windows nt 5.0') || sUserAgent.includes('windows 2000')
    if (isWin2K)
      return 'Win2000'
    const isWinXP = sUserAgent.includes('windows nt 5.1') || sUserAgent.includes('windows xp')
    if (isWinXP)
      return 'WinXP'
    const isWin2003 = sUserAgent.includes('windows nt 5.2') || sUserAgent.includes('windows 2003')
    if (isWin2003)
      return 'Win2003'
    const isWinVista = sUserAgent.includes('windows nt 6.0') || sUserAgent.includes('windows vista')
    if (isWinVista)
      return 'WinVista'
    const isWin7 = sUserAgent.includes('windows nt 6.1') || sUserAgent.includes('windows 7')
    if (isWin7)
      return 'Win7'
    return 'Windows'
  }
  if (sUserAgent.includes('android'))
    return 'Android'
  if (sUserAgent.includes('iphone'))
    return 'iPhone'
  if (sUserAgent.includes('symbianos'))
    return 'SymbianOS'
  if (sUserAgent.includes('windows phone'))
    return 'Windows Phone'
  if (sUserAgent.includes('ipad'))
    return 'iPad'
  if (sUserAgent.includes('ipod'))
    return 'iPod'
  return 'others'
}

// getAddress
// {/*<script src="http://pv.sohu.com/cityjson?ie=utf-8"></script>*/}
// {/*export function GetAddress () {*/}
//   {/*return returnCitySN*/}
// {/*}*/}
