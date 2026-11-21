import request from '../utils/request.js';
import upload from '../utils/upload.js';

// 获取验证码 ※※※
export const veriCodeReq = data => request({
	url: 'v1/account/smsLogin/code',
	data
});

// 登录 ※※※
export const loginReq = data => request({
	url: 'v1/account/smsLogin',
	data,
	isTips: true
});

// 首页banner ※※※
export const bannerListReq = () => request({
	url: 'v1/banner/list',
	method: 'GET'
});

// 首页销售排行 ※※※
// export const saleListReq = () => request({
// 	url: 'v1/authors/ranking',
// 	method: 'GET'
// });

export const saleListReq = () => request({
	url: 'v1/product/saleRankList',
	method: 'GET'
});

// 首页 推荐 ※※※
export const recommListReq = () => request({
	url: 'v1/product/recommend/list',
	method: 'GET'
});

// 首页推荐点击全部进入页面的数据  ※※※
export const allNftsReq = data => request({
	url: 'v1/product/list',
	data,
	method: 'GET'
});

// 个人账户详细信息
export const accountInfoReq = () => request({
	url: 'v1/account/info',
	method: 'GET'
});

// 个人BSN地址二维码
export const bsnAddressReq = () => request({
	url: 'v1/account/qrcode?date=9',
	method: 'GET'
});

// 更新手机号获取验证码
export const updateVeriCodeReq = data => request({
	url: 'v1/account/updatePhone/code',
	data,
	isTips: true
});

// 更新手机号
export const updatePhoneReq = data => request({
	url: 'v1/account/updatePhone/code',
	data,
	isTips: true
});

// 用户实名认证
export const accountAuthReq = data => request({
	url: 'v1/account/authentication',
	data,
	isTips: true
});

//银行列表
export const bankListReq = () => request({
	url: 'v1/bank/list',
	method: 'GET'
});



// 用户银行卡列表
export const accountBankListReq = data => request({
	url: 'v1/bank/account/list',
	data,
	method: 'GET'
});

// 绑定银行卡
export const cardBindReq = data => request({
	url: 'v1/bank/account/create',
	data
});

// 解绑银行卡
export const cardUnbindReq = data => request({
	url: `v1/bank/account/unbound/${data}`
});

// 创建作品
export const prodMCreateReq = data => request({
	url: 'v1/product/create',
	data
});

// 点赞和取消
export const prodLikeReq = data => request({
	url: `v1/product/like/${data}`,
	isTips: true
});

//作品详情
export const prodDetailReq = data => request({
	url: `v1/product/detail/${data}`,
	method: 'GET'
});

// 更新用户信息
export const accountUpdateReq = data => request({
	url: 'v1/account/update',
	data,
	isTips: true
});

// 上传作品  
export const prodUploadReq = filePath => upload({
	url: 'v1/storage/productUpload',
	filePath
});

// 上传头像  ※※※
export const avatarUploadReq = filePath => upload({
	url: 'v1/storage/avatarUpload',
	filePath
});

// 上传身份证
export const idCardUploadReq = filePath => upload({
	url: 'v1/storage/idcardUpload',
	filePath
});


// 我发布的作品
export const publishListRqe = data => request({
	url: 'v1/product/publishList',
	data,
	method: 'GET'
});

// 我卖出的作品
export const likeListReq = data => request({
	url: 'v1/product/saleList',
	data,
	method: 'GET'
});

// 我购买的作品
export const buyListReq = data => request({
	url: 'v1/product/buyList',
	data,
	method: 'GET'
});

// 提现验证码
export const withdrawalCode = data => request({
	url: 'v1/withdraw/code',
	data
})

// 提现

export const withdrawal = data => request({
	url: 'v1/withdraw/create',
	data
})

//交易记录

export const tradeHistoryReq = data => request({

	url: 'v1/account/tradeHistory',
	data,
	method: 'GET'
})

// 别的用户信息

export const otherProfileReq = data => request({
	url: `v1/account/info/${data}`,
	method: 'GET'
})

// 别人已卖出列表  
export const otherSaleListReq = data => request({
	url: `v1/product/userSaleList/${data.id}`,
	data: {
		page: data.page,
		page_size: data.page_size
	},
	method: 'GET'
})

export const otherHistoryListReq = data => request({
	url: `v1/product/salesHistory/${data.id}`,
	data: {
		page: data.page,
		page_size: data.page_size
	},
	method: 'GET'
})

export const otherProdListReq = data=>request({
	url:`v1/product/productList/${data.id}`,
	data: {
		page: data.page,
		page_size: data.page_size
	},
	method:'GET'
})


export const radioInfoReq = () => request({
	url: 'v1/notice',
	method: 'GET'
})

export const noticeTypeReq = () => request({
	url: 'v1/notice/type',
	method: 'GET'
})

export const allNoticeReq = ()=>request({
	url:'v1/notice/type/all',
	method:'GET'
})

export const singleNoticeReq = (data)=>request({
	//url:`v1/notice/type/list?nid=${data.nid}`,
	url:`v1/notice/type/list/${data.nid}`,
	method: 'GET'
})


export const artistsReq =()=>request({
	url:'v1/account/artists',
	method: 'GET'
})

export const noticeDetailReq = (data)=>request({
	url:`v1/notice/detail/${data.id}`,
	method: 'GET'
})

export const buyReq = (data)=>request({
	url:'v1/order/commit',
	data,
})


export const handselListReq = (data)=>request({
	url:'v1/handsel/list',
	data,
	method:'GET'
})

export const handsel2OtherReq = (data)=>request({
	url:'v1/works/handsel',
	data
})

export const orderListReq = (data)=>request({
	url:`v1/order/list/${data.typ}`,
	method:'GET'
})

// export const orderListReq = (data)=>request({
// 	url:`v1/order/list`,
// 	data,
// 	method:'GET'
// })

export const versionReq = ()=>request({
	url:'v1/app/version',
	method:'GET'
})

// 以下为二期接口

export const market = (data)=>request({
	url:'v1/market',
	data,
	method:'GET'
})

export const marketdetail = (data)=>request({
	url:'v1/market/detail',
	data,
	method:'GET'
})

// 已购藏品


export const nfts = (data)=>request({
	url:'v1/account/nfts',
	data,
	method:'GET'
})

// 已购藏品明细

export const nftsseries = (data)=>request({
	url:'v1/account/nfts/series',
	data,
	method:'GET'
})

// 我的邀请


export const invitelist = (data)=>request({
	url:'v1/account/invite/list',
	data,
	method:'GET'
})

export const inviteTop = () => request({
	url: 'v1/account/invite/top',
	method: 'GET'
})

// 钱包开通
export const walletcreate = ()=>request({
	url:'v1/wallet/create',	
	method:'POST'
})

// 钱包开通
export const walletinfo = ()=>request({
	url:'v1/wallet/info',	
	method:'GET'
})

// 下单 nftbuy
export const nftbuy = (data)=>request({
	url:'v1/market/nft/buy',
	data,	
	method:'POST'
})

// 上架

export const nfton = (data)=>request({
	url:'v1/market/nft/on',
	data,	
	method:'POST'
})

// 下架


export const nftoff = (data)=>request({
	url:'v1/market/nft/off',
	data,	
	method:'POST'
})

//我的寄售

export const nftsells = ()=>request({
	url:'v1/account/nfts/sells',	
	method:'GET'
})



// 单个作品流通列表
export const nftscirculating = (data)=>request({
	url:'v1/market/nfts/circulating',
	data,	
	method:'GET'
})