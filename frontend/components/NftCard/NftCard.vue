<template>
	<view class="nft-card_container" @click="navigatorToDetail">
		<!-- <view class="drop_tag" v-if="info.is_air_drop" >
			空投
		</view> -->
		<!-- 	<view class="countdown_container" v-if="info.countdown" >
			<u-icon name="clock" size="16" color="#fff"></u-icon>
			<u-count-down  :time="info.countdown*1000" format="HH:mm:ss"></u-count-down>
		</view> -->
		<view class="header">
			<view style="display: flex;">
				<u-avatar :src="info.author_avatar" fontSize="22" size="19"></u-avatar>
				<u-text class="author" :text="`${info.author_name}`" size="12" color="#EADDFF"></u-text>
			</view>
			<view style="display: flex;" @click.stop="previewImage">
				<u-icon name="/static/full.png" size="16"></u-icon>
			</view>
		</view>
		<u-gap height="16"></u-gap>



		<u-image bgColor="#6F55A2" height="100%" width="100%" :src="info.image" :fade="true" duration="450" :radius="4">
		</u-image>
		<view class="all_nft-info">

			<view style="width: 200rpx;">
				<u-text class="title" :text="`${info.name}${info.name}${info.name}${info.name}${info.name}${info.name}${info.name}`" size="14" color="#fff" lines="1"></u-text>
			</view>
			<!-- <view>
				<u-text class="title" :text="info.name" size="14" color="#fff" lines="1"></u-text>
				<u-text text="344/3500" size="11" color="#6171FB"></u-text>

			</view> -->

			<view style="display: flex;align-items: baseline;justify-content: space-between;">
				<u-text class="price" :text="`¥ ${info.price}`" size="12" color="#fff"></u-text>
				<view class="countdown_container_new" v-if="info.countdown">
					<u-icon name="clock" size="16" color="#fff"></u-icon>
					<u-count-down :time="info.countdown*1000" format="HH:mm:ss"></u-count-down>
				</view>

				<view @click.stop="handselHandler" v-if="isHandsel && info.chain_status == 1">
					<u-button color="#37419C" shape="circle" text="转赠" size="mini"></u-button>
				</view>

				<view v-if="isHandsel && info.chain_status != 1">
					<u-button color="#37419C" shape="circle" text="上链中" size="mini"></u-button>
				</view>
			</view>


		</view>


	</view>
</template>

<script>
	// import Like from '../Like/Like.vue';
	export default {
		name: 'NftCard',
		props: ['info', 'isHandsel'],
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
				const {
					isHandsel
				} = this;
				uni.navigateTo({
					url: '/pages/nftDetail/nftDetail?id=' + id + '&isHandsel=' + isHandsel + '&chain_status=' +
						chain_status
				});
			},
			previewImage() {
				uni.previewImage({
					urls: [this.info.image]
				});
			},
			handselHandler() {
				const {
					id
				} = this.info;
				// this.handsel(id)
				this.$emit('handsel', id);
			}
		}
	};
</script>

<style>
	.nft-card_container {
		/* background: #FEFEFE;
		border: 1px solid #FEFEFE; */
		background: #180240;
		border: 1px solid #180240;
		border-radius: 8px;
		padding: 16rpx;
		/* box-shadow: 0px 5px 5px #a2a2b621,
			-5px 5px 10px #a2a2b621,
			5px 5px 10px #a2a2b621; */
		margin-left: 10rpx;
		margin-right: 10rpx;
		position: relative;
		overflow: hidden;
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
	}

	/* .nft-card_container >>> .u-image__image {
		
		height: 336rpx!important;
	} */

	>>>.u-text__value {
		width: 100%;
	}

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
		margin-top: 15px;
	}

	.nft-card_container .footer {
		display: flex;
	}

	.nft-card_container .footer .price_name>>>span {
		font-weight: 400;
		line-height: 18px;
		letter-spacing: 0.005em;
		color: #D0BCFF;
	}

	.header {
		display: flex;
		justify-content: space-between;
	}

	>>>.u-count-down__text {
		color: #ffffff;
		font-size: 14px;
		margin-left: 20rpx;
	}

	.countdown_container {
		display: flex;
		background: #FF225F;
		padding: 6px 16px;
		position: absolute;
		top: 84rpx;
		right: 16rpx;
		z-index: 99;
		border-top-right-radius: 6px;
		border-bottom-left-radius: 3px;
	}


	.countdown_container_new {
		display: flex;
	}
</style>
