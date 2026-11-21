<template>
	<view class="m-flex-w sp-m">
		<view class="sp-m-v1" v-for="(item, index) in list" :key="index" @click="click(item.id)">
			<view class="m-flex-sb  sp-m-v">
				<view class="m-flex">
					<u-avatar size="22" :src="item.author_avatar" class="avatar"></u-avatar>
					
					<view class="sp-m-t2">{{item.author_name}}</view>

				</view>
				
				<view class="sp-m-t2 m-color-m2" style="margin-right: 16rpx;">#{{item.work_no}}</view>
				<!-- <u--text :text="item.price"></u--text> -->
			</view>
			
			<view class="m-flex-c">
				<!-- <u--image :src="item.image" width="150" height="150" ></u--image> -->
				<view class="image">
					<image class="sp-m-img" :src="item.image" ></image>
					<text class="tag" v-if="item.status === 0 && item.is_synthetic_material === 1">可合成</text>
				</view>
			</view>
			<view class="m-flex-sb sp-m-v">
			  <view class="sp-m-t3">{{item.name}}</view>
			  <view class="m-flex-w sp-l">
				 <view class="sp-m-t5" v-if="item.chain_status==0">上链中</view>
				 <view class="sp-m-t4" >¥{{item.price}}</view> 
				 
			  </view>
			  
			  
			
			</view> 
			
						
		</view>

	</view>
</template>

<script>
	export default {
		name: 'Purchased',
		data() {
			return {

			}
		},

		props: {
			// opuslist: [],
			list:{
				type:Array,
				default:[],
			}
		},
		watch: {

		},
		computed: {

		},
		methods: {
           click(i){
			
				uni.navigateTo({
					url: '/pages/CollectionDetails/collectiondetails?id='+i
				});
			
			 // this.$emit("ItemCall",i+1)  
		   },

		},

	}
</script>

<style lang="scss" scoped>
	.sp-m{
		padding:16rpx;
	}
	.sp-m-v{
		width: 680rpx;
		padding: 16rpx;
		align-items: center;
		
	}
	.sp-m-v1{
		background: #180240;
		border-radius: 20rpx;
		padding: 16rpx;
		width: 680rpx;
		
		margin-bottom: 24rpx;
	}
	.sp-m-t1{
		font-size: 16rpx;
	}
	.sp-m-t3{
		font-size: 36rpx;
		padding: 8rpx 8rpx 8rpx 0;
	}
	.sp-m-t2{
		font-size: 32rpx;
		padding-right: 22rpx;
	}
	.sp-m-t4{
		font-size: 32rpx;
		color: #7B61FF;
		width: 120rpx;
		text-align: right;
	}
	.sp-m-img{
		width: 650rpx;
		height: 650rpx;
		border-radius: 40rpx;
	}
	.avatar {
		margin-right: 6rpx;
	
	
	}
	.sp-l{
		display: flex;
		justify-content: right;
		margin-right: 16rpx;
		width: 150rpx;
	}
	.sp-m-t5{
		font-size: 25rpx;
		width: 120rpx;
		text-align: right;
		
		color: #888888;
		// margin-right: 8rpx;
	}
	.sp-m-t6{
		font-size: 24rpx;
		padding: 8rpx 0 8rpx 0;
		color: #7B61FF;;
	}
	
	.image {
		position: relative;
		
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
</style>
