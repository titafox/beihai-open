<template>
	<view>
		<u-navbar :titleStyle="{'color': '#fff'}" bgColor="#21005D" title="上传作品" :placeholder="true" class="nav_bar"
			:autoBack="true" leftIconColor="#151517">
			<u-icon name="arrow-leftward" color="#fff" slot="left"></u-icon>
		</u-navbar>
		<view class="mint_container">
			<u-upload :sizeType="['original']" :fileList="fileList" @afterRead="afterRead" @delete="deletePic" name="6"
				:maxCount="1" capture="album" width="320" height="189">
				<view class="custom-upload_container" :class="{'isError': ImageErr}">
					<view>
						<u-text text="作品格式PNG, JPG, WEBP，最大 50MB" size="11" color="#7F67BE"></u-text>

						<u-gap height="28"></u-gap>

						<u-button
							color="linear-gradient(0deg, rgba(0, 0, 0, 0.2), rgba(0, 0, 0, 0.2)), linear-gradient(98.46deg, #4F75FF -0.18%, #6954FF 48.62%, #892CFF 100.99%, #892CFF 103.22%)"
							text="选择作品文件" size="large" shape="circle"></u-button>
					</view>
				</view>
			</u-upload>

			<u-form labelPosition="top" :model="mintInfo" :rules="rules" ref="form">
				<u-form-item label="作品名称" prop="name" labelWidth="auto">
					<u-input :adjustPosition="false" v-model="mintInfo.name" placeholder="Pln uranus #06"></u-input>
				</u-form-item>
				<u-form-item label="销售价格" prop="price" labelWidth="auto">
					<u-input :adjustPosition="false" v-model="mintInfo.price" placeholder="¥ 12" type="number">
					</u-input>
				</u-form-item>
				<u-form-item label="销售份数" prop="count" labelWidth="auto">
					<u-input v-model="mintInfo.count" placeholder="" type="number"></u-input>
				</u-form-item>
				<u-form-item label="创作故事" prop="desc" labelWidth="auto">
					<u-textarea v-model="mintInfo.desc" placeholder="请输入内容" confirmType="done">
					</u-textarea>
				</u-form-item>

				<u-form-item label="销售时间" prop="desc" labelWidth="auto">
					<!-- <u-input v-model="mintInfo.time" placeholder=""   @focus="timeFocus"></u-input> -->
					<view class="custom_input" @click="timeFocus">
						{{mintInfo.time}}
					</view>
				</u-form-item>
			</u-form>
			<u-gap height="100"></u-gap>

		</view>
		<u-picker :show="showTime" :columns="times" keyName="name" @cancel="timeCancelHandler"
					@confirm="timeConfirmHandler" @change="timeChangeHandler"></u-picker>
		<view class="btn_group">
			
			<u-text text="将收取上链服务费9.9元" size="12" color="#7F67BE" align="center"></u-text>
			<u-gap height="10"></u-gap>
		
			<u-button
				color="linear-gradient(0deg, rgba(0, 0, 0, 0.2), rgba(0, 0, 0, 0.2)), linear-gradient(98.46deg, #4F75FF -0.18%, #6954FF 48.62%, #892CFF 100.99%, #892CFF 103.22%)"
				text="上链发售" size="large" @click="orderPayHandler" shape="circle"></u-button>
				
		</view>

		<u-action-sheet :show="sheetShow" @close="sheetShow = false" :actions="actions" cancelText="取消">
		</u-action-sheet>
		<u-popup :safeAreaInsetBottom="true" :safeAreaInsetTop="true" mode="center" :show="show" round="10"
			:overlay="true" borderRadius="8">
			<view class="popup_container">
				<u-image class="image" height="100" width="134" src="/static/popup_icon/mint_success.png" :fade="true"
					duration="450" mode="widthFix" :radius="4"></u-image>
				<u-gap height="32"></u-gap>
				<u-text class="title_1" size="18" color="#fff" text="上传成功" align="center"></u-text>
				<u-gap height="4"></u-gap>
				<u-text class="title_2" size="12" color="#D0BCFF" text="目标成为最大的服务型中文Crypto交易者社区" align="center">
				</u-text>
				<u-text class="title_2" size="12" color="#D0BCFF" text="由兔子和粉丝们共同发起成立。" align="center"></u-text>
				<u-gap height="16"></u-gap>
				<u-text class="title_3" size="12" color="#fff" text="作品链接" align="center"></u-text>
				<u-text class="title_4" size="10" color="#D0BCFF" text="https://www.www.com/akjakjsdjlsjd"
					align="center"></u-text>
				<u-gap height="8"></u-gap>
				<view class="copy_container">
					<u-image src="/static/popup_icon/copy.png" width="13" height="13" :fade="true" duration="450"
						mode="aspectFit">
					</u-image>
					<u-text text="复制链接"></u-text>
				</view>
				<u-gap height="32"></u-gap>
				<u-button color="#2E2F3A" text="上传更多" size="large" @click="resetInfo"></u-button>
				<u-gap height="8"></u-gap>
				<u-button type="info" text="进入作品页" size="large" @click="linkToCollection"></u-button>


			</view>
		</u-popup>

	</view>
