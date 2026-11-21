<script>
	import {
		mapState,
		mapMutations
	} from 'vuex';
	import {
		versionReq,
		accountInfoReq
	} from '@/api/index.js'
	export default {
		computed: {
			...mapState({
				value: state => state.tabbar.active,
				tablist: state => state.tabbar.list
			}),
		},
		globalData: {
			switchId: 0 // 默认参数 也可设置为空
		},
		onLaunch: function(option) {
			// #ifdef APP
			this.checkArguments(); // 检测启动参数
			this.checkToken(option)
			this.checkUpdate()
			plus.globalEvent.addEventListener('newintent', (e) => {
				this.checkArguments(); // 检测启动参数  
				this.checkToken(option)
				this.checkUpdate()
			});
			// #endif
			// #ifdef APP-PLUS

			if (uni.getSystemInfoSync().platform === 'android') {
				const Color = plus.android.importClass("android.graphics.Color");
				plus.android.importClass("android.view.Window");
				const mainActivity = plus.android.runtimeMainActivity();
				const window_android = mainActivity.getWindow();
				window_android.setNavigationBarColor(Color.parseColor("#000000"));
			}
			//  判断是否有更新



			// #endif

		},
		onShow: function(option) {
			this.getUserInfo();
		},
		onHide: function() {},
		methods: {
			...mapMutations(['changeActiveIndex']),
			checkUpdate() {
				const version = plus.runtime.versionCode;
				versionReq().then(res => {
					const {
						newversion,
						enforce
					} = res;


					// enforce 0 不强制
					if (newversion > version) {
						let showCancel = false;
						if (!enforce)
							showCancel = true
						uni.showModal({
							title: '提示',
							content: '新版本已经准备好，是否立刻更新？',
							showCancel,
							success: (res) => {
								if (res.confirm) {
									// console.log('用户点击确定');
									// 强制更新
									const url = 'https://download.mareboreum.com'
									plus.runtime.openURL(url, () => {
										uni.$u.toast('打开下载页面失败！');
									});
									if (!showCancel)
										this.checkUpdate()
								} else if (res.cancel) {
									// console.log('用户点击取消');
								}
							}
						});
					}
				}).catch(err => {
					console.log(err)
				})

			},
			checkToken(option) {
				const token = uni.getStorageSync('token');
				if (!token) {
					uni.navigateTo({
						url: '/pages/welcome/welcome',
						fail: res => {
							console.log(res);
						}
					});
					this.changeActiveIndex(0);
				} else {
					let routes = ['pages/index/index', 'pages/market/market', 'pages/collection/collection']
					let index = routes.indexOf(option.path);
					const {
						value,
						tablist
					} = this;
					if (index > -1 && index != value) {

						const {
							url
						} = tablist[value];
						uni.switchTab({
							url
						});
					}
				}
			},
			checkArguments() {
				var args = plus.runtime.arguments;
				// store.commit('pay/buyStateHandler',1)
				// this.buyStateHandler(1);
				if (args && args.indexOf('beihai://') > -1) {

					// 处理args参数，如直达到某新页面等  
					let id = 0;
					if (args.indexOf('?active') > -1) {
						id = args.split('?active=')[1]
					}
					getApp().globalData.switchId = id
					const url = args.substr(8)
					if (url === '/pages/mint/mint') {
						// 没有铸件之后这个不存在了
						uni.switchTab({
							url: '/pages/collection/collection'
						})
					} else {
						uni.switchTab({
							url: '/pages/collection/collection?active=2',
							success: function(res) {
								getApp().globalData.paramsData = {
									active: 0
								};
							}
						})
					}

					this.changeActiveIndex(2)


				}

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
						"is_wallet": is_wallet,
						"id": id,
						"is_handsel": is_handsel,
						"authentication": authentication
					};
					this.$storea.commit('updataUser', userinfo);
				});
			},
		}
	};
</script>

<style lang="scss">
	/* 注意要写在第一行，同时给style标签加入lang="scss"属性 */
	@import "@/uni_modules/uview-ui/index.scss";
	@import "@/common/css.scss";
</style>

<style>
	/*每个页面公共css */
	page {
		font-family: 'Inter';
		font-style: normal;
		background-color: #21005D;
		color: #fff;
	}

	/* navbar */


	.nav_bar>>>.u-navbar__content__left {
		background: transparent;
		/* border: 1px solid #F5F5F5; */
		border-radius: 50%;
		padding: 0 !important;
		width: 46px;
		height: 46px;
		/* left: 48rpx !important; */
		top: 0px !important;
		z-index: 971;

	}


	.nav_bar>>>.u-navbar__content__left .u-icon {
		margin: auto;
	}

	.nav_bar>>>.u-navbar__content__right {
		z-index: 971;
	}

	.custom-primary_btn {
		background: #2E2F3A;
		border-radius: 4px;
	}

	/* >>>.u-button {
		border-radius: 8px !important;
	} */

	.custom_container {
		padding: 48rpx 48rpx calc(env(safe-area-inset-bottom) / 2 + 24rpx);
	}

	.u-divider {
		width: 100%;
	}

	>>>.u-form-item__body__left__content__label {
		color: #fff;
	}

	>>>.u-input {
		padding: 32rpx !important;
		border: 1px solid #EBEBF0 !important;
		border-radius: 16rpx !important;
		background-color: transparent !important;
	}

	>>>.u-textarea {
		padding: 16px !important;
		border: 1px solid #EBEBF0 !important;
		border-radius: 8px !important;
		background-color: transparent !important;
	}


	>>>.uni-input-input,
	>>>.u-textarea__field {
		color: #fff !important;
	}

	>>>.u-form-item__body__left__content__label {
		color: #fff !important;
	}

	.custom_input {
		padding: 32rpx !important;
		border: 1px solid #EBEBF0 !important;
		border-radius: 16rpx !important;
		background-color: transparent !important;
		color: #fff;
		width: 100%;
	}

	>>>.u-popup__content {
		width: 80%;
		background-color: transparent;
	}

	>>>.u-popup, >>>.u-fade-zoom-enter-active {
		background-color: transparent;
	}
</style>
