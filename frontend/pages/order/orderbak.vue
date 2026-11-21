<template>
	<view>
		<u-navbar :titleStyle="{'color': '#fff'}" bgColor="#21005D" title="订单" :placeholder="true" class="nav_bar"
			:autoBack="true" leftIconColor="#151517">
			<u-icon name="arrow-leftward" color="#fff" slot="left"></u-icon>
		</u-navbar>
		<u-sticky bgColor="#21005D" offset-top="30">
			<u-tabs :activeStyle="activeTabStyle" :inactiveStyle="inactiveTabStyle" :scrollable="false"
				:list="tablist" @click="tabClickHandler" lineColor="transparent" :current="activeTab"></u-tabs>
		</u-sticky>
		<view class="custom_container">
			
			<!-- <u-gap height="28"></u-gap> -->

			<view v-show="activeTab === 0">
				<u-empty v-if="nftList_t.length === 0" mode="favor" text="啥都还没买过"></u-empty>
				<u-list v-else class="nfts_list" height="calc(100vh - 200px)">
					<!-- @scrolltolower="getTotalList" -->
					<u-list-item v-for="(nft,i) in nftList_t" :key="i">
						<view class="card">
							<view class="card_header">
								<u-avatar :src="nft.image" shape="square"></u-avatar>
								<view style="flex: auto;margin-left: 28rpx;">
									<view style="display: flex;">
										<u-text :text="nft.name" color="#fff" size="18"></u-text>
										<u-text :text="`支付金额：¥ ${nft.money}`" color="#fff" size="12" align="right">
										</u-text>
									</view>
									<view style="display: flex;">
										<u-text :text="nft.no" color="#51EC44" size="14"></u-text>
										<u-text :text="`${$u.timeFormat(nft.create_time, 'yyyy/mm/dd')}`"
											color="#9A82DB" size="12" align="right"></u-text>
									</view>
								</view>
							</view>
							<view class="card_body"
								style="display: flex;align-items: center;justify-content: space-between;">
								
								
								<u-text v-if="!nft.countdown" :text="nft.status_text" color="#9A82DB" size="12">
								</u-text>


								<view v-else style="flex: auto;display:flex;align-items: center;">
									<text style="color: #fff;font-size: 12px;">{{nft.status_text}}</text>
									<u-count-down @finish="countdownFinished('total',i)" :time="nft.countdown*1000"
										format="HH:mm:ss"></u-count-down>
									<text style="color: #fff;font-size: 12px;">后自动取消</text>
								</view>

								<u-button shape="circle" v-if="nft.status === 1"
									color="linear-gradient(0deg, rgba(0, 0, 0, 0.2), rgba(0, 0, 0, 0.2)), linear-gradient(98.46deg, #4F75FF -0.18%, #6954FF 48.62%, #892CFF 100.99%, #892CFF 103.22%)"
									text="去支付" @click="toPayHandler(nft.pay_url)" size="mini"
									:customStyle="{'width':'40px'}"></u-button>

							</view>
						</view>
						<!-- <NftCard class="nft_card" :info="nft" :isHandsel="true" @handsel="handselInfo" /> -->
					</u-list-item>
					<u-divider lineColor="#fff" text="没有更多了..." :hairline="true" v-show="pageSize_t*page_t >= count_t">
					</u-divider>
				</u-list>

			</view>
			<view v-show="activeTab === 2">

				<u-empty v-if="nftList_f.length === 0" mode="favor" text="啥都还没买过"></u-empty>
				<u-list v-else class="nfts_list"  height="calc(100vh - 200px)">
					<u-list-item v-for="(nft,i) in nftList_f" :key="i">
						<!-- @scrolltolower="getFinishedList" -->
						<!-- <NftCard class="nft_card" :info="nft" /> -->
						<view class="card">
							<view class="card_header">
								<u-avatar :src="nft.image" shape="square"></u-avatar>
								<view style="flex: auto;margin-left: 28rpx;">
									<view style="display: flex;">
										<u-text :text="nft.name" color="#fff" size="18"></u-text>
										<u-text :text="`支付金额：¥ ${nft.money}`" color="#fff" size="12" align="right">
										</u-text>
									</view>
									<view style="display: flex;">
										<u-text :text="nft.no" color="#51EC44" size="14"></u-text>
										<u-text :text="`${$u.timeFormat(nft.create_time, 'yyyy/mm/dd')}`"
											color="#9A82DB" size="12" align="right"></u-text>
									</view>
								</view>
							</view>
							<view class="card_body">
								<u-text :text="nft.status_text" color="#9A82DB" size="12"></u-text>
							</view>
						</view>
					</u-list-item>
					<u-divider lineColor="#fff" text="没有更多了..." :hairline="true" v-show="pageSize_f*page_f >= count_f">
					</u-divider>
				</u-list>


			</view>

			<view v-show="activeTab === 1">
				<u-empty v-if="nftList_l.length === 0" mode="favor" text="啥都还没买过"></u-empty>
				<!-- @scrolltolower="getCurList" -->
				<u-list v-else class="nfts_list"  height="calc(100vh - 200px)">
					<u-list-item v-for="(nft,i) in nftList_l" :key="i">
						<!-- <NftCard class="nft_card" :info="nft" /> -->
						<view class="card">
							<view class="card_header">
								<u-avatar :src="nft.image" shape="square"></u-avatar>
								<view style="flex: auto;margin-left: 28rpx;">
									<view style="display: flex;">
										<u-text :text="nft.name" color="#fff" size="18"></u-text>
										<u-text :text="`支付金额：¥ ${nft.money}`" color="#fff" size="12" align="right">
										</u-text>
									</view>
									<view style="display: flex;">
										<u-text :text="nft.no" color="#51EC44" size="14"></u-text>
										<u-text :text="`${$u.timeFormat(nft.create_time, 'yyyy/mm/dd')}`"
											color="#9A82DB" size="12" align="right"></u-text>
									</view>
								</view>
							</view>
							<view class="card_body"
								style="display: flex;align-items: center;justify-content: space-between;">
								<u-text v-if="!nft.countdown" :text="nft.status_text" color="#9A82DB" size="12">
								</u-text>


								<view v-else style="flex: auto;display:flex;align-items: center;">
									<text style="color: #fff;font-size: 12px;">{{nft.status_text}}</text>
									<u-count-down @finish="countdownFinished('cur',i)" :time="nft.countdown*1000"
										format="HH:mm:ss"></u-count-down>
									<text style="color: #fff;font-size: 12px;">后自动取消</text>
								</view>

								<u-button shape="circle"
								v-if="nft.status === 1"
									color="linear-gradient(0deg, rgba(0, 0, 0, 0.2), rgba(0, 0, 0, 0.2)), linear-gradient(98.46deg, #4F75FF -0.18%, #6954FF 48.62%, #892CFF 100.99%, #892CFF 103.22%)"
									text="去支付" @click="toPayHandler(nft.pay_url)" size="mini"
									:customStyle="{'width':'40px'}"></u-button>

							</view>
						</view>
					</u-list-item>
					<u-divider lineColor="#fff" text="没有更多了..." :hairline="true" v-show="pageSize_l*page_l >= count_l">
					</u-divider>
				</u-list>


			</view>

			<view v-show="activeTab === 3">
				<u-empty v-if="nftList_c.length === 0" mode="favor" text="啥都还没买过"></u-empty>
				<u-list v-else class="nfts_list"  height="calc(100vh - 200px)">
					<!-- @scrolltolower="getCancelList" -->
					<u-list-item v-for="(nft,i) in nftList_c" :key="i">
						<!-- <NftCard class="nft_card" :info="nft" /> -->
						<view class="card">
							<view class="card_header">
								<u-avatar :src="nft.image" shape="square"></u-avatar>
								<view style="flex: auto;margin-left: 28rpx;">
									<view style="display: flex;">
										<u-text :text="nft.name" color="#fff" size="18"></u-text>
										<u-text :text="`支付金额：¥ ${nft.money}`" color="#fff" size="12" align="right">
										</u-text>
									</view>
									<view style="display: flex;">
										<u-text :text="nft.no" color="#51EC44" size="14"></u-text>
										<u-text :text="`${$u.timeFormat(nft.create_time, 'yyyy/mm/dd')}`"
											color="#9A82DB" size="12" align="right"></u-text>
									</view>
								</view>
							</view>
							<view class="card_body">
								<u-text :text="nft.status_text" color="#9A82DB" size="12"></u-text>
							</view>
						</view>
					</u-list-item>
					<u-divider lineColor="#fff" text="没有更多了..." :hairline="true" v-show="pageSize_c*page_c >= count_c">
					</u-divider>
				</u-list>


			</view>

		</view>
	</view>