</template>

<script>
	// const UPPay = uni.requireNativePlugin('UP-Pay');
import {
		mapState
		// mapMutations
	} from 'vuex';
	import {
		prodUploadReq,
		prodMCreateReq
	} from '../../api/index.js';

	import {
		BASEURL
	} from '../../utils/constant.js';
	export default {
		data() {
			return {
				fileList: [],
				ImageErr: false,
				mintInfo: {
					name: '',
					price: '',
					desc: '',
					count: '',
					file: '',
					time:'马上发售',
					timeId:0
				},
				rules: {
					name: {
						type: 'string',
						required: true,
						message: '请填写作品名称',
						trigger: ['blur', 'change']
					},
					price: {
						type: 'number',
						required: true,
						message: '请填写销售价格',
						trigger: ['blur', 'change']
					},
				},
				show: false,
				sheetShow: false,
				actions: [{
						name: '微信支付',
					},
					{
						name: '支付宝支付',
					},
					{
						name: '选项3',
					},
				],
				showTime: false,
				times: [
					[{
						name:'马上发售',
						value:0
					}, {
						name:'明天中午12点',
						value:1
					}, {
						name:'后天中午12点',
						value:2
					}]
				],
			};
		},
		onLoad() {
			// UPPay.initialize({
			// 	appId: "8a81c1be818ca2270181f6c390a204c5", //开放平台appid
			// 	universalLink: "https://app.gyxxx.com/", //微信通用连接(ios端),
			// 	aliMiniPayScheme: "gyscheme", //支付宝小程序AppScheme(ios端),
			// 	cloudPayScheme: "ysfupppay" //云闪付UrlScheme(ios端)
			// });
		},
		onShow() {
			// this.upYunpay();
		},
		computed: {
			...mapState({
				m_count: state => state.pay.m_count,
			}),
		},
		watch:{
			m_count(newval,oldval){
				if(newval>oldval){
					uni.navigateTo({
						url:'/pages/collection/collection?active=2'
					})
				}
			}
		},
		methods: {
			// 删除图片
			deletePic(event) {
				this.fileList.splice(event.index, 1);
			},
			async afterRead(event) {

				// 当设置 mutiple 为 true 时, file 为数组格式，否则为对象格式
				let lists = [].concat(event.file);
				let fileListLen = this.fileList.length;
				lists.map(item => {
					this.fileList.push({
						...item,
						status: 'uploading',
						message: '上传中'
					});
				});
				for (let i = 0; i < lists.length; i++) {
					const result = await this.uploadFilePromise(lists[i].url);
					let item = this.fileList[fileListLen];
					this.fileList.splice(fileListLen, 1, Object.assign(item, {
						status: 'success',
						message: '',
						url: result
					}));
					this.ImageErr = false;
					fileListLen++;
				}
			},
			uploadFilePromise(url) {


				return new Promise((resolve, reject) => {
					let a = prodUploadReq(url).then(res => {
						this.mintInfo = {
							...this.mintInfo,
							file: res,
						};
						resolve(res);
					})
				});
			},
			resetInfo() {
				this.fileList = [];
				this.mintInfo = {
					name: '',
					price: '',
					desc: '',
					count: '',
					file: ''
				};
				this.show = false;
			},
			linkToCollection() {
				uni.switchTab({
					url: '/pages/collection/collection'
				});
			},
			timeFocus(){
				uni.hideKeyboard();
				this.showTime = true;
			},
			timeChangeHandler(e) {
				const {
					value,
					name
				} = e.value[0];
			
				this.mintInfo = {
					...this.mintInfo,
					time: name,
					timeId: value
				}
				this.showTime = false;
			
			},
			timeConfirmHandler(e) {
				const {
					value,
					name
				} = e.value[0];
							
				this.mintInfo = {
					...this.mintInfo,
					time: name,
					timeId: value
				}
				this.showTime = false;
			
			},
			timeCancelHandler() {
				this.shwoTime = false;
			},
			orderPayHandler() {
				const {
					fileList
				} = this;

				if (!(fileList.length > 0 && fileList[0].status === 'success')) {
					this.ImageErr = true;
					return;
				}


				this.$refs.form.validate().then(res => {
					const {
						name,
						price,
						desc,
						count,
						file,
						timeId
					} = this.mintInfo;


					prodMCreateReq({
						name,
						price: parseFloat(price),
						image: file,
						count: parseInt(count),
						description: desc,
						sale_time:timeId,
						pay_type: 2
					}).then((res) => {
						
						let {
							params
						} = res;
						params = JSON.parse(params)
						if(params.resp_code === 'C00000'){
							const {pay_url} = params
							plus.runtime.openWeb(pay_url,()=>{
								// 调用失败 callback
								uni.$u.toast('调起支付失败，请稍后再试！');
							}); 
						}else {
							uni.$u.toast(params.resp_desc);
						}
						
						
 						// TODO 跳转到个人发售页面
						
						// this.pay();
						
						// uni.switchTab({
						// 	url:'/pages/collection/collection?active=2'
						// })
						
						
						// this.show = true;
						// this.upYunpay()
						// this.sheetShow = true;
						// console.log(res)
						// const url= res.appPayRequest.qrCode
						// plus.runtime.openURL(url,()=>{
						// 	// 调用失败 callback
						// });  

						// const url= res.appPayRequest.qrCode
						// plus.runtime.openWeb(url,()=>{
						// 	// 调用失败 callback
						// }); 

						// this.upWxPay(res.appPayRequest)



					})
				}).catch(errors => {})
				// uni.navigateTo({ url: '/pages/pay/pay' });
			},
			pay(){
				// uni.switchTab({
				// 	url:'/pages/collection/collection?active=2'
				// })
				
				 
				
				// const url= res.appPayRequest.qrCode
				
			},
			pay1(){
			const url = 'http://127.0.0.1/test.html'	
			plus.runtime.openURL(url,()=>{
				// 调用失败 callback
			}); 
			},
			pay2(){
				const url='https://hfpay.testpnr.com/h5/pages/cashier/index?uuid=hfpwallet6666000100035059pay53618ba6-68f4-4515-88d1-4290ec95af6f'
				plus.runtime.openWeb(url,()=>{
					// 调用失败 callback
				}); 
			}
			// upWxPay(appPayRequest) {
			// 	//银联综合支付(微信)
			// 	console.log(appPayRequest)

			// 	// "appid": "",
			// 	// "noncestr": "KoILqTtUlXZvlMYMlBVOtomPrOQItWsq",
			// 	// "package": "Sign=WXPay",
			// 	// "partnerid": "526196802",
			// 	// "prepayid": "ori=32FY3d96ce61d2bd41e-1",
			// 	// "sign": "42BD220CDD11CA73E8A4D5D79B6C349D",
			// 	// "timestamp": "20220822194520"

			// 	var params = {
			// 		// "package":"Sign=WXPay",
			// 		// "sign":"F45D3849A6421E5ED49CE2C29F83C8CF",
			// 		// "partnerid":"466003135",
			// 		// "prepayid":"ori=130FU20220311101444518",
			// 		// "noncestr":"BgqaiZiuJYwawlXcStbQgVzLIeOQQWQQ",
			// 		// "timestamp":"20220311221444",

			// 		// "appid":"wx5789f26e37b6ae3b",
			// 		// "miniuser":"gh_744d2ebca056",//固定不变
			// 		// "minipath":"pages/appPay/index"//固定不变

			// 		...appPayRequest
			// 	};
			// 	var payChannel = "01"; //支付渠道 01微信 04支付宝小程序
			// 	var payData = JSON.stringify(params); //服务端下单获取
			// 	UPPay.upPay(payChannel, payData, function(res) {
			// 		console.log(res);
			// 		var returnCode = res.resultCode;
			// 		var resultInfo = res.resultInfo;
			// 	});
			// },
			// upAliPay(appPayRequest) {
			// 	//银联综合支付(支付宝)
			// 	var params = {

			// 		// "msgType":"trade.appPreOrder",
			// 		// "noncestr":"lSvXoGuhZptpBuHLtCsKjqDIqUDYTDCu",
			// 		// "package":"Sign=ALI",
			// 		// "prepayid":"ori=11VC022041311135464687j",
			// 		// "sign":"B763B59FC79A308F3390CC18D444B0FE",
			// 		// "timestamp":"20220413112539"


			// 		// "miniuser":"2019010762862511",//固定不变
			// 		// "minipath":"pages/appPay/index/index",//固定不变
			// 		// "appScheme":"iOS:gyscheme;Android:gyscheme",
			// 		...appPayRequest
			// 	};
			// 	var payChannel = "04"; //支付渠道 01微信 04支付宝小程序
			// 	var payData = JSON.stringify(params); //服务端下单获取
			// 	UPPay.upPay(payChannel, payData, function(res) {
			// 		console.log(res);
			// 		var returnCode = res.resultCode;
			// 		var resultInfo = res.resultInfo;
			// 	});
			// },
			// upYunpay() {
			// 	var params = {
			// 		tn: "499192794496866394405"
			// 	};
			// 	var payData = JSON.stringify(params); //服务端下单获取
			// 	UPPay.ysfPay(payData, function(res) {
			// 		console.log(res);
			// 		var returnCode = res.resultCode;
			// 		var resultInfo = res.resultInfo;
			// 	});
			// }

		}
	};
