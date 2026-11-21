<template>
	<u-popup
		round="12"
		bgColor="#21242D"
		:customStyle="{overflow: 'hidden'}"
		:show="show" 
		@close="close"
	>
		<view class="main">
			<text class="label">选择条件{{ conditionNum + 1 }} {{ condition.name }}的藏品</text>
			<z-paging
				ref="paging" 
				:fixed="false"  
				v-model="list" 
				@query="query"
				style="width: 100%;height: 380rpx;"
			>
			<view class="content">
				<view class="item" v-for="(item, index) in list" :key="index">
					<view class="item-wrapper" :class="{selected: selectedList.includes(item.id)}" @click="select(item.id)">
						<view class="cover">
							<image :src="item.image" mode="aspectFill" />
							<!-- <text class="tag">原料12</text> -->
						</view>
						<view class="mate">
							<text>{{ item.name }}</text>
							<text>#{{ item.work_no }}</text>
						</view>
					</view>
				</view>
			</view>
			</z-paging>
		</view>
		<view class="footer">
			<text class="button" @click="change">确认选择</text>
		</view>
	</u-popup>
</template>

<script>
	import { getCompositeMaterial } from '@/api/merge.js'
	
	export default {
		name: 'rawListPopup',
		data () {
			return {
				show: false,
				id: 0,
				wid: 0,
				list: [],
				selectedList: [],
				conditionNum: 0,
				condition: '',
				selectIds: []
			}
		},
		methods: {
			close () {
				this.show = false
			},
			async open () {
				this.show = true
				if (this.condition.wid !== this.wid) {
					this.selectedList = []
				}
				this.wid = this.condition.wid
			},
			async getCompositeMaterial (pageNum) {
				const data = await getCompositeMaterial({
					id: this.condition.wid,
					page: pageNum
				})
				
				if (pageNum === 1) {
					return data.fail_card.concat(data.list)
				} else {
					return data.list
				}
			},
			async query (pageNum) {
				const list = await this.getCompositeMaterial(pageNum)
				this.$refs.paging.complete(list.filter(item => !this.selectIds.includes(item.id)))
			},
			select (id) {
				if (this.selectedList.includes(id)) {
					this.selectedList = this.selectedList.filter(item => item !== id)
				} else {
					this.selectedList.push(id)
				}
			},
			change () {
				if (this.selectedList.length !== this.condition.nums) {
					uni.$u.toast(`请选择${this.condition.nums}个原料`)
					return
				}
				
				this.$emit('change', this.conditionNum, this.list.filter(item => {
					return this.selectedList.includes(item.id)
				}))
				this.show = false
				this.selectedList = []
			}
		}
		
	}
</script>

<style lang="less" scoped>	
	/deep/ .u-popup__content {
		width: auto !important;
	}
	
	.main {
		display: flex;
		flex-direction: column;
		padding: 54.35rpx;
		background-color: #21242D;
		
		.label {
			font-size: 18px;
			margin-bottom: 40rpx;
		}
		
		.content {
			display: flex;
			flex-wrap: wrap;
			margin: 0 -5px;
			
			.item {
				position: relative;
				display: flex;
				flex-direction: column;
				align-items: center;
				justify-content: center;
				margin-bottom: 28rpx;
				flex-basis: 33.333%;
				padding-bottom: 39%;
				
				.selected {
					border: 2px solid #D0BCFF;
				}
				
				.item-wrapper {
					position: absolute;
					left: 5px;
					right: 5px;
					top: 0;
					bottom: 5px;
					padding: 24rpx;
					background: #180240;
					border-radius: 20px;
					
					.cover {
						position: relative;
						border-radius: 16px;
						overflow: hidden;
						margin-bottom: 10rpx;
						aspect-ratio: 1;
						
						image {
							display: block;
							width: 100%;
							height: 100%;
							border-radius: 16px;
						}
						
						.tag {
							position: absolute;
							left: 0;
							top: 0;
							padding: 5px 8px;
							background: #381E72;
							border-radius: 16px 0px;
							font-size: 10px;
							/* #ifdef H5 */
							transform: scale(0.9);
							transform-origin: left top;
							/* #endif */
						}
					}
					
					.mate {
						display: flex;
						// flex-direction: column;
						align-items: center;
						justify-content: space-between;
						padding: 0 5rpx;
						
						text:first-child {
							font-size: 12px;
						}
						
						text:nth-child(2) {
							color: #888888;
							font-size: 10px;
							/* #ifdef H5 */
							transform: scale(0.9);
							transform-origin: left center;
							/* #endif */
						}
					}
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