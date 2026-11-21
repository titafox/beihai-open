<template>
	<view style="padding: 2rpx;">
		<u-navbar :titleStyle="{'color': '#fff'}" bgColor="#ffffff00" class="nav_bar" :autoBack="true"
			leftIconColor="transparent">
			<!-- <u-icon name="arrow-leftward" color="#fff" slot="left"></u-icon> -->
			<u-icon name="close-circle-fill" size="33" color="rgba(0, 0, 0, 0.8)" slot="right" @click="goBack"></u-icon>

		</u-navbar>
		<!-- :title="nftInfo.name" -->
		<u-image @load="imageLoad" :height="height" width="748rpx" :src="Info.image" :fade="true" duration="450"
			mode="widthFix">
			<template v-slot:loading>
				<view
					style="width: 100vw;height: 100vw;display: flex;align-items: center;justify-content: center;background-color: #21005D;">
					<u-loading-icon color="white"></u-loading-icon>
				</view>
			</template>
		</u-image>
		<view class="cd-v1">
			<view class="m-flex-sb">
				<view class="cd-v1-t1">{{Info.name}}</view>
				<view class="cd-v1-d1">#{{Info.work_no}}</view>
			</view>
			<view class="m-flex" style="padding:8rpx 8rpx 16rpx 8rpx;align-items: center;">
				<view class="cd-v1-d2">发行</view>
				<text class="cd-v1-t2">{{Info.stock}}份</text>
				<view class="cd-v1-d2" @click="circulation">流通</view>
				<text class="cd-v1-t2" @click="circulation">{{Info.circulation}}份</text>
			</view>
		</view>

		<view class="cd-title">IP介绍</view>
		<view class="cd-t1">{{Info.author_desc}}</view>
		<view class="cd-title">链上数据</view>
		<view class="m-flex  cd-t1">
			<text class="cd-t2" style="color: white;">交易哈希：</text>
			<text class="cd-t2">{{Info.hash}}</text>
		</view>
		<view class="m-flex  cd-t1">
			<text class="cd-t2" style="color: white;">链上标识：</text>
			<text class="cd-t2">#{{Info.nft_id}}</text>
		</view>
		<view class="cd-title">藏家须知</view>


		<u-parse class="cd-t1" :content="content"></u-parse>

		<view class="m-flex-w cd-v3">
			<view class="m-flex m-flex-cl" style="width: 600rpx;">
				<text class="cd-t4">创作者</text>
				<u-avatar :src="Info.author_avatar" size="30"></u-avatar>
				<text class="cd-t4" style="padding-left: 32rpx;">{{Info.author_name}}</text>

			</view>

			<view class="m-flex m-flex-cl" style="width: 600rpx;">
				<text class="cd-t4">持有者</text>
				<u-avatar :src="Info.owner_avatar" size="30"></u-avatar>
				<text class="cd-t4" style="padding-left: 32rpx;">{{Info.owner_name}}</text>

			</view>

		</view>

		<view class="cd-title">寄售记录</view>

		<view class="cd-v2 m-flex-sb" v-for="(item, index) in logs" :key="index">
			<text class="cd-t3" style="width: 200rpx;">{{item.nickname}}</text>
			<text class="cd-t3" style="width: 180rpx;">{{item.typ_text}} ￥{{item.price}}</text>
			<text class="cd-t3" style="width: 320rpx;text-align: right;">{{$u.timeFormat(item.create_time, 'yyyy/mm/dd hh:MM:ss')}}</text>
		</view>

		<u-gap height="100"></u-gap>
		<!-- v-if="tstatus.btnV!=-1" -->
		<view class="m-flex-sb m-flex-c cd-bottom" >
			<view class="cd-bottom-v2" style="color:#7B61FF;margin-left: 32rpx;" @click="merge"  v-if="ntfstatus === 0 && Info.is_synthetic_material === 1">合成</view>
			<view class="cd-bottom-v1" v-else>{{tstatus.title}}</view>
			<view class="cd-bottom-v2" :style="{ color: tstatus.color}" @click="submit"> {{tstatus.btnTxt}}</view>
		</view>

		<u-popup 
			:show="show" 
			:round="2" 
			mode="center" 
			@close="close"
		>
			<view class="m-flex-w cd-popup">
				<view class="m-flex">
					<view class="cd-popup-t1" v-if="popstatus==0">寄售</view>
					<view class="cd-popup-t2" v-if="popstatus==0">{{Info.name}}</view>
					<view class="cd-popup-t1" v-if="popstatus==1">上架成功</view>
				</view>
				<u-gap height="60" v-if="popstatus==0"></u-gap>

				<u--input 
					v-if="popstatus==0" 
					type="number" 
					prefixIcon="tags" 
					placeholder="请输入销售价格" 
					shape="circle"
					v-model="value"
				></u--input>

				<view class="m-flex" v-if="popstatus==1" style="margin-bottom: 80rpx;">
					<view class="cd-popup-t4">请注意关注你在寄售市场的销售情况</view>
					<image class="cd-popup-img1" src="../../static/two/ok.png"></image>
				</view>

				<view class="cd-popup-t3" v-if="popstatus==0">成功出售后，将直接转入到购买者的账户 寄售成功将收取4%手续费。</view>
				<view class="m-flex-sb" style="width: 550rpx;">
					<view v-if="popstatus==0" class="cd-popup-v2" @click="close">取消</view>

					<view v-if="popstatus==0" class="cd-popup-v2" style="background: #01C5A1;" @click="putshelves">确认上架
					</view>
					<view v-if="popstatus==1" style="min-width: none;10rpx"></view>
					<view v-if="popstatus==1" class="cd-popup-v2" style="background: #01C5A1;" @click="close">确定</view>

				</view>
			</view>
		</u-popup>
		
		<mergePopup 
			ref="mergePopup"
			@next="goMerge"
		/>

	</view>
