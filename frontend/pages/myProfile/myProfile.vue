<template>
	<view>
		<u-navbar :titleStyle="{'color': '#fff'}" bgColor="#21005D" title="设置" :placeholder="true" class="nav_bar"
			:autoBack="true" leftIconColor="#151517">
			<u-icon name="arrow-leftward" color="#fff" slot="left"></u-icon>
		</u-navbar>

		<view class="profile_container person_container">
			<navigator class="avatar" url="/pages/personalEdit/personalEdit">
				<u-avatar size="100" :src="accountInfo.avatar"></u-avatar>
				<view class="pen">
					<u-icon name="/static/profile_icon/pen.png"></u-icon>
				</view>
			</navigator>
			<u-gap height="20"></u-gap>
			<u-text class="author_name" :text="accountInfo.nickname" size="18" color="#fff" align="center"></u-text>
			<view class="inv_container" @click="copyUid">
				<u-text :text="`邀请码：${accountInfo.uid}`" size="12" color="#B69DF8" align="center"></u-text>
				<u-icon class="icon" name="../../static/popup_icon/copy.png"></u-icon>
			</view>
			<u-gap height="38"></u-gap>
		</view>

		<view class="profile_container">
			<u-cell-group>
				<u-cell v-for="cell in celllist" :stop="true" :key="cell.title" :name="cell.url" :title="cell.title"
					@click="cellClickHandler">
					<u-icon slot="value" :name="cell.icon"></u-icon>

				</u-cell>
				<u-cell title="联系客服" @click="toggleService"></u-cell>

				<u-cell title="区块链信息查询" @click="openBrower"></u-cell>

				<u-cell title="关于我们" url="/pages/aboutUs/aboutUs">
					<!-- <view slot="value" style="position: relative;">
						<u-text class="about_btn" text="新版本" align="right" size="10" color="#fff">
						</u-text>
						<u-badge :isDot="true" type="error" :absolute="true" :offset="[-2,-2]"></u-badge>
					</view> -->
				</u-cell>
				<u-cell title="邀请好友" @click="copyLink"></u-cell>
			</u-cell-group>
			<u-gap height="114"></u-gap>

			<u-text @click="logout" class="log_out" text="退出登录" size="16" color="#ffffff" align="center"></u-text>
		</view>

		<!-- service -->
		<u-popup :safeAreaInsetBottom="false" round="20" :show="serviceShow" mode="center" :closeOnClickOverlay="true"
			@close="toggleService">
			<view style="background-color: #21005D;padding: 20rpx 15rpx;
    display: flex;
    flex-direction: column;
    align-items: center;border-radius: 16px;" @longtap="saveImage">
				<text style="font-weight: 600;font-size: 20px;line-height: 28px;color: #FFFFFF;">北海客服</text>

				<u-gap height="20"></u-gap>

				<u-image src="/static/qrcode.png" width="480rpx" height="480rpx"></u-image>

				<u-gap height="20"></u-gap>
				<text
					style="font-weight: 400;font-size: 12px;line-height: 16px;color: #D0BCFF;">使用微信扫码，也可以长按图片保存二维码到手机后添加客服微信。</text>

				<u-gap height="30"></u-gap>

				<u-button color="#37419C" shape="circle" text="关闭" @click="toggleService"
					:customStyle="{'width':'50%'}"></u-button>
			</view>
		</u-popup>

	</view>
</template>

