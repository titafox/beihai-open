<template>
	<view class="profile_container">
		<u-navbar :titleStyle="{'color': '#fff'}" bgColor="#21005D" title="我的邀请" :placeholder="true" class="nav_bar"
			:autoBack="true" leftIconColor="#151517">
			<u-icon name="arrow-leftward" color="#fff" slot="left"></u-icon>
		</u-navbar>
		<u-gap height="20"></u-gap>

		<view class="image-wrapper">
			<image src="../../static/yaoqing.png" mode="widthFix" />
		</view>

		<u-sticky bgColor="#21005D" offset-top="30">
			<view style="padding: 0 20%;">
				<u-tabs
					:activeStyle="activeTabStyle" 
					:inactiveStyle="inactiveTabStyle" 
					:scrollable="false" 
					:list="tablist"
					@click="tabClickHandler" 
					lineColor="transparent" 
					:current="activeTab" 
				/>
			</view>
		</u-sticky>
		<u-gap height="20"></u-gap>

		<template>
			<z-paging 
				v-show="activeTab == 1"
				ref="paging" 
				:fixed="false"  
				:hide-empty-view="true"
				v-model="dblist" 
				@query="queryList"
				style="width: 100%;"
			>
			<view style="background: #381E72;" slot="top">
				<u-row>
					<u-col span="2">
						<view class="in-v1">序号</view>
					</u-col>
					<u-col span="5">
						<view class="in-v1">被邀请人手机号</view>
					</u-col>
					<u-col span="5">
						<view class="in-v1">被邀请时间</view>
					</u-col>
				</u-row>
			</view>
			<!-- <u-gap height="5"></u-gap> -->
			<u-row customStyle="margin-bottom: 10rpx" v-for="(item, index) in dblist" :key="index">
				<u-col span="2">
					<view class="in-v2">{{index+1}}</view>
				</u-col>
				<u-col span="5">
					<view class="in-v2">{{item.nickname}}</view>
				</u-col>
				<u-col span="5">
					<view class="in-v4">{{$u.timeFormat(item.create_time, 'yyyy/mm/dd hh:MM:ss')}}</view>
				</u-col>
			</u-row>
			<!-- <u-loadmore v-if="dblist.length>0" :status="status" @loadmore="loadmore" nomore-text="没有更多了..." /> -->
			<view class="in-v3">总邀请人数 {{nums}} 人，继续加油哦～</view>
			</z-paging>
		</template>

		<template v-if="activeTab == 0">
			<view style="background: #381E72;">
				<u-row>
					<u-col span="2">
						<view class="in-v1">排行</view>
					</u-col>
					<u-col span="5">
						<view class="in-v1">用户</view>
					</u-col>
					<u-col span="5">
						<view class="in-v1">邀请人数</view>
					</u-col>
				</u-row>
			</view>
			<u-gap height="5"></u-gap>
			<view>
				<u-row customStyle="margin-bottom: 10rpx" v-for="(item, index) in rankingList" :key="index">
					<u-col span="2">
						<view class="in-v2">{{ index + 1 }}</view>
					</u-col>
					<u-col span="5">
						<view class="in-v2">{{ item.nickname }}</view>
					</u-col>
					<u-col span="5">
						<view class="in-v4">{{ item.nums }}</view>
					</u-col>
				</u-row>
			</view>
		</template>


	</view>
</template>

<script>
	// import ProfileHeader from '../../components/ProfileHeader/ProfileHeader.vue';



	import {
		invitelist,
		inviteTop
	} from '../../api/index.js';
	export default {
		components: {
			// ProfileHeader,


		},
		data() {
			return {
				activeTabStyle: {
					color: '#fff',
					background: '#6750A4',
					borderRadius: '6px',
					padding: '6px 15px',
					whiteSpace: 'nowrap'
				},
				inactiveTabStyle: {
					color: '#D0BCFF',
					padding: '6px 15px',
					border: '1px solid #6750A4',
					borderRadius: '6px',
					whiteSpace: 'nowrap'
				},
				status: 'loadmore',
				total_page: 0,
				nums: 0,
				dblist: [],
				Parameters: {
					page: 1,
				},
				rankingList: [],
				tablist: [{
						name: '排行榜',
					},
					{
						name: '我的邀请',
					}
				],
				activeTab: 0
			};
		},

		onLoad(options) {

			// this.getDb();
			this.getInviteTop()


		},

		// onReachBottom() {
		// 	console.log(this.Parameters);
		// 	if (this.Parameters.page >= this.total_page) return;
		// 	this.status = 'loading';
		// 	this.Parameters.page = ++this.Parameters.page;

		// 	invitelist(this.Parameters).then(res => {

		// 		if (res) {

		// 			let arr = res.list;
		// 			this.dblist = this.dblist.concat(arr);

		// 			if (this.Parameters.page >= this.total_page) this.status = 'nomore';
		// 			else this.status = 'loading';



		// 		}
		// 	});
		// },

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
			async queryList (pageNum) {
				try{
					this.Parameters.page = pageNum
					const res = await invitelist(this.Parameters)
					this.$refs.paging.complete(res.list)
				}catch(e){
					this.$refs.paging.complete(false)
				}
			},
			tabClickHandler({ index }) {
				this.activeTab = index
				console.log(index);
			},
			loadmore() {
				if (this.Parameters.page >= this.total_page) return;
				this.status = 'loading';
				this.Parameters.page = ++this.Parameters.page;

				invitelist(this.Parameters).then(res => {

					if (res) {

						let arr = res.list;
						this.dblist = this.dblist.concat(arr);

						if (this.Parameters.page >= this.total_page) this.status = 'nomore';
						else this.status = 'loading';



					}
				});
			},

			getDb() {
				invitelist(this.Parameters).then(res => {
					console.log(res);
					if (res) {
						this.dblist = res.list;
						this.total_page = res.total_page;
						this.nums = res.nums;
						if (this.Parameters.page >= this.total_page) {};
						console.log(this.total_page);
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

			},
			async getInviteTop () {
				const res = await inviteTop()
				this.rankingList = res
			}
		}
	};
</script>

<style>
	page {
		height: 100vh;

	}
	.profile_container {
		height: 100vh;
		display: flex;
		flex-direction: column;
	}

	.in-v1 {
		text-align: center;
		padding: 16rpx;
		font-family: 'PingFang SC';
		font-style: normal;
		font-weight: 400;
		font-size: 32rpx;
	}

	.in-v2 {
		text-align: center;
		padding: 16rpx;
		font-family: 'PingFang SC';
		font-style: normal;
		font-size: 32rpx;
	}

	.in-v4 {
		text-align: center;
		padding: 16rpx;
		font-family: 'PingFang SC';
		font-style: normal;
		font-size: 29rpx;
		color: #7B61FF;
	}

	.in-v3 {
		width: 100%;
		margin-top: 30rpx;
		text-align: center;
		font-size: 32rpx;
	}

	.image-wrapper {
		display: flex;
		justify-content: center;
	}
</style>
