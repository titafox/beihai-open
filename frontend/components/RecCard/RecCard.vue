<template>
	<view class="nft-card_container" @click="navigatorToDetail">
		<view class="footer">
			<view style="display: flex;">
				<u-avatar :src="info.author_avatar" fontSize="22" size="36"></u-avatar>
				<u-text class="author" :text="`${info.author_name}`" size="12" color="#EADDFF"></u-text>
			</view>
			<view style="display: flex;" v-if="info.countdown">
				<u-icon name="clock" size="16" color="#fff"></u-icon>
				<u-count-down :time="info.countdown*1000" format="HH:mm:ss"></u-count-down>
			</view>
			<view style="display: flex;" v-else>
				<!-- <u-icon name="clock" size="16" color="#fff"></u-icon> -->
				<text>{{ info.sold_out ? '已售罄' : '出售中' }}</text>
			</view>

			<!-- <u-button v-if="info.sold_out" text="已售罄" color="#21242D" shape="circle" size="mini"
				:customStyle="{'width':'40px','marginRight':0}">
			</u-button> -->
		</view>
		<u-gap height="10"></u-gap>
		<u-image bgColor="#6F55A2" height="100%" width="100%" :src="info.image" :fade="true" duration="450"></u-image>
		<view class="all_nft-info_new">
			<view class="info_container">
				<view class="top_container">
					<u-text class="title" :text="info.name" size="18" color="#fff" lines="1"></u-text>
					<u-text class="price" :text="`¥ ${info.price}`" align="right" size="16"></u-text>
				</view>
			</view>
		</view>
	</view>
</template>

<script>
	// import Like from '../Like/Like.vue'
	export default {
		name: 'RecCard',
		props: ['info', 'home'],
		// components: {
		// 	Like
		// },
		data() {
			return {};
		},
		computed: {
			likeInfo() {
				const {
					likes,
					liked,
					id
				} = this.info;
				return {
					isLike: liked,
					count: likes,
					id
				};
			}
		},
		methods: {
			navigatorToDetail() {
				const {
					id,
					chain_status
				} = this.info;
				uni.navigateTo({
					// url: '/pages/nftDetail/nftDetail?id=' + id + "&chain_status=" + chain_status
					url: '/pages/nftDetail/nftDetail?id=' + id 
				});
			}
		}
	};
</script>

<style>
	.nft-card_container {
		/* background: #FEFEFE;
		border: 1px solid #FEFEFE; */
		/* background: #ffffff; */
		background-color: #180240;
		border: 1px solid #180240;
		border-radius: 12px;
		/* padding: 16rpx;
		box-shadow: 0px 5px 5px #a2a2b621,
			-5px 5px 10px #a2a2b621,
			5px 5px 10px #a2a2b621; */
		/* margin-left: 10rpx;
		margin-right: 10rpx; */
		position: relative;
		overflow: hidden;
		padding: 40rpx;
		box-sizing: border-box;
	}

	.drop_tag {
		position: absolute;
		right: -50rpx;
		top: 10rpx;
		width: 150rpx;
		font-size: 12px;
		text-align: center;
		background-color: #bc42ff;
		color: #ffffff;
		padding: 10rpx;
		/* border-radius: 10rpx; */
		transform: rotate(45deg);
		z-index: 12
	}

	.all_nft-info {
		padding: 20rpx;
		position: absolute;
		width: calc(100% - 68rpx);
		left: 0;
		bottom: -5px;
		padding: 36rpx;
		background: linear-gradient(180deg, rgba(0, 0, 0, 0) 0%, #000000 53.12%);
		border-radius: 0px 0px 8px 8px;
	}

	.info_container {
		background: #180240;
		border-radius: 8px;
		padding: 36rpx 0 0;
		box-sizing: border-box;
		width: 100%;

	}

	/* .nft-card_container >>> .u-image__image {
		width: 284rpx!important;
		height: 336rpx!important;
	} */

	.nft-card_container .title>>>span {
		font-weight: 600;
		line-height: 22px;
		letter-spacing: 0.005em;
		color: #ffffff;
	}

	.nft-card_container .author>>>span {
		font-weight: 400;
		line-height: 20px;
		letter-spacing: 0.005em;
		margin-left: 20rpx;
	}

	.nft-card_container .price>>>span {
		font-weight: 600;
		line-height: 20px;
		letter-spacing: 0.005em;
		background: linear-gradient(98.46deg, #00F0FF -0.18%, #51EC44 106.55%);
		-webkit-background-clip: text;
		-webkit-text-fill-color: transparent;
		background-clip: text;
		text-fill-color: transparent;

	}

	.footer {
		display: flex;
		align-items: center;
		justify-content: space-between;
	}

	.footer>>>.u-count-down__text {
		color: #ffffff;
		font-size: 14px;
		margin-left: 20rpx;
	}


	.top_container {
		display: flex;
		justify-content: space-between;
		align-items: center;
		/* margin-bottom: 20rpx; */
	}
</style>