<script>
	import { accountInfoReq } from '../../api/index.js';


	import { INVURL } from '../../utils/constant.js';
	export default {
		data() {
			return {
				accountInfo: {
					nickname: '',
					avatar: '',
					uid: ''
				},
				celllist: [{
						title: '手机号',
						icon: '/static/profile_icon/user.png',
						url: '/pages/contactInfo/contactInfo'
					},
					{
						title: '区块链地址',
						icon: '/static/profile_icon/lock.png',
						url: '/pages/BSNAddress/BSNAddress'
					},
					{
						title: '实名认证',
						icon: '/static/profile_icon/bell.png',
						url: '/pages/realNameAuth/realNameAuth',
					},
					// {
					// 	title: '银行卡管理',
					// 	icon: '/static/profile_icon/setting.png',
					// 	url: '/pages/cardManage/cardManage'
					// },
					// {
					// 	title: '我的钱包',
					// 	icon: '/static/profile_icon/wallet.png',
					// 	url: '/pages/myAccount/myAccount'
					// },
					// {
					// 	title: '关于我们',
					// 	icon: '',
					// 	url: '/pages/aboutUs/aboutUs'
					// },

				],
				serviceShow: false
			};
		},
		onLoad() {

		},
		onShow() {
			this.getAccountInfo();
		},

		methods: {
			toggleService() {
				this.serviceShow = !this.serviceShow;
			},
			openBrower() {
				// 区块查询 网址 	
				const url = 'https://tianzhou.wenchang.bianjie.ai/#/home';
				plus.runtime.openWeb(url, () => {
					// 调用失败 callback
					uni.$u.toast('调起区块链查询失败，请稍后再试！');
				});
			},
			saveImage() {
				let url = require('../../static/qrcode.png');

				uni.saveImageToPhotosAlbum({
					filePath: url, //传入临时地址
					success: () => {
						uni.$u.toast('保存成功'); //封装的提示
						this.serviceShow = false;
					},
					fail: () => {
						uni.$u.toast('保存失败');
					}
				});

				return;
				// TODO 这个url是下载接口
				uni.downloadFile({ //获得二维码的临时地址
					url,
					header: { token: uni.getStorageSync('token') },
					success: res => {
						//console.log('获取url',res)
						if (res.statusCode == 200) {
							uni.saveImageToPhotosAlbum({
								filePath: res.tempFilePath, //传入临时地址
								success: () => {
									uni.$u.toast('保存成功'); //封装的提示
								},
								fail: () => {
									uni.$u.toast('保存失败');
								}
							});
						}
					},
					fail: err => {
						console.log(err);
					}
				});
			},
			cellClickHandler(e) {
				// id_card_num
				const { accountInfo } = this;
				const { authentication } = accountInfo;
				const url = e.name;
				if (url === '/pages/realNameAuth/realNameAuth' && authentication) {
					uni.$u.toast('您已通过实名认证！');
					return;
				}

				uni.navigateTo({ url });

			},
			copyLink() {
				const { uid } = this.accountInfo;
				uni.setClipboardData({
					data: INVURL + '#/pages/login/login?code=' + uid,
					success: () => {
						uni.$u.toast('复制邀请链接成功！');
					}
				});
			},
			getAccountInfo() {
				accountInfoReq().then(res => {
					const {
						avatar,
						nickname,
						bsn_address,
						phone,
						name,
						id_card_num,
						id_card_negative_image_url,
						id_card_positive_image_url,
						amount,
						uid,
						description,
						authentication
					} = res;
					const accountInfo = {
						avatar,
						nickname,
						bsn_address,
						phone,
						name,
						id_card_num,
						id_card_negative_image_url,
						id_card_positive_image_url,
						amount,
						uid,
						description,
						authentication
					};

					uni.setStorageSync('account', JSON.stringify(accountInfo));
					this.accountInfo = accountInfo;

					// avatar: ""
					// bsn_address: "0x6805aa98e54229"
					// create_time: 1660014081
					// del_time: 0
					// id: 2
					// id_card_negative_image_url: ""
					// id_card_num: ""
					// id_card_positive_image_url: ""
					// name: ""
					// nickname: ""
					// phone: "13800000000"
					// pid: 0
					// role: 0
					// status: 0
					// uid: "sTPwgnjS"
					// update_time: 0
				});
			},
			logout() {
				try {
					uni.clearStorageSync();
					uni.reLaunch({ url: '/pages/login/login' });
				} catch (e) {
					// error
				}

			},
			copyUid() {
				const { uid } = this.accountInfo;

				uni.setClipboardData({
					data: uid,
					success: () => {
						uni.$u.toast('复制成功！');
					}
				});
			}
		}
	};
</script>

<style scoped>
	.nav_bar>>>.u-navbar__content__right {
		right: 48rpx;
		top: 0px;
		padding: 0 !important;
	}

	.u-nav-slot {
		display: flex;
		align-items: center;
		justify-content: center;
		background: #FF7438;
		border: 1px solid #21005D;
		border-radius: 50%;
		padding: 0 !important;
		width: 32px;
		height: 32px;
	}

	.profile_container {
		padding: 48rpx 48rpx calc(env(safe-area-inset-bottom) / 2 + 24rpx);
	}

	.person_container {
		display: flex;
		flex-direction: column;
		align-items: center;
	}

	.profile_container .author_name {
		font-weight: 500;
		line-height: 26px;
		letter-spacing: 0.005em;
	}

	>>>.u-cell {
		background: #180240;
		border-radius: 8px;
		padding: 0 56rpx;
		margin-bottom: 36rpx;
	}

	>>>.u-cell-group .u-cell__body {
		padding: 16px 0;
	}

	>>>.u-cell-group .u-cell__body .u-icon__img {
		height: 20px !important;
		margin-right: 36rpx;
	}

	/* >>>.u-cell__right-icon-wrap {
		display: none;
	} */

	.avatar {
		position: relative;
	}

	.pen {
		position: absolute;
		width: 32px;
		height: 32px;
		border: 1px solid #efefef;
		border-radius: 50%;
		background-color: #01C5A1;
		display: flex;
		justify-content: center;
		align-items: center;
		right: 0;
		bottom: 0;
	}

	>>>.u-line {
		display: none;
	}

	.profile_container .log_out {
		font-weight: 500;
		line-height: 24px;
		letter-spacing: 0.005em;
	}

	.inv_container {
		display: flex;
	}

	.inv_container .icon {
		margin-left: 16px;
	}

	>>>.u-cell__title-text {
		color: #fff;
	}

	.about_btn {
		background-color: #1B2374;
		flex: 0 0 auto;
		width: auto;
		padding: 10rpx 20rpx;
		border-radius: 8px;
	}
</style>
