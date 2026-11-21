<template>
	<view class="padding_0_24">
		<u-gap height="44"></u-gap>
		<view class="block-title_container">
			<u-row customStyle="margin-bottom: 10px">
				<u-col span="11">
					<view class="bh-layout bg-purple-light">

						<u-text text="Market" size="21" color="#EADDFF"></u-text>

					</view>
				</u-col>
				<u-col span="1">
					<u-image :src="profileInfo.avatar" :lazy-load="true" shape="circle" width="30" height="30"
						align="right" @click="gotonfts"> 
					</u-image>
				</u-col>
			</u-row>
		</view>
		<u-gap height="8"></u-gap>
		<view class="search padding_0_24">
			<u-search 
				placeholder="请输入藏品名称"
				bgColor="#381E72"
				color="#ffffff"
				placeholderColor="#D0BCFF"
				searchIconColor="#ffffff"
				height="80rpx"
				v-model="Parameters.keyword"
				@search="search"
				@custom="search"
			/>
		</view>
		<view class="m-flex m-flex-c tab-m">
			<view 
				class="m-flex m-flex-c tab-ma" 
				:class="current==index?'tab-ma-av':''" 
				v-for="(item, index) in tabList"
				:key="index" 
				@click="chang(index)"
			>
				<text style="padding-right: 8rpx;">{{item.name}}</text>
				<u-icon 
					:name="current!=index?'arrow-up':'arrow-down'"
					:color="current==index?'#D0BCFF':'#ffffff'" 
					size="18"
				/>
			</view>
		</view>
		<ShopCard :list="shopList"></ShopCard>
		<u-loadmore :status="status" nomore-text="没有更多了..." />
		<!-- 选择 -->

		<u-empty v-if="shopList.length<1" mode="car" icon="http://cdn.uviewui.com/uview/empty/car.png" text="敬请期待...">
		</u-empty>

		<Tabbar />
		<u-overlay :show="show" @click="show = false">
			<view class="popup-a"></view>
			<view class="popup" @tap.stop>

				<view class="popup-m " v-for="(item, index) in selectList" :key="index" @click="selectTag(item)">
					{{item.name}}
				</view>

			</view>
			<view class="popup-b"></view>
		</u-overlay>

	</view>
</template>

