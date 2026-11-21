<template>
	<view class="custom_container">
		<u-navbar 
			:titleStyle="{'color': '#fff'}" 
			bgColor="#21005D" 
			title="资料编辑" 
			:placeholder="true" 
			class="nav_bar"
			:autoBack="true" 
			leftIconColor="#151517"
		>
			<u-icon name="arrow-leftward" color="#fff" slot="left"></u-icon>
		</u-navbar>
		
		<view>
			<template
				v-for="(item, index) in menulist"
			>
				<view
					v-if='item.type === 1'
					class="m-flex-sb nav_b" 
					@click="doClick(item.click)"
				>
					<text style="width: 566rpx;">{{item.title}}</text>
					<image style="width: 36rpx;height: 36rpx;margin-right: 10rpx;" :src="item.image"></image>
					<u-icon name="arrow-right" color="#fff" size="18"></u-icon>
				</view>
				<view
					v-else
					class="m-flex-sb nav_b"
					@click="gotoPage(item)"
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
						"title": "个人信息",
						"image": '',
						"page": "pages/personalEdit/personalEdit"
					},
					{
						"id": 2,
						"type": 0,
						"title": "手机号",
						"page": "/pages/contactInfo/contactInfo"
					},
					{
						"id": 3,
						"type": 0,
						"title": "实名认证",
						"page": "/pages/realNameAuth/realNameAuth",
						"isclick":false,
						
					}
				],
				profileInfo: {},
			}
		},
		onShow(options) {
			this.getUserInfo();
			
			
		},
		methods: {
			getUserInfo() {
				this.profileInfo = this.$storea.state.userInfo;
				if(this.profileInfo.authentication==1){
				 this.menulist[2].title="已实名";	
				 this.menulist[2].isclick=true;
				};
				this.menulist[0].image = this.profileInfo.avatar
				console.log(this.profileInfo);
			
			},
			gotoPage(im) {
				if(im.isclick){
					return false;
				};
				uni.$u.route(im.page);
			},
			doClick (method) {
				this[method]()
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
