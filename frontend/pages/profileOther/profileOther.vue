<template>
	<view class="custom_container">
		<u-navbar :titleStyle="{'color': '#fff'}" bgColor="#21005D" :title="`${profileInfo.nickname}`"
			:placeholder="true" class="nav_bar" :autoBack="true" leftIconColor="#151517">
			<u-icon name="arrow-leftward" color="#fff" slot="left"></u-icon>
		</u-navbar>




		<view class="author_container">
			<view class="author_header">
				<view style="display: flex;">
					<u-avatar size="56" :src="profileInfo.avatar"></u-avatar>
					<view style="display: flex;flex-direction: column;margin-left: 10px;">
						<u-text class="preson_title" text="发行方" size="12" color="#8C8585"></u-text>
						<u-text class="person_name" :text="`@${profileInfo.nickname}`" size="12" color="#ffffff">
						</u-text>
					</view>
				</view>


			</view>
			<u-gap height="16"></u-gap>
			<view class="author_footer">
				<u-gap height="20"></u-gap>
				<u-text text="发行方介绍介绍介绍" color="#ffffff" size="12"></u-text>
			</view>
		</view>
		<u-gap height="2"></u-gap>
		<view class="profile_container">
			<u-sticky bgColor="#21005D" offset-top="90">
				<u-gap height="10"></u-gap>
				<u-tabs :activeStyle="activeTabStyle" :inactiveStyle="inactiveTabStyle"
					:list="tablist" @click="tabClickHandler" lineColor="transparent"></u-tabs>
					 <!-- :scrollable="false" -->
			</u-sticky>
			<u-gap height="12"></u-gap>
			
			<view v-show="activeTab === 0">
				
				<u-empty v-if="nftList_p.length === 0" mode="favor" text="啥都还没买过"></u-empty>
				
				<u-list v-else @scrolltolower="getPublishList" class="nfts_list">
					<u-list-item v-for="nft in nftList_p" :key="nft.id">
						<NftCard class="nft_card" :info="nft" />
					</u-list-item>
					<u-divider  lineColor="#fff" text="没有更多了..." :hairline="true"
						v-show="pageSize_p*page_p >= count_p"></u-divider>
				</u-list>
			
			</view>
			
			<view v-show="activeTab === 1">
				<u-empty v-if="nftList_s.length === 0" mode="favor" text="啥都还没卖过"></u-empty>
				
				<u-list v-else @scrolltolower="getSaleList" class="nfts_list">
					<u-list-item v-for="nft in nftList_s" :key="nft.id">
						<NftCard class="nft_card" :info="nft" />
					</u-list-item>
					<u-divider text="没有更多了..." :hairline="true" v-show="pageSize_s*page_s >= count_s"></u-divider>
				</u-list>

			</view>

			<view v-show="activeTab === 2">
				<u-empty v-if="nftList_h.length === 0" mode="favor" text="啥都还没发布过"></u-empty>
				
				<u-list v-else @scrolltolower="getHistoryList" class="nfts_list">
					<u-list-item v-for="nft in nftList_h" :key="nft.id">
						<NftCard class="nft_card" :info="nft" />
					</u-list-item>
					<u-divider text="没有更多了..." :hairline="true" v-show="pageSize_h*page_h >= count_h"></u-divider>
				</u-list>

			</view>
		</view>
	</view>
</template>

