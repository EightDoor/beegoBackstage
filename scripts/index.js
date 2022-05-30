const inquirer = require("inquirer")
const shelljs = require("shelljs")
const ssh2 = require("ssh2")
const SftpClient = require("ssh2-sftp-client")
const setting = require("./setting.json")
const { LogUtil } = require("zhoukai_utils")
const path = require("path");

const log = new LogUtil()

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
    frontFun()
  }
  else if (value === 'back') {
    backFun()
  }
}).catch((error) => {
  log.e(error, 'error')
})

function frontFun() {
  log.d("前端打包执行中")
  // 打包
  // const status = shelljs.exec("cd ../front && npm run build")
  // if(status !== 0) {
  //   log.i(status, "status")
  //   log.s("打包完成")
  //   log.d("连接ssh，上传文件")
  //
  // }else {
  //   log.e(status, "打包执行失败")
  // }
  uploadFile('front').then(res=>{
    log.s(res, "上传文件成功")
  })
}

function backFun() {

}


async function uploadFile(status) {
  const sftp = new SftpClient("upload-test")
  return new Promise((resolve, reject)=>{
    // 先打包dist
    shelljs.exec("tar -zcvf dist.tar.gz ../front/dist", async function (code, stdout, stderr) {
      console.log('Exit code:', code);
      console.log('Program output:', stdout);
      console.log('Program stderr:', stderr);
      console.log(setting[status].host)
      console.log(setting[status].user)
      console.log(setting[status].passwd)
      sftp.connect({
        host: setting[status].host,
        username: setting[status].user,
        password: setting[status].passwd
      }).then(async (res)=>{
        const localPath = path.join(__dirname, "dist.tar.gz")
         sftp.fastPut(localPath, setting[status].remotePath).then(res=>{
          log.d(res, 'res')
           resolve(res);
         }).catch(err=>{
           reject(err.message)
         })
      }).then(p => {
        console.log(`Remote working directory is ${p}`);
        return sftp.end();
      }).catch(err => {
        reject(err.message)
      });
    })
  })
}

/**
 * 连接远程电脑
 * @param status back | front
 * @constructor
 */
function Connect(status) {
  log.i(`主机: ${setting[status].host},账号:${setting[status].user},密码:${setting[status].passwd}`)
  return new Promise((resolve, reject)=>{
    const conn = new ssh2.Client()
    conn.on("ready", function () {
      resolve(conn)
    }).on("error", function (err) {
      log.e(err, "conn 连接失败")
    }).on("end", function () {
      log.d("conn 连接end")
    }).on("close", function () {
      log.d("conn 连接关闭")
    }).connect({
      ...genConfig(status),
      readyTimeout: 60000,
      debug: console.log,
      tryKeyboard: true,
    })
  })
}
