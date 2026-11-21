<template>
	<view class="profile_container">
		<u-navbar bgColor="#21005D" :titleStyle="{'color': '#fff'}" title="Mine" :placeholder="true" class="nav_bar"
			:autoBack="true" leftIconColor="#151517">
			<view slot="right" class="nav_slot_right">
				<u-icon name="/static/profile_icon/share.png" :customStyle="{'marginRight':'20px'}" color="#fff"
					slot="left" @click="shareShow = true"></u-icon>
				<u-icon name="setting" color="#fff" slot="left" @click="link2Profile"></u-icon>
			</view>
		</u-navbar>
		
		
		<view style="display: flex;">
			<u-avatar size="40" :src="profileInfo.avatar" class="avatar"></u-avatar>
			<view>
				<u-text class="person_name" :text="`${profileInfo.nickname}的藏品馆`" size="18" color="#ffffff">
				</u-text>
				<u-text class="intr_info" text="此处查看和我有关的藏品" size="12" color="#D0BCFF"></u-text>
			</view>
			<view @click="link2mint" class="mint">
				<u-icon name="/static/mint.png" size="48"></u-icon>
				<u-text text="铸件" size="12" color="#9A82DB" align="center"></u-text>
			</view>
		</view>
		
		
		<u-gap height="20"></u-gap>
		<u-tabs :activeStyle="activeTabStyle" :inactiveStyle="inactiveTabStyle" :scrollable="false" :list="tablist"
			@click="tabClickHandler" lineColor="transparent" :current="activeTab"></u-tabs>
		<u-empty v-if="nftList_b.length === 0" mode="favor" text="啥都还没买过"></u-empty>
		<u-list v-else @scrolltolower="getBuyList" class="nfts_list" height="calc(100vh - 300px)">
			<u-list-item v-for="(nft,i) in nftList_b" :key="i">
				<NftCard class="nft_card" :info="nft" :isHandsel="true" @handsel="handselInfo" />
			</u-list-item>
			<u-divider lineColor="#fff" text="没有更多了..." :hairline="true" v-show="pageSize_b*page_b >= count_b">
			</u-divider>
		</u-list>
		

		<!-- <view class="profile_container" style="background-color: #21005D;">
		

			<view v-show="activeTab === 2">

				<u-empty v-if="nftList_p.length === 0" mode="favor" text="啥都还没发布过"></u-empty>
				<u-list v-else @scrolltolower="getPublishList" class="nfts_list">
					<u-list-item v-for="nft in nftList_p" :key="nft.id">
						<NftCard class="nft_card" :info="nft" />
					</u-list-item>
					<u-divider lineColor="#fff" text="没有更多了..." :hairline="true" v-show="pageSize_p*page_p >= count_p">
					</u-divider>
				</u-list>
			</view>

			<view v-show="activeTab === 1">
				<u-empty v-if="nftList_l.length === 0" mode="favor" text="啥都还没卖过"></u-empty>
				<u-list v-else @scrolltolower="getLikeList" class="nfts_list">
					<u-list-item v-for="nft in nftList_l" :key="nft.id">
						<NftCard class="nft_card" :info="nft" />
					</u-list-item>
					<u-divider lineColor="#fff" text="没有更多了..." :hairline="true" v-show="pageSize_l*page_l >= count_l">
					</u-divider>
				</u-list>
			</view> 
		</view>-->

		<u-popup :show="handselShow" @close="toggleHandsel" :closeOnClickOverlay="true">


			<view style="background-color: #21005D;padding: 30px 40px;width:100%">
				<view>
					<text
						style="font-weight: 400;font-size: 22px;line-height: 28px;color: #FFFFFF;margin-right: 10px;">转赠</text>
					<text
						style="font-weight: 400;font-size: 12px;line-height: 16px;color: #D0BCFF;">请注意：一旦转赠成功将无法回退</text>
				</view>

				<u-gap height="45"></u-gap>

				<u-input v-model="handselData.phone" placeholder="请输入接收方手机号码" prefixIcon="gift"
					prefixIconStyle="font-size: 22px;color: #fff">
				</u-input>

				<u-gap height="100"></u-gap>

				<view style="display:flex">
					<u-button color="#37419C" shape="circle" text="取消" @click="toggleHandsel"
						:customStyle="{'marginRight':'25px'}"></u-button>
					<u-button color="#01C5A1" shape="circle" text="确认转赠" @click="handselHandler"></u-button>
				</view>
			</view>
		</u-popup>

		<u-popup :safeAreaInsetBottom="false" :round="20" :show="shareShow" mode="center" @close="toggleShare"
			:closeOnClickOverlay="true">
			<view style="background-color: #21005D;padding: 36rpx 16rpx 20rpx;width:100%;
		display: flex;
		flex-direction: column;
		align-items: center;border-radius: 16px;" @longtap="saveImage">
				<u-image src="/static/qrcode2.png" width="294" height="449"></u-image>
			</view>
		</u-popup>

		<Tabbar />
	</view>
