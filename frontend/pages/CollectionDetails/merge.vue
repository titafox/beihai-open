<template>
	<view class="profile_container">
		<u-navbar 
			:titleStyle="{'color': '#fff'}" 
			bgColor="#21005D" 
			title="选择藏品" 
			:placeholder="true" 
			class="nav_bar"
			:autoBack="true" 
			leftIconColor="#151517"
		>
			<u-icon name="arrow-leftward" color="#fff" slot="left"></u-icon>
		</u-navbar>
		
		<u-gap height="20"></u-gap>
		
		<view class="container">
			<view class="header">
				<text>{{ detail.description }}</text>
				<text>提醒：寄售中的藏品无法参与合成。</text>
			</view>
			
			<view class="condition" v-for="(item, index) of detail.material_data" :key="index">
				<view class="header">
					<text class="label">条件{{ index + 1 }} {{ item.name }}（请选择{{ item.nums }}个藏品）</text>
					<text class="button" @click="resetSelect(item, index)">重新选择</text>
				</view>
				<view class="content">
					<template v-if="selectedRaw[index].length !== 0">
						<view class="item" v-for="(item, itemIdx) in selectedRaw[index]" :key="itemIdx">
							<view class="item-wrapper">
								<view class="cover">
									<image :src="item.image" mode="aspectFill" />
									<text class="tag">原料{{ itemIdx + 1 }}</text>
								</view>
								<view class="mate">
									<text>{{ item.name }}</text>
									<text>#{{ item.work_no }}</text>
								</view>
							</view>
						</view>
					</template>
					<template v-else>
						<view class="item" v-for="itemIdx of item.nums" :key="itemIdx">
							<view class="item-wrapper empty" @click="selectRaw(item, index)">
								<u-icon name="plus" size="50" color="#EADDFF" bold />
							</view>
						</view>
					</template>
				</view>
			</view>
		
			<view class="footer">
				<view class="left">
					<text>已选择 <text>{{ selectArr.length }}</text> 个</text>
					<text>合成失败几率：{{ detail.failure_rate }}%</text>
				</view>
				<view class="right">
					<text @click="doMerge">立即合成</text>
				</view>
			</view>
		</view>
		
		<rawListPopup 
			ref="rawListPopup"
			@change="changeSelect"
		/>
	</view>
</template>

<script>
	import rawListPopup from './modules/rawListPopup.vue'
	import { getCompositeList, doComposite } from '@/api/merge.js'
	
	export default {
		components: {
			rawListPopup
		},
		data() {
			return {
				id: 0,
				collectionId: 0,
				detail: '',
				selectedRaw: [],
				selectArr: [],
				postData: null,
				selectRawIds: []
			}
		},
		onLoad (params) {
			this.id = parseInt(params.id)
			this.collectionId = parseInt(params.collectionId)
			this.init()
		},
		methods: {
			init () {
				this.getCompositeDetail()
			},
			async getCompositeDetail () {
				const data = await getCompositeList()
				this.detail = data.filter(item => item.id === this.id)[0]
				this.detail.material_data.map(() => {
					this.selectedRaw.push([])
				})
			},
			selectRaw (item, index) {
				const selectRawIds = this.selectRawIds.filter(item => !this.selectedRaw[index].map(item => item.id).includes(item))
				this.$refs.rawListPopup.selectIds = selectRawIds
				this.$refs.rawListPopup.conditionNum = index
				this.$refs.rawListPopup.condition = item
				this.$refs.rawListPopup.open()
			},
			resetSelect (item, index) {
				this.selectRaw(item, index)
			},
			async doMerge () {
				try {
					this.checkCondition()
				} catch (e) {
					uni.$u.toast(e.message)
					return
				}
				
				try {
					const data = await doComposite({
						composite_id: this.id,
						data: this.postData
					})
					uni.$u.route({
						url: 'pages/CollectionDetails/mergeResult',
						params: data
					})
				} catch (e) {
					//TODO handle the exception
				}
			},
			checkCondition () {
				this.detail.material_data.forEach((item, index) => {
					if (this.selectedRaw[index].length < item.nums) {
						throw Error(`条件${index + 1}未满足${item.nums}个原料`)
					}
				})
			},
			generatePostData (arr) {
				let data = []
				const wid = arr[0].wid
				const item = { wid, ids: [] }
				let newArr = []
				for (let i in arr) {
					if (arr[i].wid === wid) {
						item.ids.push(arr[i].id)
					} else {
						const index = parseInt(i)
						newArr = newArr.concat(arr.slice(index, index + 1).filter(item => !newArr.includes(item)))
					}
				}
				
				if (newArr.length > 0) {
					data = data.concat(this.generatePostData(newArr))
				}
				data.push(item)
				return data
			},
			changeSelect (index, list) {
				this.selectedRaw[index] = list
				this.selectArr = this.selectedRaw.reduce((prev, cur) => {
					return prev.concat(cur)
				}, [])
				this.selectRawIds = this.selectArr.map(item => item.id)
				this.postData = this.generatePostData(this.selectArr)
				this.$forceUpdate()
			}
		}
	}
</script>

<style lang="less" scoped>
	.container {
		padding: 0 24rpx 200px;
		
		& > .header {
			display: flex;
			flex-direction: column;
			line-height: 24px;
			
			text:fist-child {
				font-size: 16px;
			}
			
			text:nth-child(2) {
				font-size: 16px;
				color: #D0BCFF;
			}
		}
	
		.condition {
			padding: 40rpx 0;
			border-bottom: 1px solid #381E72;
			
			& > .header {
				display: flex;
				justify-content: space-between;
				align-items: center;
				padding: 0 5px;
				margin-bottom: 40rpx;
				
				.button {
					padding: 6px 22px;
					background: #21242D;
					border-radius: 20px;
					color: #01C5A1;
					font-size: 14px;
				}
			}
			
			.content {
				display: flex;
				flex-wrap: wrap;
				// justify-content: space-between;
				
				.item {
					position: relative;
					display: flex;
					flex-direction: column;
					align-items: center;
					justify-content: center;
					margin-bottom: 28rpx;
					flex-basis: 33.333%;
					padding-bottom: 39%;
					
					
					.item-wrapper {
						position: absolute;
						left: 5px;
						right: 5px;
						top: 0;
						bottom: 0;
						padding: 24rpx 24rpx 0;
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
							flex-direction: column;
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
				
					.item-wrapper.empty {
						display: flex;
						align-items: center;
						justify-content: center;
						padding: 0;
					}
				}
			}
		}
	
		.footer {
			position: fixed;
			bottom: 0;
			left: 0;
			right: 0;
			display: flex;
			align-items: center;
			justify-content: space-between;
			padding: 18px 30px;
			background: #0E0D0D;
			
			.left {
				display: flex;
				flex-direction: column;
				
				text:first-child {
					font-size: 14px;
					
					text {
						margin: 0 5px;
						font-size: 28px;
					}
				}
				
				text:nth-child(2) {
					color: #888888;
					font-size: 12px;
					/* #ifdef H5 */
					transform: scale(0.9);
					transform-origin: left center;
					/* #endif */
				}
			}
			
			.right text {
				padding: 8px 20px;
				background: #21242D;
				color: #7B61FF;
				font-size: 18px;
				border-radius: 99px;
			}
		}
	}
</style>
