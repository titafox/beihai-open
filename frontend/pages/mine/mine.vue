<template>
	<view class="padding_0_24">
		<u-gap height="44"></u-gap>
		<view class="block-title_container">
			<u-row customStyle="margin-bottom: 10px">
				<u-col span="11">
					<view class="bh-layout bg-purple-light">
						<u-text text="Mine" size="21" color="#EADDFF"></u-text>
					</view>
				</u-col>
				<u-col span="1">
					<u-image src="/static/profile_icon/share.png" :lazy-load="true" width="30" height="30" align="right"
						@click="$refs.share.shareShow = true">
					</u-image>
				</u-col>
			</u-row>
		</view>


		<view style="display: flex;">
			<u-avatar size="50" :src="profileInfo.avatar" class="avatar"></u-avatar>
			<view>
				<view class="m-flex-c" @click="goPersonalEdit">
					<u-text class="person_name" :text="`${profileInfo.nickname}`" size="18" color="#ffffff">
					</u-text>
					<u-icon class="arrow_right" name="arrow-right" color="#fff" size="18"></u-icon>
				</view>

				<view class="m-flex">

					<!-- <u-text class="intr_info" :text="'邀请码: '+profileInfo.uid" size="12"></u-text> -->
					<text class="intr_txt m-color-m2">邀请码: {{profileInfo.uid}}</text>
					<image class="intr_copy" src="../../static/popup_icon/copy.png" @click="copyLink"></image>
				</view>

			</view>

			<view @click="link2mint" class="mint">
				<!-- <u-icon name="/static/mint.png" size="48"></u-icon> -->
				<image src="../../static/mint.png" style="width: 126rpx;height: 126rpx;margin-bottom: -47rpx"></image>


				<u-text text="铸件" size="12" color="#9A82DB" align="center"></u-text>

			</view>


		</view>


		<u-gap height="20"></u-gap>
		<view class="m-flex-sb">
			<view class="nav_a" @click="walleta">
				<image class="nav_a_img" src="../../static/two/Uniona.png"></image>
				<text v-if="profileInfo.is_wallet==1">我的钱包</text>
				<text v-else>开通钱包</text>
			</view>
			<view class="nav_a" @click="goBlock">
				<image class="nav_a_img" src="../../static/two/Unionb.png"></image>
				<text>区块信息</text>
			</view>
			<!-- 提交 -->

		</view>
		<u-gap height="20"></u-gap>
		<view class="m-flex-w">
			<template v-for="(item, index) in menulist">
				<view v-if='item.type === 1' class="m-flex-sb nav_b" @click="doClick(item.click)">
					<text style="width: 566rpx;">{{item.title}}</text>
					<u-icon name="arrow-right" color="#fff" size="18"></u-icon>
				</view>
				<view v-else class="m-flex-sb nav_b" @click="gotoPage(item.page)">
					<text style="width: 566rpx;">{{item.title}}</text>
					<u-icon name="arrow-right" color="#fff" size="18"></u-icon>
				</view>
			</template>

		</view>

		<u-gap height="30"></u-gap>
		<view class="log_out">
			<text @click="logout" color="#ffffff" align="center">
				切换账号
			</text>
		</view>

		<Tabbar />

		<customer ref="customer" />

		<share ref="share" />
	</view>
</template>

