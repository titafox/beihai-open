<template>
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
</template>

<script>
	export default {
		name: 'customer',
		data () {
			return {
				serviceShow: false
			}
		},
		methods: {
			toggleService() {
				this.serviceShow = !this.serviceShow;
			},
			saveImage() {
				let url = require('../../../static/qrcode.png');
			
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
		}
	}
</script>

<style>
</style>