</script>

<style scoped>
	.mint_container {
		padding: 48rpx 48rpx calc(env(safe-area-inset-bottom) + 24rpx);
	}

	.u-upload {
		align-items: center;
	}

	.custom-upload_container {
		width: 320px;
		height: 189px;
		/* background: #F2F2F5; */
		border: 1px solid #EBEBF0;
		border-radius: 8px;
		display: flex;
		align-items: center;
		justify-content: center;
	}

	.custom-upload_container>view {
		display: flex;
		flex-direction: column-reverse;
		/* border: 1px dashed #00C566; */
		border-radius: 5px;
		padding: 8px 12px;
	}

	.custom-upload_container .tips>>>span {
		font-weight: 500;
		line-height: 18px;
		letter-spacing: -0.025em;
		margin-left: 4px;

	}

	.u-form>>>.u-form-item__body__left__content__label {
		font-weight: 600;
		font-size: 14px;
		line-height: 22px;
		letter-spacing: 0.005em;
		color: #2E2F3A;
	}

	.btn_group {
		/* position: fixed; */
		display: flex;
		flex-direction: column;
		align-items: center;
		justify-content: center;
		/* bottom: 0; */
		width: calc(100vw - 98rpx);
		/* background-color: #fefefe; */
		padding: 48rpx 48rpx calc(env(safe-area-inset-bottom) / 2 + 24rpx);
	}

	.popup_container {
		width: 642rpx;
		height: 500px;
		/* display: flex; 
		flex-direction: column;
		align-items: center;
		justify-content: center; */
		text-align: center;
		padding: 47px 0 0;
		background-color: #000;
	}

	.popup_container .image {
		display: flex;
		align-items: center;
	}

	.title_1>>>span {
		font-weight: 600;
		line-height: 26px;
		letter-spacing: 0.005em;
	}

	.title_2>>>span {
		font-weight: 400;
		line-height: 20px;
		letter-spacing: 0.005em;
	}

	.title_3>>>span {
		font-weight: 500;
		line-height: 20px;
		letter-spacing: 0.005em;
	}

	.title_4>>>span {
		font-weight: 400;
		line-height: 18px;
		letter-spacing: 0.005em;
	}

	.copy_container {
		display: flex;
		align-items: center;
		justify-content: center;
		background: #F2F2F5;
		border-radius: 8px;
		padding: 4px 8px;
		margin: 0 auto;
		width: 86px;
	}

	.copy_container>>>.u-image {
		margin-right: 6px;
	}

	.u-button {
		width: 206px;
		margin: 0 auto;
	}

	.u-button:last-child {
		border-color: #2e2f3a;
	}

	.isError {
		border-color: red;
	}

	.u-textarea {
		background-color: transparent;
	}
	
	
</style>