<script>
	// import ProfileHeader from '../../components/ProfileHeader/ProfileHeader.vue';
	import NftCard from '../../components/NftCard/NftCard.vue';

	import {
		otherProfileReq,
		otherSaleListReq,
		otherHistoryListReq,
		otherProdListReq
	} from '../../api/index.js'
	export default {
		components: {
			// ProfileHeader,
			NftCard
		},
		data() {
			return {
				tablist: [{
						name: '藏品',
					},
					// {
					// 	name: '',
					// },
					// {
					// 	name: '',
					// }
				],
				activeTabStyle: {
					color: '#fff',
					background: '#6750A4',
					borderRadius: '6px',
					padding: '6px 20px',
					whiteSpace: 'nowrap'
				},
				inactiveTabStyle: {
					color: '#D0BCFF',
					padding: '6px 20px',
					border: '1px solid #6750A4',
					borderRadius: '6px',
					whiteSpace: 'nowrap'
				},

				profileInfo: {
					avatar: '',
					nickname: ''
				},
				id: '',
				activeTab: 0,

				page_s: 1,
				pageSize_s: 20,
				isReq_s: false,
				count_s: 21,
				nftList_s: [],

				page_h: 1,
				pageSize_h: 20,
				isReq_h: false,
				count_h: 21,
				nftList_h: [],
				
				page_p: 1,
				pageSize_p: 20,
				isReq_p: false,
				count_p: 21,
				nftList_p: [],
			}
		},
		onLoad(options) {
			const {
				id
			} = options;
			this.userId = id;
			this.getOtherProfile();
			switch (this.activeTab) {
				case 0:
					this.getProductList();
					break;
				case 2:
					this.getHistoryList();
					break;
				default:
					this.getSaleList();
					break;
			}
		},
		methods: {
			getOtherProfile() {
				otherProfileReq(this.userId).then(res => {
					this.profileInfo = {
						avatar: res.avatar,
						nickname: res.nickname
					};
				})
			},
			getSaleList() {
				const {
					page_s,
					pageSize_s,
					isReq_s,
					count_s,
					userId
				} = this;
				if (pageSize_s * page_s >= count_s) return;
				if (isReq_s) return;

				this.isReq_s = true;
				otherSaleListReq({
					page: page_s,
					page_size: pageSize_s,
					id: userId
				}).then(res => {
					const {
						result,
						count
					} = res;
					this.count_s = count;
					if (result && result.length > 0) {
						if (page_s === 1) {
							this.nftList_s = result;
						} else {
							this.nftList_s.splice(this.nftList_s.length, 0, ...result);
						}
					}
					this.page_s++;
					this.isReq_s = false;
				}).catch(() => {
					this.isReq_s = false;
				});
			},
			getHistoryList() {

				const {
					page_h,
					pageSize_h,
					isReq_h,
					count_h,
					userId
				} = this;
				if (pageSize_h * page_h >= count_h) return;
				if (isReq_h) return;

				this.isReq_h = true;
				otherHistoryListReq({
					page: page_h,
					page_size: pageSize_h,
					id: userId
				}).then(res => {
					const {
						result,
						count
					} = res;
					this.count_h = count;
					if (result && result.length > 0) {
						if (page_h === 1) {
							this.nftList_h = result;
						} else {
							this.nftList_h.splice(this.nftList_h.length, 0, ...result);
						}
					}
					this.page_h++;
					this.isReq_h = false;
				}).catch(() => {
					this.isReq_h = false;
				});
			},
			getProductList() {

				const {
					page_p,
					pageSize_p,
					isReq_p,
					count_p,
					userId
				} = this;
				if (pageSize_p * page_p >= count_p) return;
				if (isReq_p) return;

				this.isReq_p = true;
				otherProdListReq({
					page: page_p,
					page_size: pageSize_p,
					id: userId
				}).then(res => {
					const {
						result,
						count
					} = res;
					this.count_p = count;
					if (result && result.length > 0) {
						if (page_p === 1) {
							this.nftList_p = result;
						} else {
							this.nftList_p.splice(this.nftList_p.length, 0, ...result);
						}
					}
					this.page_p++;
					this.isReq_p = false;
				}).catch(() => {
					this.isReq_p = false;
				});
			},

			tabClickHandler(e) {
				const {
					index
				} = e;
				this.activeTab = index;
				switch (index) {
					case 0:
						this.getProductList();
						break;
					case 2:
						this.getHistoryList();
						break;
					default:
						this.getSaleList();
						break;
				}
			}
		}
	}
</script>

<style scoped>
	/* .profile_container {
		padding: 48rpx 38rpx calc(env(safe-area-inset-bottom) / 2 + 24rpx);
	} */

	.intr_info>>>span {
		text-align: center;
	}


	.nft_card {
		margin-bottom: 24rpx;
	}


	.nft_card>>>.u-transition {
		width: 264rpx;
		height: 312rpx;
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

	.author_container {
		background: #180240;
		border-radius: 12px;
		padding: 20rpx 40rpx;
		margin-bottom: 20rpx;
	}

	.author_container .author_header {
		display: flex;
		justify-content: space-between;

	}

	.author_container .author_footer {
		border-top: 1px solid #381E72;
	}
</style>
