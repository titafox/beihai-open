<template>
	<view class="profile_container">
		<u-navbar :titleStyle="{'color': '#fff'}" bgColor="#21005D" title="已购藏品" :placeholder="true" class="nav_bar"
			:autoBack="true" leftIconColor="#151517">
			<u-icon name="arrow-leftward" color="#fff" slot="left"></u-icon>
		</u-navbar>


		
	</view>
</template>

<script>
	// import ProfileHeader from '../../components/ProfileHeader/ProfileHeader.vue';
	
	

	import {
		nfts,
	} from '../../api/index.js';
	export default {
		components: {
			// ProfileHeader,

			
		},
		data() {
			return {
				
			};
		},

		onLoad(options) {

			



		},
		onShow(options) {

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
			


		}
	};
</script>

<style scoped>
	
</style>
