// pages/manageuserinfo/manageuserinfo.ts
Page({

    /**
     * 页面的初始数据
     */
    data: {
        username:null,
        password:null,
        email:null,
        root:null,
        id:null,
        users:[]
    },

    change:function(){
        wx.navigateTo({url:"../machangeinfo/machangeinfo?id=" + this.data.id})
    },

    searchall:function() {
        var that = this
        wx.request({
            url: "http://localhost:8086/user/allUser",
            method: "GET",
            header: {
                'content-type': 'application/x-www-form-urlencoded'//注意个人使用application/json获取不到数据
            },
            success: function (res) {//res.data.XXX是取到后端的数据如果和admin是等的就显示登录成功跳转的相应的小程序页面
                // console.log(res.data)
                that.setData({ users: res.data })
                app.globalData.users = that.data.users
                if (1) {
                    wx.showToast({
                        title: '查询成功',
                        icon: 'success',
                        duration: 20000
                    })
                    that.setData({ users: res.data })
                    app.globalData.users = that.data.users
                    // console.log(that.data.books)
                    setTimeout(function () {
                        wx.hideToast();
                    })

                } else {
                    wx.showToast({
                        title: '查询错误',
                        icon: 'loading',
                        duration: 2000
                    })
                }
            }
        })
    },

    searchByUserName:function() {

    },

    /**
     * 生命周期函数--监听页面加载
     */
    onLoad() {
        wx.request({
            url: "http://localhost:8081/Examination_System/wxlogin",
            data: {
              'id': username,
              'username': password,
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
                })
                setTimeout(function(){
                  wx.hideToast();
                }),
                  wx.navigateTo({
                    url: '/pages/admin/admin/index?username=' + res.data.username,
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