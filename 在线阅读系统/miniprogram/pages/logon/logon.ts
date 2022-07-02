// pages/logon/logon.ts
var app=getApp();

Page({

    /**
     * 页面的初始数据
     */
    data: {
        username:null,
        password:null,
        msg:null,
        id:null,
        root:null
    },

    input1:function(e){
        // console.log(e)
        this.setData({username:e.detail.value})
        console.log(this.data.username)
    },

    input2:function(e){
        this.setData({password:e.detail.value})
        console.log(this.data.password)
    },

    tologon1:function(){
        wx.navigateTo({
            url:"../logon1/logon1"
        })
    },

    tonewlogon:function(){
        wx.navigateTo({
            url:"../newlogon/newlogon"
        })
    },

    tomine:function(){
        var that = this;
        
        wx.request({
            url: "http://localhost:8086/user/login",
            data: {
              'username': this.data.username,
              'password': this.data.password,
            },
            method: "POST",
            header: {
              'content-type': 'application/x-www-form-urlencoded'//注意个人使用application/json获取不到数据
            },
            success: function (res) {//res.data.XXX是取到后端的数据如果和admin是等的就显示登录成功跳转的相应的小程序页面
              if (res.data.msg == "") {
                wx.showToast({
                  title: '登录成功',
                  icon: 'success',
                  duration: 20000
                }),
                that.setData({username:res.data.username,
                    id:res.data.id,root:res.data.root})
                console.log(that.data.username,that.data.password,
                    that.data.id,that.data.root)
                app.globalData.userinfo={username:that.data.username,password:that.data.password,id:that.data.id,root:that.data.root}
                console.log(app.globalData.userinfo)
                setTimeout(function(){
                  wx.hideToast();
                }),
                  wx.redirectTo({
                    url: '../mine/mine?username=' + res.data.username + "&root=" + res.data.root,
                  })
              } else {
                wx.showToast({
                  title: '账号或密码错误',
                  icon: 'loading',
                  duration: 2000
                })
              }
            }
          })
        
        // if(app.globalData.userinfo.username!=null && app.globalData.userinfo.password!=null){
        //     wx.redirectTo({url:"../mine/mine"})
        // }
        
    },

    /**
     * 生命周期函数--监听页面加载
     */
    onLoad() {
        if(app.globalData.userinfo!=null){
            wx.redirectTo({url:"../mine/mine"})
        }
    },

    /**
     * 生命周期函数--监听页面初次渲染完成
     */
    onReady() {

    },

    /**
     * 生命周期函数--监听页面显示
     */
    onShow() {

    },

    /**
     * 生命周期函数--监听页面隐藏
     */
    onHide() {

    },

    /**
     * 生命周期函数--监听页面卸载
     */
    onUnload() {

    },

    /**
     * 页面相关事件处理函数--监听用户下拉动作
     */
    onPullDownRefresh() {

    },

    /**
     * 页面上拉触底事件的处理函数
     */
    onReachBottom() {

    },

    /**
     * 用户点击右上角分享
     */
    onShareAppMessage() {

    }
})