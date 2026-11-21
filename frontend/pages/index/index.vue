<template>
	<view class="padding_0_24">
		<u-gap height="44"></u-gap>

		<view class="block-title_container">
			<u-row customStyle="margin-bottom: 10px">
				<u-col span="11">
					<view class="bh-layout bg-purple-light">

						<u-text text="北海·NORTH" size="21" color="#EADDFF"></u-text>
						<u-notice-bar :icon="radio" bgColor="transparent" :text="notice" color="#6750A4"
							direction="column" @click="radioClick"></u-notice-bar>
					</view>
				</u-col>
				<u-col span="1">
					<u-image :src="avatar" :lazy-load="true" shape="circle" width="30" height="30" align="right" @click="gonft">
					</u-image>
				</u-col>
			</u-row>
		</view>
		<u-gap height="6"></u-gap>
		<view class="ads">
			<view class="block-title_container">
				<u-swiper :list="bannerList" :height="150" imgMode="heightFix" indicator indicatorMode="line" circular
					@click="bannerClickHanlder"></u-swiper>
			</view>

			<u-gap height="21"></u-gap>
			<view class="new_btn-group">

				<view class="new_btn" @click="linkToNotice">
					<view style="display: flex;align-items: center;justify-content: space-between;">
						<u-icon name="/static/notice.png" size="25"></u-icon>
						<u-icon name="arrow-right" color="#D3C5EE"></u-icon>
					</view>
					<u-gap height="30"></u-gap>
					<u-text text="官方公告" color="#EADDFF" size="14"></u-text>
				</view>

				<view class="new_btn" @click="linkToArtist">
					<view style="display: flex;align-items: center;justify-content: space-between;">
						<u-icon name="/static/artist.png" size="25"></u-icon>
						<u-icon name="arrow-right" color="#D3C5EE"></u-icon>
					</view>
					<u-gap height="30"></u-gap>
					<u-text text="ARTIST" color="#EADDFF" size="14"></u-text>
				</view>
			</view>

			<!-- 
			<view class="block-title_container btn_group">
				<u-button color="#D0BCFF" shape="circle" text="官方公告" @click="linkToNotice"></u-button>
				<u-button color="#D0BCFF" shape="circle" text="ARTIST" @click="linkToArtist"></u-button>
			</view> -->
		</view>

		<u-gap height="15"></u-gap>

		<view class="block-title_container">
			<view class="nfts_list">
				<template v-for="(item, index) in recommList">
					<!-- 	<RecCard v-if="index === 0" :info="item" :key="index" class="rec_card" />
					<NftCard v-else :info="item" :key="index" class="nft_card" /> -->
					<RecCard :info="item" :key="index" class="rec_card" />
				</template>
			</view>
		</view>
		<Tabbar />
	</view>

</template>

<script>
	import Tabbar from '../../components/Tabbar/Tabbar.vue';
	import NftCard from '../../components/NftCard/NftCard.vue';
	import RecCard from '../../components/RecCard/RecCard.vue';

	import {
		bannerListReq,
		saleListReq,
		recommListReq,
		allNftsReq,
		radioInfoReq,
		accountInfoReq
	} from '../../api/index.js';
	
	// import {
	// 	mapMutations
	// } from 'vuex';
	export default {
		components: {
			Tabbar,
			NftCard,
			RecCard
		},
		data() {
			return {
				bannerList: [],
				saleList: [],
				recommList: [],
				noticeList: [],
				radio: require('@/static/pi.png'),
				avatar: '/static/fake_avatar.png'
			};
		},
		computed: {
			notice() {
				let list = [];
				if (this.noticeList.length > 0) {
					this.noticeList.map(item => {
						list.push(item.title);
					});

				}

				return list;
			}
		},
		onShow() {
			uni.hideTabBar();
			this.getBannerList();
			// this.getSaleList();
			this.getRecommList();

			this.getNoticeList();


			// this.getUserInfo();
		},
		methods: {
			gonft(){
			   uni.navigateTo({
			   	url: '/pages/nfts/nfts'
			   });	
			 	
			},
			
			getUserInfo() {
				accountInfoReq().then(res => {
					// const {
					// 	avatar
					// } = res;
					// this.avatar = avatar;
					console.log(res);

					const {
						avatar,
						nickname,
						id_card_num,
						bsn_address,
						uid,
						role,
						is_wallet,
						id,
						is_handsel,
						authentication,
					} = res;
					this.avatar = avatar;
					let userinfo = {
						"avatar": avatar,
						"nickname": nickname,
						"uid": uid,
						"bsn_address": bsn_address,
						"is_wallet":is_wallet,
						"id":id,
						"is_handsel":is_handsel,
						"authentication":authentication
					};
					this.$storea.commit('updataUser', userinfo);
					
									
									
					
					// store.state.user.token
					
					

				});

			},
			getBannerList() {
				bannerListReq().then(res => {

					if (res) {
						let bannerList = [];
						res.map(banner => {
							bannerList.push({
								url: banner.image,
								id: banner.id,
								type: banner.banner_type,
								banner_type: banner.banner_type,
								val: banner.val
							});
						});
						if (bannerList.length > 0) this.bannerList = bannerList;
					}
				});
			},
			getSaleList() {
				saleListReq().then(res => {
					if (res) {
						this.saleList = res;
					}
				});
			},
			getRecommList() {
				allNftsReq({
					page: 1,
					page_size: 10,
					sort: '-create_time'
				}).then(res => {
					// console.log(res)
					const {
						result
					} = res;
					if (result && result.length > 0) {
						this.recommList = result;
					}
				});
			},
			bannerClickHanlder(e) {
				const {
					bannerList
				} = this;

				const {
					banner_type,
					val
				} = bannerList[e];
				//0 商品 1 公告 2 url（直接打开）
				// banner_type
				// val string 

				if (banner_type === 0) {
					uni.navigateTo({
						url: '/pages/nftDetail/nftDetail?id=' + val
					});
				} else if (banner_type === 1) {
					uni.navigateTo({
						url: '/pages/notice/noticeDetail?id=' + val
					});
				} else if (banner_type === 2) {
					plus.runtime.openWeb(val);
				} else {
					// 只是展示图片

				}

			},
			getNoticeList() {
				radioInfoReq().then(res => {
					this.noticeList = res;
				});
			},
			radioClick(e) {
				const {
					id
				} = this.noticeList[e];
				uni.navigateTo({
					url: '/pages/notice/noticeDetail?id=' + id
				});
			},
			linkToNotice() {
				uni.navigateTo({
					url: '/pages/notice/notice'
				});
			},
			linkToArtist() {
				uni.navigateTo({
					url: '/pages/artist/artist'
				});
			}
		}
	};
