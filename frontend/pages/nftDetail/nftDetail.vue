<template>
	<view>
		<u-navbar :titleStyle="{'color': '#fff'}" bgColor="#ffffff00" class="nav_bar" :autoBack="true"
			leftIconColor="transparent">
			<!-- <u-icon name="arrow-leftward" color="#fff" slot="left"></u-icon> -->
			<u-icon name="close-circle-fill" size="33" color="rgba(0, 0, 0, 0.8)" slot="right" @click="goBack"></u-icon>

		</u-navbar>
		<!-- :title="nftInfo.name" -->
		<u-image @load="imageLoad" :height="height" width="100vw" :src="nftInfo.image" :fade="true" duration="450"
			mode="widthFix">
			<template v-slot:loading>
				<view
					style="width: 100vw;height: 100vw;display: flex;align-items: center;justify-content: center;background-color: #21005D;">
					<u-loading-icon color="white"></u-loading-icon>
				</view>
			</template>
		</u-image>
		<view>
			<view class="title_container">
				<u-text class="title" :text="nftInfo.name" size="28" color="#FFFFFF"></u-text>

				<view class="count_container">
					<u-text :text="`发行：${nftInfo.stock}份`" color="#7B61FF" size="14" align="center"></u-text>
				</view>
			</view>
		</view>
		<view class="desc_container">


			<u-gap height="23"></u-gap>

			<u-text class="main_title" text="创作故事" size="16" color="#FFFFFF"></u-text>
			<u-gap height="20"></u-gap>
			<!-- <u-text class="main_content" :text="nftInfo.description" size="14" color="#D0BCFF"></u-text> -->
			<view v-html="nftInfo.description"></view>
			<u-gap height="33"></u-gap>

			<template v-if="nftInfo.hash">
				<u-text text="链上数据" color="#FFFFFF" size="16"></u-text>
				<u-text :text="`交易哈希：${nftInfo.hash.substr(0,8)}...${nftInfo.hash.substr(50)}`" color="#D0BCFF"
					size="14"></u-text>
				<u-text :text="`链上标识：${nftInfo.token_id}`" color="#D0BCFF" size="14"></u-text>
				<u-gap height="33"></u-gap>
			</template>

			<u-text class="main_title" text="藏家须知" size="16" color="#EADDFF"></u-text>
			<u-gap height="8"></u-gap>
			<u-text class="main_content" text="1. 数字藏品为虚拟数字商品，并非实物，仅限年满18周岁的用户购买或收藏。" size="14" color="#D0BCFF">
			</u-text>
			<u-text class="main_content" text="2. 数字藏品一经售出，不支持退换，也不提供源文件下载。如不能接受，请勿购买或收藏。" size="14" color="#D0BCFF">
			</u-text>
			<u-text class="main_content" text="3. 请勿对数字藏品进行任何形式的炒作、场外交易、欺诈，或以任何其他非法方式进行使用。" size="14" color="#D0BCFF">
			</u-text>
			<u-text class="main_content" text="4. 数字藏品的版权由创作者或发行方拥有，除另行取得版权拥有者书面同意外，用户不得将数字藏品用于任何商业用途。" size="14"
				color="#D0BCFF"></u-text>


			<u-gap height="23"></u-gap>





		</view>
		<view class="author_container">
			<!-- <u-avatar size="56" :src="nftInfo.author_avatar"></u-avatar> -->
			<u-image :src="nftInfo.author_avatar" width="56" height="56" radius="50%"></u-image>
			<u-gap height="18"></u-gap>
			<u-text class="person_name" align="center" :text="`${nftInfo.author_name}`" size="14" color="#EADDFF">
			</u-text>
			<u-gap height="18"></u-gap>

			<template v-if="nftInfo.author_desc">
				<view v-html="nftInfo.author_desc"></view>
			</template>

			<u-text v-else :text="nftInfo.author_desc?nftInfo.author_desc:'ta真的是个神秘的艺术家呢！'" color="#7F67BE" size="12"
				lineHeight="16">
			</u-text>


		</view>
		<u-gap height="200"></u-gap>
		<l-painter v-if="isShowPainter" :isRenderImage="true" :board="drawJson" @success="renderImage"
			custom-style="position: fixed; left: 200%;" useCORS></l-painter>

		<u-popup :show="share_popup" mode="center" @close="close">
			<view class="share_img_container">
				<u-image mode="widthFix" :showLoading="true" :src="share_image" width="375px" height="557px"></u-image>
			</view>

		</u-popup>

		<u-popup :show="share_popup" :round="10">
			<view class="share_way_container">
				<view v-for="way in shareWayList" :key="way.name">
					<u-icon :name="way.icon" size="28" :color="way.color"></u-icon>
					<text>{{way.name}}</text>
				</view>
			</view>

		</u-popup>

		<!-- <template v-if="!isHandsel"> -->
			<!-- <view class="btn_group" v-if="!isReal">
				<u-button text="实名认证后方可购买,请先进行实名认证" color="#01C5A1" shape="circle" size="large" @click="link2real">
				</u-button>
			</view> -->

			<template v-if="isReal">
				<view class="btn_container" v-if="!isSoldOut">
					<template v-if="nftInfo.countdown <= 0">
						<template v-if="nftInfo.max_supply > 0">
							<u-number-box v-model="value" :max="max">
								<view slot="minus" class="minus">
									<u-icon name="minus" size="14" color="#FFFFFF"></u-icon>
								</view>
								<text slot="input" style="width: 50px;text-align: center;" class="input">{{value}}</text>
								<view slot="plus" class="plus">
									<u-icon name="plus" color="#FFFFFF" size="14"></u-icon>
								</view>
							</u-number-box>
							<u-gap height="25"></u-gap>
						</template>
						<view class="btn_group_new">
							<u-text :text="btn_text" color="#7E7695" size="14"></u-text>
							<u-button 
								color="#21242D" 
								shape="circle" 
								:disabled="btn_disabled || nftInfo.max_supply === 0"
								:text="(nftInfo.price*value).toFixed(2)" 
								size="normal" 
								@click="buyHandler"
								:customStyle="{'color':'#7B61FF','fontSize':'18px'}"
							></u-button>
						</view>
					</template>
					<view class="btn_group_new" v-else>
						<u-text text="抢购倒计时" color="#7E7695" size="14"></u-text>
						<view class="countdown_container_new">
							<u-icon name="clock" size="16" color="#7B61FF"></u-icon>
							<u-count-down :time="nftInfo.countdown*1000" format="HH:mm:ss"></u-count-down>
						</view>
					</view>
				</view>

				<view class="btn_group" v-else>
					<u-button text="已售罄" color="#21242D" shape="circle" size="normal">
					</u-button>
				</view>

			</template>


		<!-- </template> -->


		<!-- <template v-else>
			<view class="btn_group" v-if="!!isHandsel && chain_status == 1">
				<u-button color="#01C5A1" shape="circle" text="转赠" size="large" @click="toggleHandsel"></u-button>
			</view>
		</template> -->
		<!--  #ifdef  APP-PLUS -->
		<!-- <u-button color="#8b8b8b" shape="circle" text="商品锁定中，待支付" size="large"></u-button> -->
		<!-- <u-button color="#B00000" shape="circle" text="下架作品" size="large"></u-button> -->
		<!--  #endif -->
		<!--  #ifndef  APP-PLUS -->

		<!-- <wx-open-launch-app id="launch-btn" appid="wxb24995f3cd67c573" extinfo="">
			<script type="text/wxtag-template">
				<style>.btn { padding: 12px }</style>
				    <button class="btn">App内查看</button>
				  </script>
		</wx-open-launch-app> -->
		<!--  #endif -->

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

	</view>
