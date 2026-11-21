<template>
	<view class="container" @click="goDetail()">
		<view class="image">
			<image :src="data.image" mode="aspectFill" />
			<text class="tag" v-if="data.is_synthetic_material === 1">可合成</text>
		</view>
		<view class="mate">
			<view class="top">
				<text class="title">{{ data.name }}</text>
				<!-- <text class="price">¥12.00</text> -->
			</view>
			<view class="bottom">
				<u-avatar :src="data.author_avatar" size="34rpx"/>
				<text class="author">{{ data.author_name }}</text>
				<text class="num">x{{ data.nums }}</text>
			</view>
		</view>
	</view>
</template>

<script>
	export default {
		name: 'item',
		props: {
			data: Object
		},
		methods: {
			goDetail () {
				uni.$u.route({
					url: '/pages/nftsDetails/ntfsDetails',
					params: {
						id: this.data.id,
						name: this.data.name
					}
				})
			}
		}
	}
</script>

<style lang="less" scoped>
	.container {
		padding: 24rpx;
		background: #180240;
		border-radius: 20px;
		
		.image {
			position: relative;
			aspect-ratio: 1;
			overflow: hidden;
			border-radius: 32rpx;
			margin-bottom: 20rpx;
			
			image {
				display: block;
				width: 100%;
				height: 100%;
			}
			
			.tag {
				position: absolute;
				top: 0;
				right: 0;
				padding: 6rpx 24rpx;
				font-size: 28rpx;
				background: #FF3434;
				border-radius: 99px;
			}
		}
	
		.mate {
			padding: 0 10rpx;
			
			.top {
				display: flex;
				justify-content: space-between;
				align-items: center;
				padding-bottom: 20rpx;
				margin-bottom: 20rpx;
				border-bottom: 1px solid #393556;
				
				.title {
					font-size: 28rpx;
					font-weight: 600;
				}
				
				.price {
					font-size: 28rpx;
					font-weight: 500;
					color: #FFD166;
				}
			}
			
			.bottom {
				display: flex;
				align-items: center;
				
				.author {
					margin-left: 14rpx;
					flex-grow: 1;
					font-size: 20rpx;
					/* #ifdef H5 */
					transform: scale(0.833);
					transform-origin: left center;
					/* #endif */
				}
				
				.num {
					color: #888888;
					font-size: 20rpx;
					/* #ifdef H5 */
					transform: scale(0.833);
					transform-origin: left center;
					/* #endif */
				}
			}
		}
	}
</style>