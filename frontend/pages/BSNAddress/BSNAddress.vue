<template>
	<view>
		<u-navbar :titleStyle="{'color': '#fff'}" bgColor="#21005D" title="区块链地址" :placeholder="true" class="nav_bar"
			:autoBack="true" leftIconColor="#151517">
			<u-icon name="arrow-leftward" color="#fff" slot="left"></u-icon>
		</u-navbar>
		<u-gap height="20"></u-gap>
		<view class="bsn_container">
			<view class="card_container">
				<view class="card_header">
					<view class="avatar_container">
						<u-avatar size="60" :src="accountInfo.avatar" class="avatar"></u-avatar>
					</view>
					<u-gap height="30"></u-gap>
					<u-text :text="`${accountInfo.nickname}的区块链地址`" align="center" color="#ffffff"></u-text>

				</view>
				<!-- @touchstart="barcodeTouchStartHandler" -->
				<view class="card_body" @longtap="barcodeTouchEndHandler">
					<u-gap height="25"></u-gap>
					<u-text text="扫一扫下方的二维码图案,获取区块链地址" size="14" align="center" color="#756C6C"></u-text>
					<u-gap height="25"></u-gap>
					<u-image :src="bsnQrcode" width="200" height="200">
						<view slot="error" style="font-size: 24rpx;">加载失败</view>
					</u-image>
					<u-gap height="15"></u-gap>
					<u-text text="(长按二维码保存至相册)" size="14" align="center" color="#756C6C"></u-text>
					<u-gap height="15"></u-gap>
				</view>

				<view class="card_footer">
					<view>
						<view class="card_footer_title">
							<u-text text="区块链地址" color="#E8DEF8" size="15"></u-text>
							<u-icon name="../../static/popup_icon/copy.png" color="#cfcfcf" @click="copyToClipboard">
							</u-icon>
						</view>
						<u-gap height="10"></u-gap>
						<u-text :text="accountInfo.bsn_address" size="11" color="#D0BCFF"></u-text>
					</view>
				</view>
			</view>
			<u-gap height="15"></u-gap>
			<view style="text-indent: 2em;">
				<u-text text="BSN是由国家信息中心牵头，会同中国移动、中国银联等单位联合发起并建立，是基于区块链技术的全球性基础设施网络，是可信、可控、可扩展的联盟链。" size="12"
					color="#FFFFFF7D"></u-text>
				<u-text text="BSN地址是您在BSN区块链服务网络上唯一的身份证明也是个人数字资产的存储地址，一旦遗失无人可寻回。" size="12" color="#FFFFFF7D"></u-text>
			</view>






		</view>
	</view>
</template>

<script>
	import {
		bsnAddressReq
	} from '../../api/index.js';
	export default {
		data() {
			return {
				accountInfo: {
					avatar: '',
					nickname: '',
					bsn_address: ''
				},
				startTime: 0,
				bsnQrcode: ''
			};
		},

		onShow() {
			// let accountInfo = uni.getStorageSync('account');
			let accountInfo = this.$storea.state.userInfo;
			// accountInfo = JSON.parse(accountInfo);
			const {
				avatar,
				nickname,
				
				bsn_address
			} = accountInfo;
			this.accountInfo = {
				avatar,
				nickname,
				bsn_address
			};

			this.getQrCode();
		},
		methods: {
			getQrCode() {
				console.log(11);
				bsnAddressReq().then(res => {
					console.log(res);
					this.bsnQrcode = 'data:image/png;base64,' + res
				})
			},
			barcodeTouchStartHandler() {
				this.startTime = Date.now();
			},
			barcodeTouchEndHandler() {
				const base64 = this.bsnQrcode;
				const bitmap = new plus.nativeObj.Bitmap("test");
				bitmap.loadBase64Data(base64, function() {
					const url = "_doc/" + new Date().getTime() + ".png"; // url为时间戳命名方式
					console.log('saveHeadImgFile', url)
					bitmap.save(url, {
						// overwrite: true, // 是否覆盖
						// quality: 'quality'  // 图片清晰度
					}, (i) => {
						uni.saveImageToPhotosAlbum({
							filePath: url,
							success: function() {
								uni.$u.toast('保存成功！');
								bitmap.clear()
							}
						});
					}, (e) => {
						uni.$u.toast('保存失败！');
						bitmap.clear()
					});
				}, (e) => {
					uni.$u.toast('保存失败！');
					bitmap.clear()
				});
			},
			// barcodeTouchEndHandler() {
			// 	// const endTime = Date.now();
			// 	// if (endTime - this.startTime < 2000) {
			// 	// 	return;
			// 	// }

			// 	uni.saveImageToPhotosAlbum({
			// 		filePath: this.bsnQrcode,//'https://cdn.uviewui.com/uview/goods/1.jpg',
			// 		success: res => {
			// 			uni.$u.toast('保存成功！');
			// 		},
			// 		fail: (err) => {
			// 			console.log(err)
			// 		}
			// 	});
			// },
			copyToClipboard() {
				const {
					bsn_address
				} = this.accountInfo;
				uni.setClipboardData({
					data: bsn_address,
					success: () => {
						uni.$u.toast('复制成功！');
					}
				});
			}
		}
	};
</script>

<style>
	.bsn_container {
		padding: 48rpx 48rpx calc(env(safe-area-inset-bottom) / 2 + 24rpx);
	}

	.card_container {
		border-radius: 8px;
		/* box-shadow: 0 0 3px #efefef; */
	}

	.card_header {
		position: relative;
		padding: 20rpx 40rpx;
		border: 1px solid #4952CC;
		background-color: #4952CC;
		border-top-left-radius: 8px;
		border-top-right-radius: 8px;
	}


	.avatar_container {
		position: absolute;
		width: 60px;
		height: 60px;
		left: 0;
		top: -60px;
		right: 0;
		bottom: 0;
		margin: auto;
		border: 3px solid #4952CC;
		border-radius: 50%;
		display: flex;

	}

	.card_body {

		display: flex;
		flex-direction: column;
		align-items: center;
		justify-content: center;
		background-color: #ffffff;
		border: 1px solid #efefef;
		/* 		box-shadow: 0 0 3px #efefef;
 */
		border-bottom-width: 0;
	}

	.card_footer {

		border-bottom-left-radius: 8px;
		border-bottom-right-radius: 8px;
		padding: 10px;
		background-color: #180240;
		border: 1px solid #180240;
		border-top-width: 0;
	}

	.card_footer>view {
		background-color: #180240;
		border-color: 1px solid #180240;

		border-radius: 8px;
		padding: 10px 20px;
	}

	.card_footer_title {
		display: flex;
		justify-content: space-between;
	}
</style>
