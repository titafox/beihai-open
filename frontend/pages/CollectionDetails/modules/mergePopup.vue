<template>
	<u-popup
		round="12"
		bgColor="#21242D"
		:customStyle="{overflow: 'hidden'}"
		:show="show" 
		@close="close"
	>
		<view class="main">
			<text class="label">选择合成活动</text>
			<text class="tips">请选择一下活动中的一个</text>
			<u-radio-group 
				v-model="value"
				placement="column"
				iconPlacement="right"
			>
				<u-radio 
					labelColor="#ffffff"
					activeColor="#01C5A1" 
					v-for="(item, index) in list"
					:label="item.title" 
					:key="index"
					:name="item.id"
				/>
			</u-radio-group>
		</view>
		<view class="footer">
			<text class="button" @click="goMerge">下一步</text>
		</view>
	</u-popup>
</template>

<script>
	import { getCompositeList } from '@/api/merge.js'
	
	export default {
		name: 'mergePopup',
		data () {
			return {
				show: false,
				value: 0,
				list: ''
			}
		},
		methods: {
			close () {
				this.show = false
			},
			async open () {
				await this.getCompositeList()
				this.show = true
			},
			goMerge () {
				if (this.value === 0) {
					uni.$u.toast('请选择合成活动！')
					return
				}
				
				// uni.$u.route('pages/CollectionDetails/merge', {
				// 	id: this.value
				// })
				
				this.$emit('next', this.value)
			},
			async getCompositeList () {
				const data = await getCompositeList()
				this.list = data
			}
		}
		
	}
</script>

<style lang="less" scoped>
	/deep/ .u-popup__content {
		width: auto !important;
	}
	
	/deep/ .u-radio {
		margin-bottom: 30rpx;
	}
	
	.main {
		display: flex;
		flex-direction: column;
		padding: 54.35rpx;
		background-color: #21242D;
		
		.label {
			font-size: 18px;
			margin-bottom: 10rpx;
		}
		
		.tips {
			margin-bottom: 80rpx;
			font-size: 12px;
			color: #888888;
		}
		
		.list {
			.item {
				margin-bottom: 20rpx;
				
				text {
					font-size: 14px;
					letter-spacing: 0.1px;
				}
			}
		}
	}
	
	.footer {
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
</style>