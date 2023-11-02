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
        <p style="font-size: 10px;">©穿云箭 2018-2023 使用协议</p>
        <p style="font-size: 10px;">version 1.0.1</p>
      </div>
    </el-row>
  </div>
</template>
<script>
import { Run, Stop } from '../wailsjs/go/proxy/Options.js'
export default {
  name: "App",
  data() {
    return {
      enableSpeaker: "0",
      localhost: "127.0.0.1:8080",
      remoteAddr: "",
      mode: 1,
      recording: false
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
  flex-direction:column;
}
</style>