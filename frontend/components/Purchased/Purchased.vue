<template>
	<view class="m-flex-wsb sp-m">
		<view class="sp-m-v1" v-for="(item, index) in list" :key="index" @click="click(item)">
			<view class="m-flex-sb sp-m-v">
				<view class="m-flex" style="align-items: center;">
					<u-avatar size="24" :src="item.author_avatar" class="avatar"></u-avatar>
					
					<view class="sp-m-t3" style="margin-left: 5px;">{{item.author_name}}</view>

				</view>
				<!-- <view class="sp-m-t2">¥{{item.price}}</view> -->
				<!-- <u--text :text="item.price"></u--text> -->
			</view>
			
			<view>
				<!-- <u--image :src="item.image" width="150" height="150" ></u--image> -->
                <image class="sp-m-img" :src="item.image" ></image>
			</view>
			<view class="m-flex-sb sp-m-v">
			  <view class="sp-m-t3">{{item.name}}</view>
			  <view class="sp-m-t3 m-color-m2">x{{item.nums}}</view>
			
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
			    console.log(i);
				uni.navigateTo({
					 url: '/pages/nftsDetails/ntfsDetails?id='+i.id+'&name='+i.name
				});
			
			 // this.$emit("ItemCall",i+1)  
		   },

		},

	}
</script>

<style lang="scss" scoped>
	.sp-m{
		padding: 8rpx;
	}
	.sp-m-v{
		width: 310rpx;
		padding: 8rpx;
		
	}
	.sp-m-v1{
		background: #180240;
		border-radius: 20rpx;
		padding: 16rpx;
		width: 310rpx;
		
		margin-bottom: 24rpx;
	}
	.sp-m-t1{
		font-size: 16rpx;
	}
	.sp-m-t3{
		font-size: 24rpx;
		padding: 8rpx 8rpx 8rpx 0;
	}
	.sp-m-t2{
		font-size: 18rpx;
		padding-right: 16rpx;
	}
	.sp-m-img{
		width: 300rpx;
		height: 300rpx;
		border-radius: 30rpx;
	}
	.avatar {
		margin-right: 6rpx;
	
	
	}
</style>
