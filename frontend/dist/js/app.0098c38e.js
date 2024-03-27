(function(){"use strict";var t={7490:function(t,e,n){var r=n(6369),o=function(){var t=this,e=t._self._c;return e("div",[e("el-row",[e("div",{staticClass:"logo-div"},[e("span",{staticStyle:{"font-size":"30px","margin-left":"20px","font-weight":"bold"}},[t._v("穿云箭")])])]),e("el-row",[e("span",{staticStyle:{"margin-left":"20px"}},[t._v("拨开云雾见青天")])]),e("el-row",{staticStyle:{"margin-top":"100px"}},[e("el-link",{attrs:{underline:!1,type:"primary"}},[t._v("配置")])],1),e("el-row",[e("el-row",[e("div",{staticClass:"setting-div"},[e("el-input",{attrs:{"prefix-icon":"el-icon-connection",placeholder:"请输入需要穿透的本地服务地址,例如(127.0.0.1:8081)"},on:{input:t.handleInput},model:{value:t.localhost,callback:function(e){t.localhost=e},expression:"localhost"}})],1)]),""!=t.remoteAddr?e("el-row",{staticStyle:{"margin-top":"20px"}},[e("el-link",{attrs:{underline:!1,type:"primary"}},[t._v("外网地址")])],1):t._e(),""!=t.remoteAddr?e("el-row",{staticStyle:{"margin-top":"20px"}},[e("div",{directives:[{name:"clipboard",rawName:"v-clipboard:copy",value:t.remoteAddr,expression:"remoteAddr",arg:"copy"},{name:"clipboard",rawName:"v-clipboard:success",value:t.copySuccess,expression:"copySuccess",arg:"success"}],staticClass:"setting-div"},[e("el-link",{attrs:{underline:!1}},[t._v(t._s(t.remoteAddr)),e("i",{staticClass:"el-icon-copy-document"})])],1)]):t._e(),t.recording?e("div",{staticClass:"start-div"},[e("el-button",{staticStyle:{width:"100%","margin-top":"20px"},attrs:{type:"danger"},on:{click:function(e){return t.close()}}},[t._v("停止穿透")])],1):e("div",{staticClass:"start-div"},[e("el-button",{staticStyle:{width:"100%","margin-top":"20px"},attrs:{type:"primary"},on:{click:t.startRecording}},[t._v("开始穿透")])],1)],1),e("el-row",{staticStyle:{"margin-top":"20px"}},[e("div",{staticClass:"copyright-div"},[e("p",{staticStyle:{"font-size":"10px"}},[t._v("©穿云箭 2018-2024 使用协议")]),e("p",{staticStyle:{"font-size":"10px"}},[t._v("version 1.0.3")]),e("el-link",{attrs:{underline:!1,type:"primary"},on:{click:function(e){t.dialogVisible=!0}}},[t._v("使用指南")]),e("br"),e("el-link",{attrs:{underline:!1,target:"_blank",type:"primary"},on:{click:function(e){return t.gotoBrowser("https://qm.qq.com/q/lhm1yvQQN4")}}},[t._v("交流Q群")])],1)]),e("el-dialog",{attrs:{title:"常见问题指南",visible:t.dialogVisible,width:"90%"},on:{"update:visible":function(e){t.dialogVisible=e}}},[e("h5",[t._v("1.问:什么是穿云箭")]),e("span",[t._v("答:为爱发电，内网穿透工具，适用于本地服务的外网访问")]),e("h5",[t._v("2.问:可以翻墙么")]),e("span",[t._v("答:不可以，并非VPN，只是一种内网穿透工具")]),e("h5",[t._v("3.问:浏览器直接访问提示ERR_SSL_PROTOCOL_ERROR")]),e("span",[t._v("答:检查地址栏是不是https://开头（注意：部分浏览器会自动添加https），一定要使用http访问，因为大部分本地服务都是非https，此工具默认只代理http的服务")]),e("h5",[t._v("4.问:有使用案例么")]),e("span",[t._v("第一步 例如本地有个服务端口是12345 这样代理地址就填写127.0.0.1:12345")]),e("br"),e("span",[t._v("第二步 点击[开始穿透]")]),e("br"),e("span",[t._v("第三步 下方会出现类似于 c.0a0a.cn:11111 这样的地址，这是穿透后的外网地址，这个地址任何地点的有网络的设备、浏览器都可以直接进行访问")]),e("span",{staticClass:"dialog-footer",attrs:{slot:"footer"},slot:"footer"},[e("el-button",{attrs:{type:"primary"},on:{click:function(e){t.dialogVisible=!1}}},[t._v("原来是这样")])],1)])],1)},i=[];function s(t){return window["go"]["proxy"]["Options"]["Run"](t)}function a(){return window["go"]["proxy"]["Options"]["Stop"]()}function l(t){window.runtime.BrowserOpenURL(t)}var c={name:"App",data(){return{enableSpeaker:"0",localhost:"127.0.0.1:8080",remoteAddr:"",mode:1,recording:!1,dialogVisible:!0}},methods:{handleInput(t){const e=t.replace(/[^0-9:.]/g,"");this.localhost=e},async startRecording(){this.recording=!0;try{const t=await s(this.localhost);this.remoteAddr=t,this.$message.success("穿透成功")}catch(t){this.$message.error(t)}},close(){this.recording=!1,a(),this.remoteAddr="",this.$message.success("停止成功")},copySuccess(){this.$message.success("复制成功")},copyError(){this.$message.error("复制失败")},gotoBrowser(t){l(t)}}},u=c,p=n(1001),d=(0,p.Z)(u,o,i,!1,null,null,null),f=d.exports,v=n(8499),h=n.n(v),g=n(5428),m=n.n(g);r["default"].use(h()),r["default"].config.productionTip=!1,r["default"].use(m()),new r["default"]({render:t=>t(f)}).$mount("#app")}},e={};function n(r){var o=e[r];if(void 0!==o)return o.exports;var i=e[r]={id:r,loaded:!1,exports:{}};return t[r].call(i.exports,i,i.exports,n),i.loaded=!0,i.exports}n.m=t,function(){n.amdO={}}(),function(){var t=[];n.O=function(e,r,o,i){if(!r){var s=1/0;for(u=0;u<t.length;u++){r=t[u][0],o=t[u][1],i=t[u][2];for(var a=!0,l=0;l<r.length;l++)(!1&i||s>=i)&&Object.keys(n.O).every((function(t){return n.O[t](r[l])}))?r.splice(l--,1):(a=!1,i<s&&(s=i));if(a){t.splice(u--,1);var c=o();void 0!==c&&(e=c)}}return e}i=i||0;for(var u=t.length;u>0&&t[u-1][2]>i;u--)t[u]=t[u-1];t[u]=[r,o,i]}}(),function(){n.n=function(t){var e=t&&t.__esModule?function(){return t["default"]}:function(){return t};return n.d(e,{a:e}),e}}(),function(){n.d=function(t,e){for(var r in e)n.o(e,r)&&!n.o(t,r)&&Object.defineProperty(t,r,{enumerable:!0,get:e[r]})}}(),function(){n.g=function(){if("object"===typeof globalThis)return globalThis;try{return this||new Function("return this")()}catch(t){if("object"===typeof window)return window}}()}(),function(){n.o=function(t,e){return Object.prototype.hasOwnProperty.call(t,e)}}(),function(){n.r=function(t){"undefined"!==typeof Symbol&&Symbol.toStringTag&&Object.defineProperty(t,Symbol.toStringTag,{value:"Module"}),Object.defineProperty(t,"__esModule",{value:!0})}}(),function(){n.nmd=function(t){return t.paths=[],t.children||(t.children=[]),t}}(),function(){var t={143:0};n.O.j=function(e){return 0===t[e]};var e=function(e,r){var o,i,s=r[0],a=r[1],l=r[2],c=0;if(s.some((function(e){return 0!==t[e]}))){for(o in a)n.o(a,o)&&(n.m[o]=a[o]);if(l)var u=l(n)}for(e&&e(r);c<s.length;c++)i=s[c],n.o(t,i)&&t[i]&&t[i][0](),t[i]=0;return n.O(u)},r=self["webpackChunkqhtool_ui"]=self["webpackChunkqhtool_ui"]||[];r.forEach(e.bind(null,0)),r.push=e.bind(null,r.push.bind(r))}();var r=n.O(void 0,[998],(function(){return n(7490)}));r=n.O(r)})();
//# sourceMappingURL=app.0098c38e.js.map