</template>

<script>
	import {
		mapState
		// mapMutations
	} from 'vuex';
	import Like from '../../components/Like/Like.vue';
	import {
		prodDetailReq,
		prodLikeReq,
		buyReq,
		accountInfoReq,
		handsel2OtherReq
	} from '../../api/index.js';

	import {
		getDrawJson
	}
	from '../../utils/utils.js';

	import lPainter from '../../components/LimePainter/index.vue';

	// import wx from 'weixin-js-sdk'
	export default {
		components: {
			Like,
			lPainter
		},
		data() {
			return {
				id: null,
				isHandsel: false,
				chain_status: null,
				height: '100vw',
				nftInfo: {
					author_avatar: '',
					author_id: '',
					author_name: '',
					create_time: '',
					description: '',
					id: '',
					image: '',
					liked: '',
					likes: '',
					name: '',
					owner_avatar: '',
					owner_id: '',
					owner_name: '',
					price: '',
					sales: '',
					stock: '',
					token_id: '',
					can_buy: false,
					max_supply: null
				},
				isShowPainter: false,
				drawJson: null,
				share_popup: false,
				share_image: '',
				shareWayList: [{
						name: '保存图片',
						icon: 'download',
						color: '#333333f1'
					},
					{
						name: '微信',
						icon: 'weixin-fill',
						color: '#09bb07'
					},
					{
						name: '朋友圈',
						icon: require('@/static/share/timeline.png'),
						color: ''
					},
					{
						name: '口令',
						icon: 'tags-fill',
						color: '#bc42ff'
					}
				],
				isReal: false,
				value: 1,
				btn_text: '',
				btn_disabled: false,


				handselShow: false,
				handselData: {
					phone: '',
					id: null
				},

				isSoldOut: false, //是否售罄


			};
		},
		computed: {
			likeInfo() {
				const {
					likes,
					liked,
					id
				} = this.nftInfo;
				return {
					isLike: liked,
					count: likes,
					id
				};
			},
			max() {
				const {
					max_supply
				} = this.nftInfo;

				if (max_supply && max_supply > 5) {
					return 5;
				} else if (max_supply && max_supply <= 5) {
					return max_supply;
				}

				return 0;
			},
			...mapState({
				b_count: state => state.pay.b_count,
			}),
		},
		watch: {
			b_count(newval, oldval) {
				if (newval > oldval) {
					uni.navigateTo({
						url: '/pages/collection/collection'
					});
				}
			}
		},
		onLoad(options) {
			let userinfo = this.$storea.state.userInfo;
			// console.log(tmp);
			console.log(options);
			const {
				id,
				// isHandsel,
				chain_status
			} = options;


			this.id = id;

			// if (isHandsel == 'true')
			// 	this.isHandsel = true;
			this.isHandsel = userinfo.is_handsel;
			if (chain_status == 1) {
				this.chain_status = parseInt(chain_status);
			}
			this.getNftDetail();
			this.getHasReal();
		},
		onShow() {
			uni.$once('auth', () => {
				this.getNftDetail();
				this.getHasReal();
			});
			this.btn_text = '';
			this.btn_disabled = false;
		},
		methods: {
			getHasReal() {
				accountInfoReq().then(res => {
					const {
						authentication
					} = res;
					if (authentication) {
						this.isReal = true;
					}

				});
			},
			close() {
				this.share_popup = false;
			},
			imageLoad() {
				this.height = 'auto';
			},
			getNftDetail() {
				const {
					id
				} = this;
				prodDetailReq(id).then(res => {
					console.log(res);
					this.nftInfo = res;
					this.isSoldOut = res.sold_out;
					
					this.btn_text = `限购${res.max_supply}份`;
				});
			},
			heartClickHandler(type) {
				const {
					id,
					liked,
					likes
				} = this.nftInfo;
				if (type) {
					// if(liked)this.nftInfo.likes=likes-1;
					// else this.nftInfo.likes=likes+1;
					this.nftInfo.liked = !liked;
				} else {
					prodLikeReq(id).then(() => {
						if (liked) this.nftInfo.likes = likes - 1;
						else this.nftInfo.likes = likes + 1;
						this.nftInfo.liked = !liked;
					});
				}
			},
			buyHandler() {
				this.btn_text = '商品锁定中';
				this.btn_disabled = true;

				const {
					id,
					value
				} = this;

				buyReq({
					// pid: parseInt(id),
					// pay_type: 1,
					id: parseInt(id),
					num: value
				}).then(res => {
					const pay_url = res;
					plus.runtime.openURL(pay_url, () => {
						// plus.runtime.openWeb(pay_url, () => {
						// 调用失败 callback
						uni.$u.toast('调起支付失败，请稍后再试！');
						this.btn_text = `限购${this.nftInfo.max_supply}份`;
						this.btn_disabled = false;
						this.getNftDetail();
					});
					// let {
					// 	params
					// } = res;
					// params = JSON.parse(params)
					// if (params.resp_code === 'C00000') {
					// 	const {
					// 		pay_url
					// 	} = params
					// 	plus.runtime.openWeb(pay_url, () => {
					// 		// 调用失败 callback
					// 		uni.$u.toast('调起支付失败，请稍后再试！');
					// 	});
					// } else {
					// 	uni.$u.toast(params.resp_desc);
					// }
				}).catch(() => {
					this.btn_text = '购买';
					this.btn_disabled = false;
					this.getNftDetail();
				});
			},
			shareHandler() {
				this.drawJson = getDrawJson({
					avatar: this.nftInfo.author_avatar, //'../../static/test/avatar.jpeg',
					nickname: this.nftInfo.author_name,
					cover: this.nftInfo.image, //'../../static/test/cover.png',
					price: this.nftInfo.price,
					name: this.nftInfo.name,
					mp_qrcode: '../../static/test/avatar.jpeg',

				});
				this.isShowPainter = true;
				// uni.share({
				// 	provider: "weixin",
				// 	scene: "WXSceneSession",
				// 	type: 2,
				// 	imageUrl: "https://bjetxgzv.cdn.bspapp.com/VKCEYUGU-uni-app-doc/d8590190-4f28-11eb-b680-7980c8a877b8.png",
				// 	success: function(res) {
				// 		console.log("success:" + JSON.stringify(res));
				// 	},
				// 	fail: function(err) {
				// 		console.log("fail:" + JSON.stringify(err));
				// 	}
				// });
			},
			renderImage(e) {
				this.share_image = e;
				this.isShowPainter = false;

				this.share_popup = true;
				console.log(e);
			},
			link2real() {
				uni.navigateTo({
					url: '/pages/realNameAuth/realNameAuth'
				});
			},
			toggleHandsel() {
				this.handselShow = !this.handselShow;
			},
			handselHandler() {
				const {
					phone
				} = this.handselData;
				let {
					id
				} = this;
				id = parseInt(id);
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
					this.getNftDetail();
				});
			},
			goBack() {
				uni.navigateBack()
			}
		}
	};