<script>
	// import ProfileHeader from '../../components/ProfileHeader/ProfileHeader.vue';

	import Tabbar from '../../components/Tabbar/Tabbar.vue';
	import customer from './modules/customer.vue';
	import share from './modules/share.vue';
	import {
		INVURL
	} from '../../utils/constant.js';

	import {
		accountInfoReq,
		walletcreate,
		walletinfo,


	} from '../../api/index.js';
	export default {
		components: {
			// ProfileHeader,
			Tabbar,
			customer,
			share
		},
		data() {
			return {
				menulist: [{
						"id": 1,
						"type": 0,
						"title": "我的仓库",
						"page": "/pages/nfts/nfts"
					},
					{
						"id": 2,
						"type": 0,
						"title": "我的寄售",
						"page": "/pages/handselList/handselList"
					},
					{
						"id": 3,
						"type": 0,
						"title": "订单查询",
						"page": "/pages/order/order"
					},
					{
						"id": 4,
						"type": 0,
						"title": "我的邀请",
						"page": "/pages/invitation/invitation"
					},
					{
						"id": 5,
						"type": 1,
						"title": "联系客服",
						"click": "showCustomer"
					},
					{
						"id": 6,
						"type": 0,
						"title": "关于我们",
						"page": "/pages/aboutUs/aboutUs"
					},
					// {
					// 	"id": 7,
					// 	"title": "隐私政策",
					// 	"page": ""
					// },
					// {
					// 	"id": 8,
					// 	"title": "安全中心",
					// 	"page": ""
					// },
				],


				profileInfo: {
					avatar: '',
					nickname: '',
					role: '',
					bsn_address: '',
					is_wallet: ''

				}

			};
		},

		onLoad(options) {





		},
		onShow(options) {
			uni.hideTabBar();
			this.getUserInfo();
		},
		watch: {

			// activeTab(newval) {

			// 	switch (newval) {
			// 		// case 2:
			// 		// 	this.getPublishList() // 发布
			// 		// 	break;
			// 		case 2:
			// 			uni.navigateTo({ url: '/pages/order/order' }); // 全部订单
			// 			break;
			// 		case 0:
			// 			this.getBuyList(); //藏品
			// 			break;
			// 			// default:
			// 			// 	this.getLikeList() //卖出
			// 			// 	break;
			// 		default:
			// 			uni.navigateTo({ url: '/pages/handselList/handselList' }); //转赠记录
			// 			break;
			// 	}

			// }
		},
		methods: {
			doClick(method) {
				this[method]()
			},
			walleta() {
				if (this.profileInfo.is_wallet == 1) {
					walletinfo().then(res => {
						// console.log(res);
						// let url = "pages/webview/web?url=" + res;
						// uni.$u.route(url);
						
						plus.runtime.openURL(res);
					})


				};

				if (this.profileInfo.is_wallet != 1) {


					walletcreate().then(res => {
						plus.runtime.openURL(res);
						// let url = "pages/webview/web?url=" + res;
						// uni.$u.route(url);
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

				}
			},


			getUserInfo() {
				// this.profileInfo = this.$storea.state.userInfo;
				// uni.setStorageSync('account', JSON.stringify(this.profileInfo));
				// console.log(this.$storea.state.userInfo);
				accountInfoReq().then(res => {
					// const {
					// 	avatar
					// } = res;
					// this.avatar = avatar;
					// console.log(res);
					
					const {
						avatar,
						nickname,
						id_card_num,
						bsn_address,
						uid,
						role,
						is_wallet,
						id,
						is_handsel,
						authentication,
					} = res;
					this.avatar = avatar;
					let userinfo = {
						"avatar": avatar,
						"nickname": nickname,
						"uid": uid,
						"bsn_address": bsn_address,
						"is_wallet":is_wallet,
						"id":id,
						"is_handsel":is_handsel,
						"authentication":authentication
					};
					this.$storea.commit('updataUser', userinfo);
					this.profileInfo = userinfo;
				})
			},
			gotoPage(page) {
				uni.$u.route(page);
			},
			link2mint() {
				uni.showToast({
					icon: 'none',
					title: '敬请期待！'
				})
			},
			showCustomer() {
				this.$refs.customer.serviceShow = true;
			},
			copyLink() {
				const {
					uid
				} = this.profileInfo;
				uni.setClipboardData({
					data: INVURL + '#/pages/login/login?code=' + uid,
					success: () => {
						uni.$u.toast('复制邀请链接成功！');
					}
				});
			},
			logout() {
				try {
					uni.clearStorageSync();
					uni.reLaunch({
						url: '/pages/login/login'
					});
				} catch (e) {
					// error
				}

			},
			goBlock() {
				uni.$u.route('/pages/block/block')
			},
			goPersonalEdit() {
				uni.$u.route('/pages/personal/personal')
			}
		}

	};
</script>

<style lang="less" scoped>
	.padding_0_24 {
		padding: 0 28rpx;
	}

	.nav_bar /deep/ .u-navbar__content__left {
		display: none;
	}

	.profile_container {
		padding: 48rpx 28rpx calc(env(safe-area-inset-bottom) / 2 + 24rpx);
	}

	.intr_info /deep/ span {
		text-align: center;
		margin-top: 10rpx;
		color: #888888;
	}


	.nft_card {
		margin-bottom: 24rpx;
	}

	.nft_card /deep/ u-transition {
		width: 282rpx;
		height: 336rpx;
		margin: auto;
	}

	.nav_bar /deep/ u-navbar__content__right {
		right: 48rpx;
		top: 0px;
		padding: 0 !important;
	}

	.u-nav-slot {
		display: flex;
		align-items: center;
		justify-content: center;
		/* background: #FEFEFE; */
		/* border: 1px solid #F5F5F5; */
		border-radius: 50%;
		padding: 0 !important;
		width: 46px;
		height: 46px;
	}

	.nfts_list {
		height: calc(100vh - 140px - env(safe-area-inset-bottom)) !important;
	}

	.nfts_list /deep/ .uni-scroll-view-content>view {
		display: flex;
		flex-direction: row;
		flex-wrap: wrap;
		justify-content: space-between;
		padding-top: 20rpx;
	}

	.avatar {
		padding: 16rpx;
		margin-right: 32rpx;


	}


	.mint {
		padding-left: 20rpx;
		margin-left: auto;
		width: 130rpx;




	}


	.person_name {
		padding: 20rpx 0rpx 0rpx 0rpx;
	}

	.arrow_right {
		padding: 20rpx 0rpx 0rpx 16rpx;

	}

	.intr_copy {

		padding-top: 16rpx;
		margin-left: 8rpx;

		width: 28rpx;
		height: 28rpx;
	}

	.intr_txt {

		padding-top: 12rpx;
		font-size: 24rpx;

	}

	.nav_a {
		margin: 32rpx;
		padding: 32rpx;
		display: flex;
		background: #180240;
		border-radius: 12rpx;
		align-items: center;
	}

	.nav_a_img {
		margin-right: 36rpx;
		width: 50rpx;
		height: 50rpx;
	}

	.nav_b {
		display: flex;
		margin: 8rpx 16rpx;
		padding: 24rpx 32rpx;
		align-items: center;
		background: #180240;
		border-radius: 8rpx;
	}

	.nav_slot_right {
		display: flex;
	}


	.log_out {
		display: flex;
		justify-content: center;

		text {
			display: inline-block;
			box-sizing: border-box;
			padding: 16rpx 40rpx;
			background: #180240;
			border-radius: 40rpx;
		}
	}
</style>
