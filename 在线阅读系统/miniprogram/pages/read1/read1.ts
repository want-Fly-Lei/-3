// pages/read1/read1.ts
var app = getApp()
Page({

    /**
     * 页面的初始数据
     */
    data: {
        bid:null,
        cid:1,
        passage:null
    },

    add:function() {
        var that = this
        that.setData({
          cid: that.data.cid + 1
        })
        wx.request({
            url: "http://localhost:8086/chapter/showChapterContext?bid=" + this.data.bid +"&cid=" + this.data.cid,
            method: "GET",
            header: {
                'content-type': 'application/x-www-form-urlencoded'//注意个人使用application/json获取不到数据
            },
            success: function (res) {//res.data.XXX是取到后端的数据如果和admin是等的就显示登录成功跳转的相应的小程序页面
                console.log(res.data.context)
                that.setData({
                    bookname: res.data.bookname,
                    description: res.data.description,
                    url: res.data.cover,
                    passage: res.data.context
                })
                app.globalData.seeingbook = ({
                    bid: that.data.id,
                    books:that.data.books
                })
                // console.log(that.data.bookname)
                // console.log(app.globalData.seeingbook)
                // app.globalData.books=that.data.books
                if (1) {
                    wx.showToast({
                        title: '查询成功',
                        icon: 'success',
                        duration: 20000
                    })
                    // that.setData({books:res.data})
                    // app.globalData.books=that.data.books
                    // console.log(that.data.books)
                    setTimeout(function () {
                        wx.hideToast();
                    })

                } else {
                    wx.showToast({
                        title: '查询失败',
                        icon: 'loading',
                        duration: 2000
                    })
                }
            }
        }) 
    },

    sub:function() {
        var that = this
        that.setData({
          cid: that.data.cid - 1
        })
        wx.request({
            url: "http://localhost:8086/chapter/showChapterContext?bid=" + this.data.bid +"&cid=" + this.data.cid,
            method: "GET",
            header: {
                'content-type': 'application/x-www-form-urlencoded'//注意个人使用application/json获取不到数据
            },
            success: function (res) {//res.data.XXX是取到后端的数据如果和admin是等的就显示登录成功跳转的相应的小程序页面
                console.log(res.data.context)
                that.setData({
                    bookname: res.data.bookname,
                    description: res.data.description,
                    url: res.data.cover,
                    passage: res.data.context
                })
                app.globalData.seeingbook = ({
                    bid: that.data.id,
                    books:that.data.books
                })
                // console.log(that.data.bookname)
                // console.log(app.globalData.seeingbook)
                // app.globalData.books=that.data.books
                if (1) {
                    wx.showToast({
                        title: '查询成功',
                        icon: 'success',
                        duration: 20000
                    })
                    // that.setData({books:res.data})
                    // app.globalData.books=that.data.books
                    // console.log(that.data.books)
                    setTimeout(function () {
                        wx.hideToast();
                    })

                } else {
                    wx.showToast({
                        title: '查询失败',
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
        var that = this
        this.setData({bid:app.globalData.now_bookinfo.id})
        console.log(this.data.bid)
        wx.request({
            url: "http://localhost:8086/chapter/showChapterContext?bid=" + this.data.bid +"&cid=" + this.data.cid,
            method: "GET",
            header: {
                'content-type': 'application/x-www-form-urlencoded'//注意个人使用application/json获取不到数据
            },
            success: function (res) {//res.data.XXX是取到后端的数据如果和admin是等的就显示登录成功跳转的相应的小程序页面
                console.log(res.data.context)
                that.setData({
                    bookname: res.data.bookname,
                    description: res.data.description,
                    url: res.data.cover,
                    passage: res.data.context
                })
                app.globalData.seeingbook = ({
                    bid: that.data.id,
                    books:that.data.books
                })
                // console.log(that.data.bookname)
                // console.log(app.globalData.seeingbook)
                // app.globalData.books=that.data.books
                if (1) {
                    wx.showToast({
                        title: '登录成功',
                        icon: 'success',
                        duration: 20000
                    })
                    // that.setData({books:res.data})
                    // app.globalData.books=that.data.books
                    // console.log(that.data.books)
                    setTimeout(function () {
                        wx.hideToast();
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