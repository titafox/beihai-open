<template>
	<view class="profile_container">
		<u-navbar 
			:titleStyle="{'color': '#fff'}" 
			bgColor="#21005D" 
			title="合成中" 
			:placeholder="true" 
			class="nav_bar"
			leftIconColor="#151517"
			@leftClick="backMineNfts"
		>
			<u-icon name="arrow-leftward" color="#fff" slot="left"></u-icon>
		</u-navbar>
		
		<u-gap height="20"></u-gap>
		
		<view class="container">
			<view class="cover">
				<view class="image">
					<image :src="image" mode="aspectFill" />
					<view class="loading">
						<image src="../../static/loading.gif" mode="widthFix" />
					</view>
				</view>
				<view class="mate">
					<text>{{ name }}</text>
					<!-- <text>#1234/1000</text> -->
				</view>
			</view>
			
			<view class="content">
				<text>您的藏品正在合成中，预计需要10分钟，您可在 “已购藏品” 中查看进度。如合成失败，将赠送“合成失败权益卡”。</text>
			</view>
			
			<view class="footer">
				<text class="button" @click="backMineNfts">返回我的仓库</text>
			</view>
		</view>
		
	</view>
</template>

<script>
	export default {
		data() {
			return {
				image: '',
				name: ''
			}
		},
		onLoad (params) {
			this.image = params.image
			this.name = params.name
		},
		onBackPress ({ from }) {
			if (from === 'navigateBack') {
				return false
			}
			this.backMineNfts()
			return true
		},
		methods: {
			backMineNfts () {
				uni.navigateBack({
					delta: 4,
					fail (e) {
						console.log(e)
					}
				})
			}
		}
	}
</script>

<style lang="less" scoped>
	.container {
		padding: 0 70rpx;
		
		.cover {
			display: flex;
			flex-direction: column;
			margin-bottom: 50rpx;
			
			.image {
				position: relative;
				overflow: hidden;
				border-radius: 16px;
				aspect-ratio: 1;
				margin-bottom: 30rpx;
				
				image {
					width: 100%;
					height: 100%;
				}
				
				.loading {
					position: absolute;
					top: 0;
					bottom: 0;
					left: 0;
					right: 0;
					display: flex;
					align-items: center;
					justify-content: center;
					padding: 120rpx;
					background-color: rgba(0,0,0,0.6);
					
					image {
						width: 100%;
					}
				}
			}
			
			.mate {
				display: flex;
				flex-direction: column;
				align-items: center;
				
				text:first-child {
					font-size: 20px;
				}
				
				text:nth-child(2) {
					font-size: 16px;
					color: #888888;
					font-weight: 500;
				}
			}
		}
		
		.content {
			text {
				font-size: 16px;
				line-height: 24px;
			}
		}
		
		.footer {
			position: fixed;
			bottom: 0;
			left: 0;
			right: 0;
			display: flex;
			justify-content: center;
			padding: 50rpx 0;
			background: #0E0D0D;
			
			.button {
				padding: 20rpx 140rpx;
				background: #21242D;
				border-radius: 99px;
				color: #01C5A1;
				font-size: 14px;
			}
		}
	}
</style>
