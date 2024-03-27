<template>
  <div>
    <el-row>
      <div class="logo-div">
        <span style="font-size: 30px;margin-left: 20px;font-weight: bold;">穿云箭</span>
      </div>
    </el-row>
    <el-row>
      <span style="margin-left: 20px;">拨开云雾见青天</span>
    </el-row>
    <el-row style="margin-top: 100px;">
      <el-link :underline="false" type="primary">配置</el-link>
    </el-row>
    <el-row>
      <el-row>
        <div class="setting-div">
          <el-input prefix-icon="el-icon-connection" @input="handleInput" v-model="localhost"
            placeholder="请输入需要穿透的本地服务地址,例如(127.0.0.1:8081)"></el-input>
        </div>
      </el-row>
      <el-row v-if="remoteAddr != ''" style="margin-top: 20px;">
        <el-link :underline="false" type="primary">外网地址</el-link>
      </el-row>
      <el-row v-if="remoteAddr != ''" style="margin-top: 20px;">
        <div class="setting-div" v-clipboard:copy="remoteAddr" v-clipboard:success="copySuccess">
          <el-link :underline="false">{{ remoteAddr }}<i class="el-icon-copy-document"></i></el-link>
        </div>
      </el-row>
      <div class="start-div" v-if="recording">
        <el-button type="danger" style="width: 100%;margin-top: 20px;" @click="close()">停止穿透</el-button>
      </div>
      <div class="start-div" v-else>
        <el-button type="primary" style="width: 100%;margin-top: 20px;" @click="startRecording">开始穿透</el-button>
      </div>
    </el-row>
    <el-row style="margin-top: 20px;">
      <div class="copyright-div">
        <p style="font-size: 10px;">©穿云箭 2018-2024 使用协议</p>
        <p style="font-size: 10px;">version 1.0.3</p>
        <el-link :underline="false" @click="dialogVisible = true" type="primary">使用指南</el-link>
        <br>
        <el-link :underline="false" target="_blank" @click="gotoBrowser('https://qm.qq.com/q/lhm1yvQQN4')"
          type="primary">交流Q群</el-link>
      </div>
    </el-row>
    <el-dialog title="常见问题指南" :visible.sync="dialogVisible" width="90%">
      <h5>1.问:什么是穿云箭</h5>
      <span>答:为爱发电，内网穿透工具，适用于本地服务的外网访问</span>
      <h5>2.问:可以翻墙么</h5>
      <span>答:不可以，并非VPN，只是一种内网穿透工具</span>
      <h5>3.问:浏览器直接访问提示ERR_SSL_PROTOCOL_ERROR</h5>
      <span>答:检查地址栏是不是https://开头（注意：部分浏览器会自动添加https），一定要使用http访问，因为大部分本地服务都是非https，此工具默认只代理http的服务</span>
      <h5>4.问:有使用案例么</h5>
      <span>第一步 例如本地有个服务端口是12345 这样代理地址就填写127.0.0.1:12345</span>
      <br>
      <span>第二步 点击[开始穿透]</span>
      <br>
      <span>第三步 下方会出现类似于 c.0a0a.cn:11111 这样的地址，这是穿透后的外网地址，这个地址任何地点的有网络的设备、浏览器都可以直接进行访问</span>
      <span slot="footer" class="dialog-footer">
        <el-button type="primary" @click="dialogVisible = false">原来是这样</el-button>
      </span>
    </el-dialog>
  </div>
</template>   
<script>
import { Run, Stop } from '../wailsjs/go/proxy/Options.js'
import { BrowserOpenURL } from '../wailsjs/runtime/runtime.js'
export default {
  name: "App",
  data() {
    return {
      enableSpeaker: "0",
      localhost: "127.0.0.1:8080",
      remoteAddr: "",
      mode: 1,
      recording: false,
      dialogVisible: true
    }
  },
  methods: {
    handleInput(value) {
      // 限制只能输入数字和冒号  
      const validValue = value.replace(/[^0-9:.]/g, ''); // 移除所有非数字和非冒号字符  
      this.localhost = validValue; // 更新绑定的值  
    },
    async startRecording() {
      this.recording = true
      try {
        const response = await Run(this.localhost); // 调用异步函数并传递参数  
        this.remoteAddr = response; // 设置返回值 
        this.$message.success("穿透成功")
      } catch (error) {
        this.$message.error(error)
      }
    },
    close() {
      this.recording = false
      Stop()
      this.remoteAddr = ""
      this.$message.success("停止成功")
    },
    copySuccess() {
      this.$message.success("复制成功")
    },
    copyError() {
      this.$message.error("复制失败")
    },
    gotoBrowser(url) {
      BrowserOpenURL(url)
    }
  }
};
</script>
<style>
* {
  user-select: none;
}

.box-card {
  height: 180px;
  width: 200px;
}

.logo-div {
  display: flex;
  align-items: center;
}

.start-div {
  display: flex;
  align-items: center;
  justify-content: center;
}

.transcribe-div {
  display: flex;
  align-items: center;
  justify-content: center;
  height: 100px;
  border-radius: 10px;
  border: 2px solid rgb(207, 207, 207);
}

.transcribe-div:hover {
  display: flex;
  align-items: center;
  justify-content: center;
  height: 100px;
  border: 2px;
  border-radius: 10px;
  border: 2px solid rgb(75, 119, 238);
}


.transcribe-div-select {
  display: flex;
  align-items: center;
  justify-content: center;
  height: 100px;
  border: 2px;
  border-radius: 10px;
  border: 2px solid rgb(75, 119, 238);
}

.setting-div {
  display: flex;
  align-items: center;
  justify-content: center;
}

.copyright-div {
  display: flex;
  justify-content: center;
  text-align: center;
  flex-direction: column;
}
</style>