</template>

<script>
	// import ProfileHeader from '../../components/ProfileHeader/ProfileHeader.vue';
	import NftCard from '../../components/NftCard/NftCard.vue';
	import Tabbar from '../../components/Tabbar/Tabbar.vue';

	import {
		accountInfoReq,
		publishListRqe,
		likeListReq,
		buyListReq,
		handsel2OtherReq
	} from '../../api/index.js';
	export default {
		components: {
			// ProfileHeader,
			NftCard,
			Tabbar
		},
		data() {
			return {
				tablist: [{
						name: '藏品',
					},
					// {
					// 	name: '已卖出',
					// },
					// {
					// 	name: '已发布',
					// }
					{
						name: '转赠记录',
					},
					{
						name: '全部订单',
					}
				],
				activeTabStyle: {
					color: '#fff',
					background: '#6750A4',
					borderRadius: '6px',
					padding: '6px 20px',
					whiteSpace: 'nowrap'
				},
				inactiveTabStyle: {
					color: '#D0BCFF',
					padding: '6px 20px',
					border: '1px solid #6750A4',
					borderRadius: '6px',
					whiteSpace: 'nowrap'
				},
				profileInfo: {
					avatar: '',
					nickname: '',
					role: ''
				},
				activeTab: 0,

				page_p: 1,
				pageSize_p: 20,
				isReq_p: false,
				count_p: 21,
				nftList_p: [],

				page_l: 1,
				pageSize_l: 20,
				isReq_l: false,
				count_l: 21,
				nftList_l: [],

				page_b: 1,
				pageSize_b: 20,
				isReq_b: false,
				count_b: 21,
				nftList_b: [],

				handselShow: false,
				handselData: {
					phone: '',
					id: null
				},
				shareShow: false

			};
		},

		onLoad(options) {

			if (options.active) {
				let active = parseInt(options.active);
				this.activeTab = active;
			}



		},
		onShow(options) {
			this.getAccountInfo();
			let active = 0;
			if (getApp().globalData.paramsData) {

				let paramsData = getApp().globalData.paramsData;
				let {
					active
				} = paramsData;
				if (paramsData.active) {
					active = parseInt(active);
				}
			}

			this.activeTab = active;

			switch (active) {
				// case 2:
				// 	this.getPublishList() // 发布
				// 	break;
				case 2:
					uni.navigateTo({
						url: '/pages/order/order'
					}); // 全部订单
					break;
				case 0:
					this.page_b = 1;
					this.pageSize_b = 20;
					this.isReq_b = false;
					this.count_b = 21;
					this.getBuyList();
					break;
					// default:
					// 	this.getLikeList() //卖出
					// 	break;
				default:
					uni.navigateTo({
						url: '/pages/handselList/handselList'
					}); //转赠记录
					break;
			}







		},
		watch: {

			// activeTab(newval) {

			// 	switch (newval) {
			// 		// case 2:
			// 		// 	this.getPublishList() // 发布
			// 		// 	break;
			// 		case 2:
			// 			uni.navigateTo({ url: '/pages/order/order' }); // 全部订单
			// 			break;
			// 		case 0:
			// 			this.getBuyList(); //藏品
			// 			break;
			// 			// default:
			// 			// 	this.getLikeList() //卖出
			// 			// 	break;
			// 		default:
			// 			uni.navigateTo({ url: '/pages/handselList/handselList' }); //转赠记录
			// 			break;
			// 	}

			// }
		},
		methods: {
			toggleShare() {
				this.shareShow = !this.shareShow;
			},
			saveImage() {
				let url = require('../../static/qrcode2.png');


				uni.saveImageToPhotosAlbum({
					filePath: url, //传入临时地址
					success: () => {
						uni.$u.toast('保存成功'); //封装的提示
						this.shareShow = false;
					},
					fail: () => {
						uni.$u.toast('保存失败');
					}
				});

				return;
				// TODO 这个url是下载接口


				uni.downloadFile({ //获得二维码的临时地址
					url,
					header: {
						token: uni.getStorageSync('token')
					},
					success: res => {
						//console.log('获取url',res)
						if (res.statusCode == 200) {
							uni.saveImageToPhotosAlbum({
								filePath: res.tempFilePath, //传入临时地址
								success: () => {
									uni.$u.toast('保存成功'); //封装的提示
								},
								fail: () => {
									uni.$u.toast('保存失败');
								}
							});
						}
					},
					fail: err => {
						console.log(err);
					}
				});
			},
			toggleHandsel() {
				this.handselShow = !this.handselShow;
			},
			handselHandler() {
				const {
					phone,
					id
				} = this.handselData;
				if (!phone || !uni.$u.test.mobile(phone)) {
					uni.$u.toast('手机号不正确！');
					return;
				}

				handsel2OtherReq({
					phone,
					id
				}).then(res => {
					this.handselShow = false;
					uni.$u.toast('转赠成功！');

					const {
						nftList_b
					} = this;

					let index = -1;
					nftList_b.map((nft, i) => {
						if (nft.id === id) {
							index = i;
						}
					});

					if (index > -1) this.nftList_b.splice(index, 1);
					// this.page_b = 1;
					// this.pageSize_b = 20;
					// this.isReq_b = false;
					// this.count_b = 21;
					// this.getBuyList();
				});
			},
			handselInfo(id) {
				this.handselData = {
					...this.handselData,
					id
				};

				this.handselShow = true;
			},
			getAccountInfo() {
				accountInfoReq().then(res => {
					const {
						avatar,
						nickname,
						id_card_num,
						role
					} = res;
					this.profileInfo = {
						nickname,
						avatar,
						id_card_num,
						role
					};
				});
			},

			getPublishList() {

				const {
					page_p,
					pageSize_p,
					isReq_p,
					count_p
				} = this;
				if (pageSize_p * page_p >= count_p) return;
				if (isReq_p) return;

				this.isReq_p = true;
				publishListRqe({
					page: page_p,
					page_size: pageSize_p,
				}).then(res => {
					const {
						result,
						count
					} = res;
					this.count_p = count;
					if (result && result.length > 0) {
						if (page_p === 1) {
							this.nftList_p = result;
						} else {
							this.nftList_p.splice(this.nftList_p.length, 0, ...result);
						}
					}
					this.page_p++;
					this.isReq_p = false;
				}).catch(() => {
					this.isReq_p = false;
				});
			},
			getLikeList() {
				const {
					page_l,
					pageSize_l,
					isReq_l,
					count_l
				} = this;
				if (pageSize_l * page_l >= count_l) return;
				if (isReq_l) return;

				this.isReq_l = true;
				likeListReq({
					page: page_l,
					page_size: pageSize_l,
				}).then(res => {
					const {
						result,
						count
					} = res;
					this.count_l = count;
					if (result && result.length > 0) {
						if (page_l === 1) {
							this.nftList_l = result;
						} else {
							this.nftList_l.splice(this.nftList_l.length, 0, ...result);
						}
					}
					this.page_l++;
					this.isReq_l = false;
				}).catch(() => {
					this.isReq_l = false;
				});
			},
			getBuyList() {
				const {
					page_b,
					pageSize_b,
					isReq_b,
					count_b
				} = this;
				if (pageSize_b * page_b >= count_b) return;
				if (isReq_b) return;

				this.isReq_b = true;
				buyListReq({
					page: page_b,
					page_size: pageSize_b,
				}).then(res => {
					const {
						result,
						count
					} = res;
					this.count_b = count;
					if (result && result.length > 0) {
						// console.log(result)
						this.nftList_b = result;

						
						// if (page_b === 1) {
						// 	this.nftList_b = result;
						// } else {
						// 	this.nftList_b.splice(this.nftList_b.length, 0, ...result);
						// }
					}
					this.page_b++;
					this.isReq_b = false;
				}).catch(() => {
					this.isReq_b = false;
				});
			},

			tabClickHandler(e) {
				const {
					index
				} = e;
				this.activeTab = index;
				// switch (index) {
				// 	// case 2:
				// 	// 	this.getPublishList() // 发布
				// 	// 	break;
				// 	case 0:
				// 		this.getBuyList(); //藏品
				// 		break;
				// 		// default:
				// 		// 	this.getLikeList() //卖出
				// 		// 	break;
				// }
				switch (index) {
					// case 2:
					// 	this.getPublishList() // 发布
					// 	break;
					case 2:
						uni.navigateTo({
							url: '/pages/order/order'
						}); // 全部订单
						break;
					case 0:
						this.getBuyList(); //藏品
						break;
						// default:
						// 	this.getLikeList() //卖出
						// 	break;
					default:
						uni.navigateTo({
							url: '/pages/handselList/handselList'
						}); //转赠记录
						break;
				}

			},
			link2mint() {
				uni.$u.toast('敬请期待！');
				// const {
				// 	id_card_num
				// } = this.profileInfo;
				// if (id_card_num) {
				// 	uni.navigateTo({
				// 		url: '/pages/mint/mint'
				// 	})
				// } else {
				// 	uni.$u.toast('实名认证后方可铸件，请先进行实名认证！');
				// 	setTimeout(() => {
				// 		uni.navigateTo({
				// 			url: '/pages/realNameAuth/realNameAuth'
				// 		})
				// 	}, 2000)
				// }
			},
			link2Profile() {
				uni.navigateTo({
					url: '/pages/myProfile/myProfile'
				});
			}

		}
	};
