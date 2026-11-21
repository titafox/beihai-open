<template>
	<view class="profile_container">
		<u-navbar 
			:titleStyle="{'color': '#fff'}" bgColor="#21005D" 
			:title="title" 
			:placeholder="true" 
			class="nav_bar"
			:autoBack="true" 
			leftIconColor="#151517"
		>
			<u-icon name="arrow-leftward" color="#fff" slot="left"></u-icon>
		</u-navbar>

		<view class="m-flex-sb ci-v1">
			<view 
				class="ci-v2" 
				:class="current == index ? '' : 'ci-v2-av'" 
				v-for="(item, index) in tablist" 
				:key="index"
				@click="chang(index)"
			>
				{{item.name}}
			</view>

		</view>

		<u-empty v-if="dblist.length === 0" mode="favor" text="啥都还没买过"></u-empty>
		<ShopCard :list="dblist"></ShopCard>

		<u-loadmore v-if="dblist.length>0" :status="status" @loadmore="loadmore" nomore-text="没有更多了..." />
		
	</view>
</template>

<script>
	// import ProfileHeader from '../../components/ProfileHeader/ProfileHeader.vue';

	import ShopCard from '../../components/ShopCard2/ShopCard.vue'
	import {
		nftscirculating
	} from '../../api/index.js';
	export default {
		components: {
			ShopCard // ProfileHeader,


		},
		data() {
			return {
				status: 'loadmore',
				total_page: 0,
				current: 0,
				Parameters: {
					page: 1,
					sort: 'price',
					wid: 0,
				},
				dblist: [],
				title: "北海",
				tablist: [{
					"id": 0,
					"name": "按价格排序",
					"sort": 'price',
				}, {
					"id": 1,
					"name": "按编号排序",
					"sort": 'no',
				}],
			};
		},

		onLoad(options) {
			console.log(options);
			this.title = options.name;
			this.Parameters.wid = options.wid;



		},
		onShow(options) {
			this.getdb(this.Parameters);
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

		onReachBottom() {
			console.log(this.Parameters);
			if (this.Parameters.page >= this.total_page) return;
			this.status = 'loading';
			this.Parameters.page = ++this.Parameters.page;

			nftscirculating(this.Parameters).then(res => {

				if (res) {

					let arr = res.list;
					this.dblist = this.dblist.concat(arr);

					if (this.Parameters.page >= this.total_page) this.status = 'nomore';
					else this.status = 'loading';



				}
			});
		},
		methods: {
			chang(i) {
				this.current = i;
				this.Parameters.page = 1;
				this.Parameters.sort = this.tablist[i].sort;
				this.getdb(this.Parameters);


			},

			loadmore() {
				if (this.Parameters.page >= this.total_page) return;
				this.status = 'loading';
				this.Parameters.page = ++this.Parameters.page;

				nftscirculating(this.Parameters).then(res => {

					if (res) {

						let arr = res.list;
						this.dblist = this.dblist.concat(arr);

						if (this.Parameters.page >= this.total_page) this.status = 'nomore';
						else this.status = 'loading';



					}
				});
			},



			getdb(db) {
				nftscirculating(db).then(res => {
					console.log(res);
					this.dblist = res.list;
					this.total_page = res.total_page;
					if (this.Parameters.page >= this.total_page) this.status = 'nomore';
				}).catch(() => {

				});
			},


		}
	};
</script>

<style>
	.ci-v1 {
		margin: 32rpx 118rpx;
	}

	.ci-v2 {
		background: #6750A4;
		border-radius: 12rpx;
		width: 228rpx;
		height: 65rpx;
		text-align: center;
		line-height: 65rpx;
	}

	.ci-v2-av {
		border: 2rpx #6750A4 solid;
		background: #21005D;
	}

	.ci-v3 {
		width: 100%;
		background-color: white;
	}
</style>