</template>

<script>
	import {
		marketdetail,
		nftbuy,
		walletcreate,
		nfton,
		nftoff
	} from '../../api/index.js';
	import mergePopup from './modules/mergePopup.vue'




	// import wx from 'weixin-js-sdk'
	export default {
		components: {
			mergePopup
		},
		data() {
			return {
				poptitle: "寄售",
				show: false,
				popstatus: 0,
				value: "",
				id: null,
				userInfo: {},
				height: '100vw',
				ntfstatus: -1,
				logs: [],
				Info: {


					author_avatar: '',
					author_desc: '',
					author_name: '',
					chain_status: '',
					circulation: '',
					stock: '',
					hash: '',
					id: '',
					image: '',
					work_no: '',
					name: '',
					owner_avatar: '',
					owner_name: '',
					price: 0,
					owner_id: 0,

				},

				tstatus: {
					"title": ' ',
					"color": '#7B61FF',
					"btnTxt": '',
					"btnV": -1
				},

				content: `	1. 数字藏品为虚拟数字商品，并非实物，仅限年满18周岁的用户购买或收藏。<br/><br/>
					2. 数字藏品一经售出，不支持退换，也不提供源文件下载。如不能接受，请勿购买或收藏。<br/><br/>
					3. 请勿对数字藏品进行任何形式的炒作、场外交易、欺诈，或以任何其他非法方式进行使用。<br/><br/>
					4. 数字藏品的版权由创作者或发行方拥有，除另行取得版权拥有者书面同意外，用户不得将数字藏品用于任何商业用途。<br/>
									
								`,




			};
		},
		computed: {

		},
		watch: {
			ntfstatus: {
				immediate: true,
				handler (v) {
					console.log(this.Info.owner_id, this.userInfo.id)
					switch (v) {
						case 0: {
							if (this.Info.owner_id == this.userInfo.id) {
								this.tstatus.title = "在我的仓库中";
								this.tstatus.color = "#01C5A1";
								this.tstatus.btnTxt = '上架';
								this.tstatus.btnV = 0;
								return false
							};
					
							// this.tstatus.btnTxt = 'xxxxx';
							break;
						}
					
						case 1: {
							if (this.Info.owner_id == this.userInfo.id) {
								this.tstatus.title = '售价￥' + this.Info.price;
								this.tstatus.color = "#FF225F";
								this.tstatus.btnTxt = "下架"
								this.tstatus.btnV = 3;
					
							};
							if (this.Info.owner_id != this.userInfo.id) {
								this.tstatus.title = "寄售中";
								this.tstatus.color = "#7B61FF";
								this.tstatus.btnTxt = '￥' + this.Info.price;
								this.tstatus.btnV = 1;
								return false
							};
							break;
						}
						case 2: {
							this.tstatus.title = "商品锁定中";
							this.tstatus.color = "#fff";
					
					
							if (this.Info.pay_url === '') {
								this.tstatus.btnTxt = '正在支付中';
								this.tstatus.btnV = 2;
							} else {
								this.tstatus.title = '售价￥' + this.Info.price;
								this.tstatus.btnTxt = '继续支付';
								this.tstatus.btnV = 5;
							}
					
							break;
						}
					
					};
				}
			}
		},
		onLoad(options) {
			const {
				id,

			} = options;

			this.id = id;
			// this.id = 19205;
			let dt = {
				"id": this.id
			}
		},
		onShow() {
			this.userInfo = this.$storea.state.userInfo;
			console.log(this.userInfo)
			this.getDb({
				id: this.id
			})
		},
		methods: {
			circulation() {
				// uni.$u.route({
				// 	url: '/pages/market/market',
				// 	type: "switchTab"
				// })

				uni.$u.route('/pages/circulation/circulation?name=' + this.Info.name+"&wid="+this.Info.wid);

			},

			close() {
				this.show = false
				// console.log('close');
			},

			putshelves() {
				if (this.value == "" || parseInt(this.value) < 1 || !this.value) {
					uni.showToast({
						title: '请输入销售价格',
						duration: 1000,
						icon: "none"
					});
					return false
				};
				let db = {
					"id": parseInt(this.id),
					"price": parseFloat(this.value)
				}

				nfton(db).then(res => {
					console.log(res + 'AAAAAAAAAAAAAA上架');
					// if(res.code==0){
					console.log()
					// uni.showToast({
					// 	title: '上架成功',
					// 	duration: 1000,
					// 	icon: "none"
					// });
					this.popstatus = 1;
					let dt = {
						"id": this.id
					}

					this.getDb(dt);
					// this.close();

					// }
				})
			},

			submit() {
				if (this.tstatus.btnV == 1) {
					if (this.userInfo.is_wallet == 0) {
						walletcreate().then(res => {
							plus.runtime.openURL(res);
							// uni.showModal({
							// 	title: '提示',
							// 	content: '确认前往官网注册成功后再重新登录”',
							// 	success: function(res) {
							// 		if (res.confirm) {
							// 			plus.runtime.openURL(res);
							// 		} else if (res.cancel) {
							// 			return false
							// 		}
							// 	}
							// });

							// let url="pages/webview/web?url="+res;
							// console.log(url);


						})
						return
					};

					let db = {};
					db['id'] = parseInt(this.Info.id);
					nftbuy(db).then(res => {
						console.log(res);
						if (res) {
							this.ntfstatus = 2;
							// let url = "pages/webview/web?url=" + res;
							// uni.$u.route(url);
							
							const pay_url = res;
							plus.runtime.openURL(pay_url, () => {
								// plus.runtime.openWeb(pay_url, () => {
								// 调用失败 callback
								uni.$u.toast('调起支付失败，请稍后再试！');
								
							});

						}
					})



				};
				// 上架
				if (this.tstatus.btnV == 0) {

					this.show = true;
				};
				if (this.tstatus.btnV == 2) {
					return false

				};
				// 继续支付
				if (this.tstatus.btnV == 5) {

					plus.runtime.openWeb(this.Info.pay_url, () => {
						// 调用失败 callback
						uni.$u.toast('调起支付失败，请稍后再试！');
					});

					return false

				};
				// 下架
				if (this.tstatus.btnV == 3) {
					let db = {
						"id": parseInt(this.id)
					}
					nftoff(db).then(res => {
						console.log(res + 'AAAAAAAAAAAAAA下架');
						// if(res.code==0){

						uni.showToast({
							title: '操作成功',
							duration: 2000,
							icon: "none"
						});
						this.popstatus = 0;
						let dt = {
							"id": this.id
						}

						this.getDb(dt);
						// this.close();

						// }
					})
					return false

				};





			},

			imageLoad() {
				this.height = 'auto';
			},

			getDb(db) {
				marketdetail(db).then(res => {
					console.log(res);
					this.logs = res.logs;
					if (res) {

						const {
							author_avatar,
							author_desc,
							author_name,
							chain_status,
							circulation,
							hash,
							id,
							image,
							work_no,
							stock,
							name,
							owner_id,
							owner_avatar,
							owner_name,
							price,
							pay_url
						} = res;
						this.Info = res;

						this.ntfstatus = res.status;

						// console.log('aaaaaaaaaaaaaaaa' + this.ntfstatus);


					}
				});
			},

			goBack() {
				uni.navigateBack()
			},
			merge () {
				this.$refs.mergePopup.open ()
			},
			goMerge (e) {
				uni.$u.route('pages/CollectionDetails/merge', {
					id: e,
					collectionId: this.id
				})
			}
		}
	};
