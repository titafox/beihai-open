<template>
	<view>
		<u-navbar :titleStyle="{'color': '#fff'}" bgColor="#21005D" title="订单" :placeholder="true" class="nav_bar"
			:autoBack="true" leftIconColor="#151517">
			<u-icon name="arrow-leftward" color="#fff" slot="left"></u-icon>
		</u-navbar>
		<u-sticky bgColor="#21005D" offset-top="30">
			<u-tabs :activeStyle="activeTabStyle" :inactiveStyle="inactiveTabStyle" :scrollable="false" :list="tablist"
				@click="tabClickHandler" lineColor="transparent" :current="activeTab"></u-tabs>
		</u-sticky>
		<view class="custom_container">
			<u-gap height="28"></u-gap>


			<u-empty v-if="nftList.length === 0" mode="favor" text="啥都还没买过"></u-empty>
			<u-list v-else class="nfts_list" height="calc(100vh - 200px)">
				<!-- @scrolltolower="getTotalList" -->
				<template v-for="(nft, i) in nftList">
					<u-list-item :key="i" v-if="!(nft.countdown <= 0 && nft.status === 1)">
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
										<u-text :text="`${$u.timeFormat(nft.create_time, 'yyyy/mm/dd')}`" color="#9A82DB"
											size="12" align="right"></u-text>
									</view>
								</view>
							</view>
							<view class="card_body"
								style="display: flex;align-items: center;justify-content: space-between;">


								<u-text v-if="!nft.countdown" :text="nft.status_text" color="#9A82DB" size="12">
								</u-text>


								<view v-else style="flex: auto;display:flex;align-items: center;">
									<text style="color: #fff;font-size: 12px;">{{nft.status_text}}</text>
									<u-count-down  @finish="countdownFinished('total',i)"
										:time="nft.countdown*1000" format="HH:mm:ss"></u-count-down>
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
				</template>
				<u-divider lineColor="#fff" text="没有更多了..." :hairline="true">
				</u-divider>
			</u-list>




		</view>
	</view>
</template>

<script>
	// import debounce from 'lodash.debounce'
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

				typ: 0,

				nftList: [],
			};
		},
		onLoad() {
			this.getTotalList(0);
		},
		methods: {
			tabClickHandler(e) {
				const {
					index
				} = e;
				this.activeTab = index;
				this.getTotalList(index);

			},

			getTotalList(i) {



				orderListReq({
					typ: i,

					// page_size: pageSize_t,

				}).then(res => {
					console.log(res);
					this.nftList = res.filter(item => {
						return !(item.countdown <= 0 && item.status === 1)
					});
					// const {

				}).catch(() => {

				});
			},

			toPayHandler(pay_url) {
				// debounce(() => {
				// 	plus.runtime.openWeb(pay_url, () => {
				// 		// 调用失败 callback
				// 		uni.$u.toast('调起支付失败，请稍后再试！');
				// 	});
				// }, 150)
				
				plus.runtime.openWeb(pay_url, () => {
					// 调用失败 callback
					uni.$u.toast('调起支付失败，请稍后再试！');
				});
			},

			countdownFinished(type, index) {
				this.nftList = this.nftList.filter(item => {
					return item.id !== this.nftList[index].id
				})
				// if (type === 'total') {
				// 	const {
				// 		nftList_t
				// 	} = this

				// 	let item = nftList_t[index];
				// 	item = {
				// 		...item,
				// 		countdown: 0,
				// 		status: 0
				// 	}
				// 	this.nftList_t.splice(index, 1)
				// } else {
				// 	const {
				// 		nftList_l
				// 	} = this;

				// 	let item = nftList_l[index];
				// 	item = {
				// 		...item,
				// 		countdown: 0,
				// 		status: 0
				// 	}
				// 	this.nftList_l.splice(index, 1)
				// }
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
