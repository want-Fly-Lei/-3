// pages/search/search.ts
var app =  getApp();


Page({

    /**
     * 页面的初始数据
     */
    data: {
        data_input:null,
        searchbook:[]
    },

    input:function(e){
        this.setData({data_input:e.detail.value})
        console.log(this.data.data_input)
    },

    toseebook:function(e){
      var ids=e.target.id
      this.setData({id:ids})
      app.globalData.now_bookinfo=({id:this.data.id,c_id:this.data.c_id})
    //   console.log(app.globalData.now_bookinfo)
    //   console.log(this.data.id)
    //   console.log(e.target)
      wx.navigateTo({url:"../read/read?bid="+this.data.id})
      console.log(url)
  },

    submitSearch:function(){
        var that=this
        wx.request({
            url: "http://localhost:8086/book/selectByBookNameMoHu?bookname=" + this.data.data_input ,
            method: "GET",
            header: {
                'content-type': 'application/x-www-form-urlencoded'//注意个人使用application/json获取不到数据
            },
            success: function (res) {//res.data.XXX是取到后端的数据如果和admin是等的就显示登录成功跳转的相应的小程序页面
                // console.log(res.data.context)
                // that.setData({
                //     bookname: res.data.bookname,
                //     description: res.data.description,
                //     url: res.data.cover,
                //     passage: res.data.context
                // })
               
                // console.log(that.data.bookname)
                // console.log(app.globalData.seeingbook)
                // app.globalData.books=that.data.books
                console.log(res.data.msg)
                if (res.data.msg=='') {
                    wx.showToast({
                        title: '登录成功',
                        icon: 'success',
                        duration: 20000
                    })
                    app.globalData.booksearch=res.data.books
                    that.setData({searchbook:app.globalData.booksearch})
                    console.log(app.globalData.booksearch)
                    // that.setData({books:res.data})
                    // app.globalData.books=that.data.books
                    console.log(that.data.books)
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
        // console.log(this.data.data_input)
    },

    /**
     * 生命周期函数--监听页面加载
     */
    onLoad() {

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