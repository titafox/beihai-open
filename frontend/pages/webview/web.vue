<template>
	<view class="profile_container">
		<u-navbar :titleStyle="{'color': '#fff'}" bgColor="#21005D" title="北海数字藏品" :placeholder="true" class="nav_bar"
			 leftIconColor="#151517">
			<u-icon name="close" color="#fff" slot="left" @click="close"></u-icon>
		</u-navbar>

		<web-view :webview-styles="webviewStyles" :src="weburl"></web-view>

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
				weburl: "",
				webviewStyles: {
					progress: {
						color: '#FF3333'
					}
				}
			};
		},

		onLoad(options) {
			console.log(options);
			this.weburl = options.url;



		},
		onShow(options) {

			// this.weburl = options.url;
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
           close(){
			   uni.navigateBack({
			   	delta: 1
			   });
		   }


		}
	};
</script>

<style scoped>

</style>