</script>

<style scoped>
	.padding_0_24 {
		padding: 0 28rpx;
	}

	.block-title_container .block-title>>>span {
		font-weight: 600;
		line-height: 26px;
		text-align: center;
		letter-spacing: 0.005em;
	}

	.block-title_container .block-title {
		padding: 26rpx 28rpx;
		width: auto;
	}

	.new_btn-group {
		display: flex;
		align-items: center;
		justify-content: space-between;
	}

	.new_btn {
		display: flex;
		flex-direction: column;
		background: #381E72;
		border-radius: 20px;
		padding: 32rpx 28rpx;
		flex: auto;
	}

	.new_btn-group .new_btn:first-child {
		margin-right: 15px;
	}

	.new_btn-group .new_btn:last-child {
		margin-left: 15px;
	}

	.mini_card {
		display: flex;
		background: #FEFEFE;
		border: 1px solid #EBEBF0;
		border-radius: 24px;
		padding: 16rpx;
		margin-right: 26rpx;
		min-width: 143px;
	}

	.mini_card .u-avatar {
		margin-right: 10px;
	}

	.mini_card .main_title>>>span,
	.mini_card .sub_title>>>span {
		font-weight: 500;
		line-height: 20px;
		letter-spacing: 0.005em;
	}

	.block-title_group {
		display: flex;
		padding: 26rpx 28rpx;
	}

	.block-title_group .u-text>>>span {
		text-align: center;
		letter-spacing: 0.005em;
	}

	.block-title_group .u-text:first-child>>>span {
		line-height: 26px;
		font-weight: 600;
	}

	.block-title_group .u-text:last-child>>>span {
		line-height: 22px;
		font-weight: 500;
	}

	.block-title_group .u-text:last-child {
		justify-content: flex-end !important;
	}

	.block-title_group .u-text:last-child>>>.u-icon__icon {
		color: #bc42ff !important;
	}

	.nft_card {
		margin-bottom: 48rpx;
	}

	.nft_card>>>.u-transition {
		width: 284rpx;
		height: 326rpx;
		margin: auto;
	}

	.nfts_list {
		display: flex;
		flex-direction: row;
		flex-wrap: wrap;
		justify-content: space-between;
		padding-top: 20rpx;
	}

	.ads {
		background: #180240;
		border-radius: 12px;
		padding: 30rpx;
		min-width: 143px;
	}

	.rec_card {
		margin-bottom: 48rpx;
	}

	.rec_card>>>.u-transition {
		/* width: 684rpx;
		height: 620rpx; */
		width: 620rpx;
		height: 620rpx;
		border-radius: 12px;
		margin: auto;
	}



	.btn_group {
		display: flex;
	}

	.btn_group .u-button {
		width: 30%;
		margin-left: 0;
		margin-right: 64rpx;
	}

	.btn_group>>>.u-button__text {
		color: #381E72;
	}

	.bh-layout {

		border-radius: 4px;
	}
</style>