</script>

<style scoped>
	.nav_bar>>>.u-navbar__content__left {
		display: none;
	}

	.profile_container {
		padding: 48rpx 28rpx calc(env(safe-area-inset-bottom) / 2 + 24rpx);
	}

	.intr_info>>>span {
		text-align: center;
		margin-top: 10rpx;
	}


	.nft_card {
		margin-bottom: 24rpx;
	}

	.nft_card>>>.u-transition {
		width: 282rpx;
		height: 336rpx;
		margin: auto;
	}

	.nav_bar>>>.u-navbar__content__right {
		right: 48rpx;
		top: 0px;
		padding: 0 !important;
	}

	.u-nav-slot {
		display: flex;
		align-items: center;
		justify-content: center;
		/* background: #FEFEFE; */
		/* border: 1px solid #F5F5F5; */
		border-radius: 50%;
		padding: 0 !important;
		width: 46px;
		height: 46px;
	}

	.nfts_list {
		height: calc(100vh - 140px - env(safe-area-inset-bottom)) !important;
	}

	.nfts_list>>>.uni-scroll-view-content>view {
		display: flex;
		flex-direction: row;
		flex-wrap: wrap;
		justify-content: space-between;
		padding-top: 20rpx;
	}

	.avatar {
		padding: 30rpx 20rpx 0rpx 20rpx;


	}

	.mint {
		padding-left: 20rpx;
		margin-left: auto;

	}

	.person_name {
		padding: 20rpx 0rpx 0rpx 0rpx;
	}

	.nav_slot_right {
		display: flex;
	}
</style>
