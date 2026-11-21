<template>
	<view class="custom_container">
		<u-navbar 
			:titleStyle="{'color': '#fff'}" 
			bgColor="#21005D" 
			title="区块信息" 
			:placeholder="true" 
			class="nav_bar"
			:autoBack="true" 
			leftIconColor="#151517"
		>
			<u-icon name="arrow-leftward" color="#fff" slot="left"></u-icon>
		</u-navbar>
		
		<view>
			<template v-for="(item, index) in menulist">
				<view
					v-if='item.type === 1'
					class="m-flex-sb nav_b" 
					@click="doClick(item.click)"
				>
					<text style="flex-grow: 1">{{item.title}}</text>
					<image style="width: 36rpx;height: 36rpx;margin-right: 10rpx;" :src="item.image"></image>
					<u-icon name="arrow-right" color="#fff" size="18"></u-icon>
				</view>
				<view
					v-else
					class="m-flex-sb nav_b"
					@click="gotoPage(item.page)"
				>
					<text style="flex-grow: 1">{{item.title}}</text>
					<image style="width: 36rpx;height: 36rpx;margin-right: 10rpx;" :src="item.image"></image>
					<u-icon name="arrow-right" color="#fff" size="18"></u-icon>
				</view>
			</template>
			
		</view>
	</view>
</template>

<script>
	export default {
		data() {
			return {
				menulist: [{
						"id": 1,
						"type": 0,
						"title": "我的区块链地址",
						"image": "../../static/profile_icon/qrcode.png",
						"page": "pages/BSNAddress/BSNAddress"
					},
					{
						"id": 2,
						"type": 1,
						"title": "区块链信息查询",
						"click": "openBrower"
					}
				]
			}
		},
		methods: {
			doClick (method) {
				this[method]()
			},
			gotoPage(page) {
				uni.$u.route(page);
			},
			openBrower() {
				// 区块查询 网址 	
				const url = 'https://tianzhou.wenchang.bianjie.ai/#/home';
				plus.runtime.openWeb(url, () => {
					// 调用失败 callback
					uni.$u.toast('调起区块链查询失败，请稍后再试！');
				});
			},
		}
	}
</script>

<style scoped>
.custom_container {
	padding: 48rpx 28rpx calc(env(safe-area-inset-bottom) / 2 + 24rpx);
}

.nav_b {
	display: flex;
	margin: 8rpx 16rpx 22rpx 16rpx;
	padding: 24rpx 32rpx 24rpx 32rpx;
	align-items: center;
	background: #180240;
	border-radius: 8rpx;
}
</style>
