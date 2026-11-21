<template>
	<view class="all_nft-container">
		<u-navbar :titleStyle="{'color': '#fff'}" bgColor="#21005D" title="官方公告" :placeholder="true" class="nav_bar"
			:autoBack="true" leftIconColor="#151517">
			<u-icon name="arrow-leftward" color="#fff" slot="left"></u-icon>
		</u-navbar>
		<u-sticky bgColor="#21005D" offset-top="95">
			<u-gap height="5"></u-gap>
			<u-tabs :activeStyle="activeTabStyle" :inactiveStyle="inactiveTabStyle" :list="tablist"
				@click="tabClickHandler" lineColor="#EADDFF" :scrollable="false" :current="activeTab">
				<view slot="center" style="padding-left: 4px;">
					<u-icon name="arrow-up-fill" color="#BC42FF" size="10"></u-icon>
					<u-icon name="arrow-down-fill" color="#606266C2" size="10"></u-icon>
				</view>
			</u-tabs>
		</u-sticky>
		<u-gap height="12"></u-gap>
		<view v-show="activeTab === 0">
			<template v-for="all in allList">
				<view class="title_container" :key="all.name" v-if="all.child.length>0">
					<u-text :text="all.name" color="#FFFFFF"></u-text>
					<u-button v-if="all.child.length>3" color="#6750A4" text="更多" size="normal" @click="activeTab = 1"
						shape="circle" class="more_btn"></u-button>
				</view>
				<navigator :url="`/pages/notice/noticeDetail?id=${item.id}`" class="line_container"
					v-for="item in all.child" :key="item.id">
					<u-icon name="/static/radio.png" size="20"></u-icon>
					<view class="content">
						<u-text :lines="4" color="#D0BCFF" size="12" :text="item.title">
						</u-text>
						<u-gap height="15"></u-gap>
						<u-text :text="`${$u.timeFormat(item.create_time, 'yyyy-mm-dd hh:MM')}`" color="#FFFFFF"
							size="14"></u-text>
					</view>
				</navigator>
			</template>

		</view>

		<u-list @scrolltolower="scrolltolower" class="nfts_list" v-show="activeTab !== 0">
			<u-list-item v-for="nft in nftList" :key="nft.id">

				<navigator :url="`/pages/notice/noticeDetail?id=${nft.id}`" class="line_container">
					<u-icon name="/static/radio.png" size="20"></u-icon>
					<view class="content">
						<u-text :lines="4" color="#D0BCFF" size="12" :text="nft.title">
						</u-text>
						<u-gap height="15"></u-gap>
						<u-text :text="`${$u.timeFormat(nft.create_time, 'yyyy-mm-dd hh:MM')}`" color="#FFFFFF"
							size="14"></u-text>
					</view>
				</navigator>
			</u-list-item>
			<u-divider text="没有更多了..." :hairline="true" v-show="pageSize*page >= count"></u-divider>
		</u-list>
	</view>
</template>

