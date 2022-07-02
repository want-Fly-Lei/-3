// pages/changeinfo/changeinfo.ts
var app=getApp()

Page({

    /**
     * 页面的初始数据
     */
    data: {
        username:null,
        password:null,
        email:null,
        id:null
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

    input3:function(e){
        this.setData({email:e.detail.value})
        console.log(this.data.email)
    },

    change:function(){
        var that=this
        // if(this.data.username!=null){app.globalData.userinfo.username=this.data.username}
        // if(this.data.password!=null){app.globalData.userinfo.password=this.data.password}
        // if(this.data.email!=null){app.globalData.userinfo.email=this.data.email}
        // console.log(this.data.email)
        // console.log(this.data.id)
        // console.log(this.data.password)
        // console.log(this.data.username)
        wx.request({
            url: "http://localhost:8086/user/reset",
            data: {
                'id':this.data.id,
                'username': this.data.username,
                'password': this.data.password,
                'email':this.data.email
            },
            method: "POST",
            header: {
              'content-type': 'application/x-www-form-urlencoded'//注意个人使用application/json获取不到数据
            },
            
            success: function (res) {//res.data.XXX是取到后端的数据如果和admin是等的就显示登录成功跳转的相应的小程序页面
                console.log(that.data.email)
                console.log(that.data.id)
                console.log(that.data.password)
                console.log(that.data.username)
              if (res.data.msg == "ok") {
                wx.showToast({
                  title: '修改成功',
                  icon: 'success',
                  duration: 20000
                })
                setTimeout(function(){
                  wx.hideToast();
                }),
                //   wx.redirectTo({
                //     url: '../mine/mine?username=' + res.data.username,
                //   })
                wx.navigateBack({delta:1})
              } else {
                wx.showToast({
                  title: '账号或密码错误',
                  icon: 'loading',
                  duration: 2000
                })
              }
            }
          })
        
    },

    /**
     * 生命周期函数--监听页面加载
     */
    onLoad() {
        this.setData({id:app.globalData.userinfo.id})
        console.log(this.data.id)
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