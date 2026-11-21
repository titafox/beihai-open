<template>
	<view class="profile_container">
		<u-navbar 
			:titleStyle="{'color': '#fff'}" 
			bgColor="#21005D" 
			title="我的仓库" 
			:placeholder="true" 
			class="nav_bar"
			:autoBack="true" 
			leftIconColor="#151517"
		>
			<u-icon name="arrow-leftward" color="#fff" slot="left"></u-icon>
		</u-navbar>
		<u-gap height="20"></u-gap>
		<view class="list">
			<view class="item" v-for="(item, index) in dblist" :key="index">
				<item :data="item" />
			</view>
		</view>
		<!-- <Purchased :list="dblist"></Purchased> -->

		<u-empty v-if="dblist.length<1" mode="car" icon="http://cdn.uviewui.com/uview/empty/car.png" text="没有藏品..." />
		
		<u-gap height="20"></u-gap>
	</view>
</template>

<script>
	// import ProfileHeader from '../../components/ProfileHeader/ProfileHeader.vue';

	import Purchased from '../../components/Purchased/Purchased.vue'
	import item from './modules/item.vue'

	import {
		nfts,
	} from '../../api/index.js';
	export default {
		components: {
			Purchased,
			item
		},
		data() {
			return {
				dblist: [],
				// Parameters: {
				// 	page: 1,

				// },
			};
		},

		onLoad(options) {
			this.getDb();
		},
		onShow(options) {

		},
		async onPullDownRefresh () {
			await this.getDb()
			uni.stopPullDownRefresh()
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
			async getDb() {
				const res = await nfts()
				if (res) {
					this.dblist=res;
				}
			}
		}
	};
</script>

<style scoped lang="less">
	.list {
		display: flex;
		flex-wrap: wrap;
		padding: 0 24rpx;
		
		.item {
			width: 50%;
			padding: 0 12rpx 24rpx;
			box-sizing: border-box;
		}
	}
</style>
