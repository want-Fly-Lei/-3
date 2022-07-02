// pages/read/read.ts
var app = getApp()
Page({

    /**
     * 页面的初始数据
     */
    data: {
        bookname: null,
        description: null,
        url: null,
        id: null,
        books: []
    },

    toread: function (e) {
        wx.navigateTo({ url: "../read1/read1" })
    },

    /**
     * 生命周期函数--监听页面加载
     */
    onLoad() {
        var that = this
        this.setData({ id: app.globalData.now_bookinfo.id })
        // console.log(this.data.id)
        wx.request({
            url: "http://localhost:8086/book/allBook",
            method: "GET",
            data: {
                'id': this.data.id
            },
            header: {
                'content-type': 'application/x-www-form-urlencoded'//注意个人使用application/json获取不到数据
            },
            success: function (res) {//res.data.XXX是取到后端的数据如果和admin是等的就显示登录成功跳转的相应的小程序页面
                console.log(res.data)
                that.setData({
                    bookname: res.data.bookname,
                    description: res.data.description,
                    url: res.data.cover,
                    books: res.data
                })
                app.globalData.seeingbook = ({
                    bookname: that.data.bookname, url: that.data.url,
                    description: that.data.description, id: that.data.id,
                    books:that.data.books
                })
                console.log(that.data.bookname)
                console.log(app.globalData.seeingbook)
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
        // this.setData({bookname:app.globalData.now_bookinfo.bookname,
        //     description:app.globalData.now_bookinfo.description,
        // url:app.globalData.now_bookinfo.cover})
        // console.log(this.data.bookname)
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