<script>
	import NftCard from '../../components/NftCard/NftCard.vue';
	import {
		noticeTypeReq,
		allNoticeReq,
		singleNoticeReq
	} from '../../api/index.js';

	export default {
		components: {
			// Tabbar,
			NftCard
		},
		data() {
			return {
				tablist: [],
				activeTabStyle: {
					color: '#fff'
				},
				inactiveTabStyle: {
					color: '#D0BCFF'
				},
				iconStyle: {
					color: '#fff'
				},
				allList: [],
				page: 1,
				pageSize: 20,
				isReq: false,
				count: 21,
				nftList: [],
				activeTab: 0,
			};
		},
		watch: {
			activeTab(newval, oldval) {
				if (newval !== oldval && newval !== 0) {
					this.page = 1;
					this.count = 21;
					this.nftList = [];
					this.getSingleNotice(newval)
				}
			}
		},
		onShow() {
			this.page = 1;
			// this.getAllNftList();

			this.getNoticeType();

			this.getAllNotices()
		},
		methods: {
			getSingleNotice(index) {



				const {
					page,
					pageSize,
					isReq,
					count,
					activeTab,
					tablist
				} = this;

				if (pageSize * page >= count) return;
				if (isReq) return;




				let nid = this.tablist[index].id

				this.isReq = true;
				singleNoticeReq({
					nid,
					page,
					page_size: pageSize,
				}).then(res => {
					const {
						result,
						count
					} = res;
					this.count = count;
					if (result && result.length > 0) {
						if (page === 1) {
							this.nftList = result;
						} else {
							this.nftList.splice(this.nftList.length, 0, ...result);
						}
					}
					this.page++;
					this.isReq = false;
				}).catch((err) => {
					this.isReq = false;


				});
			},
			getNoticeType() {
				noticeTypeReq().then(res => {
					let tablist = [{
						name: '综合',
						id: -1
					}];

					tablist = tablist.concat(res)

					this.tablist = tablist;
				})
			},
			// getAllNftList() {

			// 	const {
			// 		page,
			// 		pageSize,
			// 		isReq,
			// 		count,
			// 		activeTab,
			// 		tablist
			// 	} = this;

			// 	if (pageSize * page >= count) return;
			// 	if (isReq) return;

			// 	const sortList = ['index', 'likes', 'price', 'create_time'];

			// 	let sort = sortList[activeTab];

			// 	const operator = tablist[activeTab].sort;
			// 	if (operator === '-')
			// 		sort = operator + sort;


			// 	this.isReq = true;
			// 	allNftsReq({
			// 		page,
			// 		page_size: pageSize,
			// 		sort
			// 	}).then(res => {
			// 		const {
			// 			result,
			// 			count
			// 		} = res;
			// 		this.count = count;
			// 		if (result && result.length > 0) {
			// 			if (page === 1) {
			// 				this.nftList = result;
			// 			} else {
			// 				this.nftList.splice(this.nftList.length, 0, ...result);
			// 			}
			// 		}
			// 		this.page++;
			// 		this.isReq = false;
			// 	}).catch((err) => {
			// 		this.isReq = false;


			// 	});

			// },

			scrolltolower() {
				// this.getAllNftList();
				this.getSingleNotice()
			},
			getAllNotices() {
				allNoticeReq().then(res => {
					let allList = [];
					for (let key in res) {
						let option = {
							name: key,
							child: res[key]
						}

						allList.push(option)
					}

					this.allList = allList;
				})
			},
			tabClickHandler(e) {
				const {
					index
				} = e;


				let {
					tablist
				} = this;


				tablist = tablist.map(item => {
					item.isShow = false
					return item
				})
				this.activeTab = index;
				let tab = tablist[index];
				this.tablist.splice(index, 1, tab)



				this.page = 1;
				this.count = 21;
				if (tab.id === -1) {
					this.getAllNotices()
				} else {

				}
			}
		}
	};
</script>

<style scoped>
	.all_nft-container {
		padding: 48rpx 38rpx calc(env(safe-area-inset-bottom) / 2 + 24rpx);
	}

	/* .scroll-Y {
		height: 300rpx;
		display: flex;
		flex-wrap: wrap;
	} */
	.nft_card {
		margin-bottom: 24rpx;
	}

	.nft_card>>>.u-transition {
		width: 284rpx;
		height: 336rpx;
		margin: auto;
	}

	.nfts_list {
		height: calc(100vh - 164px - env(safe-area-inset-bottom)) !important;
	}

	.nfts_list>>>.uni-scroll-view-content>view {
		display: flex;
		flex-direction: row;
		flex-wrap: wrap;
		justify-content: space-between;
		padding-top: 20rpx;
	}

	>>>.u-list-item {
		width: 100%;
	}

	>>>.u-cell__title-text {
		color: #fff;
	}

	>>>.u-cell__label {
		color: #D0BCFF;
	}

	.title_container {
		display: flex;
		padding: 0 11px;
	}

	.more_btn {
		width: 130rpx;
		padding: 2px 5px;
		height: 60rpx;
	}

	.line_container {
		display: flex;
		align-items: flex-start;
		padding: 32rpx 0;
		border-bottom: 1px solid rgba(255, 255, 255, 0.08);
	}

	.line_container .content {
		margin-left: 26rpx;
	}
</style>