</script>

<style scoped>
	.u-nav-slot {
		display: flex;
	}

	.cd-popup {

		padding: 32rpx;

		background: #21005D;
	}

	.cd-popup-v1 {
		border: 2rpx solid #4F378B;
		border-radius: 66rpx;
		margin: 32rpx 16rpx;
		/* width: 95%; */
		height: 100rpx;
		padding: 0 8rpx;
	}

	.cd-popup-v2 {
		width: 200rpx;
		height: 80rpx;
		line-height: 80rpx;
		text-align: center;

		/* Token-Brand-Colours/03 */
		background: #37419C;
		border-radius: 100rpx;
	}

	.cd-popup-t1 {
		font-size: 45rpx;
	}

	.cd-popup-t2 {
		font-size: 27rpx;
		margin: 15rpx 0 0 8rpx;
		width: 400rpx;
		padding-top: 10rpx;

	}

	.cd-popup-t3 {
		font-size: 27rpx;
		margin: 50rpx 8rpx;
		width: 500rpx;
		color: #D0BCFF;


	}

	.cd-popup-t4 {
		font-size: 27rpx;
		margin: 50rpx 8rpx;

		color: #D0BCFF;


	}

	.cd-popup-img1 {
		width: 66rpx;
		height: 66rpx;
		margin: 30rpx 0 32rpx 16rpx;

	}


	.cd-v3 {
		/* width: 99%; */
		padding: 24rpx;
		background: #381E72;
	}

	.cd-v1 {
		/* width: 99%; */

		padding: 32rpx 24rpx 16rpx 24rpx;

		background: #180240;
		border-radius: 0rpx 0rpx 12rpx 12rpx;
	}

	.cd-v1-t1 {
		font-weight: 400;
		font-size: 56rpx;
	}

	.cd-v1-t2 {

		font-size: 28rpx;
		padding: 8rpx 24rpx 8rpx 24rpx;
		background-color: #000;
	}

	.cd-v1-d1 {
		padding: 12rpx 15rpx;
		gap: 10px;
		margin-right: 50rpx;

		width: 120rpx;
		height: 39rpx;
		line-height: 39rpx;
		text-align: center;
		background: #21242D;
		border-radius: 20rpx;
		font-size: 26rpx;

		color: #7B61FF;


	}

	.cd-v1-d2 {
		padding: 2rpx 6rpx 2rpx 6rpx;
		height: 36rpx;
		font-size: 20rpx;
		color: black;
		background: #7B61FF;
	}

	.cd-title {
		font-weight: 400;
		font-size: 36rpx;
		line-height: 36rpx;
		margin: 24rpx;
	}

	.cd-v2 {
		padding: 8rpx 24rpx;
		/* width: 99%; */
		margin-bottom: 16rpx;
	}

	.cd-t1 {
		color: #D0BCFF;
		padding: 8rpx 24rpx;
		font-size: 30rpx;

	}

	.cd-t2 {
		color: #D0BCFF;
		font-size: 28rpx;

	}

	.cd-t3 {
		color: #D0BCFF;
		font-size: 29rpx;

	}

	.cd-t4 {
		margin: 32rpx 0rpx;
		font-size: 32rpx;
		padding-right: 32rpx;
	}

	.cd-bottom {
		position: fixed;
		right: 0rpx;
		left: 0;
		right: 0;
		bottom: 0rpx;
		padding: 16rpx;
		z-index: 9;
		font-size: 36rpx;
		color: white;
		font-weight: 500;
		/* width: 99%; */
		height: 168rpx;
		background: #000;


	}

	.cd-bottom-v1 {
		margin-left: 68rpx;
		color: #7E7695;
	}

	.cd-bottom-v2 {
		margin-right: 32rpx;
		padding: 16px;
		gap: 10px;
		width: 180rpx;
		height: 33rpx;
		text-align: center;
		line-height: 33rpx;
		background: #21242D;
		border-radius: 40rpx;
		font-size: 30rpx;
	}
</style>
