(window["webpackJsonp"]=window["webpackJsonp"]||[]).push([["chunk-0053cef9"],{"0a9f":function(t,e,s){t.exports=s.p+"static/img/blueBg.8f4dd5e7.png"},"1a16":function(t,e,s){},2893:function(t,e,s){"use strict";var n=function(){var t=this,e=t.$createElement,s=t._self._c||e;return s("div",{staticClass:"perWord"},[s("van-tabs",{attrs:{animated:"",sticky:""},on:{click:t.changeBrow},model:{value:t.active,callback:function(e){t.active=e},expression:"active"}},[s("van-tab",{attrs:{title:"作品"}},[s("Browse",{directives:[{name:"show",rawName:"v-show",value:t.allMsgInfo,expression:"allMsgInfo"}],attrs:{info:t.allMsgInfo}}),s("van-empty",{directives:[{name:"show",rawName:"v-show",value:!t.allMsgInfo,expression:"!allMsgInfo"}],attrs:{description:"您还没有作品哦,快来发布！"}})],1),s("van-tab",{attrs:{title:"收藏"}},[s("van-empty",{attrs:{description:"您还没有收藏的哦"}})],1),s("van-tab",{attrs:{title:"点赞"}},[s("Browse",{directives:[{name:"show",rawName:"v-show",value:t.likeMsgInfo,expression:"likeMsgInfo"}],attrs:{info:t.likeMsgInfo}}),s("van-empty",{directives:[{name:"show",rawName:"v-show",value:!t.likeMsgInfo,expression:"!likeMsgInfo"}],attrs:{description:"您还没有喜欢的哦"}})],1)],1)],1)},i=[],o=s("1da1"),a=(s("4de4"),s("d3b7"),s("96cf"),s("a83c")),r=s("fa7d"),c={data:function(){return{active:0,msg:null,likeMsg:null,allMsgInfo:null,likeMsgInfo:null}},components:{Browse:a["a"]},props:{id:{required:!1}},mounted:function(){this.getMsgInfo()},methods:{changeBrow:function(t){switch(t){case 0:this.getMsgInfo();break;case 1:break;case 2:this.getGiveLike();break}},getMsgInfo:function(){var t=this;return Object(o["a"])(regeneratorRuntime.mark((function e(){var s,n,i;return regeneratorRuntime.wrap((function(e){while(1)switch(e.prev=e.next){case 0:return s=window.sessionStorage.getItem("userid"),e.next=3,t.$http.get("/article/getAllArticle",{params:{userid:t.id?t.id:s}});case 3:n=e.sent,i=n.data,"success"===i.status&&(t.msg=i.articleInfoList?Object(r["a"])(i.articleInfoList):i.articleInfoList),t.allMsgInfo=t.msg;case 7:case"end":return e.stop()}}),e)})))()},getGiveLike:function(){var t=this;return Object(o["a"])(regeneratorRuntime.mark((function e(){var s,n,i;return regeneratorRuntime.wrap((function(e){while(1)switch(e.prev=e.next){case 0:return s=window.sessionStorage.getItem("userid"),e.next=3,t.$http.get("/article/getGiveLikeArticle",{params:{userid:t.id?t.id:s}});case 3:n=e.sent,i=n.data,"success"===i.status&&(t.likeMsg=i.giveLikeArticleList?Object(r["a"])(i.giveLikeArticleList):i.giveLikeArticleList,t.likeMsg=t.likeMsg.filter((function(t){return null!==t})),0!==t.likeMsg.length&&null!==t.likeMsg||(t.likeMsg=null),t.likeMsgInfo=t.likeMsg);case 6:case"end":return e.stop()}}),e)})))()}}},l=c,u=s("2877"),f=Object(u["a"])(l,n,i,!1,null,null,null);e["a"]=f.exports},"31c4":function(t,e,s){},"4de4":function(t,e,s){"use strict";var n=s("23e7"),i=s("b727").filter,o=s("1dde"),a=o("filter");n({target:"Array",proto:!0,forced:!a},{filter:function(t){return i(this,t,arguments.length>1?arguments[1]:void 0)}})},5909:function(t,e,s){"use strict";s.r(e);var n=function(){var t=this,e=t.$createElement,s=t._self._c||e;return s("div",[s("van-nav-bar",{attrs:{"left-text":"返回","left-arrow":""},on:{"click-left":t.onClickLeft}}),s("div",[s("div",{staticClass:"showpart"},[s("div",{staticClass:"userInfo"},[s("div",{staticClass:"infoView"},[s("InfoView",{attrs:{info:t.userInfo}})],1),s("div",{staticClass:"attentionInfo"},[s("div",[s("p",[t._v(t._s(t.follow_count))]),s("p",[t._v("关注")])]),s("div",[s("p",[t._v(t._s(t.fans_count))]),s("p",[t._v("粉丝")])]),t._m(0),s("div",{staticClass:"focuspart"},[s("div",{directives:[{name:"show",rawName:"v-show",value:t.isNotSelf,expression:"isNotSelf"}]},[s("van-button",{directives:[{name:"show",rawName:"v-show",value:!t.ifFocus,expression:"!ifFocus"}],staticClass:"focus",attrs:{round:"",hairline:""},on:{click:t.focus}},[t._v("关注")]),s("van-button",{directives:[{name:"show",rawName:"v-show",value:t.ifFocus,expression:"ifFocus"}],attrs:{round:"",hairline:""},on:{click:t.gotoMessage}},[t._v("发消息")]),s("van-button",{directives:[{name:"show",rawName:"v-show",value:!t.ifFocus,expression:"!ifFocus"}],attrs:{size:"small",round:"",hairline:""},on:{click:function(e){return t.gotoMessage()}}},[s("van-icon",{attrs:{name:"comment-o",size:"20"}})],1),s("van-button",{directives:[{name:"show",rawName:"v-show",value:t.ifFocus,expression:"ifFocus"}],attrs:{size:"small",round:"",hairline:""},on:{click:t.removeFocus}},[s("van-icon",{attrs:{name:"passed",size:"20"}})],1)],1)])]),s("PersonWork",{attrs:{id:t.userid}})],1)])])],1)},i=[function(){var t=this,e=t.$createElement,s=t._self._c||e;return s("div",[s("p",[t._v("0")]),s("p",[t._v("获赞与收藏")])])}],o=s("e38a"),a=s("2893"),r=s("fa7d"),c={data:function(){return{userInfo:{},isNotSelf:!0,ifFocus:!1,userid:this.$route.query.userId,follow_count:0,fans_count:0}},components:{InfoView:o["a"],PersonWork:a["a"]},methods:{onClickLeft:function(){this.$router.back()},getRelation:function(){var t=this,e={relation_id:this.userid};this.$http.get("/relationInfo/getRelation",{params:e}).then((function(e){if(e=e.data,"success"===e.status){var s=e.relationInfo;t.userInfo=Object(r["a"])(s.user_info),t.userid==window.sessionStorage.getItem("userid")?t.isNotSelf=!1:t.ifFocus=1===s.status||3===s.status}}),(function(t){console.log(t)}))},getCount:function(){var t=this;this.$http.get("/relationInfo/getFollowAndFansCount",{params:{userid:String(this.userid)}}).then((function(e,s){var n=e.data.follow_fans_count;t.follow_count=n.follow_count,t.fans_count=n.fans_count})).catch((function(t){console.log(t)}))},focus:function(){var t=this;this.$http.put("/relationInfo/changeRelationStatus?relation_id=".concat(this.userid)).then((function(e){"success"===e.data.status&&(t.getRelation(),t.getCount(),t.$Toast.success("关注成功"))}),(function(t){console.log(t)}))},gotoMessage:function(){this.$router.push({name:"Chatpage",query:{userId:this.userid}})},removeFocus:function(){var t=this;this.$dialog.confirm({title:"取消关注",message:"确定要取消关注吗？"}).then((function(){t.$http.put("/relationInfo/changeRelationStatus?relation_id=".concat(t.userid)).then((function(e){"success"===e.data.status&&(t.$Toast.success("取关成功"),t.getRelation(),t.getCount())}),(function(t){console.log(t)}))}))}},mounted:function(){this.userid=this.$route.query.userId,this.getRelation(),this.getCount()}},l=c,u=(s("71be"),s("2877")),f=Object(u["a"])(l,n,i,!1,null,"2d968911",null);e["default"]=f.exports},"71be":function(t,e,s){"use strict";s("31c4")},"7b27":function(t,e,s){"use strict";s("7fd2")},"7fd2":function(t,e,s){},"9a45":function(t,e,s){"use strict";s("1a16")},a83c:function(t,e,s){"use strict";var n=function(){var t=this,e=t.$createElement,s=t._self._c||e;return s("div",{directives:[{name:"show",rawName:"v-show",value:t.info,expression:"info"}],staticClass:"index"},[s("div",{ref:"partBox",staticClass:"part",attrs:{id:"partBox"},on:{scroll:t.scrollFn}},[s("div",{staticClass:"scrollBox"},t._l(t.info,(function(e,n){return s("div",{key:n,staticClass:"show"},[s("img",{directives:[{name:"lazy",rawName:"v-lazy",value:t.coverImg(e),expression:"coverImg(item)"}],attrs:{src:t.defalute,alt:"简介图片"},on:{click:function(s){return t.gotoDetial(e)}}}),s("p",[t._v(t._s(e&&e.title?e.title:""))]),s("div",{staticClass:"authorInfo"},[s("van-image",{attrs:{height:"20",width:"20",radius:"50%",src:e.author_info.head_photo,fit:"fill"}}),s("span",{staticClass:"authname"},[t._v(t._s(e.author_info.username))]),s("ShowInfoAction",{attrs:{msgItem:e,likeCount:e.give_like_count}})],1)])})),0)])])},i=[],o=s("01ac"),a={data:function(){return{defalute:s("0a9f"),name:"ShowPart"}},mounted:function(){this.isAll&&window.removeEventListener("scroll",this.scrollFn)},computed:{scrollFn:function(){return this.throttle(this.listenBottomOut)},coverImg:function(){var t=this;return function(e){if(!e)return"";if(1===e.is_video)return e.cover_url?e.cover_url:t.defalute;var s=e.resource_url?e.resource_url[0]:t.defalute;return e.cover_url?e.cover_url:s}}},components:{ShowInfoAction:o["a"]},props:["info","isAll"],methods:{gotoDetial:function(t){0===t.is_video?this.$router.push({name:"Detial",query:{messageId:t.article_id}}):this.$router.push({name:"video",query:{messageId:t.article_id}})},listenBottomOut:function(){var t=this;this.$nextTick((function(){var e=document.getElementById("partBox"),s=document.getElementsByClassName("scrollBox")[0],n=e.scrollTop||document.body.scrollTop,i=e.clientHeight,o=s.scrollHeight;n+i>=o&&(t.isAll?window.removeEventListener("scroll",t.scrollFn):t.$bus.$emit("getAllarticle"))}))},throttle:function(t){var e=this,s=arguments.length>1&&void 0!==arguments[1]?arguments[1]:500,n=0;return function(){var i=(new Date).getTime();if(i-n>s){n=i;for(var o=arguments.length,a=new Array(o),r=0;r<o;r++)a[r]=arguments[r];t.apply(e,a)}}}},destroyed:function(){window.removeEventListener("scroll",this.scrollFn)}},r=a,c=(s("7b27"),s("2877")),l=Object(c["a"])(r,n,i,!1,null,"451281d3",null);e["a"]=l.exports},e38a:function(t,e,s){"use strict";var n=function(){var t=this,e=t.$createElement,s=t._self._c||e;return s("div",{staticClass:"userInfo"},[s("img",{staticClass:"bgImg",attrs:{src:t.info.back_photo,alt:""}}),s("div",{staticClass:" perInfo"},[s("img",{staticClass:"perfile",attrs:{src:t.info.head_photo,alt:""}}),s("div",{staticClass:"userName"},[t._v(t._s(t.info.username))])]),s("div",{staticClass:"signature"},[s("p",[t._v(t._s(t.info.info))])])])},i=[],o={data:function(){return{perInfo:null}},props:{info:{type:Object,require:!1}},mounted:function(){},methods:{},computed:{}},a=o,r=(s("9a45"),s("2877")),c=Object(r["a"])(a,n,i,!1,null,"e7a8cd54",null);e["a"]=c.exports}}]);
//# sourceMappingURL=chunk-0053cef9.045e1317.js.map