</template>

<script>
	import {
		orderListReq
	} from '@/api/index.js';
	export default {
		data() {
			return {
				tablist: [{
						name: '全部',
					},
					{
						name: '进行中',
					},
					{
						name: '已完成',
					},
					{
						name: '已取消',
					}
				],
				activeTabStyle: {
					color: '#fff',
					background: '#6750A4',
					borderRadius: '6px',
					padding: '6px 10px',
					whiteSpace: 'nowrap'
				},
				inactiveTabStyle: {
					color: '#D0BCFF',
					padding: '6px 10px',
					border: '1px solid #6750A4',
					borderRadius: '6px',
					whiteSpace: 'nowrap'
				},
				activeTab: 0,

				// 全部
				page_t: 1,
				pageSize_t: 20,
				isReq_t: false,
				count_t: 21,
				nftList_t: [],

				// 进行中
				page_l: 1,
				pageSize_l: 20,
				isReq_l: false,
				count_l: 21,
				nftList_l: [],
				// 已完成
				page_f: 1,
				pageSize_f: 20,
				isReq_f: false,
				count_f: 21,
				nftList_f: [],

				// 已取消
				page_c: 1,
				pageSize_c: 20,
				isReq_c: false,
				count_c: 21,
				nftList_c: [],
			};
		},
		onLoad() {
			this.getTotalList();
		},
		methods: {
			tabClickHandler(e) {
				const {
					index
				} = e;
				this.activeTab = index;
				switch (index) {
					case 0:
						this.getTotalList();
						break;
					case 1:
						this.getCurList();
						break;
					case 2:
						this.getFinishedList();
						break;
					default:
						this.getCancelList();

				}
			},

			getTotalList() {
				const {
					page_t,
					pageSize_t,
					isReq_t,
					count_t
				} = this;
				if (pageSize_t * page_t >= count_t) return;
				if (isReq_t) return;

				this.isReq_t = true;
				orderListReq({
					typ: 0,
					page: page_t,
					// page_size: pageSize_t,
					
				}).then(res => {
					console.log(res);
					// const {
					// 	result,
					// 	count
					// } = res;
					// this.count_t = count;
					const result = res;
					// console.log(res)
					if (result && result.length > 0) {
						if (page_t === 1) {
							this.nftList_t = result;
						} else {
							this.nftList_t.splice(this.nftList_t.length, 0, ...result);
						}
					}
					this.page_t++;
					this.isReq_t = false;
				}).catch(() => {
					this.isReq_t = false;
				});
			},

			getCurList() {
				const {
					page_l,
					pageSize_l,
					isReq_l,
					count_l
				} = this;
				if (pageSize_l * page_l >= count_l) return;
				if (isReq_l) return;

				this.isReq_l = true;
				orderListReq({
					page: page_l,
					page_size: pageSize_l,
					type: 1
				}).then(res => {
					// const {
					// 	result,
					// 	count
					// } = res;
					// this.count_l = count;

					const result = res;
					if (result && result.length > 0) {
						if (page_l === 1) {
							this.nftList_l = result;
						} else {
							this.nftList_l.splice(this.nftList_l.length, 0, ...result);
						}
					}
					this.page_l++;
					this.isReq_l = false;
				}).catch(() => {
					this.isReq_l = false;
				});
			},
			getFinishedList() {
				const {
					page_f,
					pageSize_f,
					isReq_f,
					count_f
				} = this;
				if (pageSize_f * page_f >= count_f) return;
				if (isReq_f) return;

				this.isReq_f = true;
				orderListReq({
					page: page_f,
					page_size: pageSize_f,
					type: 2
				}).then(res => {
					// const {
					// 	result,
					// 	count
					// } = res;
					// this.count_f = count;
					const result = res;

					if (result && result.length > 0) {
						if (page_f === 1) {
							this.nftList_f = result;
						} else {
							this.nftList_f.splice(this.nftList_f.length, 0, ...result);
						}
					}
					this.page_f++;
					this.isReq_f = false;
				}).catch(() => {
					this.isReq_f = false;
				});
			},

			getCancelList() {
				const {
					page_c,
					pageSize_c,
					isReq_c,
					count_c
				} = this;
				if (pageSize_c * page_c >= count_c) return;
				if (isReq_c) return;

				this.isReq_c = true;
				orderListReq({
					page: page_c,
					page_size: pageSize_c,
					type: 3
				}).then(res => {
					// const {
					// 	result,
					// 	count
					// } = res;
					// this.count_f = count;
					const result = res;

					if (result && result.length > 0) {
						if (page_c === 1) {
							this.nftList_c = result;
						} else {
							this.nftList_c.splice(this.nftList_c.length, 0, ...result);
						}
					}
					this.page_c++;
					this.isReq_c = false;
				}).catch(() => {
					this.isReq_c = false;
				});
			},

			toPayHandler(pay_url) {
				plus.runtime.openWeb(pay_url, () => {
					// 调用失败 callback
					uni.$u.toast('调起支付失败，请稍后再试！');
				});
			},
			countdownFinished(type, index) {
				if (type === 'total') {
					const {
						nftList_t
					} = this

					let item = nftList_t[index];
					item = {
						...item,
						countdown: 0,
						status:0
					}
					this.nftList_t.splice(index, 1)
				} else {
					const {
						nftList_l
					} = this;

					let item = nftList_l[index];
					item = {
						...item,
						countdown: 0,
						status:0
					}
					this.nftList_l.splice(index, 1)
				}
			}

		}
	};
</script>

<style scoped>
	.card {
		background: #180240;
		border-radius: 12px;
		margin-bottom: 56rpx;
	}

	.card .card_header {
		padding: 24rpx 20rpx;
		border-bottom: 1px solid #381E72;
		display: flex;
	}

	.card .card_body {
		padding: 32rpx 28rpx;
	}

	>>>.u-count-down__text {
		color: #ffffff;
		font-size: 12px;
		/* margin-left: 20rpx; */
	}
</style>
