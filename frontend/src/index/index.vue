<template>
    <div class="item">
        <t-row>
            <div class="logo-div">
                <span style="font-size: 30px;font-weight: bold;">穿云箭</span>
            </div>
        </t-row>
        <t-row style="margin-top: 10px;">
            <span>拨开云雾见青天</span>
        </t-row>
        <t-row style="margin-top: 100px;">
            <span>内网地址</span>
        </t-row>
        <t-row>
            <t-col :span="24">
                <div>
                    <t-input size="large" @input="handleInput" v-model="localhost"
                        placeholder="请输入需要穿透的本地服务地址,例如(127.0.0.1:8081)"></t-input>
                </div>
            </t-col>
        </t-row>
        <t-row style="margin-top: 10px;">
            <span>映射协议</span>
        </t-row>
        <t-row>
            <t-col :span="24">
                <div>
                    <t-radio-group variant="default-filled" default-value="4" v-model="proxyType">
                        <t-radio-button v-for="item in proxyTypeList" :value="item">{{ item }}</t-radio-button>
                    </t-radio-group>
                </div>
            </t-col>
        </t-row>
        <t-row v-if="remoteAddr != ''" style="margin-top: 20px;">
            <span>外网地址</span>
        </t-row>
        <t-row v-if="remoteAddr != ''">
            <t-input size="large" status="success" v-model="remoteAddr">
                <template #suffixIcon>
                    <copy-icon @click="copy(remoteAddr)" />
                </template>
            </t-input>
        </t-row>
        <t-row>
            <div class="start-div" v-if="recording">
                <t-button size="large" theme="danger" style="width: 100%;margin-top: 20px;"
                    @click="close()">停止穿透</t-button>
            </div>
            <div class="start-div" v-else>
                <t-button size="large" theme="primary" style="width: 100%;margin-top: 20px;"
                    @click="startRecording">开始穿透</t-button>
            </div>
        </t-row>
        <t-row style=" position: absolute;bottom: 2%; left: 0;right: 0;">
            <div class="copyright-div">
                <div>
                    <t-link style="margin: 5px;" theme="primary" @click="dialogVisible = true"
                        hover="color">使用指南</t-link>
                    <t-link style="margin: 5px;" theme="primary" hover="color"
                        @click="gotoBrowser('http://c.0a0a.cn')">软件官网</t-link>
                    <t-link style="margin: 5px;" theme="primary" target="_blank"
                        @click="gotoBrowser('https://qm.qq.com/q/lhm1yvQQN4')" hover="color">交流Q群</t-link>
                </div>
                <p style="font-size: 10px;">©穿云箭 2018-2024 使用协议</p>
                <p style="font-size: 10px;">version 1.0.5</p>
            </div>
        </t-row>
        <t-dialog title="常见问题指南" v-model:visible="dialogVisible" width="90%"
            :on-confirm="() => (dialogVisible = false)">
            <h5>1.问:什么是穿云箭</h5>
            <span>答:为爱发电，内网穿透工具，适用于本地服务的外网访问</span>
            <h5>2.问:可以翻墙么</h5>
            <span>答:不可以，并非VPN，只是一种内网穿透工具</span>
            <h5>3.问:浏览器直接访问提示net::ERR_CERT_AUTHORITY_INVALID</h5>
            <span>答:因为证书问题导致，不影响使用，点击【高级】-【继续前往】</span>
            <h5>4.问:有使用案例么</h5>
            <span>第一步 例如本地有个服务端口是12345 这样代理地址就填写127.0.0.1:12345</span>
            <br>
            <span>第二步 点击[开始穿透]</span>
            <br>
            <span>第三步 下方会出现类似于 c.0a0a.cn:11111 这样的地址，这是穿透后的外网地址，这任何地点的有网络的设备、浏览器都可以直接进行访问这个地址</span>
        </t-dialog>
    </div>
</template>
<script setup>
import { ref, reactive } from 'vue'
import { Run, Stop } from '../../wailsjs/go/proxy/Options.js'
import { BrowserOpenURL } from '../../wailsjs/runtime/runtime.js'
import { MessagePlugin } from 'tdesign-vue-next'
import useClipboard from 'vue-clipboard3'
import { setBodyBg } from '@/utils/utils';
const localhost = ref("127.0.0.1:8080")
const remoteAddr = ref("")
const proxyType = ref("http")
const proxyTypeList = ref(["http", "https"])
const recording = ref(false)
const dialogVisible = ref(false)
import {
    CopyIcon,
} from 'tdesign-icons-vue-next';

const handleInput = (value) => {
    // 限制只能输入数字和冒号  
    localhost.value = value.replace(/[^0-9:.]/g, '') // 更新绑定的值  
}
const startRecording = () => {
    recording.value = true
    const response = Run(localhost.value, proxyType.value);
    response.then((res) => {
        console.log(res)
        remoteAddr.value = res; // 设置返回值 
        MessagePlugin.success("穿透成功")
        //todo 更改为流动特效
        // setBodyBg("linear-gradient(-225deg,#FFFEFF ,#c9ffbf 100%)")
    }).catch((error) => {
        MessagePlugin.error(error)
    })
}
const close = () => {
    recording.value = false
    Stop()
    remoteAddr.value = ""
    MessagePlugin.success("停止成功")
    //todo 更改为流动特效
    // setBodyBg("linear-gradient(-225deg, #E3FDF5, #FFE6FA 100%)")

}

const { toClipboard } = useClipboard()
const copy = async (text) => {
    try {
        await toClipboard(text) // 实现复制
        MessagePlugin.success('复制成功')
    } catch (e) {
        console.error(e)
    }
}

function gotoBrowser(url) {
    BrowserOpenURL(url)
}

</script>
<style scoped>
* {
    user-select: none;
}

.t-row {
    display: block;
}

.t-link {
    display: inline;
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
    width: 100%;
    display: flex;
    align-items: center;
    justify-content: center;
}

.copyright-div {
    position: relative;
    display: flex;
    justify-content: center;
    text-align: center;
    flex-direction: column;
}
</style>
<style>
body {
    background-image: linear-gradient(-225deg, #E3FDF5, #FFE6FA 100%)
}
</style>