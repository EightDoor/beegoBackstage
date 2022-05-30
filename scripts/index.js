const inquirer = require("inquirer")
const shelljs = require("shelljs")
const setting = require("./setting.json")
const { LogUtil } = require("zhoukai_utils")
const path = require("path");
const zipper = require("zip-local")
const NodeSSH = require("node-ssh").NodeSSH

const SSH = new NodeSSH()
const log = new LogUtil()
const distDir = path.join(__dirname, "../front", setting.front.distDir)

// 选择项目
inquirer.prompt([
  {
    type: 'list',
    message: '请选择部署项目',
    name: 'value',
    choices: [
      {
        name: '前端-front',
        value: 'front',
      },
      {
        name: '后台-back',
        value: 'back',
      },
    ],
  },
]).then((answers) => {
  const value = answers.value
  log.i(value, '选择项目')
  if (value === 'front') {
    frontFun(value)
  }
  else if (value === 'back') {
    backFun(value)
  }
}).catch((error) => {
  log.e(error, 'error')
})

async function frontFun(status) {
  log.d("前端打包执行中")
  // 打包
  const result = shelljs.exec("cd ../front && npm run build")
  if(result.code === 0) {
    try {
      log.d("开始压缩文件")
      await zipper.sync.zip(distDir).compress().save(setting[status].remoteFile)
      log.s("文件压缩成功")
      await connectSSh(status)
    }catch (e) {
      throw new Error("文件压缩失败")
    }
  }else {
    throw new Error("前端打包失败")
  }

}

function backFun(status) {
  // 先执行打包命令
  log.i("开始打包")
  const result = shelljs.exec("cd ../back && bee pack -be GOOS=linux -be GOARCH=amd64")
  if (result.code === 0) {
    const mv = shelljs.mv("../back/back.tar.gz", ".")
    if (mv.code === 0) {
      connectSSh(status)
    }else {
      log.e(status + "移动文件失败")
    }

  }else {
    log.e(status + "打包失败")
  }
}


async function uploadFile(status) {
  try {
    await SSH.putFile(setting[status].remoteFile, setting[status].remoteCwd + setting[status].remoteFile)
    log.s("完成上传")
    log.i("完成上传，开始解压缩")

    if(status === 'back') {
      await runCommand(`tar -zxvf ${setting[status].remoteFile}`, status)
    }else {
      await runCommand(`unzip -o ${setting[status].remoteFile}`, status)
    }
    log.s("部署完成")
    process.exit(0)
  }catch (e) {
    throw new Error("ssh 上传失败")
  }
}

async function runCommand(command, status) {
  log.d(setting[status].remoteCwd, "remoteCwd")
  const cwd = setting[status].remoteCwd
  const result = await SSH.execCommand(command,  {cwd})
  log.d(result, "result")
  // back运行
  if(status === 'back') {
    try {
      await SSH.execCommand("chmod -R 755 ./ssh", {cwd})
      await SSH.execCommand("./ssh/stop.sh", {cwd})
      await SSH.execCommand("./ssh/start.sh",{cwd})
      log.s("back 启动成功")
    }catch (e) {
      throw new Error( e + "运行back失败")
    }
  }
}

async function connectSSh(status) {
  log.i("正在连接ssh")
  log.i(`主机: ${setting[status].host},用户名: ${setting[status].user},密码: ${setting[status].passwd}`)
  SSH.connect({
    host: setting[status].host,
    username: setting[status].user,
    password: setting[status].passwd
  }).then(async ()=>{
    log.s("ssh 连接成功")
    await uploadFile(status)
  }).catch(err=>{
    log.e(err, "ssh 连接失败")
    process.exit(0)
  })
}
