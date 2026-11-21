// 首页 获取推荐作者和作品1
const recommendAuthor  = () => request({
	url:'index/recommend/author'
})

//首页轮播图推荐2
const recommendCarousel = () => request({
	url:'index/recommend/carousel'
})


//喜欢 3
const castingLike = (data)=> request({
	url:'index/like',
	method:'post',
	data:{
		...data
	}
});

// 用户未登录 创建用户4
const createUser = (data) =>  request({
	url:'user/login',
	method:'post',
	data: {
		...data
	}
});

//获取别人的信息或者登录时需要的签名信息5
const userProfile = (data) => request({
	url:'user/profile',
	data:{
		...data
	}
})

//获取铸件详情6
const worksDetail = (data) => request({
	url:'works/detail',
	data:{
		...data
	}
});


//获取类别7
const categoryList = (data) => request({
	url:'works/item-category',
	data:{
		...data
	}
});


//获取铸件分类对应的铸件8
const castingList = (data)=> request({
	url:'index/item',
	data: {
		...data
	}
});

//9
const userSign = (data) => request({
	method: 'post',
	url:'user/sign',
	data:{
		...data
	}
});

//详情页面价格信息10
const castingTransaction = (data) => request({
	url:'casting/transaction',
	data: {
		...data
	}
});


//购买11
const tradeBuy = (data) => request({
	method: 'post',
	url:'trade/buy',
	data:{
		...data
	}
});

//购买之后12
const tradeBuyBefore = (data) => request({
	url:'trade/buy-before',
	method: 'post',
	data: {
		...data
	}
});


//获取推荐铸件13
const worksRecommend = (data) => request({
	url: 'works/recommend',
	data: {
		...data
	}
});

//获取登录用户的信息14
const userCenter = (data) => request({
	url:'user/center',
	method:'post',
	data:{
		...data
	}
});


//我拥有的作品15
const worksOwner = (data) => request({
	url:'works/owner',
	data: {
		...data
	}
});

//别人拥有的16
const otherWorksOwner = (data) => request({
	url:'works/other/owner',
	data: {
		...data
	}
});


//我创作的作品17
const authorCasting = (data) => request({
	url: 'casting/author',
	data: {
		...data
	}
})

//别人创作的18
const otherAuthorCasting = (data) => request({
	url:'casting/other/author',
	data: {
		...data
	}
})


// 获取铸件合集里面的作品19
const collectionWorks = (data) => request({
	url:'collection/works',
	data: {
		...data
	}
})


//修改登录用户信息20
const setUserProfile = (data) => request({
	url:'user/set-profile',
	method:'post',
	data: {
		...data
	}
});


//是否有铸件权限21
const isUpload = (data) => request({
	url:'user/is_upload',
	data: {
		...data
	}
})

//发布单个铸件22
const saveCasting = (data) => request({
	method:'post',
	url:'casting/save-casting',
	data: {
		...data
	}
});

//编辑单个铸件23
const worksEdit = (data) => request({
	method:'post',
	url:'works/edit',
	data: {
		...data
	}
});

//创建合集24
const createCollection = (data) => request({
	method:'post',
	url: 'collection/create',
	data: {
		...data
	}
})


// 批量上传并且生成合集25
const worksUpload = (data) => request({
	method:'post',
	url: 'works/upload',
	data: {
		...data
	}
})


//在已有的合集里面批量上传26
const worksUploadAdd = (data) => request({
	method:'post',
	url: 'works/upload-add',
	data: {
		...data
	}
})


//获取合集详情27
const getCollection = (data) => request({
		url: 'collection/get',
		data:{
			...data
		}
	}

)

//编辑合集28
const editCollection = (data) => request({
	method:'post',
	url: 'collection/edit',
	data:{
		...data
	}
})

const importPre = (data) => request({
	method:'post',
	url:'collection/import-pre',
	data:{
		...data
	}
})

const importCollection = (data) => request({
	method:'post',
	url:'collection/import',
	data:{
		...data
	}
})


const createSubCasting = (data) => request({
	method:'post',
	url:'casting/create-sub',
	data:{
		...data
	}
})

//绑定推荐人

const bindUser = (data) => request({
	method:'post',
	url:'user/bind',
	data:{
		...data
	}
})





const balanceOfCTS = (data) => request({
	method: 'get',
	url: 'sc/balanceOf',
	data: {
		...data
	}
})


const getApprovalInfo = (data) => request({
	methods: 'get',
	url: 'sc/approve',
	data: {
		...data
	}
})


const getBuyContract = (data) => request({
	methods: 'get',
	url: 'sc/buy',
	data: {
		...data
	}
})



const getBuyTokenInfo = (data) => request({
	methods: 'get',
	url: 'sc/buy',
	data: {
		...data
	}
})

const getIsApprovalForAllInfo = (data) => request({
	methods: 'get',
	url: 'sc/isApprovedForAll',
	data: {
		...data
	}
})


const getsetApprovalForAllInfo = (data) => request({
	method: 'get',
	url: 'sc/setApprovalForAll',
	data: {
		...data
	}
})


const getTicketSellInfo = (data) => request({
	method: 'get',
	url: 'sc/sell',
	data: {
		...data
	}
})



