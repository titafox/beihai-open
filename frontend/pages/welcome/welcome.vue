<template>
	<view @click="linkToLogin" >
		
		
		<!-- <view class="welcome_container">
			<u-image src="/static/welcome/welcome.png" width="100%" height="100%" :fade="true" duration="450"
				mode="aspectFill">
			</u-image>
		</view>
		<view class="info_container">
			<u-text class="title_1" text="北海" color="#FFFFFF" size="57"></u-text>
			<u-text class="title_1" text="NORTH" color="#FFFFFF" size="57"></u-text>
			<u-gap height="48"></u-gap>
			<u-text class="title_2" text="在北海，由社区一起开启一扇全新的" color="#D0BCFF" size="14"></u-text>
			<u-text class="title_2" text="数字艺术大门。" color="#D0BCFF" size="14"></u-text>
		</view> -->
		<view class="circle_container">
			<u-gap height="40"></u-gap>
			<u-image src="/static/welcome/group8.png" height="90vh" :fade="true" duration="450"
				mode="widthFix">
			</u-image>
		</view>
		
	<!-- 	<view class="title_container">
			<u-image src="/static/welcome/title.png"  height="100vh" :fade="true" duration="450"
				mode="heightFix" >
			</u-image>
		</view> -->
	</view>
</template>

<script>
	export default {

		data() {
			return {

			}
		},
		methods: {
			linkToLogin() {
				uni.navigateTo({
					url: '/pages/login/login'
				})
			}
		}
	}
</script>

<style>
	.welcome_container>>>.u-transition {
		width: 100vw;
		height: 700rpx;
	}
	
	
	.info_container {
		padding-top: 234rpx;
		padding-left: 106rpx;
	}


	.title_1>>>span {
		font-weight: 200;
		line-height: 64px;
		letter-spacing: -0.25px;
	}



	.title_2>>>span {
		font-weight: 500;
		line-height: 20px;
		letter-spacing: 0.1px;
	}
	
	.circle_container {
		position: absolute;
		right: 0;
	}
	
	/* .circle_container >>> .u-transition {
		width: 100vw;
		height: auto;
	} */
	
	.title_container {
		position: absolute;
		top: 0;
		right: -20vw;}
		
</style>