</script>

<style scoped>
	/* .nav_bar>>>.u-navbar__content__left {
		background: linear-gradient(98.21deg, rgba(255, 255, 255, 0.264) -2.97%, rgba(255, 255, 255, 0.066) 103.09%);
		backdrop-filter: blur(20px);
		width: 48px;
		height: 48px;
		border-width: 0;
	}

	.nav_bar>>>.u-navbar__content__right {
		right: 48rpx;
		top: 0px;
		padding: 0 !important;
	} */

	.u-nav-slot {
		display: flex;
	}

	.icon_container {
		background: linear-gradient(98.21deg, rgba(255, 255, 255, 0.264) -2.97%, rgba(255, 255, 255, 0.066) 103.09%);
		backdrop-filter: blur(20px);
		width: 48px;
		height: 48px;
		border-width: 0;
		border-radius: 50%;
		padding: 0 !important;
		display: flex;
		align-items: center;
		justify-content: center;
	}

	.icon_container:first-child {
		margin-right: 16rpx;
	}

	.icon_container:last-child {
		transform: rotate(90deg);
	}

	.desc_container {
		padding: 32rpx 28rpx calc(env(safe-area-inset-bottom) / 2 + 24rpx);
	}

	.title_container,
	.title-right_container {
		display: flex;
		align-items: flex-start;
	}

	.title_container {
		background: #180240;
		border-radius: 0px 0px 12px 12px;
		padding: 25px 10px;
	}

	.title_container .title {
		flex: auto;
	}

	.title_container .title>>>span {
		font-weight: 600;
		line-height: 32px;
		letter-spacing: 0.005em;
	}

	.title-right_container .u-button {
		margin-left: 32rpx;
	}

	.count_container {
		display: flex;
		background: #21242D;
		border-radius: 20px;
		padding: 2px 15px;
		box-sizing: border-box;
		min-width: 120px;
	}

	.count_container>>>.u-text:nth-child(2n) {
		background-color: #A38DD0;
		margin: 0 20rpx !important;
		padding: 0 10rpx;
	}

	.count_container>>>.u-text {
		flex: 0 0 auto;
		width: auto;
		margin-right: 10rpx !important;
	}

	.link_container {
		background: #180240;
		border-radius: 12px;
		padding: 20rpx 40rpx;
		margin-bottom: 20rpx;
	}

	.author_container {
		background: #180240;
		/* border-radius: 12px; */
		margin-bottom: 20rpx;
		padding: 20px;
	}

	.author_container .author_header {
		display: flex;
		justify-content: space-between;

	}

	.author_container .author_footer {
		border-top: 1px solid #381E72;
	}

	.author_container>>>.u-transition {
		display: flex;
		align-items: center;
	}



	.preson_title>>>span {
		font-weight: 400;
		line-height: 20px;
		letter-spacing: 0.005em;
	}

	.person_container {
		display: flex;
		align-items: center;
	}

	.person_container .person_name>>>span {
		font-weight: 500;
		line-height: 20px;
		letter-spacing: 0.005em;
		margin-left: 16rpx;
	}

	.main_title>>>span {
		font-weight: 600;
		line-height: 22px;
		letter-spacing: 0.005em;
	}

	.main_content>>>span {
		font-weight: 400;
		line-height: 22px;
		letter-spacing: 0.005em;

	}

	.btn_group {
		position: fixed;
		display: flex;
		bottom: 0;
		left: 0;
		width: 100vw;
		/* width: calc(100vw - 1rpx); */
		background-color: #21005D;
	/* 	background-color: #ffff; */
		justify-content: center;
		padding:16rpx;
		/* padding: 48rpx 48rpx calc(env(safe-area-inset-bottom) / 2 + 24rpx); */
	}

	.btn_container {
		position: fixed;
		display: flex;
		flex-direction: column;
		bottom: 0;
		left: 0;
		width: 100vw;
		justify-content: center;
		/* width: calc(100vw - 98rpx); */
		background-color: #21005D;
		padding:16rpx;
		/* padding: 48rpx 48rpx calc(env(safe-area-inset-bottom) / 2 + 24rpx); */

	}

	.btn_container .u-number-box {
		justify-content: center;
	}

	.btn_container .minus,
	.btn_container .plus {
		background: #37419C;
		width: 30px;
		height: 30px;
		border-radius: 50%;
		display: flex;
		align-items: center;
		justify-content: center;
	}

	.btn_group_new {
		display: flex;
		padding-right: 24rpx;
	}

	.btn_group uni-text,
	.btn_group_new uni-text {
		flex: 1;
		margin-right: 16px;
		display: flex;
		align-items: center;
		justify-content: center;




		background: linear-gradient(98.46deg, #00F0FF -0.18%, #51EC44 106.55%);

		-webkit-background-clip: text;
		-webkit-text-fill-color: rgba(255, 255, 255, 0);
		background-clip: text;
		text-fill-color: rgba(255, 255, 255, 0);
		font-size: 22px;
		white-space: nowrap;


	}

	.btn_group .u-button:first-child,
	.btn_group_new .u-button:first-child {
		flex: 1;
		margin-right: 16px;
	}

	.btn_group .u-button:last-child,
	.btn_group_new .u-button:last-child {
		flex: 1;
	}

	.share_img_container {}

	.share_way_container {
		display: flex;
		padding: 40rpx;
		justify-content: space-around
	}

	.share_way_container>view {
		display: flex;
		flex-direction: column;
		align-items: center;
	}

	.share_way_container>view>>>.u-icon {
		width: 80rpx;
		height: 80rpx;
		border-radius: 50%;
		background-color: #ebebebf1;
		display: flex;
		align-items: center;
		justify-content: center;
	}

	.countdown_container_new {
		display: flex;
		flex: 2;
		justify-content: center;
		align-items: center;
		background-color: #21242D;
		border-radius: 100px;
		padding: 30rpx;


	}

	>>>.u-count-down__text {
		color: #7B61FF;
		font-size: 14px;
		margin-left: 20rpx;
	}
</style>