<script>
	import Tabbar from '../../components/Tabbar/Tabbar.vue';
	import ShopCard from '../../components/ShopCard2/ShopCard.vue'
	import {
		market
	} from '../../api/index.js';

	export default {
		components: {
			Tabbar,
			ShopCard
		},
		data() {
			return {
				scrollTop: 0,
				current: 1,
				show: false,
				status: 'loadmore',
				total_page: 0,
				profileInfo: {},
				tabList: [{
					"id": 1,
					"name": "藏品",

				}, {
					"id": 2,
					"name": "全部系列",

				}, {
					"id": 3,
					"name": "最新上架",

				}, ],
				shopList: [],
				selectList: [],
				selectLista: [{
					"id": 1,
					"name": "藏品"
				}, {
					"id": 2,
					"name": "盲盒"
				}],
				selectListb: [],


				selectListc: [{
					"id": 1,
					"name": "最新上架",
					"sort": "time_desc",
				}, {
					"id": 2,
					"name": "价格升序",
					"sort": "price_asc",
				}, {
					"id": 3,
					"name": "价格降序",
					"sort": "price_desc",

				}],
				Parameters: {
					page: 1,
					typ: 1,
					sort: 'time_desc',
					wid: 0,
					keyword: ''
				},
			}
		},
		onLoad() {
			this.profileInfo = this.$storea.state.userInfo;
			let db = this.Parameters;
			this.loadData(db);
		},
		onReachBottom() {
			if (this.Parameters.page >= this.total_page) return;
			this.status = 'loading';
			this.Parameters.page = ++this.Parameters.page;

			market(this.Parameters).then(res => {

				if (res) {

					let arr = res.result;
					this.shopList = this.shopList.concat(arr);
					
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
		async onPullDownRefresh () {
			this.Parameters.page = 1
			await this.loadData(this.Parameters)
			uni.stopPullDownRefresh()
		},
		onShow() {
			uni.hideTabBar();
		},
		onPageScroll(e) {
			this.scrollTop = e.scrollTop;
		},
		watch: {


			show: function(e) {
				if (e) {
					this.stop();
				} else {
					this.move();
				}
			}

		},
		methods: {
			gotonfts(){
				uni.$u.route("/pages/nfts/nfts");
			},
			
			selectTag(it) {
				this.Parameters.page = 1;
				let i = this.current;
				this.tabList[i].name = it.name;
				switch (i) {
					case 0: {
						this.Parameters.typ = it.id;
						break;
					}
					case 1: {
						this.Parameters.wid = it.id;
						break;
					}
					case 2: {
						this.Parameters.sort = it.sort;
						break;
					}

				};

				let db = this.Parameters;
				this.loadData(db);
				this.show = false;

			},

			top() { //回到顶部  
				uni.pageScrollTo({
					scrollTop: 0,
					duration: 200
				});
			},
			chang(i) {
				this.top();
				console.log(i);
				this.current = i;
				switch (i) {
					case 0: {
						this.selectList = this.selectLista;
						break;
					}
					case 1: {
						this.selectList = this.selectListb;
						break;
					}
					case 2: {
						this.selectList = this.selectListc;
						break;
					}

				}

				this.show = true;

			},
			async loadData(db) {
				const res = await market(db)
				if (res) {
					this.shopList = res.result;
					this.total_page = res.total_page;
					if (this.Parameters.page >= this.total_page) this.status = 'nomore';
					console.log(res);
					let arr = [{
						"id": 0,
						"name": "全部系列"
					}];
					let tmp = arr.concat(res.series);
					this.selectListb = JSON.parse(JSON.stringify(tmp));
				}
			},
			search () {
				this.loadData(this.Parameters)
			}
		}
	}
</script>

<style lang="less">
	.padding_0_24 {
		padding: 0 16rpx;
	}

	.block-title_container .block-title \deep\ span {
		font-weight: 600;
		line-height: 26px;
		text-align: center;
		letter-spacing: 0.005em;
	}

	.block-title_container .block-title {
		padding: 26rpx 28rpx;
		width: auto;
	}

	.tab-m {
		padding: 8rpx;
		margin: 8rpx 0 16rpx;
		/* background-color: white; */
	}

	.tab-ma {
		justify-content: space-between;
		flex: 1;
		/* border-bottom: 3rpx #7E7E7E solid; */
		color: #ffffff;
		padding: 12rpx 20rpx;
		margin: 0 10rpx;
		font-weight: 600;
		font-size: 28rpx;
		border-radius: 99px;
		border: 1px solid #9A82DB;
	}

	.tab-ma-av {
		/* border-bottom: 5rpx #D0BCFF solid; */
		color: #D0BCFF;
	}

	.popup {
		display: flex;
		flex-wrap: wrap;
		justify-content: space-between;
		padding: 0rpx 32rpx 0rpx 32rpx;
		background-color: #21005D;
		border-radius: 0 0 16rpx 16rpx;
		opacity: 0.9;

	}

	.popup-a {
		/* background-color: #21005D; */
		min-height: 400rpx;
		opacity: 0.2;
	}

	.popup-b {
		background-color: #21005D;
		min-height: 30rpx;
		opacity: 0.9;
	}


	.popup-m {
		width: 300rpx;
		height: 68rpx;
		color: white;
		line-height: 68rpx;
		text-align: center;
		background: #7B61FF;
		border-radius: 32rpx;
		margin-top: 32rpx;
	}
	
	.search {
		margin-bottom: 30rpx;
		
		/deep/ .u-search {
			.u-search__content {
				border-radius: 30px 0 0 30px !important;
			}
			
			.u-search__action {
				display: flex;
				align-items: center;
				justify-content: center;
				background: #7B61FF !important;
				border-radius: 0 30px 30px 0 !important;
				margin-left: 0 !important;
				color: #ffffff !important;
				width: 65px !important;
				align-self: stretch !important;
			}
		}
	}
</style>
