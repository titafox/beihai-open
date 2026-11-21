<template>
	<view>
		<u-navbar :titleStyle="{'color': '#fff'}" bgColor="#21005D" title="我的寄售" :placeholder="true" class="nav_far"
			:autoBack="true" leftIconColor="#151517">
			<u-icon name="arrow-leftward" color="#fff" slot="left"></u-icon>
		</u-navbar>
		
		<view class="custom_container">
			<u-empty v-if="nftList_t.length === 0" mode="favor" text="啥都还没买过"></u-empty>
			<u-list v-else @scrolltolower="getTotalList" class="nfts_list">
				<u-list-item v-for="(nft,i) in nftList_t" :key="i">
					<view class="card">
						<view class="card_header">
							<u-avatar :src="nft.image" shape="square"></u-avatar>
							<view style="flex: auto;margin-left: 28rpx;">
								<view style="display: flex;">
									<u-text :text="nft.name" color="#fff" size="18"></u-text>
									<!-- <u-text :text="`支付金额：¥ ${nft.money}`" color="#fff" size="12" align="right"></u-text> -->
								</view>
								<view style="display: flex;">
									<u-text :text="nft.no" color="#51EC44" size="14"></u-text>
									<u-text :text="`${$u.timeFormat(nft.create_time, 'yyyy/mm/dd')}`" color="#9A82DB" size="12" align="right"></u-text>
								</view>
							</view>
						</view>
						<view class="card_body">
							<u-text :text="nft.des" color="#9A82DB" size="12"></u-text>
							
<!-- 							<u-text :text="`已成功转让给${nft.mobile}`" color="#9A82DB" size="12"></u-text>
 -->						</view>
					</view>
					<!-- <NftCard class="nft_card" :info="nft" :isHandsel="true" @handsel="handselInfo" /> -->
				</u-list-item>
				<u-divider lineColor="#fff" text="没有更多了..." :hairline="true" v-show="pageSize_t*page_t >= count_t">
				</u-divider>
			</u-list>
			
		</view>
		
	</view>
</template>

<script>
	import { handselListReq } from '@/api/index.js';
	export default {
		data() {
			return {
				
				// 全部
				page_t: 1,
				pageSize_t: 20,
				isReq_t: false,
				count_t: 21,
				nftList_t: [],
				
			};
		},
		onShow() {
			this.getTotalList();	
		},
		methods: {
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
				handselListReq({
					page: page_t,
					page_size: pageSize_t,
				}).then(res => {
					// const {
					// 	result,
					// 	count
					// } = res;
					// this.count_t = count;
					const result = res;
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
		padding: 24rpx 20rpx ;
		border-bottom: 1px solid #381E72;
		display: flex;
	}

	.card .card_body {
		padding: 32rpx 28rpx ;
	}
</style>