<template>
	<view class="profile_container">
		<u-navbar :titleStyle="{'color': '#fff'}" bgColor="#21005D" :title="title" :placeholder="true" class="nav_bar"
			:autoBack="true" leftIconColor="#151517">
			<u-icon name="arrow-leftward" color="#fff" slot="left"></u-icon>
		</u-navbar>
		<u-gap height="20"></u-gap>
		<PurchasedBig :list="dblist"></PurchasedBig>
		<u-loadmore :status="status" nomore-text="没有更多了..." />
		<u-empty 
			v-if="dblist.length < 1" 
			mode="car" 
			icon="http://cdn.uviewui.com/uview/empty/car.png" 
			text="没有藏品..."
		/>
		<u-gap height="20"></u-gap>
	</view>
</template>

<script>
	// import ProfileHeader from '../../components/ProfileHeader/ProfileHeader.vue';

	import PurchasedBig from "../../components/PurchasedBig/PurchasedBig.vue"

	import {
		nftsseries,
	} from '../../api/index.js';
	export default {
		components: {
			PurchasedBig,


		},
		data() {
			return {
				title: '藏品',
				status: 'loadmore',
				total_page: 0,
				dblist: [],
				Parameters: {
					page: 1,
					id: 0,

				},
			};
		},

		onReachBottom() {
			if (this.Parameters.page >= this.total_page) return;
			this.status = 'loading';
			this.Parameters.page = ++this.Parameters.page;

			nftsseries(this.Parameters).then(res => {

				if (res) {

					let arr = res.list;
					this.dblist = this.dblist.concat(arr);

					if (this.Parameters.page >= this.total_page) this.status = 'nomore';
					else this.status = 'loading';



				}
			});


			// setTimeout(() => {
			// 	this.list += 10;
			// 	if (this.page >= 3) this.status = 'nomore';
			// 	else this.status = 'loading';
			// }, 2000)
		},

		onLoad(options) {
			this.Parameters.id = options.id;
			this.title = options.name;
			this.getDb(this.Parameters);



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
			getDb(db) {
				nftsseries(db).then(res => {
					console.log(res);
					if (res) {
						this.dblist = res.list;
						this.total_page = res.total_page;
						// this.shopList = res.result;
						// this.total_page = res.total_page;
						// console.log(res);
						// let arr = [{
						// 	"id": 0,
						// 	"name": "全部系列"
						// }];
						// let tmp = arr.concat(res.series);
						// this.selectListb = JSON.parse(JSON.stringify(tmp));

					}
				});

			}


		}
	};
</script>

<style scoped>

</style>
