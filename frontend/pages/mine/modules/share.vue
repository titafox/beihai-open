<template>
	<u-popup 
		:safeAreaInsetBottom="false" 
		:round="20" 
		:show="shareShow" 
		mode="center" 
		@close="toggleShare"
		:closeOnClickOverlay="true"
	>
		<view class="container" @longtap="saveImage">
			<u-image src="/static/qrcode2.png" width="294" height="449"></u-image>
		</view>
	</u-popup>
</template>

<script>
	export default {
		name: 'share',
		data () {
			return {
				shareShow: false
			}
		},
		methods: {
			toggleShare() {
				this.shareShow = !this.shareShow;
			},
			saveImage() {
				let url = require('static/qrcode2.png');
			
				uni.saveImageToPhotosAlbum({
					filePath: url, //传入临时地址
					success: () => {
						uni.$u.toast('保存成功'); //封装的提示
						this.shareShow = false;
					},
					fail: () => {
						uni.$u.toast('保存失败');
					}
				});
			
				return;
				
				// TODO 这个url是下载接口
				
				uni.downloadFile({ //获得二维码的临时地址
					url,
					header: {
						token: uni.getStorageSync('token')
					},
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

<style lang="less" scoped>
	.container {
		background-color: #21005D;
		padding: 36rpx 16rpx 20rpx;
		width:100%;
		display: flex;
		flex-direction: column;
		align-items: center;
		border-radius: 16px;
	}
</style>